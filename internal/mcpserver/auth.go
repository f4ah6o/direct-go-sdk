package mcpserver

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const authInfoKey contextKey = "mcp-auth-info"

type AuthInfo struct {
	Subject string
	Scopes  map[string]bool
}

type Authenticator struct {
	cfg        MCPConfig
	httpClient *http.Client
	now        func() time.Time
	mu         sync.Mutex
	keys       map[string]jwk
	fetchedAt  time.Time
	ttl        time.Duration
	// inflight deduplicates concurrent JWKS fetches: the first caller fetches
	// while the mutex is released, and waiters share the result.
	inflight *jwksFetch
}

type jwksFetch struct {
	done chan struct{}
	keys map[string]jwk
	err  error
}

type tokenClaims struct {
	Scope string `json:"scope,omitempty"`
	Scp   string `json:"scp,omitempty"`
	jwt.RegisteredClaims
}

type jwksResponse struct {
	Keys []jwk `json:"keys"`
}

type jwk struct {
	KID string   `json:"kid"`
	X5T string   `json:"x5t"`
	KTY string   `json:"kty"`
	Use string   `json:"use"`
	Alg string   `json:"alg"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5C []string `json:"x5c"`
}

func NewAuthenticator(cfg MCPConfig) *Authenticator {
	return &Authenticator{
		cfg:        cfg,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		now:        time.Now,
		keys:       map[string]jwk{},
		ttl:        defaultJWKSCacheTTL,
	}
}

func AuthInfoFromContext(ctx context.Context) (*AuthInfo, bool) {
	info, ok := ctx.Value(authInfoKey).(*AuthInfo)
	return info, ok
}

func RequireScope(ctx context.Context, scope string) error {
	info, ok := AuthInfoFromContext(ctx)
	if !ok {
		return ErrUnauthorized
	}
	if scope == "" || info.Scopes[scope] {
		return nil
	}
	return ErrForbidden
}

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

func (a *Authenticator) Authenticate(ctx context.Context, r *http.Request) (context.Context, error) {
	raw, err := bearerToken(r)
	if err != nil {
		return ctx, err
	}
	claims := &tokenClaims{}
	parser := jwt.NewParser(
		jwt.WithAudience(a.cfg.JWTAudience),
		jwt.WithIssuer(a.cfg.JWTIssuer),
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(2*time.Minute),
		jwt.WithTimeFunc(a.now),
	)
	token, err := parser.ParseWithClaims(raw, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodRS256.Alg() &&
			token.Method.Alg() != jwt.SigningMethodRS384.Alg() &&
			token.Method.Alg() != jwt.SigningMethodRS512.Alg() {
			return nil, fmt.Errorf("unsupported signing alg %q", token.Method.Alg())
		}
		key, err := a.keyForToken(ctx, token, false)
		if err != nil {
			key, err = a.keyForToken(ctx, token, true)
		}
		if err != nil {
			return nil, err
		}
		if key.Alg != "" && key.Alg != token.Method.Alg() {
			return nil, fmt.Errorf("key alg %q does not match token alg %q", key.Alg, token.Method.Alg())
		}
		return key.publicKey()
	})
	if err != nil || !token.Valid {
		if err == nil {
			err = ErrUnauthorized
		}
		return ctx, fmt.Errorf("%w: %v", ErrUnauthorized, err)
	}
	info := &AuthInfo{Subject: claims.Subject, Scopes: parseScopes(claims.Scope, claims.Scp)}
	return context.WithValue(ctx, authInfoKey, info), nil
}

func bearerToken(r *http.Request) (string, error) {
	auth := strings.TrimSpace(r.Header.Get("Authorization"))
	if auth == "" {
		return "", ErrUnauthorized
	}
	typ, token, ok := strings.Cut(auth, " ")
	if !ok || !strings.EqualFold(typ, "Bearer") || strings.TrimSpace(token) == "" {
		return "", ErrUnauthorized
	}
	return strings.TrimSpace(token), nil
}

func parseScopes(values ...string) map[string]bool {
	out := map[string]bool{}
	for _, value := range values {
		for _, scope := range strings.Fields(value) {
			out[scope] = true
		}
	}
	return out
}

func (a *Authenticator) keyForToken(ctx context.Context, token *jwt.Token, refresh bool) (jwk, error) {
	keys, err := a.loadKeys(ctx, refresh)
	if err != nil {
		return jwk{}, err
	}
	for _, headerName := range []string{"kid", "x5t"} {
		value, _ := token.Header[headerName].(string)
		if value == "" {
			continue
		}
		if key, ok := keys[value]; ok {
			return key, nil
		}
	}
	return jwk{}, errors.New("signing key not found")
}

const (
	defaultJWKSCacheTTL = time.Hour
	minJWKSCacheTTL     = time.Minute
	maxJWKSCacheTTL     = time.Hour
)

// loadKeys serves cached keys when fresh, otherwise fetches the JWKS once
// per refresh round even under concurrent callers. A failed refresh keeps
// the last-known-good key set: callers still fail closed on unknown kids
// because the stale set simply will not contain them.
func (a *Authenticator) loadKeys(ctx context.Context, refresh bool) (map[string]jwk, error) {
	a.mu.Lock()
	if len(a.keys) > 0 && !refresh && a.now().Before(a.fetchedAt.Add(a.ttl)) {
		keys := a.keys
		a.mu.Unlock()
		return keys, nil
	}
	if a.inflight != nil {
		f := a.inflight
		a.mu.Unlock()
		select {
		case <-f.done:
			if f.err != nil {
				if stale := a.staleKeys(); len(stale) > 0 {
					return stale, nil
				}
				return nil, f.err
			}
			return f.keys, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	f := &jwksFetch{done: make(chan struct{})}
	a.inflight = f
	a.mu.Unlock()

	keys, ttl, err := a.fetchKeys(ctx)

	a.mu.Lock()
	a.inflight = nil
	f.keys = keys
	f.err = err
	if err == nil {
		a.keys = keys
		a.fetchedAt = a.now()
		a.ttl = ttl
	}
	a.mu.Unlock()
	close(f.done)

	if err != nil {
		if stale := a.staleKeys(); len(stale) > 0 {
			return stale, nil
		}
		return nil, err
	}
	return keys, nil
}

func (a *Authenticator) staleKeys() map[string]jwk {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.keys
}

// fetchKeys retrieves and validates the JWKS document. The response body is
// bounded by mcp.max_jwks_bytes and the resulting key set must be non-empty,
// free of duplicate identifiers, and contain only usable RSA signing keys.
func (a *Authenticator) fetchKeys(ctx context.Context) (map[string]jwk, time.Duration, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.cfg.JWKSURL, nil)
	if err != nil {
		return nil, 0, err
	}
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, 0, fmt.Errorf("GET JWKS status=%d", resp.StatusCode)
	}
	max := a.cfg.MaxJWKSBytes
	if max <= 0 {
		max = 1 << 20
	}
	if resp.ContentLength > max {
		return nil, 0, fmt.Errorf("jwks response exceeds %d bytes", max)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, max+1))
	if err != nil {
		return nil, 0, err
	}
	if int64(len(body)) > max {
		return nil, 0, fmt.Errorf("jwks response exceeds %d bytes", max)
	}
	var jwks jwksResponse
	if err := json.Unmarshal(body, &jwks); err != nil {
		return nil, 0, fmt.Errorf("malformed jwks document: %w", err)
	}
	keys, err := indexKeys(jwks.Keys)
	if err != nil {
		return nil, 0, err
	}
	return keys, jwksCacheTTL(resp.Header), nil
}

// indexKeys indexes usable RSA signing keys by kid and x5t. Keys without an
// identifier or an unusable public key are skipped; a duplicate identifier
// rejects the whole set since key selection would be ambiguous.
func indexKeys(jwksKeys []jwk) (map[string]jwk, error) {
	keys := map[string]jwk{}
	for _, key := range jwksKeys {
		if key.KTY != "" && key.KTY != "RSA" {
			continue
		}
		if key.Use != "" && key.Use != "sig" {
			continue
		}
		if _, err := key.publicKey(); err != nil {
			continue
		}
		for _, id := range []string{key.KID, key.X5T} {
			if id == "" {
				continue
			}
			if _, dup := keys[id]; dup {
				return nil, fmt.Errorf("jwks contains duplicate key identifier")
			}
			keys[id] = key
		}
	}
	if len(keys) == 0 {
		return nil, errors.New("jwks contains no usable RSA signing keys")
	}
	return keys, nil
}

// jwksCacheTTL honors Cache-Control max-age on the JWKS response, clamped to
// [minJWKSCacheTTL, maxJWKSCacheTTL] so a misconfigured issuer can neither
// stall key rotation nor induce a fetch storm. no-store/no-cache pins the
// TTL to the minimum.
func jwksCacheTTL(h http.Header) time.Duration {
	ttl := defaultJWKSCacheTTL
	for _, part := range strings.Split(h.Get("Cache-Control"), ",") {
		part = strings.TrimSpace(strings.ToLower(part))
		switch {
		case part == "no-cache" || part == "no-store":
			ttl = minJWKSCacheTTL
		case strings.HasPrefix(part, "max-age="):
			if secs, err := strconv.Atoi(strings.TrimPrefix(part, "max-age=")); err == nil && secs > 0 {
				ttl = time.Duration(secs) * time.Second
			}
		}
	}
	if ttl < minJWKSCacheTTL {
		return minJWKSCacheTTL
	}
	if ttl > maxJWKSCacheTTL {
		return maxJWKSCacheTTL
	}
	return ttl
}

func (j jwk) publicKey() (*rsa.PublicKey, error) {
	if len(j.X5C) > 0 {
		der, err := base64.StdEncoding.DecodeString(j.X5C[0])
		if err != nil {
			return nil, err
		}
		cert, err := x509.ParseCertificate(der)
		if err != nil {
			return nil, err
		}
		pub, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("x5c certificate public key is not RSA")
		}
		return pub, nil
	}
	if j.N == "" || j.E == "" {
		return nil, errors.New("rsa jwk missing n or e")
	}
	nb, err := base64.RawURLEncoding.DecodeString(j.N)
	if err != nil {
		return nil, err
	}
	eb, err := base64.RawURLEncoding.DecodeString(j.E)
	if err != nil {
		return nil, err
	}
	e := 0
	for _, b := range eb {
		e = e<<8 + int(b)
	}
	if e == 0 {
		return nil, errors.New("rsa jwk has invalid exponent")
	}
	return &rsa.PublicKey{N: new(big.Int).SetBytes(nb), E: e}, nil
}

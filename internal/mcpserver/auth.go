package mcpserver

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
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

func (a *Authenticator) loadKeys(ctx context.Context, refresh bool) (map[string]jwk, error) {
	a.mu.Lock()
	if len(a.keys) > 0 && !refresh && a.now().Sub(a.fetchedAt) < time.Hour {
		keys := a.keys
		a.mu.Unlock()
		return keys, nil
	}
	a.mu.Unlock()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.cfg.JWKSURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("GET JWKS status=%d", resp.StatusCode)
	}
	var jwks jwksResponse
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, err
	}
	keys := map[string]jwk{}
	for _, key := range jwks.Keys {
		if key.KTY != "" && key.KTY != "RSA" {
			continue
		}
		if key.Use != "" && key.Use != "sig" {
			continue
		}
		if key.KID != "" {
			keys[key.KID] = key
		}
		if key.X5T != "" {
			keys[key.X5T] = key
		}
	}
	if len(keys) == 0 {
		return nil, errors.New("jwks contains no RSA signing keys")
	}
	a.mu.Lock()
	a.keys = keys
	a.fetchedAt = a.now()
	a.mu.Unlock()
	return keys, nil
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

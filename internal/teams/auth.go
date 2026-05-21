package teams

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
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

const (
	botFrameworkIssuer = "https://api.botframework.com"
	teamsChannelID     = "msteams"
)

var emulatorIssuers = map[string]bool{
	"https://sts.windows.net/d6d49420-f39b-4df7-a1dc-d59a935871db/":               true,
	"https://login.microsoftonline.com/d6d49420-f39b-4df7-a1dc-d59a935871db/v2.0": true,
	"https://login.microsoftonline.com/botframework.com/v2.0":                     true,
}

type AuthValidator interface {
	Validate(context.Context, *http.Request, Activity) error
}

type BotFrameworkValidator struct {
	cfg        config.BotConfig
	httpClient *http.Client
	now        func() time.Time
	mu         sync.Mutex
	metadata   map[string]openidMetadata
	keys       map[string]map[string]jwk
}

type botClaims struct {
	ServiceURL  string `json:"serviceurl,omitempty"`
	ServiceURL2 string `json:"serviceUrl,omitempty"`
	jwt.RegisteredClaims
}

type openidMetadata struct {
	JWKSURI     string   `json:"jwks_uri"`
	SigningAlgs []string `json:"id_token_signing_alg_values_supported"`
	FetchedAt   time.Time
}

type jwksResponse struct {
	Keys []jwk `json:"keys"`
}

type jwk struct {
	KID          string   `json:"kid"`
	X5T          string   `json:"x5t"`
	KTY          string   `json:"kty"`
	Use          string   `json:"use"`
	Alg          string   `json:"alg"`
	N            string   `json:"n"`
	E            string   `json:"e"`
	X5C          []string `json:"x5c"`
	Endorsements []string `json:"endorsements"`
}

func NewBotFrameworkValidator(cfg config.BotConfig) *BotFrameworkValidator {
	return &BotFrameworkValidator{
		cfg:        cfg,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		now:        time.Now,
		metadata:   map[string]openidMetadata{},
		keys:       map[string]map[string]jwk{},
	}
}

func (v *BotFrameworkValidator) Validate(ctx context.Context, r *http.Request, activity Activity) error {
	tokenText, err := bearerToken(r)
	if err != nil {
		return err
	}
	if activity.ServiceURL == "" {
		return errors.New("activity serviceUrl is required")
	}
	if activity.ChannelID == "" {
		return errors.New("activity channelId is required")
	}

	var selectedKey jwk
	claims := &botClaims{}
	parser := jwt.NewParser(
		jwt.WithAudience(v.cfg.AppID),
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(5*time.Minute),
		jwt.WithTimeFunc(v.now),
	)
	token, err := parser.ParseWithClaims(tokenText, claims, func(token *jwt.Token) (interface{}, error) {
		if !allowedSigningAlg(token.Method.Alg()) {
			return nil, fmt.Errorf("unsupported signing alg %q", token.Method.Alg())
		}
		key, err := v.keyForToken(ctx, token, false)
		if err != nil {
			key, err = v.keyForToken(ctx, token, true)
		}
		if err != nil {
			return nil, err
		}
		selectedKey = key
		return key.publicKey()
	})
	if err != nil {
		return fmt.Errorf("invalid bot framework token: %w", err)
	}
	if !token.Valid {
		return errors.New("invalid bot framework token")
	}
	if err := v.validateIssuer(claims.Issuer); err != nil {
		return err
	}
	claimServiceURL := firstNonEmpty(claims.ServiceURL, claims.ServiceURL2)
	if claimServiceURL == "" {
		return errors.New("token serviceurl claim is required")
	}
	if normalizeServiceURL(claimServiceURL) != normalizeServiceURL(activity.ServiceURL) {
		return errors.New("token serviceurl does not match activity serviceUrl")
	}
	if !v.isEmulatorIssuer(claims.Issuer) && activity.ChannelID != teamsChannelID {
		return fmt.Errorf("unsupported channelId %q", activity.ChannelID)
	}
	if len(v.cfg.AllowedServiceURLs) > 0 && !serviceURLAllowed(activity.ServiceURL, v.cfg.AllowedServiceURLs) {
		return errors.New("activity serviceUrl is not allowed")
	}
	if len(selectedKey.Endorsements) > 0 && !containsString(selectedKey.Endorsements, activity.ChannelID) {
		return fmt.Errorf("signing key is not endorsed for channelId %q", activity.ChannelID)
	}
	return nil
}

func bearerToken(r *http.Request) (string, error) {
	auth := strings.TrimSpace(r.Header.Get("Authorization"))
	if auth == "" {
		return "", errors.New("missing authorization header")
	}
	typ, token, ok := strings.Cut(auth, " ")
	if !ok || !strings.EqualFold(typ, "Bearer") || strings.TrimSpace(token) == "" {
		return "", errors.New("missing bearer token")
	}
	return strings.TrimSpace(token), nil
}

func (v *BotFrameworkValidator) validateIssuer(issuer string) error {
	if issuer == botFrameworkIssuer {
		return nil
	}
	if v.cfg.AllowEmulator && v.isEmulatorIssuer(issuer) {
		return nil
	}
	return fmt.Errorf("invalid issuer %q", issuer)
}

func (v *BotFrameworkValidator) isEmulatorIssuer(issuer string) bool {
	return emulatorIssuers[issuer]
}

func (v *BotFrameworkValidator) keyForToken(ctx context.Context, token *jwt.Token, refresh bool) (jwk, error) {
	metadataURL := v.cfg.OpenIDMetadataURL
	if v.cfg.AllowEmulator {
		if iss, _ := token.Claims.GetIssuer(); v.isEmulatorIssuer(iss) {
			metadataURL = v.cfg.EmulatorOpenIDMetadataURL
		}
	}
	meta, err := v.loadMetadata(ctx, metadataURL, refresh)
	if err != nil {
		return jwk{}, err
	}
	if len(meta.SigningAlgs) > 0 && !containsString(meta.SigningAlgs, token.Method.Alg()) {
		return jwk{}, fmt.Errorf("signing alg %q is not advertised by metadata", token.Method.Alg())
	}
	keys, err := v.loadKeys(ctx, meta.JWKSURI, refresh)
	if err != nil {
		return jwk{}, err
	}
	for _, headerName := range []string{"kid", "x5t"} {
		value, _ := token.Header[headerName].(string)
		if value == "" {
			continue
		}
		if key, ok := keys[value]; ok {
			if key.Alg != "" && key.Alg != token.Method.Alg() {
				return jwk{}, fmt.Errorf("key alg %q does not match token alg %q", key.Alg, token.Method.Alg())
			}
			return key, nil
		}
	}
	return jwk{}, errors.New("signing key not found")
}

func (v *BotFrameworkValidator) loadMetadata(ctx context.Context, metadataURL string, refresh bool) (openidMetadata, error) {
	v.mu.Lock()
	meta, ok := v.metadata[metadataURL]
	if ok && !refresh && v.now().Sub(meta.FetchedAt) < 24*time.Hour {
		v.mu.Unlock()
		return meta, nil
	}
	v.mu.Unlock()

	var out openidMetadata
	if err := getJSON(ctx, v.httpClient, metadataURL, &out); err != nil {
		return openidMetadata{}, err
	}
	if out.JWKSURI == "" {
		return openidMetadata{}, errors.New("openid metadata missing jwks_uri")
	}
	out.FetchedAt = v.now()
	v.mu.Lock()
	v.metadata[metadataURL] = out
	v.mu.Unlock()
	return out, nil
}

func (v *BotFrameworkValidator) loadKeys(ctx context.Context, jwksURL string, refresh bool) (map[string]jwk, error) {
	v.mu.Lock()
	keys, ok := v.keys[jwksURL]
	if ok && !refresh {
		v.mu.Unlock()
		return keys, nil
	}
	v.mu.Unlock()

	var resp jwksResponse
	if err := getJSON(ctx, v.httpClient, jwksURL, &resp); err != nil {
		return nil, err
	}
	out := map[string]jwk{}
	for _, key := range resp.Keys {
		if key.KTY != "" && key.KTY != "RSA" {
			continue
		}
		if key.Use != "" && key.Use != "sig" {
			continue
		}
		if key.KID != "" {
			out[key.KID] = key
		}
		if key.X5T != "" {
			out[key.X5T] = key
		}
	}
	if len(out) == 0 {
		return nil, errors.New("jwks contains no RSA signing keys")
	}
	v.mu.Lock()
	v.keys[jwksURL] = out
	v.mu.Unlock()
	return out, nil
}

func getJSON(ctx context.Context, client *http.Client, rawURL string, out interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("GET %s status=%d", rawURL, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (k jwk) publicKey() (*rsa.PublicKey, error) {
	if len(k.X5C) > 0 {
		der, err := base64.StdEncoding.DecodeString(k.X5C[0])
		if err != nil {
			return nil, err
		}
		cert, err := x509.ParseCertificate(der)
		if err != nil {
			return nil, err
		}
		pub, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("x5c certificate is not RSA")
		}
		return pub, nil
	}
	nBytes, err := base64.RawURLEncoding.DecodeString(k.N)
	if err != nil {
		return nil, err
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(k.E)
	if err != nil {
		return nil, err
	}
	e := 0
	for _, b := range eBytes {
		e = e<<8 + int(b)
	}
	if e == 0 {
		return nil, errors.New("invalid RSA exponent")
	}
	return &rsa.PublicKey{N: new(big.Int).SetBytes(nBytes), E: e}, nil
}

func allowedSigningAlg(alg string) bool {
	return alg == jwt.SigningMethodRS256.Alg() || alg == jwt.SigningMethodRS384.Alg() || alg == jwt.SigningMethodRS512.Alg()
}

func serviceURLAllowed(serviceURL string, allowed []string) bool {
	normalized := normalizeServiceURL(serviceURL)
	for _, candidate := range allowed {
		if normalizeServiceURL(candidate) == normalized {
			return true
		}
	}
	return false
}

func normalizeServiceURL(raw string) string {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return strings.TrimRight(strings.TrimSpace(raw), "/")
	}
	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)
	u.Path = strings.TrimRight(u.Path, "/")
	return u.String()
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

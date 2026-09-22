package mcpserver

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"maps"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestAuthenticatorValidatesJWTAndScopes(t *testing.T) {
	key := mustRSAKey(t)
	jwks := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(jwksResponse{Keys: []jwk{publicJWK("kid-1", &key.PublicKey)}})
	}))
	defer jwks.Close()

	cfg := MCPConfig{
		JWTIssuer:   "https://auth.example.com",
		JWTAudience: "https://mcp.example.com/mcp",
		JWKSURL:     jwks.URL,
		ReadScope:   DefaultReadScope,
		WriteScope:  DefaultWriteScope,
	}
	authn := NewAuthenticator(cfg)
	authn.now = func() time.Time { return time.Unix(1000, 0) }
	token := signToken(t, key, "kid-1", tokenClaims{
		Scope: "direct:read",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user-1",
			Issuer:    cfg.JWTIssuer,
			Audience:  jwt.ClaimStrings{cfg.JWTAudience},
			ExpiresAt: jwt.NewNumericDate(time.Unix(1600, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(900, 0)),
		},
	})
	req := httptest.NewRequest(http.MethodPost, "http://localhost/mcp", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	ctx, err := authn.Authenticate(context.Background(), req)
	if err != nil {
		t.Fatalf("Authenticate() = %v", err)
	}
	if err := RequireScope(ctx, DefaultReadScope); err != nil {
		t.Fatalf("RequireScope(read) = %v", err)
	}
	if err := RequireScope(ctx, DefaultWriteScope); err != ErrForbidden {
		t.Fatalf("RequireScope(write) = %v, want ErrForbidden", err)
	}
}

func TestAuthenticatorRejectsWrongAudience(t *testing.T) {
	key := mustRSAKey(t)
	jwks := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(jwksResponse{Keys: []jwk{publicJWK("kid-1", &key.PublicKey)}})
	}))
	defer jwks.Close()

	authn := NewAuthenticator(MCPConfig{
		JWTIssuer:   "https://auth.example.com",
		JWTAudience: "https://mcp.example.com/mcp",
		JWKSURL:     jwks.URL,
	})
	authn.now = func() time.Time { return time.Unix(1000, 0) }
	token := signToken(t, key, "kid-1", tokenClaims{
		Scope: "direct:read",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "https://auth.example.com",
			Audience:  jwt.ClaimStrings{"wrong-audience"},
			ExpiresAt: jwt.NewNumericDate(time.Unix(1600, 0)),
		},
	})
	req := httptest.NewRequest(http.MethodPost, "http://localhost/mcp", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if _, err := authn.Authenticate(context.Background(), req); err == nil {
		t.Fatalf("expected wrong audience to be rejected")
	}
}

// jwksServer serves a mutable JWKS document and counts requests.
type jwksServer struct {
	*httptest.Server
	mu      sync.Mutex
	body    []byte
	status  int
	headers map[string]string
	hits    int
	block   chan struct{}
}

func newJWKSServer(t *testing.T, body []byte) *jwksServer {
	t.Helper()
	s := &jwksServer{body: body, status: 200, headers: map[string]string{}}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		s.hits++
		block := s.block
		headers := maps.Clone(s.headers)
		status, body := s.status, s.body
		s.mu.Unlock()
		if block != nil {
			<-block
		}
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(status)
		_, _ = w.Write(body)
	}))
	return s
}

func (s *jwksServer) set(status int, body []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status, s.body = status, body
}

func (s *jwksServer) setHeader(k, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.headers[k] = v
}

func (s *jwksServer) requestCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.hits
}

func jwksBody(t *testing.T, keys ...jwk) []byte {
	t.Helper()
	out, err := json.Marshal(jwksResponse{Keys: keys})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func TestJWKSRejectsOversizedResponse(t *testing.T) {
	key := mustRSAKey(t)
	jwks := newJWKSServer(t, jwksBody(t, publicJWK("kid-1", &key.PublicKey)))
	defer jwks.Close()

	authn := NewAuthenticator(MCPConfig{JWKSURL: jwks.URL, MaxJWKSBytes: 32})
	if _, err := authn.loadKeys(context.Background(), false); err == nil ||
		!strings.Contains(err.Error(), "exceeds") {
		t.Fatalf("loadKeys oversized = %v", err)
	}
}

func TestJWKSConcurrentFetchDeduped(t *testing.T) {
	key := mustRSAKey(t)
	jwks := newJWKSServer(t, jwksBody(t, publicJWK("kid-1", &key.PublicKey)))
	defer jwks.Close()
	jwks.mu.Lock()
	jwks.block = make(chan struct{})
	jwks.mu.Unlock()

	authn := NewAuthenticator(MCPConfig{JWKSURL: jwks.URL})
	const callers = 8
	errs := make(chan error, callers)
	for i := 0; i < callers; i++ {
		go func() {
			_, err := authn.loadKeys(context.Background(), true)
			errs <- err
		}()
	}
	// Let all callers arrive, then release the single fetch.
	time.Sleep(100 * time.Millisecond)
	close(jwks.block)
	for i := 0; i < callers; i++ {
		if err := <-errs; err != nil {
			t.Fatalf("loadKeys = %v", err)
		}
	}
	if got := jwks.requestCount(); got != 1 {
		t.Fatalf("jwks hits = %d, want 1 (deduped)", got)
	}
}

func TestJWKSRotationAndUnknownKID(t *testing.T) {
	key1, key2 := mustRSAKey(t), mustRSAKey(t)
	jwks := newJWKSServer(t, jwksBody(t, publicJWK("kid-1", &key1.PublicKey)))
	defer jwks.Close()

	authn := NewAuthenticator(MCPConfig{
		JWTIssuer:   "https://auth.example.com",
		JWTAudience: "https://mcp.example.com/mcp",
		JWKSURL:     jwks.URL,
	})
	authn.now = func() time.Time { return time.Unix(1000, 0) }

	authenticate := func(key *rsa.PrivateKey, kid string) error {
		token := signToken(t, key, kid, tokenClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "https://auth.example.com",
				Audience:  jwt.ClaimStrings{"https://mcp.example.com/mcp"},
				ExpiresAt: jwt.NewNumericDate(time.Unix(1600, 0)),
			},
		})
		req := httptest.NewRequest(http.MethodPost, "http://localhost/mcp", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		_, err := authn.Authenticate(context.Background(), req)
		return err
	}

	if err := authenticate(key1, "kid-1"); err != nil {
		t.Fatalf("kid-1 = %v", err)
	}
	// Rotate: only kid-2 is served now. Unknown kids must trigger one refresh
	// and then still fail closed when absent from the fresh set.
	jwks.set(200, jwksBody(t, publicJWK("kid-2", &key2.PublicKey)))
	if err := authenticate(key2, "kid-2"); err != nil {
		t.Fatalf("kid-2 after rotation = %v", err)
	}
	if err := authenticate(key1, "kid-1"); err == nil {
		t.Fatalf("retired kid-1 must fail closed")
	}
	if err := authenticate(key1, "kid-unknown"); err == nil {
		t.Fatalf("unknown kid must fail closed")
	}
}

func TestJWKSRefreshFailureKeepsLastKnownGood(t *testing.T) {
	key1 := mustRSAKey(t)
	jwks := newJWKSServer(t, jwksBody(t, publicJWK("kid-1", &key1.PublicKey)))
	defer jwks.Close()

	authn := NewAuthenticator(MCPConfig{JWKSURL: jwks.URL})
	if _, err := authn.loadKeys(context.Background(), false); err != nil {
		t.Fatal(err)
	}
	jwks.set(500, nil)

	// A refresh triggered by an unknown kid fails; stale keys are returned
	// instead, so the caller still fails closed on the unknown kid.
	if _, err := authn.keyForToken(context.Background(), fakeTokenWithKid("kid-2"), true); err == nil {
		t.Fatalf("unknown kid must fail closed when refresh fails")
	}
	keys, err := authn.loadKeys(context.Background(), true)
	if err != nil {
		t.Fatalf("stale keys should be served on refresh failure: %v", err)
	}
	if _, ok := keys["kid-1"]; !ok {
		t.Fatalf("last-known-good keys lost: %#v", keys)
	}

	// With no prior good set, a failed fetch must error out.
	authn2 := NewAuthenticator(MCPConfig{JWKSURL: jwks.URL})
	if _, err := authn2.loadKeys(context.Background(), false); err == nil {
		t.Fatalf("empty cache + failed fetch must error")
	}
}

func fakeTokenWithKid(kid string) *jwt.Token {
	return &jwt.Token{Header: map[string]interface{}{"kid": kid}}
}

func TestJWKSCancellationHonored(t *testing.T) {
	key := mustRSAKey(t)
	jwks := newJWKSServer(t, jwksBody(t, publicJWK("kid-1", &key.PublicKey)))
	defer jwks.Close()
	block := make(chan struct{})
	defer close(block) // release the hanging handler before Close
	jwks.mu.Lock()
	jwks.block = block
	jwks.mu.Unlock()

	authn := NewAuthenticator(MCPConfig{JWKSURL: jwks.URL})
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() {
		_, err := authn.loadKeys(ctx, false)
		done <- err
	}()
	time.Sleep(50 * time.Millisecond)
	cancel()
	select {
	case err := <-done:
		if err == nil {
			t.Fatalf("expected cancellation error")
		}
	case <-time.After(2 * time.Second):
		t.Fatalf("loadKeys did not honor cancellation")
	}
}

func TestJWKSRejectsMalformedAndDuplicateSets(t *testing.T) {
	key := mustRSAKey(t)

	cases := []struct {
		name string
		body []byte
	}{
		{"malformed json", []byte("{not json")},
		{"empty set", jwksBody(t)},
		{"only non-RSA keys", jwksBody(t, jwk{KID: "ec-1", KTY: "EC"})},
		{"duplicate kid", jwksBody(t, publicJWK("kid-1", &key.PublicKey), publicJWK("kid-1", &key.PublicKey))},
		{"unusable key", jwksBody(t, jwk{KID: "bad", KTY: "RSA", N: "!!!", E: "!!!"})},
	}
	for _, tc := range cases {
		jwks := newJWKSServer(t, tc.body)
		authn := NewAuthenticator(MCPConfig{JWKSURL: jwks.URL})
		if _, err := authn.loadKeys(context.Background(), false); err == nil {
			t.Errorf("%s: expected rejection", tc.name)
		}
		jwks.Close()
	}
}

func TestJWKSCacheControlTTLBounds(t *testing.T) {
	key := mustRSAKey(t)
	jwks := newJWKSServer(t, jwksBody(t, publicJWK("kid-1", &key.PublicKey)))
	defer jwks.Close()

	authn := NewAuthenticator(MCPConfig{JWKSURL: jwks.URL})
	if _, err := authn.loadKeys(context.Background(), false); err != nil {
		t.Fatal(err)
	}
	if authn.ttl != defaultJWKSCacheTTL {
		t.Fatalf("default ttl = %v", authn.ttl)
	}

	jwks.setHeader("Cache-Control", "max-age=5")
	if _, err := authn.loadKeys(context.Background(), true); err != nil {
		t.Fatal(err)
	}
	if authn.ttl != minJWKSCacheTTL {
		t.Fatalf("clamped-down ttl = %v, want %v", authn.ttl, minJWKSCacheTTL)
	}

	jwks.setHeader("Cache-Control", "max-age=86400")
	if _, err := authn.loadKeys(context.Background(), true); err != nil {
		t.Fatal(err)
	}
	if authn.ttl != maxJWKSCacheTTL {
		t.Fatalf("clamped-up ttl = %v, want %v", authn.ttl, maxJWKSCacheTTL)
	}
}

func mustRSAKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	return key
}

func publicJWK(kid string, pub *rsa.PublicKey) jwk {
	return jwk{
		KID: kid,
		KTY: "RSA",
		Use: "sig",
		Alg: jwt.SigningMethodRS256.Alg(),
		N:   base64.RawURLEncoding.EncodeToString(pub.N.Bytes()),
		E:   base64.RawURLEncoding.EncodeToString(big.NewInt(int64(pub.E)).Bytes()),
	}
}

func signToken(t *testing.T, key *rsa.PrivateKey, kid string, claims tokenClaims) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = kid
	out, err := token.SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

package mcpserver

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
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

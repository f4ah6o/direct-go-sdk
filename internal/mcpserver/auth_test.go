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

func TestAuthenticatorExtractsPrincipalsFromClaims(t *testing.T) {
	key := mustRSAKey(t)
	jwks := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(jwksResponse{Keys: []jwk{publicJWK("kid-1", &key.PublicKey)}})
	}))
	defer jwks.Close()

	newAuth := func(subjectClaim string) *Authenticator {
		a := NewAuthenticator(MCPConfig{
			JWTIssuer:    "https://auth.example.com",
			JWTAudience:  "https://mcp.example.com/mcp",
			JWKSURL:      jwks.URL,
			SubjectClaim: subjectClaim,
		})
		a.now = func() time.Time { return time.Unix(1000, 0) }
		return a
	}

	authenticate := func(a *Authenticator, claims jwt.MapClaims) (*AuthInfo, error) {
		token := signMapToken(t, key, "kid-1", claims)
		req := httptest.NewRequest(http.MethodPost, "http://localhost/mcp", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		ctx, err := a.Authenticate(context.Background(), req)
		if err != nil {
			return nil, err
		}
		info, _ := AuthInfoFromContext(ctx)
		return info, nil
	}

	base := func() jwt.MapClaims {
		return jwt.MapClaims{
			"iss": "https://auth.example.com",
			"aud": "https://mcp.example.com/mcp",
			"exp": time.Unix(1600, 0).Unix(),
			"iat": time.Unix(900, 0).Unix(),
			"sub": "user-1",
		}
	}

	// Default claim (sub): principal is the subject.
	info, err := authenticate(newAuth(""), base())
	if err != nil {
		t.Fatal(err)
	}
	if len(info.Principals) != 1 || info.Principals[0] != "user-1" {
		t.Fatalf("sub principals = %#v", info.Principals)
	}

	// A list claim yields one principal per element.
	claims := base()
	claims["groups"] = []interface{}{"team-a", "team-b"}
	info, err = authenticate(newAuth("groups"), claims)
	if err != nil {
		t.Fatal(err)
	}
	if len(info.Principals) != 2 || info.Principals[0] != "team-a" || info.Principals[1] != "team-b" {
		t.Fatalf("groups principals = %#v", info.Principals)
	}

	// A malformed claim (non-string/list) falls back to sub, never to nil.
	claims = base()
	claims["groups"] = map[string]interface{}{"nested": true}
	info, err = authenticate(newAuth("groups"), claims)
	if err != nil {
		t.Fatal(err)
	}
	if len(info.Principals) != 1 || info.Principals[0] != "user-1" {
		t.Fatalf("malformed claim principals = %#v", info.Principals)
	}

	// Missing configured claim also falls back to sub.
	info, err = authenticate(newAuth("groups"), base())
	if err != nil {
		t.Fatal(err)
	}
	if len(info.Principals) != 1 || info.Principals[0] != "user-1" {
		t.Fatalf("missing claim principals = %#v", info.Principals)
	}
}

func TestAuthenticatorParsesScopeClaims(t *testing.T) {
	key := mustRSAKey(t)
	jwks := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(jwksResponse{Keys: []jwk{publicJWK("kid-1", &key.PublicKey)}})
	}))
	defer jwks.Close()

	authn := NewAuthenticator(MCPConfig{
		JWTIssuer:   "https://auth.example.com",
		JWTAudience: "https://mcp.example.com/mcp",
		JWKSURL:     jwks.URL,
		ReadScope:   DefaultReadScope,
		WriteScope:  DefaultWriteScope,
	})
	authn.now = func() time.Time { return time.Unix(1000, 0) }

	// scp as a list of strings (the form some issuers emit).
	claims := jwt.MapClaims{
		"iss": "https://auth.example.com",
		"aud": "https://mcp.example.com/mcp",
		"exp": time.Unix(1600, 0).Unix(),
		"sub": "user-1",
		"scp": []interface{}{"direct:read", "direct:write"},
	}
	token := signMapToken(t, key, "kid-1", claims)
	req := httptest.NewRequest(http.MethodPost, "http://localhost/mcp", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	ctx, err := authn.Authenticate(context.Background(), req)
	if err != nil {
		t.Fatalf("Authenticate() = %v", err)
	}
	if err := RequireScope(ctx, DefaultWriteScope); err != nil {
		t.Fatalf("RequireScope(write) with list scp = %v", err)
	}
}

func signMapToken(t *testing.T, key *rsa.PrivateKey, kid string, claims jwt.MapClaims) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = kid
	out, err := token.SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	return out
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

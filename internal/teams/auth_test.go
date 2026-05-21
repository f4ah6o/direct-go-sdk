package teams

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func TestBotFrameworkValidatorAcceptsValidTeamsToken(t *testing.T) {
	key, metadataURL := testJWTServer(t, "key-1", []string{"msteams"})
	now := time.Unix(1700000000, 0)
	cfg := config.BotConfig{
		AppID:             "bot-app-id",
		OpenIDMetadataURL: metadataURL,
		AllowedServiceURLs: []string{
			"https://smba.trafficmanager.net/amer/",
		},
	}
	validator := NewBotFrameworkValidator(cfg)
	validator.now = func() time.Time { return now }

	token := signBotToken(t, key, "key-1", cfg.AppID, botFrameworkIssuer, "https://smba.trafficmanager.net/amer/", now)
	req := httptest.NewRequest(http.MethodPost, "/api/messages", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	activity := Activity{Type: "message", ServiceURL: "https://smba.trafficmanager.net/amer/", ChannelID: "msteams"}

	if err := validator.Validate(context.Background(), req, activity); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestBotFrameworkValidatorRejectsServiceURLMismatch(t *testing.T) {
	key, metadataURL := testJWTServer(t, "key-1", []string{"msteams"})
	now := time.Unix(1700000000, 0)
	cfg := config.BotConfig{AppID: "bot-app-id", OpenIDMetadataURL: metadataURL}
	validator := NewBotFrameworkValidator(cfg)
	validator.now = func() time.Time { return now }

	token := signBotToken(t, key, "key-1", cfg.AppID, botFrameworkIssuer, "https://smba.trafficmanager.net/amer/", now)
	req := httptest.NewRequest(http.MethodPost, "/api/messages", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	activity := Activity{Type: "message", ServiceURL: "https://evil.example.com/", ChannelID: "msteams"}

	if err := validator.Validate(context.Background(), req, activity); err == nil {
		t.Fatalf("expected serviceUrl mismatch")
	}
}

func TestBotFrameworkValidatorRejectsMissingEndorsement(t *testing.T) {
	key, metadataURL := testJWTServer(t, "key-1", []string{"webchat"})
	now := time.Unix(1700000000, 0)
	cfg := config.BotConfig{AppID: "bot-app-id", OpenIDMetadataURL: metadataURL}
	validator := NewBotFrameworkValidator(cfg)
	validator.now = func() time.Time { return now }

	token := signBotToken(t, key, "key-1", cfg.AppID, botFrameworkIssuer, "https://smba.trafficmanager.net/amer/", now)
	req := httptest.NewRequest(http.MethodPost, "/api/messages", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	activity := Activity{Type: "message", ServiceURL: "https://smba.trafficmanager.net/amer/", ChannelID: "msteams"}

	if err := validator.Validate(context.Background(), req, activity); err == nil {
		t.Fatalf("expected endorsement rejection")
	}
}

func TestFileProxySignature(t *testing.T) {
	t.Setenv("APP_PASSWORD", "secret")
	cfg := config.BotConfig{AppPasswordEnv: "APP_PASSWORD"}
	now := time.Unix(1700000000, 0)
	rawURL := "https://api.direct4b.com/albero-app-server/files/file/token?message_id=1"
	signed, err := signedDirectFileURL("https://bridge.example.com", cfg, "1h", "account-a", rawURL, now)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, signed, nil)
	exp, err := parseInt64(req.URL.Query().Get("exp"))
	if err != nil {
		t.Fatal(err)
	}
	if err := validateDirectFileSignature(cfg, "account-a", rawURL, exp, req.URL.Query().Get("sig"), now.Add(30*time.Minute)); err != nil {
		t.Fatalf("validateDirectFileSignature() error = %v", err)
	}
	if err := validateDirectFileSignature(cfg, "account-a", rawURL, exp, req.URL.Query().Get("sig"), now.Add(2*time.Hour)); err == nil {
		t.Fatalf("expected expired URL")
	}
}

func testJWTServer(t *testing.T, kid string, endorsements []string) (*rsa.PrivateKey, string) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	var baseURL string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/metadata":
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"jwks_uri":                              baseURL + "/keys",
				"id_token_signing_alg_values_supported": []string{"RS256"},
			})
		case "/keys":
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"keys": []map[string]interface{}{{
					"kty":          "RSA",
					"use":          "sig",
					"kid":          kid,
					"alg":          "RS256",
					"n":            base64.RawURLEncoding.EncodeToString(key.PublicKey.N.Bytes()),
					"e":            base64.RawURLEncoding.EncodeToString([]byte{0x01, 0x00, 0x01}),
					"endorsements": endorsements,
				}},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)
	baseURL = server.URL
	return key, server.URL + "/metadata"
}

func signBotToken(t *testing.T, key *rsa.PrivateKey, kid, audience, issuer, serviceURL string, now time.Time) string {
	t.Helper()
	claims := botClaims{
		ServiceURL: serviceURL,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Audience:  jwt.ClaimStrings{audience},
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
			NotBefore: jwt.NewNumericDate(now.Add(-time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = kid
	out, err := token.SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

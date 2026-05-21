package teams

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
)

func signedDirectFileURL(publicBaseURL string, cfg config.BotConfig, ttlText, accountID, fileURL string, now time.Time) (string, error) {
	if strings.TrimSpace(publicBaseURL) == "" {
		return fileURL, nil
	}
	ttl, err := time.ParseDuration(ttlText)
	if err != nil {
		return "", err
	}
	exp := now.Add(ttl).Unix()
	values := url.Values{}
	values.Set("account", accountID)
	values.Set("url", fileURL)
	values.Set("exp", strconv.FormatInt(exp, 10))
	values.Set("sig", fileProxySignature(cfg, accountID, fileURL, exp))
	return strings.TrimRight(publicBaseURL, "/") + "/files/direct?" + values.Encode(), nil
}

func validateDirectFileSignature(cfg config.BotConfig, accountID, fileURL string, exp int64, sig string, now time.Time) error {
	if accountID == "" || fileURL == "" || exp == 0 || sig == "" {
		return fmt.Errorf("missing signed file proxy parameters")
	}
	if now.Unix() > exp {
		return fmt.Errorf("signed file proxy URL expired")
	}
	want := fileProxySignature(cfg, accountID, fileURL, exp)
	if !hmac.Equal([]byte(sig), []byte(want)) {
		return fmt.Errorf("invalid signed file proxy signature")
	}
	return nil
}

func fileProxySignature(cfg config.BotConfig, accountID, fileURL string, exp int64) string {
	mac := hmac.New(sha256.New, []byte(botSecret(cfg)))
	_, _ = fmt.Fprintf(mac, "%s\n%s\n%d", accountID, fileURL, exp)
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func botSecret(cfg config.BotConfig) string {
	if cfg.AppPassword != "" {
		return cfg.AppPassword
	}
	if cfg.AppPasswordEnv != "" {
		return os.Getenv(cfg.AppPasswordEnv)
	}
	return cfg.AppPasswordRef
}

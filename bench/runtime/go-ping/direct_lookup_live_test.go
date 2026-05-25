package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"gopkg.in/yaml.v3"
)

type lookupTestConfig struct {
	Op       lookupTestOPConfig  `yaml:"op"`
	Accounts []lookupTestAccount `yaml:"accounts"`
}

type lookupTestOPConfig struct {
	Binary string `yaml:"binary"`
}

type lookupTestAccount struct {
	ID       string `yaml:"id"`
	TokenEnv string `yaml:"token_env"`
	TokenRef string `yaml:"token_ref"`
	Endpoint string `yaml:"endpoint"`
	ProxyURL string `yaml:"proxy_url"`
}

func TestLiveDirectLookupResolvesUserAndRoomNames(t *testing.T) {
	if os.Getenv("DIRECT_LOOKUP_LIVE") != "1" {
		t.Skip("set DIRECT_LOOKUP_LIVE=1 to run the live Direct lookup test")
	}
	talkID := firstEnv("DIRECT_LOOKUP_TALK_ID", "DIRECT_BENCH_TALK_ID")
	if talkID == "" {
		t.Fatal("DIRECT_LOOKUP_TALK_ID or DIRECT_BENCH_TALK_ID is required")
	}
	configPath := envOr("DIRECT_LOOKUP_CONFIG", "../../../../config.yaml")
	accountID := envOr("DIRECT_LOOKUP_ACCOUNT", "bot-trial2")

	cfg, account, err := loadLookupTestAccount(configPath, accountID)
	if err != nil {
		t.Fatal(err)
	}
	token, err := resolveLookupTestToken(cfg, account)
	if err != nil {
		t.Fatal(err)
	}
	endpoint := account.Endpoint
	if endpoint == "" {
		endpoint = direct.DefaultEndpoint
	}

	client := direct.NewClient(direct.Options{
		Endpoint:    endpoint,
		AccessToken: token,
		ProxyURL:    account.ProxyURL,
		Name:        "runtime-direct-lookup-test",
	})
	ready := make(chan struct{})
	client.On(direct.EventDataRecovered, func(data interface{}) {
		select {
		case <-ready:
		default:
			close(ready)
		}
	})
	if err := client.Connect(); err != nil {
		t.Fatalf("connect failed: %v", err)
	}
	defer client.Close()
	select {
	case <-ready:
	case <-time.After(30 * time.Second):
		t.Fatal("timed out waiting for Direct notification readiness")
	}

	talk, ok := findLookupTalk(t, client, talkID)
	if !ok {
		t.Fatalf("talk %s not found in get_talks", talkID)
	}
	roomName := firstLookupValue(talk, "name", "display_name", "displayName")
	domainID := firstLookupValue(talk, "domain_id", "domainId")
	if domainID == "" {
		t.Fatalf("talk %s resolved but domain id is empty: %#v", talkID, talk)
	}

	userID := os.Getenv("DIRECT_LOOKUP_USER_ID")
	if userID == "" {
		userID = firstLookupUserID(t, client, talk)
	}
	if userID == "" {
		t.Fatalf("could not infer user id from talk %s; set DIRECT_LOOKUP_USER_ID", talkID)
	}

	usersResult, err := client.Call(direct.MethodGetUsers, []interface{}{lookupRPCID(domainID), []interface{}{lookupRPCID(userID)}})
	if err != nil {
		t.Fatalf("get_users domain=%s user=%s failed: %v", domainID, userID, err)
	}
	userName := lookupUserDisplayName(usersResult, userID)
	if userName == "" {
		t.Fatalf("user %s resolved but display/name is empty: %#v", userID, usersResult)
	}

	t.Logf("lookup ok: talk_id=%s room_name=%q domain_id=%s user_id=%s user_name=%q", talkID, roomName, domainID, userID, userName)
}

func findLookupTalk(t *testing.T, client *direct.Client, talkID string) (map[string]interface{}, bool) {
	t.Helper()
	result, err := client.Call(direct.MethodGetTalks, []interface{}{})
	if err != nil {
		t.Fatalf("get_talks failed: %v", err)
	}
	arr, ok := result.([]interface{})
	if !ok {
		t.Fatalf("get_talks returned %T, want []interface{}", result)
	}
	for _, item := range arr {
		talk, ok := lookupStringKeyMap(item)
		if !ok {
			continue
		}
		id := firstLookupValue(talk, "id", "talk_id", "talkId")
		if id == talkID {
			return talk, true
		}
	}
	return nil, false
}

func firstLookupUserID(t *testing.T, client *direct.Client, talk map[string]interface{}) string {
	t.Helper()
	selfID := lookupSelfUserID(t, client)
	for _, key := range []string{"user_ids", "userIds"} {
		value, ok := talk[key]
		if !ok {
			continue
		}
		arr, ok := value.([]interface{})
		if !ok {
			continue
		}
		for _, id := range arr {
			normalized := lookupIDString(id)
			if normalized != "" && normalized != selfID {
				return normalized
			}
		}
	}
	return selfID
}

func lookupSelfUserID(t *testing.T, client *direct.Client) string {
	t.Helper()
	result, err := client.Call(direct.MethodGetMe, []interface{}{})
	if err != nil {
		t.Fatalf("get_me failed: %v", err)
	}
	me, ok := lookupStringKeyMap(result)
	if !ok {
		t.Fatalf("get_me returned %T, want map", result)
	}
	return firstLookupValue(me, "id", "user_id", "userId")
}

func lookupUserDisplayName(result interface{}, userID string) string {
	arr, ok := result.([]interface{})
	if !ok {
		return ""
	}
	for _, item := range arr {
		user, ok := lookupStringKeyMap(item)
		if !ok {
			continue
		}
		if id := firstLookupValue(user, "id", "user_id", "userId"); id != "" && id != userID {
			continue
		}
		if name := firstLookupValue(user, "display_name", "displayName", "name"); name != "" {
			return name
		}
	}
	return ""
}

func lookupStringKeyMap(value interface{}) (map[string]interface{}, bool) {
	switch m := value.(type) {
	case map[string]interface{}:
		return m, true
	case map[interface{}]interface{}:
		out := make(map[string]interface{}, len(m))
		for key, value := range m {
			out[fmt.Sprint(key)] = value
		}
		return out, true
	default:
		return nil, false
	}
}

func firstLookupValue(m map[string]interface{}, keys ...string) string {
	for _, key := range keys {
		value, ok := m[key]
		if !ok || value == nil {
			continue
		}
		if s := strings.TrimSpace(lookupIDString(value)); s != "" {
			return s
		}
	}
	return ""
}

func lookupIDString(value interface{}) string {
	switch v := value.(type) {
	case float32:
		return strconv.FormatInt(int64(v), 10)
	case float64:
		return strconv.FormatInt(int64(v), 10)
	default:
		return fmt.Sprint(v)
	}
}

func lookupRPCID(value string) interface{} {
	if id, err := strconv.ParseUint(value, 10, 64); err == nil {
		return id
	}
	return value
}

func loadLookupTestAccount(path, id string) (lookupTestConfig, lookupTestAccount, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return lookupTestConfig{}, lookupTestAccount{}, err
	}
	var cfg lookupTestConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return lookupTestConfig{}, lookupTestAccount{}, err
	}
	for _, account := range cfg.Accounts {
		if account.ID == id {
			return cfg, account, nil
		}
	}
	return lookupTestConfig{}, lookupTestAccount{}, fmt.Errorf("account %q not found in %s", id, path)
}

func resolveLookupTestToken(cfg lookupTestConfig, account lookupTestAccount) (string, error) {
	if account.TokenEnv != "" {
		if token := os.Getenv(account.TokenEnv); token != "" {
			return token, nil
		}
	}
	if account.TokenRef == "" {
		return "", fmt.Errorf("%s is not set and account %s has no token_ref", account.TokenEnv, account.ID)
	}
	opBinary := cfg.Op.Binary
	if opBinary == "" {
		opBinary = "op"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, opBinary, "read", account.TokenRef)
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("op read failed for account %s: %s", account.ID, strings.TrimSpace(string(exitErr.Stderr)))
		}
		return "", fmt.Errorf("op read failed for account %s: %w", account.ID, err)
	}
	token := strings.TrimSpace(string(out))
	if token == "" {
		return "", fmt.Errorf("op read returned an empty token for account %s", account.ID)
	}
	return token, nil
}

func firstEnv(keys ...string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return ""
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

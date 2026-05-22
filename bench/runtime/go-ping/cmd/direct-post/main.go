package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"gopkg.in/yaml.v3"
)

type config struct {
	Op       opConfig  `yaml:"op"`
	Accounts []account `yaml:"accounts"`
}

type opConfig struct {
	Binary string `yaml:"binary"`
}

type account struct {
	ID       string `yaml:"id"`
	TokenEnv string `yaml:"token_env"`
	TokenRef string `yaml:"token_ref"`
	Endpoint string `yaml:"endpoint"`
	ProxyURL string `yaml:"proxy_url"`
}

func main() {
	configPath := flag.String("config", "config.test.yaml", "YAML config with direct accounts")
	accountID := flag.String("account", "bot-trial2", "account id from config")
	talkID := flag.String("talk-id", os.Getenv("DIRECT_BENCH_TALK_ID"), "direct talk id")
	text := flag.String("text", "ping", "message text")
	count := flag.Int("count", 1, "number of messages to send")
	interval := flag.Duration("interval", time.Second, "delay between messages")
	suffixCount := flag.Bool("suffix-count", false, "append message number when count is greater than 1")
	timeout := flag.Duration("timeout", 15*time.Second, "per-run timeout")
	flag.Parse()

	if *talkID == "" {
		log.Fatal("--talk-id or DIRECT_BENCH_TALK_ID is required")
	}
	if *count < 1 {
		log.Fatal("--count must be >= 1")
	}

	cfg, acct, err := loadAccount(*configPath, *accountID)
	if err != nil {
		log.Fatal(err)
	}
	token, err := resolveToken(cfg, acct)
	if err != nil {
		log.Fatal(err)
	}

	endpoint := acct.Endpoint
	if endpoint == "" {
		endpoint = direct.DefaultEndpoint
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	for i := 0; i < *count; i++ {
		msg := *text
		if *suffixCount && *count > 1 {
			msg = fmt.Sprintf("%s %d", *text, i+1)
		}
		start := time.Now()
		messageID, err := sendOnce(ctx, endpoint, token, acct.ProxyURL, *talkID, msg)
		if err != nil {
			log.Fatalf("send failed: %v", err)
		}
		fmt.Printf("sent account=%s talk=%s message_id=%s elapsed_ms=%d text=%q\n", acct.ID, normalizeForLog(*talkID), messageID, time.Since(start).Milliseconds(), msg)
		if i+1 < *count {
			time.Sleep(*interval)
		}
	}
}

func sendOnce(ctx context.Context, endpoint, token, proxyURL, talkID, text string) (string, error) {
	client := direct.NewClient(direct.Options{
		Endpoint:    endpoint,
		AccessToken: token,
		ProxyURL:    proxyURL,
		Name:        "runtime-direct-post",
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
		return "", fmt.Errorf("connect failed: %w", err)
	}
	defer client.Close()
	select {
	case <-ready:
	case <-ctx.Done():
		return "", fmt.Errorf("ready timeout: %w", ctx.Err())
	}
	return client.CreateTextMessageWithContext(ctx, talkID, text)
}

func loadAccount(path, id string) (config, account, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, account{}, err
	}
	var cfg config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return config{}, account{}, err
	}
	for _, acct := range cfg.Accounts {
		if acct.ID == id {
			return cfg, acct, nil
		}
	}
	ids := make([]string, 0, len(cfg.Accounts))
	for _, acct := range cfg.Accounts {
		ids = append(ids, acct.ID)
	}
	return config{}, account{}, fmt.Errorf("account %q not found in %s; available: %s", id, path, strings.Join(ids, ", "))
}

func resolveToken(cfg config, acct account) (string, error) {
	if acct.TokenEnv != "" {
		if token := os.Getenv(acct.TokenEnv); token != "" {
			return token, nil
		}
	}
	if acct.TokenRef == "" {
		return "", fmt.Errorf("%s is not set and account %s has no token_ref", acct.TokenEnv, acct.ID)
	}
	opBinary := cfg.Op.Binary
	if opBinary == "" {
		opBinary = "op"
	}
	cmd := exec.Command(opBinary, "read", acct.TokenRef)
	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("op read failed for account %s: %s", acct.ID, strings.TrimSpace(string(exitErr.Stderr)))
		}
		return "", fmt.Errorf("op read failed for account %s: %w", acct.ID, err)
	}
	token := strings.TrimSpace(string(out))
	if token == "" {
		return "", fmt.Errorf("op read returned an empty token for account %s", acct.ID)
	}
	return token, nil
}

func normalizeForLog(id string) string {
	if _, err := strconv.ParseUint(id, 10, 64); err == nil {
		return id
	}
	return id
}

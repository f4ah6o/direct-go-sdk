package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to direct as a bot account",
	Long:  `Login to the direct service using your bot account credentials.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLogin()
	},
}

func runLogin() error {
	auth := direct.NewAuth()

	// Check if already logged in
	if auth.HasToken() {
		fmt.Println("Already logged in.")
		fmt.Println("Run 'daabgo logout' first to login with a different account.")
		return nil
	}

	// Get endpoint from environment
	if err := auth.LoadEnv(); err != nil {
		fmt.Printf("Warning: could not load .env: %v\n", err)
	}

	endpoint := os.Getenv("HUBOT_DIRECT_ENDPOINT")
	if endpoint == "" {
		endpoint = direct.DefaultEndpoint
	}

	proxyURL := os.Getenv("HUBOT_DIRECT_PROXY_URL")
	if proxyURL == "" {
		proxyURL = os.Getenv("HTTPS_PROXY")
	}
	if proxyURL == "" {
		proxyURL = os.Getenv("HTTP_PROXY")
	}

	// Prompt for credentials
	email, password, err := promptCredentials()
	if err != nil {
		return fmt.Errorf("failed to read credentials: %w", err)
	}

	fmt.Println()
	fmt.Println("Connecting to direct for authentication...")
	fmt.Printf("Endpoint: %s\n", endpoint)

	// Create client for login (without token)
	client := direct.NewClient(direct.Options{
		Endpoint: endpoint,
		ProxyURL: proxyURL,
	})

	// Connect first
	if err := client.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer client.Close()

	fmt.Println("Getting access token...")

	// Call create_access_token API
	// Parameters: [email, password, device_name, device_info, ""]
	result, err := client.Call("create_access_token", []interface{}{
		email,
		password,
		"daabgo",
		"Go Bot Client",
		"",
	})

	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	// Extract token from result
	token := extractToken(result)
	if token == "" {
		return fmt.Errorf("failed to extract token from response: %v", result)
	}

	// Save token
	if err := auth.SetToken(token); err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}

	fmt.Println("Logged in successfully!")
	return nil
}

func promptCredentials() (email, password string, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Email: ")
	email, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	email = strings.TrimSpace(email)

	fmt.Print("Password: ")
	// Read password without echo
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		// Fallback to normal read if terminal read fails
		password, err = reader.ReadString('\n')
		if err != nil {
			return
		}
	} else {
		password = string(passwordBytes)
	}
	password = strings.TrimSpace(password)

	return
}

func extractToken(result interface{}) string {
	// The result could be a string or a map with access_token field
	if token, ok := result.(string); ok {
		return token
	}

	if m, ok := result.(map[string]interface{}); ok {
		if token, ok := m["access_token"].(string); ok {
			return token
		}
		// Try array format - result might be [access_token, ...]
	}

	if arr, ok := result.([]interface{}); ok && len(arr) > 0 {
		if token, ok := arr[0].(string); ok {
			return token
		}
	}

	return ""
}

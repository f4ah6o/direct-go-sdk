// auth.go provides authentication and configuration management for the direct client.
//
// The Auth type manages access tokens and environment variables, supporting both
// environment variables and .env file storage. It provides convenience methods for
// token lifecycle management.
package direct

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Authentication constants.
const (
	// TokenEnvKey is the environment variable name for the direct API access token.
	// Set this environment variable to authenticate without needing a .env file.
	TokenEnvKey = "HUBOT_DIRECT_TOKEN"

	// EnvFile is the default filename for environment variable configuration.
	// By default, Auth reads from and writes to ".env" in the current directory.
	EnvFile = ".env"
)

// Auth manages authentication tokens and environment configuration.
// It supports reading and writing tokens to .env files and environment variables.
//
// Token lookup order:
// 1. HUBOT_DIRECT_TOKEN environment variable (highest priority)
// 2. Value in the .env file (lower priority)
//
// Example:
//
//	auth := direct.NewAuth()
//	token := auth.GetToken()
//	if token == "" {
//		log.Fatal("No authentication token found")
//	}
type Auth struct {
	envFile string
}

// NewAuth creates a new Auth manager using the default .env file in the current directory.
// The Auth manager handles token storage and retrieval from environment variables and .env files.
//
// Example:
//
//	auth := direct.NewAuth()
//	token := auth.GetToken()
func NewAuth() *Auth {
	return &Auth{envFile: EnvFile}
}

// NewAuthWithFile creates a new Auth manager using a custom environment file path.
// This allows using a different file than the default ".env" for token storage.
//
// Parameters:
// - envFile: Path to the environment file (e.g., "config/.env.local", "/etc/app/.env")
//
// Example:
//
//	auth := direct.NewAuthWithFile("/home/user/mybot/.env")
//	token := auth.GetToken()
func NewAuthWithFile(envFile string) *Auth {
	return &Auth{envFile: envFile}
}

// HasToken checks if an access token is available in the environment or .env file.
// It first checks the HUBOT_DIRECT_TOKEN environment variable, then checks the .env file.
// Returns true if a token is found in either location, false otherwise.
func (a *Auth) HasToken() bool {
	// Check environment variable first
	if os.Getenv(TokenEnvKey) != "" {
		return true
	}

	// Check .env file
	token, _ := a.readTokenFromFile()
	return token != ""
}

// GetToken retrieves the access token from the environment or .env file.
// It checks sources in priority order:
// 1. HUBOT_DIRECT_TOKEN environment variable
// 2. Value in the .env file
//
// Returns an empty string if no token is found in either location.
//
// Example:
//
//	token := auth.GetToken()
//	if token == "" {
//		log.Fatal("Authentication token not found")
//	}
//	client := direct.NewClient(direct.Options{AccessToken: token})
func (a *Auth) GetToken() string {
	// Check environment variable first
	if token := os.Getenv(TokenEnvKey); token != "" {
		return token
	}

	// Check .env file
	token, _ := a.readTokenFromFile()
	return token
}

// SetToken stores or updates the access token in the .env file.
// If the token already exists in the file, its value is updated.
// If the token parameter is empty, the token entry is removed from the file.
// The file is created if it doesn't exist, with permissions 0600 (readable/writable by owner only).
//
// Parameters:
// - token: The access token to store, or empty string to remove the token
//
// Returns an error if the file cannot be read or written.
//
// Example:
//
//	err := auth.SetToken("new-access-token")
//	if err != nil {
//		log.Printf("Failed to save token: %v", err)
//	}
func (a *Auth) SetToken(token string) error {
	content, err := a.readEnvFile()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Update or add the token
	lines := strings.Split(content, "\n")
	found := false
	newLines := make([]string, 0, len(lines)+1)

	for _, line := range lines {
		if strings.HasPrefix(line, TokenEnvKey+"=") {
			if token != "" {
				newLines = append(newLines, TokenEnvKey+"="+token)
			}
			found = true
		} else if line != "" {
			newLines = append(newLines, line)
		}
	}

	if !found && token != "" {
		newLines = append(newLines, TokenEnvKey+"="+token)
	}

	// Write back
	return os.WriteFile(a.envFile, []byte(strings.Join(newLines, "\n")+"\n"), 0600)
}

// ClearToken removes the access token from the .env file.
// This is a convenience method equivalent to SetToken("").
// Returns an error if the file cannot be written.
//
// Example:
//
//	err := auth.ClearToken()
//	if err != nil {
//		log.Printf("Failed to clear token: %v", err)
//	}
func (a *Auth) ClearToken() error {
	return a.SetToken("")
}

// readEnvFile reads the entire .env file content.
func (a *Auth) readEnvFile() (string, error) {
	data, err := os.ReadFile(a.envFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// readTokenFromFile reads the token from the .env file.
func (a *Auth) readTokenFromFile() (string, error) {
	file, err := os.Open(a.envFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, TokenEnvKey+"=") {
			return strings.TrimPrefix(line, TokenEnvKey+"="), nil
		}
	}

	return "", nil
}

// LoadEnv loads environment variables from the .env file into the current process environment.
// Variables are only set if they are not already defined in the environment
// (existing environment variables are not overwritten).
// Lines starting with # are treated as comments and ignored.
// Empty lines are skipped.
// Returns nil if the .env file doesn't exist (no error).
// Returns an error if the file cannot be read or parsed.
//
// Format:
// The .env file should contain KEY=VALUE pairs, one per line:
//
//	HUBOT_DIRECT_TOKEN=your-token-here
//	HUBOT_DIRECT_ENDPOINT=wss://api.direct4b.com/...
//	# This is a comment
//
// Example:
//
//	auth := direct.NewAuth()
//	if err := auth.LoadEnv(); err != nil {
//		log.Printf("Failed to load .env file: %v", err)
//	}
//	// Now environment variables from .env are available via os.Getenv()
func (a *Auth) LoadEnv() error {
	file, err := os.Open(a.envFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			// Only set if not already set
			if os.Getenv(key) == "" {
				os.Setenv(key, value)
			}
		}
	}

	return scanner.Err()
}

// PromptCredentials prompts the user to enter email and password via standard input.
// This is typically used for interactive authentication flows where the user needs to
// provide credentials to obtain an access token.
//
// Prompts:
// - "Email: " for the user's email address
// - "Password: " for the user's password
//
// Returns:
// - email: The email address entered by the user (trimmed of whitespace)
// - password: The password entered by the user (trimmed of whitespace)
// - err: An error if reading from stdin fails
//
// Note: This function reads passwords in plain text from stdin. For production
// use, consider using a library that reads passwords securely without echo.
//
// Example:
//
//	email, password, err := direct.PromptCredentials()
//	if err != nil {
//		log.Fatal("Failed to read credentials")
//	}
//	// Use email and password to obtain an access token
func PromptCredentials() (email, password string, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Email: ")
	email, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	email = strings.TrimSpace(email)

	fmt.Print("Password: ")
	password, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	password = strings.TrimSpace(password)

	return
}

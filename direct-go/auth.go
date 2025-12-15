package direct

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// TokenEnvKey is the environment variable name for the access token.
	TokenEnvKey = "HUBOT_DIRECT_TOKEN"

	// EnvFile is the default .env file name.
	EnvFile = ".env"
)

// Auth handles token storage and retrieval.
type Auth struct {
	envFile string
}

// NewAuth creates a new Auth manager using the default .env file.
// The Auth manager handles token storage and retrieval from environment variables and .env files.
func NewAuth() *Auth {
	return &Auth{envFile: EnvFile}
}

// NewAuthWithFile creates a new Auth manager using a custom env file path.
// This allows using a different file than the default .env for token storage.
func NewAuthWithFile(envFile string) *Auth {
	return &Auth{envFile: envFile}
}

// HasToken checks if an access token exists in the environment or .env file.
// It first checks the HUBOT_DIRECT_TOKEN environment variable, then the .env file.
func (a *Auth) HasToken() bool {
	// Check environment variable first
	if os.Getenv(TokenEnvKey) != "" {
		return true
	}

	// Check .env file
	token, _ := a.readTokenFromFile()
	return token != ""
}

// GetToken retrieves the access token from environment or .env file.
// It first checks the HUBOT_DIRECT_TOKEN environment variable, then the .env file.
// Returns an empty string if no token is found.
func (a *Auth) GetToken() string {
	// Check environment variable first
	if token := os.Getenv(TokenEnvKey); token != "" {
		return token
	}

	// Check .env file
	token, _ := a.readTokenFromFile()
	return token
}

// SetToken stores the access token in the .env file.
// If the token already exists, it updates the value.
// If the token parameter is empty, it removes the token entry.
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
// This is a convenience method that calls SetToken with an empty string.
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

// LoadEnv loads environment variables from the .env file into the process environment.
// It only sets variables that are not already defined in the environment.
// Lines starting with # are treated as comments and ignored.
// Returns nil if the .env file doesn't exist.
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

// PromptCredentials prompts the user for email and password via stdin.
// This is typically used for interactive login flows.
// Returns the trimmed email and password strings, or an error if reading fails.
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

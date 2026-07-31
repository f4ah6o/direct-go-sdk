package direct

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"strings"

	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
)

const (
	// TokenEnvKey is the environment variable name for the access token.
	TokenEnvKey = "HUBOT_DIRECT_TOKEN"

	// DeviceIDEnvKey is the environment variable name for the direct device id.
	DeviceIDEnvKey = "HUBOT_DIRECT_DEVICE_ID"

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
	token, err := a.readTokenFromFile()
	if err != nil && !os.IsNotExist(err) {
		// Log unexpected errors (permission issues, etc.) but ignore "file not found"
		vlog("[WARNING] Error reading token file: %s", debuglog.SummarizePayload(err))
	}
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
	token, err := a.readTokenFromFile()
	if err != nil && !os.IsNotExist(err) {
		// Log unexpected errors (permission issues, etc.) but ignore "file not found"
		vlog("[WARNING] Error reading token file: %s", debuglog.SummarizePayload(err))
	}
	return token
}

// EnsureDeviceID retrieves or creates a stable device id in the .env file.
func (a *Auth) EnsureDeviceID() (string, error) {
	if deviceID := os.Getenv(DeviceIDEnvKey); deviceID != "" {
		return deviceID, nil
	}
	deviceID, err := a.readValueFromFile(DeviceIDEnvKey)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}
	if deviceID != "" {
		return deviceID, nil
	}
	deviceID, err = randomHex(16)
	if err != nil {
		return "", err
	}
	return deviceID, a.setValue(DeviceIDEnvKey, deviceID)
}

// SetToken stores the access token in the .env file.
// If the token already exists, it updates the value.
// If the token parameter is empty, it removes the token entry.
func (a *Auth) SetToken(token string) error {
	return a.setValue(TokenEnvKey, token)
}

func (a *Auth) setValue(key, value string) error {
	content, err := a.readEnvFile()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Update or add the token
	lines := strings.Split(content, "\n")
	found := false
	newLines := make([]string, 0, len(lines)+1)

	for _, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			if value != "" {
				newLines = append(newLines, key+"="+value)
			}
			found = true
		} else if line != "" {
			newLines = append(newLines, line)
		}
	}

	if !found && value != "" {
		newLines = append(newLines, key+"="+value)
	}

	// Write back
	// Note: 0600 permissions apply on Unix-like systems; Windows handles file permissions differently
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
	return a.readValueFromFile(TokenEnvKey)
}

func (a *Auth) readValueFromFile(key string) (string, error) {
	file, err := os.Open(a.envFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			return strings.TrimPrefix(line, key+"="), nil
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

func randomHex(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	if len(b) == 16 {
		b[6] = (b[6] & 0x0f) | 0x40
		b[8] = (b[8] & 0x3f) | 0x80
		return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
			uint32(b[0])<<24|uint32(b[1])<<16|uint32(b[2])<<8|uint32(b[3]),
			uint16(b[4])<<8|uint16(b[5]),
			uint16(b[6])<<8|uint16(b[7]),
			uint16(b[8])<<8|uint16(b[9]),
			uint64(b[10])<<40|uint64(b[11])<<32|uint64(b[12])<<24|uint64(b[13])<<16|uint64(b[14])<<8|uint64(b[15]),
		), nil
	}
	return fmt.Sprintf("%x", b), nil
}

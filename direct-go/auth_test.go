package direct

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewAuth(t *testing.T) {
	auth := NewAuth()
	if auth == nil {
		t.Fatal("Expected auth to be created")
	}
	if auth.envFile != EnvFile {
		t.Errorf("Expected envFile to be %s, got %s", EnvFile, auth.envFile)
	}
}

func TestNewAuthWithFile(t *testing.T) {
	customFile := ".env.test"
	auth := NewAuthWithFile(customFile)
	if auth.envFile != customFile {
		t.Errorf("Expected envFile to be %s, got %s", customFile, auth.envFile)
	}
}

func TestSetAndGetToken(t *testing.T) {
	// Create temporary directory for test
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")

	auth := NewAuthWithFile(envFile)

	// Test setting token
	testToken := "test-token-12345"
	err := auth.SetToken(testToken)
	if err != nil {
		t.Fatalf("SetToken failed: %v", err)
	}

	// Test getting token
	token := auth.GetToken()
	if token != testToken {
		t.Errorf("Expected token %s, got %s", testToken, token)
	}

	// Test HasToken
	if !auth.HasToken() {
		t.Error("Expected HasToken to return true")
	}

	// Verify file content
	content, err := os.ReadFile(envFile)
	if err != nil {
		t.Fatalf("Failed to read env file: %v", err)
	}

	expected := TokenEnvKey + "=" + testToken + "\n"
	if string(content) != expected {
		t.Errorf("Expected file content %q, got %q", expected, string(content))
	}
}

func TestUpdateToken(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")

	auth := NewAuthWithFile(envFile)

	// Set initial token
	err := auth.SetToken("initial-token")
	if err != nil {
		t.Fatalf("SetToken failed: %v", err)
	}

	// Update token
	newToken := "updated-token"
	err = auth.SetToken(newToken)
	if err != nil {
		t.Fatalf("SetToken failed: %v", err)
	}

	// Verify updated token
	token := auth.GetToken()
	if token != newToken {
		t.Errorf("Expected token %s, got %s", newToken, token)
	}

	// Verify there's only one token line in the file
	content, err := os.ReadFile(envFile)
	if err != nil {
		t.Fatalf("Failed to read env file: %v", err)
	}

	expected := TokenEnvKey + "=" + newToken + "\n"
	if string(content) != expected {
		t.Errorf("Expected file content %q, got %q", expected, string(content))
	}
}

func TestClearToken(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")

	auth := NewAuthWithFile(envFile)

	// Set token
	err := auth.SetToken("test-token")
	if err != nil {
		t.Fatalf("SetToken failed: %v", err)
	}

	// Clear token
	err = auth.ClearToken()
	if err != nil {
		t.Fatalf("ClearToken failed: %v", err)
	}

	// Verify token is cleared
	if auth.HasToken() {
		t.Error("Expected HasToken to return false after clearing")
	}

	token := auth.GetToken()
	if token != "" {
		t.Errorf("Expected empty token, got %s", token)
	}

	// Verify file is empty (or contains only newline)
	content, err := os.ReadFile(envFile)
	if err != nil {
		t.Fatalf("Failed to read env file: %v", err)
	}

	if string(content) != "\n" && string(content) != "" {
		t.Errorf("Expected empty file or newline, got %q", string(content))
	}
}

func TestGetTokenFromEnvironment(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")

	auth := NewAuthWithFile(envFile)

	// Set environment variable
	expectedToken := "env-token-123"
	os.Setenv(TokenEnvKey, expectedToken)
	defer os.Unsetenv(TokenEnvKey)

	// GetToken should return environment variable value
	token := auth.GetToken()
	if token != expectedToken {
		t.Errorf("Expected token from environment %s, got %s", expectedToken, token)
	}

	// HasToken should return true
	if !auth.HasToken() {
		t.Error("Expected HasToken to return true when env var is set")
	}
}

func TestLoadEnv(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")

	// Create .env file with multiple variables
	content := `HUBOT_DIRECT_TOKEN=token123
CUSTOM_VAR=value456
# This is a comment
ANOTHER_VAR=test

`
	err := os.WriteFile(envFile, []byte(content), 0600)
	if err != nil {
		t.Fatalf("Failed to create env file: %v", err)
	}

	auth := NewAuthWithFile(envFile)

	// Ensure env vars are not set
	os.Unsetenv("HUBOT_DIRECT_TOKEN")
	os.Unsetenv("CUSTOM_VAR")
	os.Unsetenv("ANOTHER_VAR")

	// Load environment
	err = auth.LoadEnv()
	if err != nil {
		t.Fatalf("LoadEnv failed: %v", err)
	}

	// Verify variables were loaded
	if os.Getenv("HUBOT_DIRECT_TOKEN") != "token123" {
		t.Errorf("Expected HUBOT_DIRECT_TOKEN=token123, got %s", os.Getenv("HUBOT_DIRECT_TOKEN"))
	}
	if os.Getenv("CUSTOM_VAR") != "value456" {
		t.Errorf("Expected CUSTOM_VAR=value456, got %s", os.Getenv("CUSTOM_VAR"))
	}
	if os.Getenv("ANOTHER_VAR") != "test" {
		t.Errorf("Expected ANOTHER_VAR=test, got %s", os.Getenv("ANOTHER_VAR"))
	}

	// Cleanup
	os.Unsetenv("HUBOT_DIRECT_TOKEN")
	os.Unsetenv("CUSTOM_VAR")
	os.Unsetenv("ANOTHER_VAR")
}

func TestLoadEnvDoesNotOverride(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")

	// Create .env file
	content := `TEST_VAR=from_file
`
	err := os.WriteFile(envFile, []byte(content), 0600)
	if err != nil {
		t.Fatalf("Failed to create env file: %v", err)
	}

	// Set environment variable before loading
	os.Setenv("TEST_VAR", "from_env")
	defer os.Unsetenv("TEST_VAR")

	auth := NewAuthWithFile(envFile)

	// Load environment
	err = auth.LoadEnv()
	if err != nil {
		t.Fatalf("LoadEnv failed: %v", err)
	}

	// Verify variable was not overridden
	if os.Getenv("TEST_VAR") != "from_env" {
		t.Errorf("Expected TEST_VAR=from_env (not overridden), got %s", os.Getenv("TEST_VAR"))
	}
}

func TestLoadEnvMissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env.missing")

	auth := NewAuthWithFile(envFile)

	// Loading missing file should not return error
	err := auth.LoadEnv()
	if err != nil {
		t.Errorf("LoadEnv with missing file should not error, got: %v", err)
	}
}

func TestHasTokenWithMissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env.missing")

	// Ensure environment variable is not set
	os.Unsetenv(TokenEnvKey)

	auth := NewAuthWithFile(envFile)

	// HasToken should return false for missing file
	if auth.HasToken() {
		t.Error("Expected HasToken to return false for missing file")
	}
}

func TestGetTokenWithMissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env.missing")

	// Ensure environment variable is not set
	os.Unsetenv(TokenEnvKey)

	auth := NewAuthWithFile(envFile)

	// GetToken should return empty string for missing file
	token := auth.GetToken()
	if token != "" {
		t.Errorf("Expected empty token for missing file, got %s", token)
	}
}

func TestSetTokenPreservesOtherVars(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env")

	// Create .env file with other variables
	initialContent := `OTHER_VAR=value1
HUBOT_DIRECT_TOKEN=old-token
ANOTHER_VAR=value2
`
	err := os.WriteFile(envFile, []byte(initialContent), 0600)
	if err != nil {
		t.Fatalf("Failed to create env file: %v", err)
	}

	auth := NewAuthWithFile(envFile)

	// Update token
	err = auth.SetToken("new-token")
	if err != nil {
		t.Fatalf("SetToken failed: %v", err)
	}

	// Read file and verify other variables are preserved
	content, err := os.ReadFile(envFile)
	if err != nil {
		t.Fatalf("Failed to read env file: %v", err)
	}

	fileContent := string(content)

	// Check that other variables are preserved
	if !contains(fileContent, "OTHER_VAR=value1") {
		t.Error("Expected OTHER_VAR to be preserved")
	}
	if !contains(fileContent, "ANOTHER_VAR=value2") {
		t.Error("Expected ANOTHER_VAR to be preserved")
	}
	if !contains(fileContent, "HUBOT_DIRECT_TOKEN=new-token") {
		t.Error("Expected HUBOT_DIRECT_TOKEN to be updated")
	}
	if contains(fileContent, "old-token") {
		t.Error("Old token should not be present")
	}
}

func TestSetTokenCreatesFileIfNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	envFile := filepath.Join(tmpDir, ".env.new")

	auth := NewAuthWithFile(envFile)

	// Set token when file doesn't exist
	err := auth.SetToken("new-token")
	if err != nil {
		t.Fatalf("SetToken failed: %v", err)
	}

	// Verify file was created
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		t.Error("Expected env file to be created")
	}

	// Verify token
	token := auth.GetToken()
	if token != "new-token" {
		t.Errorf("Expected token 'new-token', got %s", token)
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

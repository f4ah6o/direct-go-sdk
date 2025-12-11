package cli

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templateFS embed.FS

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup a new daabgo bot project",
	Long:  `Initialize a new daabgo bot project in the current directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit()
	},
}

func runInit() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	projectName := filepath.Base(cwd)
	fmt.Printf("Initializing daabgo project: %s\n", projectName)

	// Create project structure
	dirs := []string{
		"scripts",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Create main.go
	mainContent := `package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

func main() {
	// Load environment
	auth := direct.NewAuth()
	if err := auth.LoadEnv(); err != nil {
		log.Printf("Warning: could not load .env: %v", err)
	}

	token := auth.GetToken()
	if token == "" {
		log.Fatal("No access token. Run 'daabgo login' first.")
	}

	// Create client
	client := direct.NewClient(direct.Options{
		AccessToken: token,
	})

	// Register handlers
	pingPattern := regexp.MustCompile("(?i)ping$")
	echoPattern := regexp.MustCompile("(?i)echo (.+)$")

	client.OnMessage(func(msg direct.Message) {
		if pingPattern.MatchString(msg.Text) {
			client.SendText(msg.RoomID, "PONG")
		} else if matches := echoPattern.FindStringSubmatch(msg.Text); len(matches) > 1 {
			client.SendText(msg.RoomID, matches[1])
		}
	})

	// Connect
	fmt.Println("Connecting to direct...")
	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	fmt.Println("Bot is running! Press Ctrl+C to stop.")

	// Wait for interrupt
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\nShutting down...")
}
`

	if err := os.WriteFile("main.go", []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}

	// Create go.mod
	goModContent := fmt.Sprintf(`module %s

go 1.21

require github.com/f4ah6o/direct-go v0.0.0
`, projectName)

	if err := os.WriteFile("go.mod", []byte(goModContent), 0644); err != nil {
		return fmt.Errorf("failed to create go.mod: %w", err)
	}

	// Create .gitignore
	gitignoreContent := `.env
*.exe
*.dll
*.so
*.dylib
`

	if err := os.WriteFile(".gitignore", []byte(gitignoreContent), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}

	fmt.Println("daabgo initialized successfully!")
	fmt.Println("")
	fmt.Println("Next steps:")
	fmt.Println("  1. Run 'daabgo login' to authenticate with direct")
	fmt.Println("  2. Run 'daabgo run' to start the bot")
	fmt.Println("")

	return nil
}

// Template data for project generation
type ProjectData struct {
	Name        string
	Version     string
	Author      string
	Description string
}

// renderTemplate renders a template with the given data
func renderTemplate(tmplContent string, data interface{}) (string, error) {
	tmpl, err := template.New("").Parse(tmplContent)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

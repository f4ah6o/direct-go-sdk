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
	"context"
	"log"

	"github.com/f4ah6o/direct-go-sdk/daab-go/bot"
)

func main() {
	// Create a new bot instance
	robot := bot.New(
		bot.WithName("mybot"),
	)

	// Register a handler that responds when the bot is directly mentioned
	robot.Respond("ping", func(ctx context.Context, res bot.Response) {
		if err := res.Send("PONG"); err != nil {
			log.Printf("Failed to send response: %v", err)
		}
	})

	// Register a handler that listens to all messages
	robot.Hear(".*", func(ctx context.Context, res bot.Response) {
		log.Printf("[%s] %s: %s", res.RoomID(), res.UserID(), res.Text())
	})

	// Run the bot
	if err := robot.Run(context.Background()); err != nil {
		log.Fatalf("Failed to run bot: %v", err)
	}
}
`

	if err := os.WriteFile("main.go", []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}

	// Create go.mod
	goModContent := fmt.Sprintf(`module %s

go 1.21

require github.com/f4ah6o/direct-go-sdk/daab-go v0.0.0
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
	fmt.Println("  2. Run 'go get github.com/f4ah6o/direct-go-sdk/daab-go' to fetch dependencies")
	fmt.Println("  3. Run 'go run .' to start the bot")
	fmt.Println("")
	fmt.Println("The generated bot uses the high-level daab-go/bot framework.")
	fmt.Println("Customize main.go to add your bot logic (Respond, Hear handlers).")
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

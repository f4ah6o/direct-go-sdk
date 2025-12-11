package cli

import (
	"fmt"

	direct "github.com/f4ah6o/direct-go"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from the direct service",
	Long:  `Remove the stored access token and logout from the direct service.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLogout()
	},
}

func runLogout() error {
	auth := direct.NewAuth()

	if !auth.HasToken() {
		fmt.Println("Not logged in.")
		return nil
	}

	if err := auth.ClearToken(); err != nil {
		return fmt.Errorf("failed to clear token: %w", err)
	}

	fmt.Println("Logged out successfully.")
	return nil
}

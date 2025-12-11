package cli

import (
	"context"
	"fmt"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/spf13/cobra"
)

var invitesCmd = &cobra.Command{
	Use:   "invites",
	Short: "Show and accept domain invites",
	Long:  `List pending domain invites and optionally accept them.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return showInvites()
	},
}

func showInvites() error {
	auth := direct.NewAuth()

	// Load environment
	if err := auth.LoadEnv(); err != nil {
		fmt.Printf("Warning: could not load .env: %v\n", err)
	}

	// Check if logged in
	if !auth.HasToken() {
		fmt.Println("Not logged in. Run 'daabgo login' first.")
		return nil
	}

	token := auth.GetToken()

	// Create client
	client := direct.NewClient(direct.Options{
		Endpoint:    direct.DefaultEndpoint,
		AccessToken: token,
	})

	fmt.Println("Connecting to direct...")
	if err := client.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer client.Close()

	// Get domain invites
	ctx := context.Background()
	invites, err := client.GetDomainInvitesWithContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to get domain invites: %w", err)
	}

	if len(invites) == 0 {
		fmt.Println("No pending domain invites.")
		return nil
	}

	// Display invites
	fmt.Printf("Found %d pending invite(s):\n\n", len(invites))
	for i, invite := range invites {
		fmt.Printf("%d. Domain: %s\n", i+1, invite.Name)
		fmt.Printf("   ID: %v\n", invite.ID)
		if invite.UpdatedAt > 0 {
			fmt.Printf("   Updated: %d\n", invite.UpdatedAt)
		}
		fmt.Println()
	}

	// Ask user if they want to accept any
	fmt.Print("Enter invite number to accept (or 0 to skip): ")
	var choice int
	fmt.Scanln(&choice)

	if choice > 0 && choice <= len(invites) {
		invite := invites[choice-1]
		fmt.Printf("Accepting invite to domain: %s\n", invite.Name)

		_, err := client.AcceptDomainInviteWithContext(ctx, invite.ID)
		if err != nil {
			return fmt.Errorf("failed to accept invite: %w", err)
		}

		fmt.Println("Invite accepted successfully!")
	} else {
		fmt.Println("No invite accepted.")
	}

	return nil
}

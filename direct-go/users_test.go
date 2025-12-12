package direct

import (
	"context"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-go/testutil"
)

func TestGetUsers(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	// Setup mock response
	mockServer.OnSimple("get_users", []interface{}{
		map[string]interface{}{
			"id":           "user1",
			"display_name": "User One",
			"email":        "user1@example.com",
		},
		map[string]interface{}{
			"id":           "user2",
			"display_name": "User Two",
			"email":        "user2@example.com",
		},
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	userIDs := []interface{}{"user1", "user2"}
	users, err := client.GetUsers(ctx, "domain123", userIDs)
	if err != nil {
		t.Fatalf("GetUsers failed: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}

	if users[0].ID != "user1" {
		t.Errorf("Expected first user ID=user1, got %v", users[0].ID)
	}
}

func TestGetProfile(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("get_profile", map[string]interface{}{
		"user_id":      "user123",
		"display_name": "Test User",
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	profile, err := client.GetProfile(ctx, "domain123", "user123")
	if err != nil {
		t.Fatalf("GetProfile failed: %v", err)
	}

	if profile.UserID != "user123" {
		t.Errorf("Expected user_id=user123, got %v", profile.UserID)
	}

	if profile.DisplayName != "Test User" {
		t.Errorf("Expected display_name='Test User', got %v", profile.DisplayName)
	}
}

func TestUpdateProfile(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("update_profile", true)

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	updates := map[string]interface{}{
		"display_name": "New Name",
	}

	err = client.UpdateProfile(ctx, "domain123", updates)
	if err != nil {
		t.Fatalf("UpdateProfile failed: %v", err)
	}

	if mockServer.GetCallCount("update_profile") != 1 {
		t.Errorf("Expected update_profile to be called once")
	}
}

func TestGetFriends(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("get_friends", []interface{}{
		map[string]interface{}{
			"id":           "friend1",
			"display_name": "Friend One",
		},
		map[string]interface{}{
			"id":           "friend2",
			"display_name": "Friend Two",
		},
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	friends, err := client.GetFriends(ctx)
	if err != nil {
		t.Fatalf("GetFriends failed: %v", err)
	}

	if len(friends) != 2 {
		t.Errorf("Expected 2 friends, got %d", len(friends))
	}
}

func TestGetAcquaintances(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("get_acquaintances", []interface{}{
		map[string]interface{}{
			"id":           "acq1",
			"display_name": "Acquaintance One",
		},
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	acquaintances, err := client.GetAcquaintances(ctx)
	if err != nil {
		t.Fatalf("GetAcquaintances failed: %v", err)
	}

	if len(acquaintances) != 1 {
		t.Errorf("Expected 1 acquaintance, got %d", len(acquaintances))
	}
}

func TestAddAndDeleteFriend(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("add_friend", true)
	mockServer.OnSimple("delete_friend", true)

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Add friend
	err = client.AddFriend(ctx, "user123")
	if err != nil {
		t.Fatalf("AddFriend failed: %v", err)
	}

	if mockServer.GetCallCount("add_friend") != 1 {
		t.Errorf("Expected add_friend to be called once, got %d", mockServer.GetCallCount("add_friend"))
	}

	// Delete friend
	err = client.DeleteFriend(ctx, "user123")
	if err != nil {
		t.Fatalf("DeleteFriend failed: %v", err)
	}

	if mockServer.GetCallCount("delete_friend") != 1 {
		t.Errorf("Expected delete_friend to be called once, got %d", mockServer.GetCallCount("delete_friend"))
	}
}

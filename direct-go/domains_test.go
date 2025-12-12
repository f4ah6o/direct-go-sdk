package direct

import (
	"context"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-go/testutil"
)

func TestGetDomainsWithContext(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("get_domains", []interface{}{
		map[string]interface{}{
			"id":          "domain1",
			"name":        "Test Domain 1",
			"domain_code": "test-domain-1",
		},
		map[string]interface{}{
			"id":          "domain2",
			"name":        "Test Domain 2",
			"domain_code": "test-domain-2",
		},
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	domains, err := client.GetDomainsWithContext(ctx)
	if err != nil {
		t.Fatalf("GetDomainsWithContext failed: %v", err)
	}

	if len(domains) != 2 {
		t.Errorf("Expected 2 domains, got %d", len(domains))
	}
}

func TestGetDomainInvitesWithContext(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("get_domain_invites", []interface{}{
		map[string]interface{}{
			"id":         "invite1",
			"name":       "Test Domain",
			"updated_at": int64(1702345678),
		},
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	invites, err := client.GetDomainInvitesWithContext(ctx)
	if err != nil {
		t.Fatalf("GetDomainInvitesWithContext failed: %v", err)
	}

	if len(invites) != 1 {
		t.Errorf("Expected 1 invite, got %d", len(invites))
	}

	if invites[0].ID != "invite1" {
		t.Errorf("Expected invite ID=invite1, got %v", invites[0].ID)
	}
}

func TestAcceptDomainInviteWithContext(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("accept_domain_invite", map[string]interface{}{
		"id":     "domain123",
		"name":   "Test Domain",
		"closed": false,
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	result, err := client.AcceptDomainInviteWithContext(ctx, "invite123")
	if err != nil {
		t.Fatalf("AcceptDomainInviteWithContext failed: %v", err)
	}

	if result.ID != "domain123" {
		t.Errorf("Expected domain ID=domain123, got %v", result.ID)
	}
}

func TestDeleteDomainInvite(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("delete_domain_invite", true)

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	err = client.DeleteDomainInvite(ctx, "invite123")
	if err != nil {
		t.Fatalf("DeleteDomainInvite failed: %v", err)
	}

	if mockServer.GetCallCount("delete_domain_invite") != 1 {
		t.Errorf("Expected delete_domain_invite to be called once")
	}
}

func TestLeaveDomain(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("leave_domain", true)

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	err = client.LeaveDomain(ctx, "domain123")
	if err != nil {
		t.Fatalf("LeaveDomain failed: %v", err)
	}

	if mockServer.GetCallCount("leave_domain") != 1 {
		t.Errorf("Expected leave_domain to be called once")
	}
}

func TestGetDomainUsers(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("get_domain_users", []interface{}{
		map[string]interface{}{
			"id":           "user1",
			"display_name": "Domain User 1",
			"email":        "user1@domain.com",
		},
		map[string]interface{}{
			"id":           "user2",
			"display_name": "Domain User 2",
			"email":        "user2@domain.com",
		},
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	users, err := client.GetDomainUsers(ctx, "domain123")
	if err != nil {
		t.Fatalf("GetDomainUsers failed: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestSearchDomainUsers(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnDynamic("search_domain_users", func(params []interface{}) (interface{}, error) {
		// params: [domainID, query]
		return []interface{}{
			map[string]interface{}{
				"id":           "user_found",
				"display_name": "Found User",
				"email":        "found@domain.com",
			},
		}, nil
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	results, err := client.SearchDomainUsers(ctx, "domain123", "search query")
	if err != nil {
		t.Fatalf("SearchDomainUsers failed: %v", err)
	}

	if len(results) != 1 {
		t.Errorf("Expected 1 search result, got %d", len(results))
	}
}

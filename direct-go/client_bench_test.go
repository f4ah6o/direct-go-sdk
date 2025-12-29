//go:build !short
// +build !short

package direct

import (
	"testing"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-go/testutil"
)

func BenchmarkClientCall(b *testing.B) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("test_method", map[string]interface{}{
		"result": "success",
	})

	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	if err := client.Connect(); err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.Call("test_method", []interface{}{})
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkClientCallConcurrent(b *testing.B) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("test_method", map[string]interface{}{
		"result": "success",
	})

	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	if err := client.Connect(); err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		go func() {
			_, _ = client.Call("test_method", []interface{}{})
		}()
	}
	// Wait for all goroutines to complete
	time.Sleep(100 * time.Millisecond)
}

func BenchmarkClientSendText(b *testing.B) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("create_message", map[string]interface{}{
		"id":      "msg123",
		"talk_id": "room456",
	})

	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	if err := client.Connect(); err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := client.SendText("room456", "test message")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClientConnect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mockServer := testutil.NewMockServer()
		client := NewClient(Options{
			Endpoint: mockServer.URL(),
		})

		if err := client.Connect(); err != nil {
			b.Fatal(err)
		}
		client.Close()
		mockServer.Close()
	}
}

func BenchmarkGetMe(b *testing.B) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("GetMe", map[string]interface{}{
		"id":       "user123",
		"name":     "Test User",
		"icon_url": "",
	})

	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	if err := client.Connect(); err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.GetMe()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetTalks(b *testing.B) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("GetTalks", []interface{}{
		map[string]interface{}{
			"id":   int64(12345),
			"name": "Test Room",
		},
	})

	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	if err := client.Connect(); err != nil {
		b.Fatal(err)
	}
	defer client.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.GetTalks()
		if err != nil {
			b.Fatal(err)
		}
	}
}

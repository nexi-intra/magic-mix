package subscription

import (
	"testing"
	"time"
)

func TestAddSubscription(t *testing.T) {
	store := NewInMemorySubscriptionStore()
	ttl := 5
	service := NewSubscriptionService(store, ttl)

	req := map[string]interface{}{
		"action": "add",
		"id":     "test_subscription_1",
	}

	resp, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to add subscription: %v", err)
	}

	if resp.(map[string]string)["status"] != "added" {
		t.Fatalf("Expected status 'added', got: %v", resp)
	}

	sub, err := store.Get("test_subscription_1")
	if err != nil {
		t.Fatalf("Expected subscription to be added, but got error: %v", err)
	}
	if sub.ID != "test_subscription_1" {
		t.Fatalf("Expected subscription ID 'test_subscription_1', got: %v", sub.ID)
	}
}

func TestAddExistingSubscription(t *testing.T) {
	store := NewInMemorySubscriptionStore()
	ttl := 5
	service := NewSubscriptionService(store, ttl)

	// Add a subscription
	req := map[string]interface{}{
		"action": "add",
		"id":     "test_subscription_2",
	}
	_, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to add subscription: %v", err)
	}

	// Try adding the same subscription again
	_, err = service.HandleRequest(req)
	if err == nil || err.Error() != "subscription already exists" {
		t.Fatalf("Expected error 'subscription already exists', got: %v", err)
	}
}

func TestRemoveSubscription(t *testing.T) {
	store := NewInMemorySubscriptionStore()
	ttl := 5
	service := NewSubscriptionService(store, ttl)

	// Add a subscription
	req := map[string]interface{}{
		"action": "add",
		"id":     "test_subscription_3",
	}
	_, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to add subscription: %v", err)
	}

	// Remove the subscription
	req = map[string]interface{}{
		"action": "remove",
		"id":     "test_subscription_3",
	}
	resp, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to remove subscription: %v", err)
	}

	if resp.(map[string]string)["status"] != "removed" {
		t.Fatalf("Expected status 'removed', got: %v", resp)
	}

	// Verify the subscription is removed
	_, err = store.Get("test_subscription_3")
	if err == nil {
		t.Fatalf("Expected subscription to be removed, but it still exists")
	}
}

func TestUpdateSubscription(t *testing.T) {
	store := NewInMemorySubscriptionStore()
	ttl := 5
	service := NewSubscriptionService(store, ttl)

	// Add a subscription
	req := map[string]interface{}{
		"action": "add",
		"id":     "test_subscription_4",
	}
	_, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to add subscription: %v", err)
	}

	// Update the subscription's LastSeen time
	time.Sleep(1) // Simulate some delay before updating
	req = map[string]interface{}{
		"action": "update",
		"id":     "test_subscription_4",
	}
	resp, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to update subscription: %v", err)
	}

	if resp.(map[string]string)["status"] != "updated" {
		t.Fatalf("Expected status 'updated', got: %v", resp)
	}

	// Ensure the tracker updated the LastSeen time
	lastSeen, err := service.tracker.GetLastSeen("test_subscription_4")
	if err != nil {
		t.Fatalf("Expected LastSeen to be updated, but got error: %v", err)
	}

	if time.Since(lastSeen) > 2*time.Second {
		t.Fatalf("Expected LastSeen to be recently updated, but it wasn't")
	}
}

func TestListSubscriptions(t *testing.T) {
	store := NewInMemorySubscriptionStore()
	ttl := 5
	service := NewSubscriptionService(store, ttl)

	// Add subscriptions
	for i := 1; i <= 3; i++ {
		req := map[string]interface{}{
			"action": "add",
			"id":     "test_subscription_list_" + string(rune(i)),
		}
		_, err := service.HandleRequest(req)
		if err != nil {
			t.Fatalf("Failed to add subscription %d: %v", i, err)
		}
	}

	// List subscriptions
	req := map[string]interface{}{
		"action": "list",
	}
	resp, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to list subscriptions: %v", err)
	}

	subs := resp.([]*Subscription)
	if len(subs) != 3 {
		t.Fatalf("Expected 3 subscriptions, got: %d", len(subs))
	}
}

func TestSubscriptionTTL(t *testing.T) {
	store := NewInMemorySubscriptionStore()
	ttl := 2
	service := NewSubscriptionService(store, ttl)

	// Add a subscription
	req := map[string]interface{}{
		"action": "add",
		"id":     "test_subscription_ttl",
	}
	_, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to add subscription: %v", err)
	}

	// Wait for the TTL to expire
	time.Sleep(5)

	// Check if the subscription was automatically removed
	_, err = store.Get("test_subscription_ttl")
	if err == nil {
		t.Fatalf("Expected subscription to be removed due to TTL, but it still exists")
	}
}

func TestGetMessages(t *testing.T) {
	store := NewInMemorySubscriptionStore()
	ttl := 5
	service := NewSubscriptionService(store, ttl)

	// Add a subscription
	req := map[string]interface{}{
		"action": "add",
		"id":     "test_subscription_5",
	}
	_, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to add subscription: %v", err)
	}

	// Fetch messages with a timeout
	req = map[string]interface{}{
		"action":  "getMessages",
		"id":      "test_subscription_5",
		"timeout": 3.0, // 3 seconds timeout
	}
	resp, err := service.HandleRequest(req)
	if err != nil {
		t.Fatalf("Failed to get messages: %v", err)
	}

	messages := resp.(map[string]interface{})["messages"].([]string)
	if len(messages) == 0 {
		t.Fatalf("Expected some messages, but got none")
	}
}

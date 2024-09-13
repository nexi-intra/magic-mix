package drivers

import (
	"testing"

	"github.com/magicbutton/magic-mix/subscription"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func setupTestNATS(t *testing.T) (*nats.Conn, func()) {
	// For an actual test, use an embedded or local NATS server.
	// Mocking could be applied here if needed.
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Fatalf("Unable to connect to NATS: %v", err)
	}

	return nc, func() {
		nc.Close()
	}
}

func TestNewJetStreamSubscriptionStore(t *testing.T) {
	nc, cleanup := setupTestNATS(t)
	defer cleanup()

	// Create a new JetStreamSubscriptionStore
	store, err := NewJetStreamSubscriptionStore(nc, "TEST_STREAM", "test.*")
	assert.NoError(t, err, "Error should be nil when creating a new store")
	assert.NotNil(t, store, "Store should be initialized")
}

func TestSetSubscription(t *testing.T) {
	nc, cleanup := setupTestNATS(t)
	defer cleanup()

	store, err := NewJetStreamSubscriptionStore(nc, "TEST_STREAM", "test.*")
	assert.NoError(t, err)

	// Add a subscription
	sub := &subscription.Subscription{
		ID: "test-sub",
	}
	err = store.Set(sub)
	assert.NoError(t, err, "Error should be nil when setting a subscription")
}

func TestGetSubscription(t *testing.T) {
	nc, cleanup := setupTestNATS(t)
	defer cleanup()

	store, err := NewJetStreamSubscriptionStore(nc, "TEST_STREAM", "test.*")
	assert.NoError(t, err)

	// Add a subscription
	sub := &subscription.Subscription{
		ID: "test-sub",
	}
	err = store.Set(sub)
	assert.NoError(t, err)

	// Retrieve the subscription
	retrievedSub, err := store.Get("test-sub")
	assert.NoError(t, err, "Error should be nil when getting a subscription")
	assert.NotNil(t, retrievedSub, "Retrieved subscription should not be nil")
	assert.Equal(t, sub.ID, retrievedSub.ID, "Subscription ID should match")
}

func TestListSubscriptions(t *testing.T) {
	nc, cleanup := setupTestNATS(t)
	defer cleanup()

	store, err := NewJetStreamSubscriptionStore(nc, "TEST_STREAM", "test.*")
	assert.NoError(t, err)

	// Add multiple subscriptions
	sub1 := &subscription.Subscription{ID: "sub1"}
	sub2 := &subscription.Subscription{ID: "sub2"}
	err = store.Set(sub1)
	assert.NoError(t, err)
	err = store.Set(sub2)
	assert.NoError(t, err)

	// List all subscriptions
	subs, err := store.List()
	assert.NoError(t, err, "Error should be nil when listing subscriptions")
	assert.Equal(t, 2, len(subs), "There should be two subscriptions")
}

func TestRemoveSubscription(t *testing.T) {
	nc, cleanup := setupTestNATS(t)
	defer cleanup()

	store, err := NewJetStreamSubscriptionStore(nc, "TEST_STREAM", "test.*")
	assert.NoError(t, err)

	// Add and remove a subscription
	sub := &subscription.Subscription{ID: "test-sub"}
	err = store.Set(sub)
	assert.NoError(t, err)

	err = store.Remove("test-sub")
	assert.NoError(t, err, "Error should be nil when removing a subscription")

	// Try to get the removed subscription
	_, err = store.Get("test-sub")
	assert.Error(t, err, "Should return an error when trying to get a removed subscription")
}

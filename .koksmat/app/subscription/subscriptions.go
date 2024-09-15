package subscription

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// Subscription represents a basic subscriber without TTL and LastSeen directly tied to it.
type Subscription struct {
	ID string `json:"id"`
	// other relevant subscription fields can go here
}

type Message struct {
	Subject string `json:"subject"`
	Data    string `json:"data"`
}

// SubscriptionStore is the interface that must be implemented by any subscription storage mechanism.
type SubscriptionStore interface {
	Get(id string) (*Subscription, error)
	Set(sub *Subscription) error
	List() ([]*Subscription, error)
	Remove(id string) error
	ReadMessages(id string) ([]Message, error)
}

// LastSeenTracker manages the last seen timestamps for subscriptions.
type LastSeenTracker interface {
	SetLastSeen(subID string)
	GetLastSeen(subID string) (time.Time, error)
	HasExpired(subID string, ttlSeconds int) bool
}

// TTLManager handles TTL expiration for subscriptions.
type TTLManager interface {
	CheckExpiredSubscriptions(subscriptions []*Subscription)
}

// InMemoryLastSeenTracker is an example implementation of LastSeenTracker interface.
type InMemoryLastSeenTracker struct {
	lastSeen map[string]time.Time
	mu       sync.Mutex
}

// NewInMemoryLastSeenTracker creates a new instance of InMemoryLastSeenTracker.
func NewInMemoryLastSeenTracker() *InMemoryLastSeenTracker {
	return &InMemoryLastSeenTracker{
		lastSeen: make(map[string]time.Time),
	}
}

// SetLastSeen updates the last seen timestamp for a specific subscription.
func (t *InMemoryLastSeenTracker) SetLastSeen(subID string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.lastSeen[subID] = time.Now()
}

// GetLastSeen retrieves the last seen timestamp for a specific subscription.
func (t *InMemoryLastSeenTracker) GetLastSeen(subID string) (time.Time, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	lastSeen, exists := t.lastSeen[subID]
	if !exists {
		return time.Time{}, errors.New("subscription not found")
	}
	return lastSeen, nil
}

// HasExpired checks if a subscription's TTL has expired based on its last seen timestamp.
func (t *InMemoryLastSeenTracker) HasExpired(subID string, ttlSeconds int) bool {
	lastSeen, err := t.GetLastSeen(subID)
	if err != nil {
		return true // Treat missing subscriptions as expired
	}

	// Convert ttl from seconds to time.Duration
	ttl := time.Duration(ttlSeconds) * time.Second

	return time.Since(lastSeen) > ttl
}

// DefaultTTLManager is an example TTL manager that checks for expired subscriptions.
type DefaultTTLManager struct {
	tracker LastSeenTracker
	store   SubscriptionStore
	ttl     int
	mu      sync.Mutex
}

// NewDefaultTTLManager creates a new instance of DefaultTTLManager.
func NewDefaultTTLManager(tracker LastSeenTracker, store SubscriptionStore, ttlSeconds int) *DefaultTTLManager {
	return &DefaultTTLManager{
		tracker: tracker,
		store:   store,
		ttl:     ttlSeconds,
	}
}

// CheckExpiredSubscriptions checks and handles subscriptions that have exceeded their TTL.
func (m *DefaultTTLManager) CheckExpiredSubscriptions(subscriptions []*Subscription) {
	// m.mu.Lock()
	// defer m.mu.Unlock()

	// for _, sub := range subscriptions {
	// 	if m.tracker.HasExpired(sub.ID, m.ttl) {
	// 		// Remove the expired subscription from the store
	// 		err := m.store.Remove(sub.ID)
	// 		if err != nil {
	// 			fmt.Printf("Failed to remove expired subscription %s: %v\n", sub.ID, err)
	// 		} else {
	// 			fmt.Printf("Subscription %s has expired and was removed\n", sub.ID)
	// 		}
	// 	}
	// }
}

// SubscriptionService manages subscriptions and their lifecycle.
type SubscriptionService struct {
	store      SubscriptionStore
	tracker    LastSeenTracker
	ttlManager TTLManager
	ttl        int
	mu         sync.Mutex
}

// NewSubscriptionService creates a new SubscriptionService and starts TTL monitoring.
func NewSubscriptionService(store SubscriptionStore, ttlSeconds int) *SubscriptionService {
	tracker := NewInMemoryLastSeenTracker()
	ttlManager := NewDefaultTTLManager(tracker, store, ttlSeconds)

	service := &SubscriptionService{
		store:      store,
		tracker:    tracker,
		ttlManager: ttlManager,
		ttl:        ttlSeconds,
	}
	go service.monitorTTL() // Start TTL monitoring in a goroutine
	return service
}

// HandleRequest processes the incoming request based on the 'action' field and performs operations accordingly.
func (s *SubscriptionService) HandleRequest(req map[string]interface{}) (interface{}, error) {
	// Defer a function to recover from panics and handle them gracefully.
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Fatal error occurred: %v\n", r)
			// You can also log this or send it to an error tracking system
		}
	}()

	// Check if the action field exists in the request.
	action, ok := req["action"].(string)
	if !ok {
		// Log and return a custom error for missing action field
		err := fmt.Errorf("missing 'action' field")
		fmt.Println("Error:", err)
		return nil, err
	}
	log.Println("SubscriptionService.HandleRequest.Action:", action)
	// Switch statement to handle different actions.
	switch action {
	case "add":
		resp, err := s.handleAddSubscription(req)
		if err != nil {
			fmt.Printf("Error in add action: %v\n", err) // Log the error
			return nil, err
		}
		return resp, nil
	case "remove":
		resp, err := s.handleRemoveSubscription(req)
		if err != nil {
			fmt.Printf("Error in remove action: %v\n", err) // Log the error
			return nil, err
		}
		return resp, nil
	case "update":
		resp, err := s.handleUpdateSubscription(req)
		if err != nil {
			fmt.Printf("Error in update action: %v\n", err) // Log the error
			return nil, err
		}
		return resp, nil
	case "list":
		resp, err := s.handleListSubscriptions()
		if err != nil {
			fmt.Printf("Error in list action: %v\n", err) // Log the error
			return nil, err
		}
		return resp, nil
	case "getMessages":
		resp, err := s.handleGetMessages(req)
		if err != nil {
			fmt.Printf("Error in getMessages action: %v\n", err) // Log the error
			return nil, err
		}
		return resp, nil
	default:
		// Handle invalid action errors
		err := fmt.Errorf("invalid action")
		fmt.Println("Error:", err)
		return nil, err
	}
}

// handleAddSubscription adds a new subscription if it does not exist.
func (s *SubscriptionService) handleAddSubscription(req map[string]interface{}) (interface{}, error) {
	id, ok := req["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("invalid 'id' field")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	data, _ := s.store.Get(id)
	if data != nil {
		return nil, fmt.Errorf("subscription already exists")
	}

	sub := &Subscription{
		ID: id,
	}

	err := s.store.Set(sub)
	if err != nil {
		return nil, fmt.Errorf("failed to add subscription: %v", err)
	}

	// Set last seen and update tracker
	s.tracker.SetLastSeen(id)

	return map[string]string{"status": "added"}, nil
}

// handleRemoveSubscription removes a subscription by ID.
func (s *SubscriptionService) handleRemoveSubscription(req map[string]interface{}) (interface{}, error) {
	id, ok := req["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("invalid 'id' field")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.store.Remove(id)
	if err != nil {
		return nil, fmt.Errorf("failed to remove subscription: %v", err)
	}

	return map[string]string{"status": "removed"}, nil
}

// handleUpdateSubscription updates the LastSeen time of an existing subscription.
func (s *SubscriptionService) handleUpdateSubscription(req map[string]interface{}) (interface{}, error) {
	id, ok := req["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("invalid 'id' field")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.store.Get(id)
	if err != nil {
		return nil, fmt.Errorf("subscription not found")
	}

	// Update last seen
	s.tracker.SetLastSeen(id)

	return map[string]string{"status": "updated"}, nil
}

// handleGetMessages fetches new messages for a subscription with a timeout.
func (s *SubscriptionService) handleGetMessages(req map[string]interface{}) (interface{}, error) {
	id, ok := req["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("invalid 'id' field")
	}

	// timeout, ok := req["timeout"].(float64) // Assuming timeout is sent as a number of seconds
	// if !ok || timeout <= 0 {
	// 	return nil, fmt.Errorf("invalid or missing 'timeout' field")
	// }
	log.Println("Getting messages for subscription:", id)
	s.mu.Lock()
	_, err := s.store.Get(id)

	s.mu.Unlock()

	if err != nil {
		sub := &Subscription{
			ID: id,
		}

		err := s.store.Set(sub)
		if err != nil {
			return nil, fmt.Errorf("failed to add subscription: %v", err)
		}
		log.Println("Created subscription for:", id)
	}
	s.tracker.SetLastSeen(id)
	messages := []Message{} // Placeholder: Fetch new messages logic here
	messageChannel := make(chan []Message)

	go func() {
		// Simulate fetching messages (this logic would depend on your actual messaging system)
		// You would replace this with code that actually polls for messages related to the subscription
		msgs, error := s.store.ReadMessages(id)
		if error != nil {
			fmt.Println("Error reading messages:", error)
		}
		if len(msgs) > 0 {
			messageChannel <- msgs
		}
		// time.Sleep(2 * time.Second)                        // Simulate delay
		// messageChannel <- []string{"message1", "message2"} // Simulate messages received
	}()

	select {
	case messages = <-messageChannel:
		return map[string]interface{}{
			"status":   "success",
			"messages": messages,
		}, nil
		// case <-time.After(time.Duration(timeout) * time.Second):
		// 	return map[string]interface{}{
		// 		"status":   "timeout",
		// 		"messages": messages, // empty list
		// 	}, nil
	}
}

// handleListSubscriptions lists all subscriptions.
func (s *SubscriptionService) handleListSubscriptions() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	subs, err := s.store.List()
	if err != nil {
		return nil, fmt.Errorf("failed to list subscriptions: %v", err)
	}

	return subs, nil
}

// monitorTTL periodically checks for subscriptions that have exceeded their TTL and removes them.
func (s *SubscriptionService) monitorTTL() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		s.mu.Lock()

		subs, err := s.store.List()
		if err != nil {
			s.mu.Unlock()
			continue
		}

		// Check for expired subscriptions
		s.ttlManager.CheckExpiredSubscriptions(subs)

		s.mu.Unlock()
	}
}

// InMemorySubscriptionStore is an in-memory implementation of SubscriptionStore, useful for testing or simple use cases.
type InMemorySubscriptionStore struct {
	data map[string]*Subscription
	mu   sync.Mutex
}

// NewInMemorySubscriptionStore creates a new in-memory subscription store.
func NewInMemorySubscriptionStore() *InMemorySubscriptionStore {
	return &InMemorySubscriptionStore{data: make(map[string]*Subscription)}
}

// Get retrieves a subscription by its ID.
func (store *InMemorySubscriptionStore) Get(id string) (*Subscription, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	sub, exists := store.data[id]
	if !exists {
		return nil, errors.New("subscription not found")
	}
	return sub, nil
}

// Set adds or updates a subscription.
func (store *InMemorySubscriptionStore) Set(sub *Subscription) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.data[sub.ID] = sub
	return nil
}

// List returns all subscriptions.
func (store *InMemorySubscriptionStore) List() ([]*Subscription, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	var subs []*Subscription
	for _, sub := range store.data {
		subs = append(subs, sub)
	}
	return subs, nil
}

// Remove deletes a subscription by ID.
func (store *InMemorySubscriptionStore) Remove(id string) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	delete(store.data, id)
	return nil
}

func (store *InMemorySubscriptionStore) ReadMessages(id string) ([]Message, error) {
	// Placeholder: Fetch new messages logic here
	return []Message{}, nil
}

// Example usage
func main() {
	store := NewInMemorySubscriptionStore()
	ttl := 10 * 60
	service := NewSubscriptionService(store, ttl)

	// Example request for adding a subscription
	req := map[string]interface{}{
		"action": "add",
		"id":     "sub_123",
	}
	resp, err := service.HandleRequest(req)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", resp)
	}
}

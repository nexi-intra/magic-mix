package messaging

import (
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

type Message struct {
	ID      int
	Content []byte
}

type Subscriber struct {
	ID            int
	LastMessageID int
	Messages      chan []Message
}

type Service struct {
	nc          *nats.Conn
	subject     string
	messageID   int
	subscribers map[int]*Subscriber
	mu          sync.Mutex
}

func NewService(nc *nats.Conn, subject string) *Service {
	return &Service{
		nc:          nc,
		subject:     subject,
		messageID:   0,
		subscribers: make(map[int]*Subscriber),
		mu:          sync.Mutex{},
	}
}

// Boot function that sets up the NATS subscription and message handling
func (s *Service) Boot() error {
	// Subscribe to the subject and start listening for messages
	_, err := s.nc.Subscribe(s.subject, s.messageHandler)
	if err != nil {
		return fmt.Errorf("error subscribing to subject: %v", err)
	}

	// Infinite loop to keep the service alive
	for {
		time.Sleep(time.Second)
	}
}

// Message handler function that increments message ID and dispatches to subscribers
func (s *Service) messageHandler(msg *nats.Msg) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.messageID++
	message := Message{
		ID:      s.messageID,
		Content: msg.Data,
	}

	// Notify all subscribers
	for _, sub := range s.subscribers {
		if sub.LastMessageID < s.messageID {
			sub.Messages <- []Message{message}
		}
	}
}

// Add a new subscriber
func (s *Service) AddSubscriber(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.subscribers[id] = &Subscriber{
		ID:            id,
		LastMessageID: 0,                         // Initialize with 0, so they receive all messages
		Messages:      make(chan []Message, 100), // Buffered channel
	}
	fmt.Printf("Subscriber %d added\n", id)
}

// Remove a subscriber
func (s *Service) RemoveSubscriber(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.subscribers, id)
	fmt.Printf("Subscriber %d removed\n", id)
}

// List all subscribers
func (s *Service) ListSubscribers() []int {
	s.mu.Lock()
	defer s.mu.Unlock()

	ids := make([]int, 0, len(s.subscribers))
	for id := range s.subscribers {
		ids = append(ids, id)
	}
	return ids
}

// Get next messages for the subscriber based on their last message ID
func (s *Service) GetNextMessages(subscriberID int) ([]Message, error) {
	s.mu.Lock()
	sub, exists := s.subscribers[subscriberID]
	s.mu.Unlock()

	if !exists {
		return nil, fmt.Errorf("subscriber %d not found", subscriberID)
	}

	// Wait for new messages
	select {
	case messages := <-sub.Messages:
		// Update last message ID
		s.mu.Lock()
		sub.LastMessageID = messages[len(messages)-1].ID
		s.mu.Unlock()
		return messages, nil
	case <-time.After(10 * time.Second): // Timeout after 10 seconds
		return nil, fmt.Errorf("no new messages")
	}
}

// func main() {
// 	// Connect to NATS server
// 	nc, err := nats.Connect(nats.DefaultURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer nc.Close()

// 	// Initialize service
// 	service := NewService(nc, "example.subject")

// 	// Boot the service (infinite loop)
// 	go func() {
// 		err := service.Boot()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	// Example: Add and remove subscribers
// 	service.AddSubscriber(1)
// 	service.AddSubscriber(2)

// 	// Example: Fetch new messages for a subscriber
// 	go func() {
// 		for {
// 			messages, err := service.GetNextMessages(1)
// 			if err != nil {
// 				fmt.Println(err)
// 			} else {
// 				fmt.Printf("Subscriber 1 received messages: %v\n", messages)
// 			}
// 		}
// 	}()

// 	// Keep the main function running
// 	select {}
// }

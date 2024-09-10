package drivers

import (
	"fmt"

	"github.com/magicbutton/magic-mix/subscription"
	"github.com/nats-io/nats.go"
)

// type Subscription struct {
// 	ID      string
// 	Subject string
// 	Config  *nats.ConsumerConfig
// }

// type SubscriptionStore interface {
// 	Get(id string) (*Subscription, error)
// 	Set(sub *Subscription) error
// 	List() ([]*Subscription, error)
// 	Remove(id string) error
// }

type JetStreamSubscriptionStore struct {
	js      nats.JetStreamContext
	stream  string
	subject string
}

// Ensure the stream exists on instantiation
func NewJetStreamSubscriptionStore(nc *nats.Conn, streamName, subjectName string) (*JetStreamSubscriptionStore, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	// Ensure the stream exists
	_, err = js.StreamInfo(streamName)
	if err == nats.ErrStreamNotFound {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{subjectName},
		})
		if err != nil {
			return nil, fmt.Errorf("error creating stream: %v", err)
		}
	} else if err != nil {
		return nil, err
	}

	return &JetStreamSubscriptionStore{
		js:      js,
		stream:  streamName,
		subject: subjectName,
	}, nil
}

// Get a subscription by ID (maps to NATS consumer info)
func (s *JetStreamSubscriptionStore) Get(id string) (*subscription.Subscription, error) {
	consumerInfo, err := s.js.ConsumerInfo(s.stream, id)
	if err != nil {
		return nil, err
	}
	return &subscription.Subscription{
		ID: consumerInfo.Name,
		// Subject: consumerInfo.Config.FilterSubject,
		// Config:  &consumerInfo.Config,
	}, nil
}

// Set a subscription (maps to creating a durable consumer)
func (s *JetStreamSubscriptionStore) Set(sub *subscription.Subscription) error {
	_, err := s.js.AddConsumer(s.stream, &nats.ConsumerConfig{
		Durable: sub.ID,
		// FilterSubject: sub.Subject,
		// AckPolicy:     nats.AckExplicitPolicy, // example ack policy
	})
	if err != nil {
		return err
	}
	return nil
}

// List all subscriptions (maps to listing consumers)
func (s *JetStreamSubscriptionStore) List() ([]*subscription.Subscription, error) {
	// Consumers returns a channel, so we need to iterate over it
	consumerCh := s.js.Consumers(s.stream)

	var subscriptions []*subscription.Subscription

	for consumerInfo := range consumerCh {
		if consumerInfo == nil {
			break
		}
		subscriptions = append(subscriptions, &subscription.Subscription{
			ID: consumerInfo.Name,
			// Subject: consumerInfo.Config.FilterSubject,
			// Config:  &consumerInfo.Config,
		})
	}

	return subscriptions, nil
}

// Remove a subscription (maps to deleting a NATS consumer)
func (s *JetStreamSubscriptionStore) Remove(id string) error {
	err := s.js.DeleteConsumer(s.stream, id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Example usage
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// Instantiate the subscription store
	store, err := NewJetStreamSubscriptionStore(nc, "SUBSCRIPTIONS", "subscriptions.*")
	if err != nil {
		panic(err)
	}

	// Add a subscription
	sub := &subscription.Subscription{
		ID: "sub1",
		//Subject: "subscriptions.example",
	}
	err = store.Set(sub)
	if err != nil {
		fmt.Println("Error setting subscription:", err)
	}

	// List subscriptions
	subs, err := store.List()
	if err != nil {
		fmt.Println("Error listing subscriptions:", err)
	}
	for _, s := range subs {
		fmt.Printf("Subscription ID: %ss\n", s.ID)
	}

	// Get a specific subscription
	subDetails, err := store.Get("sub1")
	if err != nil {
		fmt.Println("Error getting subscription:", err)
	} else {
		fmt.Printf("Fetched Subscription: %+v\n", subDetails)
	}

	// Remove a subscription
	err = store.Remove("sub1")
	if err != nil {
		fmt.Println("Error removing subscription:", err)
	}
}

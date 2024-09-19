package endpoints

import (
	"context"
	"log"

	"github.com/magicbutton/magic-mix/drivers"
	"github.com/magicbutton/magic-mix/subscription"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"github.com/swaggest/usecase"
)

func Workflow() usecase.Interactor {
	type Request struct {
		Id string `json:"id" description:"Durable subscription ID"`
	}
	//TODO: Implement this function
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *[]subscription.Message) error {
		log.Fatal("Not implemented")
		natsServer := viper.GetString("NATS")
		if natsServer == "" {
			natsServer = "nats://127.0.0.1:4222"
		}
		log.Println("Connecting to", natsServer)
		nc, err := nats.Connect(natsServer)
		if err != nil {
			return err
		}
		defer nc.Close()

		stream, err := drivers.NewJetStreamSubscriptionStore(nc, "workflow_events", "workflow.events.*")

		if err != nil {
			return err
		}
		messages, err := stream.ReadMessages(input.Id)
		if err != nil {
			return err
		}

		newMessages := []subscription.Message{}
		for _, m := range messages {
			newMessages = append(newMessages, subscription.Message{
				Subject: m.Subject,
				Data:    m.Data,
			})
		}
		*output = append(*output, newMessages...)

		return nil

	})
	u.SetTitle("Get events")
	u.SetDescription(`This function is used to get events from the subscriptions. All events are stored in max 10 minutes. Use this function for getting events happend in the last 10 minutes, repeat calling it with the same ID to get new as they arrive. The function will return immidiately if there are new events, or after 30 seconds with an empty array if there are no <b>new events.</b> 
`)
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Subscription")
	return u
}

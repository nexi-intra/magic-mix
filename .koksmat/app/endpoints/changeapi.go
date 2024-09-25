package endpoints

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/magicbutton/magic-mix/drivers"
	"github.com/magicbutton/magic-mix/events"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"github.com/swaggest/usecase"
)

func GetChanges() usecase.Interactor {
	type Request struct {
		Id string `json:"id" description:"Durable subscription ID"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *[]events.Event) error {

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

		stream, err := drivers.NewJetStreamSubscriptionStore(nc, "database", "database.*")

		if err != nil {
			return err
		}
		messages, err := stream.ReadMessages(input.Id)
		if err != nil {
			return err
		}

		newMessages := []events.Event{}
		for _, m := range messages {
			var b64 string
			json.Unmarshal([]byte(m.Data), &b64)

			data, err := base64.StdEncoding.DecodeString(b64)
			if err != nil {
				log.Print("error when base64.StdEncoding.DecodeString:", err)
			}

			fmt.Printf("%q\n", data)
			newMessages = append(newMessages, events.Event{
				Subject: m.Subject,
				Data:    data,
			})
		}
		*output = append(*output, newMessages...)

		return nil

	})
	u.SetTitle("Get database changes")
	u.SetDescription(`This function is used to get events from the subscriptions. All events are stored in max 10 minutes. Use this function for getting events happend in the last 10 minutes, repeat calling it with the same ID to get new as they arrive. The function will return immidiately if there are new events, or after 30 seconds with an empty array if there are no <b>new events.</b> 
`)
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Subscription")
	return u
}

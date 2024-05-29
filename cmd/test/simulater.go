package test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/ttagiyeva/superstream/internal/config"
	"github.com/ttagiyeva/superstream/internal/model"
)

// SimulateEvents sends test events to the subscriber topic every 5 seconds
func SimulateEvents(publisher message.Publisher, cfg *config.Config) {
	i := 0
	for {
		e := model.Event{
			ID:        i,
			Timestamp: time.Now(),
			Data:      fmt.Sprintf("test_data_%d", i),
		}

		payload, err := json.Marshal(e)
		if err != nil {
			panic(err)
		}

		err = publisher.Publish(
			cfg.Queue.SubscriberTopic,
			message.NewMessage(watermill.NewUUID(), payload),
		)
		if err != nil {
			panic(err)
		}

		i++

		time.Sleep(5 * time.Second)
	}
}

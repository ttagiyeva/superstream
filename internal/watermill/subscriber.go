package watermill

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"

	"github.com/ttagiyeva/superstream/internal/config"
)

func NewSubscriber(log *zap.Logger, cfg *config.Config, logger watermill.LoggerAdapter, mr kafka.MarshalerUnmarshaler) message.Subscriber {
	kafkaSubscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:       cfg.Kafka.Brokers,
			Unmarshaler:   mr,
			ConsumerGroup: cfg.Queue.SubscriberName,
		},
		logger,
	)
	if err != nil {
		log.Error("failed to create subscriber", zap.Error(err))

		panic(err)
	}

	return kafkaSubscriber
}

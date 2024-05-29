package watermill

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"

	"github.com/ttagiyeva/superstream/internal/config"
)

func NewPublisher(log *zap.Logger, cfg *config.Config, logger watermill.LoggerAdapter, mr kafka.MarshalerUnmarshaler) message.Publisher {
	p, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   cfg.Kafka.Brokers,
			Marshaler: mr,
		},
		logger,
	)
	if err != nil {
		log.Error("failed to create publisher", zap.Error(err))

		panic(err)
	}

	return p
}

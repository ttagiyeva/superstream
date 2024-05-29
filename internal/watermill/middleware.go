package watermill

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"go.uber.org/zap"

	"github.com/ttagiyeva/superstream/internal/config"
)

func NewPoisonQueueMiddleware(log *zap.Logger, publisher message.Publisher, cfg *config.Config) message.HandlerMiddleware {
	pqm, err := middleware.PoisonQueue(publisher, cfg.Queue.PoisonTopic)
	if err != nil {
		log.Error("failed to create poison queue middleware", zap.Error(err))

		panic(err)
	}

	return pqm
}

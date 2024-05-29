package watermill

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/ttagiyeva/superstream/internal/config"
	"github.com/ttagiyeva/superstream/internal/superstream/handler"
)

func RegisterHandlers(
	cfg *config.Config,
	router *message.Router,
	publisher message.Publisher,
	subscriber message.Subscriber,
	poison message.HandlerMiddleware,
	logger watermill.LoggerAdapter,
	mr kafka.MarshalerUnmarshaler,

	handler *handler.Handler,
) {
	router.AddHandler(
		"superstream_test_handler",
		cfg.Queue.SubscriberTopic,
		subscriber,
		cfg.Queue.PublishTopic,
		publisher,
		func(msg *message.Message) ([]*message.Message, error) {
			return handler.Handle(msg)
		},
	).AddMiddleware(poison)
}

func NewRouter(log *zap.Logger, lc fx.Lifecycle, logger watermill.LoggerAdapter) *message.Router {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		log.Error("failed to create watermill router", zap.Error(err))

		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)

	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 1000,
			Logger:          logger,
		}.Middleware,
		middleware.Recoverer,
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errCh := make(chan error)
			go func() {
				errCh <- router.Run(context.Background())
			}()

			select {
			case <-router.Running():
				return nil
			case err := <-errCh:

				return err
			case <-ctx.Done():

				return ctx.Err()
			}
		},
		OnStop: func(ctx context.Context) error {
			err := router.Close()

			return err
		},
	})

	return router
}

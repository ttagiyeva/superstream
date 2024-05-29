package main

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"

	"github.com/ttagiyeva/superstream/cmd/test"
	"github.com/ttagiyeva/superstream/internal/config"
	"github.com/ttagiyeva/superstream/internal/log"
	"github.com/ttagiyeva/superstream/internal/superstream/handler"
	"github.com/ttagiyeva/superstream/internal/watermill"
)

// main is the entry point of the application
// fx is used for dependency injection and lifecycle management
func main() {
	fx.New(
		fx.Provide(
			config.New,
			log.NewZapLogger,
			handler.NewHandler,
			watermill.NewSubscriber,
			watermill.NewMarshaler,
			watermill.NewLogger,
			watermill.NewPublisher,
			watermill.NewPoisonQueueMiddleware,
			watermill.NewRouter,
		),
		fx.Invoke(
			watermill.RegisterHandlers,
		),
		fx.Invoke(
			func(p message.Publisher, cfg *config.Config) {
				go test.SimulateEvents(p, cfg)
			},
		),
	).Run()
}

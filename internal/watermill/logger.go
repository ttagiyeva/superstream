package watermill

import (
	"github.com/ThreeDotsLabs/watermill"
)

func NewLogger() watermill.LoggerAdapter {
	watermillLogger := watermill.NewStdLogger(false, false)

	return watermillLogger
}

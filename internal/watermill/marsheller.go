package watermill

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

func NewMarshaler() kafka.MarshalerUnmarshaler {
	return kafka.DefaultMarshaler{}
}

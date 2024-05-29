package handler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"

	"github.com/ttagiyeva/superstream/internal/model"
)

type Handler struct {
	log *zap.Logger
}

func NewHandler(l *zap.Logger) *Handler {
	return &Handler{
		log: l,
	}
}

// Handle processes the incoming message to the subscriber topic
// and returns a new message to be published to the publish topic
func (h *Handler) Handle(msg *message.Message) ([]*message.Message, error) {
	event := model.Event{}

	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		h.log.Error("failed to unmarshal message payload", zap.Error(err))

		return nil, err
	}

	if !validateEvent(event) {
		msg.Ack()
		h.log.Info("event is too old", zap.Int("event_id", event.ID))

		return nil, nil
	}

	newPayload, err := json.Marshal(newProcessedEvent(event))
	if err != nil {
		h.log.Error("failed to marshal processed event", zap.Error(err))

		return nil, err
	}

	newMessage := message.NewMessage(msg.UUID, newPayload)

	msg.Ack()

	fmt.Println("received ", msg.UUID)

	return []*message.Message{newMessage}, nil
}

func validateEvent(event model.Event) bool {
	now := time.Now()
	duration := now.Sub(event.Timestamp)

	return duration <= 24*time.Hour
}

func newProcessedEvent(event model.Event) model.ProcessedEvent {
	processedEvent := model.ProcessedEvent{
		ID:        event.ID,
		Timestamp: event.Timestamp,
		Data:      event.Data,
		Status:    model.Valid,
	}

	if len(event.Data) < 10 {
		processedEvent.Status = model.Invalid
	}

	return processedEvent
}

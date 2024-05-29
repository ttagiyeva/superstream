package handler

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	h := &Handler{}

	now := time.Now()

	testCases := []struct {
		name          string
		request       *message.Message
		response      *message.Message
		expectedError error
		expectAck     bool
	}{
		{
			name: "Valid event",
			request: &message.Message{
				UUID:    uuid.New().String(),
				Payload: []byte(fmt.Sprintf(`{"id":1,"timestamp":"%s","data":"data"}`, now.Add(-23*time.Hour).Format(time.RFC3339))),
			},
			response: &message.Message{
				UUID:    uuid.New().String(),
				Payload: []byte(fmt.Sprintf(`{"id":1,"timestamp":"%s","data":"data","status":"invalid"}`, now.Add(-23*time.Hour).Format(time.RFC3339))),
			},
			expectAck: true,
		},
		{
			name: "Invalid event",
			request: &message.Message{
				UUID:    uuid.New().String(),
				Payload: []byte(`{invalid }`),
			},
			expectedError: errors.New("invalid character 'i' looking for beginning of object key string"),
		},
		{
			name: "older than 24 hour event",
			request: &message.Message{
				UUID:    uuid.New().String(),
				Payload: []byte(fmt.Sprintf(`{"id":1,"timestamp":"%s","data":"data"}`, now.Add(-25*time.Hour).Format(time.RFC3339))),
			},
			expectAck: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := h.Handle(tc.request)

			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)

				if tc.response == nil {
					assert.Nil(t, result)
				} else {
					assert.JSONEq(t, string(result[0].Payload), string(tc.response.Payload))
				}
			}

			if tc.expectAck {
				assert.NotNil(t, tc.request.Acked())
			} else {
				assert.Nil(t, tc.request.Acked())
			}
		})
	}
}

package api

import "insider-assessment/pkg/domain/message"

// MessageResponse represents a single message in the system
// @Description Single message response structure
type MessageResponse struct {
	// The unique identifier of the message, big serial in database
	ID uint64 `json:"id" example:"1"`
	// The recipient's phone number
	PhoneNumber string `json:"recipient_phone_number" example:"+1234567890"`
	// The content of the message
	Content string `json:"content" example:"Hello, this is a test message"`
	// The current status of the message (pending, sent, failed)
	Status string `json:"status" example:"sent"`
}

// MessageListResponse represents a collection of messages
// @Description List of messages response structure
type MessageListResponse struct {
	// Array of messages
	Messages []MessageResponse `json:"messages"`
}

func newMessageResponse(m *message.Message) MessageResponse {
	return MessageResponse{
		ID:          m.ID(),
		PhoneNumber: m.RecipientPhoneNumber(),
		Content:     m.Content(),
		Status:      m.Status().String(),
	}
}

func newMessageListResponse(messages []*message.Message) MessageListResponse {
	messageResponses := make([]MessageResponse, 0, len(messages))
	for _, p := range messages {
		messageResponses = append(messageResponses, newMessageResponse(p))
	}

	return MessageListResponse{Messages: messageResponses}
}

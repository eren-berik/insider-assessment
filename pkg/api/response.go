package api

import "insider-assesment/pkg/domain/message"

type MessageResponse struct {
	ID          uint64 `json:"id"`
	PhoneNumber string `json:"recipient_phone_number"`
	Content     string `json:"content"`
	Status      string `json:"status"`
}

type MessageListResponse struct {
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

package message

import "insider-assessment/pkg/domain/message"

type Message struct {
	ID                   uint64
	Content              string
	RecipientPhoneNumber string
	Status               message.Status
}

func newMessage(m *Message) *message.Message {
	return message.NewMessage(
		m.ID,
		m.Content,
		m.RecipientPhoneNumber,
		m.Status,
	)
}

func newMessageList(p []*Message) []*message.Message {
	if len(p) == 0 {
		return []*message.Message{}
	}
	messageList := make([]*message.Message, len(p))
	for i, v := range p {
		messageList[i] = newMessage(v)
	}
	return messageList
}

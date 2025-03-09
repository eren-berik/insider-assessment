package message

import (
	"bytes"
	"encoding/gob"
	"insider-assesment/pkg/domain/message"
	"time"
)

type (
	Serializer        struct{}
	SerializedMessage struct {
		Id       string    `json:"id"`
		SentTime time.Time `json:"sent_time"`
	}
)

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s Serializer) Serialize(m message.Cache) ([]byte, error) {
	msg := NewSerializedMessage(m)
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(msg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s Serializer) Deserialize(data []byte) (message.Cache, error) {
	var m SerializedMessage
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&m); err != nil {
		return message.Cache{}, err
	}
	return Message(m), nil
}

func NewSerializedMessage(m message.Cache) SerializedMessage {
	return SerializedMessage{
		m.ID(),
		m.SentTime().UTC().Truncate(time.Microsecond),
	}
}

func Message(m SerializedMessage) message.Cache {
	return message.NewMessageCache(
		m.Id,
		m.SentTime.UTC().Truncate(time.Microsecond),
	)
}

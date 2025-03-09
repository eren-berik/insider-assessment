package app

import (
	"bytes"
	"encoding/json"
	"insider-assessment/pkg/domain/message"
	"io"
	"log"
	"net/http"
)

type MessageSender struct{}

func NewMessageSender() *MessageSender {
	return &MessageSender{}
}

type result struct {
	Message   string `json:"message"`
	MessageId string `json:"messageId"`
}

func (s *MessageSender) SendMessage(msg message.Message) (bool, result) {
	payload, _ := json.Marshal(map[string]string{
		"to":      msg.RecipientPhoneNumber(),
		"content": msg.Content(),
	})

	req, _ := http.NewRequest("POST", "https://webhook.site/d34480d9-6792-4f47-bab0-ab6c4629f039", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ins-auth-key", "INS.me1x9uMcyYGlhKKQVPoc.bO3j9aZwRTOcA2Ywo")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Failed to send message:", err)
		return false, result{}
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var res result

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return false, result{}
	}

	return resp.StatusCode == http.StatusAccepted, res
}

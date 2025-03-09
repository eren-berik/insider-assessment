package message

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"insider-assesment/pkg/domain/message"
)

type Service struct {
	client *pgxpool.Pool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{client: pool}
}

func (s Service) AllMessages(ctx context.Context) ([]*message.Message, error) {
	var messages []*Message

	query := `SELECT id, phone_number, content, status FROM "messages"`
	rows, err := s.client.Query(ctx, query)
	if err != nil {
		return newMessageList(messages), err
	}
	for rows.Next() {
		m := &Message{}
		err = rows.Scan(
			&m.ID,
			&m.RecipientPhoneNumber,
			&m.Content,
			&m.Status,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return newMessageList(messages), nil
}

func (s Service) GetMessagesByStatus(ctx context.Context, status message.Status) ([]*message.Message, error) {
	var messages []*Message

	query := `SELECT id, phone_number, content, status FROM "messages" WHERE status = $1`
	rows, err := s.client.Query(ctx, query, status)
	if err != nil {
		return newMessageList(messages), err
	}
	for rows.Next() {
		m := &Message{}
		err = rows.Scan(
			&m.ID,
			&m.RecipientPhoneNumber,
			&m.Content,
			&m.Status,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return newMessageList(messages), nil
}

func (s Service) UpdateMessageStatus(ctx context.Context, id uint64, status message.Status) error {
	query := `UPDATE messages SET status = $1 WHERE id = $2`
	_, err := s.client.Exec(ctx, query, status, id)
	return err
}

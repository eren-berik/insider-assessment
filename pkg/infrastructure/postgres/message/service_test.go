package message

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"insider-assessment/pkg/domain/message"
	"insider-assessment/pkg/infrastructure/postgres"
	"testing"
)

func TestService_AllMessages(t *testing.T) {
	conn := postgres.NewPGPool("postgres://postgres:postgres@localhost:5432/insider-test?sslmode=disable")
	defer conn.Close()
	service := Service{client: conn}

	seedTestMessages(conn, t)

	messages, err := service.AllMessages(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, messages)
	assert.Len(t, messages, 2)

	assert.Equal(t, uint64(100), messages[0].ID())
	assert.Equal(t, "1234567890", messages[0].RecipientPhoneNumber())
	assert.Equal(t, "Test content 1", messages[0].Content())
	assert.Equal(t, uint8(1), messages[0].Status())

	assert.Equal(t, uint64(101), messages[1].ID())
	assert.Equal(t, "9876543210", messages[1].RecipientPhoneNumber())
	assert.Equal(t, "Test content 2", messages[1].Content())
	assert.Equal(t, uint8(0), messages[1].Status())
}

func TestService_GetMessagesByStatus(t *testing.T) {
	conn := postgres.NewPGPool("postgres://postgres:postgres@localhost:5432/insider?sslmode=disable")
	defer conn.Close()
	service := Service{client: conn}

	seedTestMessages(conn, t)

	sentMessages, err := service.GetMessagesByStatus(context.Background(), message.Sent, 2)
	assert.NoError(t, err)
	assert.NotNil(t, sentMessages)
	assert.Len(t, sentMessages, 1)

	assert.Equal(t, uint64(100), sentMessages[0].ID())
	assert.Equal(t, "1234567890", sentMessages[0].RecipientPhoneNumber())
	assert.Equal(t, "Test content 1", sentMessages[0].Content())
	assert.Equal(t, "Sent", sentMessages[0].Status().String())

	pendingMessages, err := service.GetMessagesByStatus(context.Background(), message.Pending, 2)
	assert.NoError(t, err)
	assert.NotNil(t, pendingMessages)
	assert.Len(t, pendingMessages, 1)

	assert.Equal(t, uint64(101), pendingMessages[0].ID())
	assert.Equal(t, "9876543210", pendingMessages[0].RecipientPhoneNumber())
	assert.Equal(t, "Test content 2", pendingMessages[0].Content())
	assert.Equal(t, "Pending", pendingMessages[0].Status().String())
}

func TestService_UpdateMessageStatus(t *testing.T) {
	conn := postgres.NewPGPool("postgres://postgres:postgres@localhost:5432/insider?sslmode=disable")
	defer conn.Close()

	service := Service{client: conn}

	seedTestMessages(conn, t)

	var updatedMessage Message

	err := service.UpdateMessageStatus(context.Background(), 101, message.Sent)
	assert.NoError(t, err)
	err = conn.QueryRow(context.Background(), `SELECT id, phone_number, content, status FROM messages WHERE id = $1`, 101).
		Scan(&updatedMessage.ID, &updatedMessage.RecipientPhoneNumber, &updatedMessage.Content, &updatedMessage.Status)
	assert.NoError(t, err)
	assert.Equal(t, message.Sent.String(), updatedMessage.Status.String())

	err = service.UpdateMessageStatus(context.Background(), 100, message.Pending)
	assert.NoError(t, err)
	err = conn.QueryRow(context.Background(), `SELECT id, phone_number, content, status FROM messages WHERE id = $1`, 100).
		Scan(&updatedMessage.ID, &updatedMessage.RecipientPhoneNumber, &updatedMessage.Content, &updatedMessage.Status)
	assert.NoError(t, err)
	assert.Equal(t, message.Pending.String(), updatedMessage.Status.String())
}

func seedTestMessages(conn *pgxpool.Pool, t *testing.T) {
	deleteQuery := `DELETE FROM messages;`
	_, err := conn.Exec(context.Background(), deleteQuery)
	if err != nil {
		t.Fatalf("Error cleaning up the database: %v", err)
	}

	_, err = conn.Exec(context.Background(), `
		INSERT INTO "messages" (id, phone_number, content, status)
		VALUES
			(100, '1234567890', 'Test content 1', 1),
			(101, '9876543210', 'Test content 2', 0)
	`)
	if err != nil {
		t.Fatalf("Error inserting test data: %v", err)
	}
}

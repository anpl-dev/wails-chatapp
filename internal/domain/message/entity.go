package message

import (
	"context"
	"strings"
	"wails-chatapp/internal/domain"

	"github.com/google/uuid"
)

type (
	MessageRepository interface {
		Save(ctx context.Context, msg *Message) error
		FindByID(ctx context.Context, id uuid.UUID) (*Message, error)
	}

	Message struct {
		ID       uuid.UUID
		Contents string
	}
)

func NewMessage(cont string) (*Message, error) {
	if strings.TrimSpace(cont) == "" {
		return nil, domain.ErrInvalidMessage
	}
	return &Message{
		ID:       uuid.New(),
		Contents: cont,
	}, nil
}

package user

import (
	"context"

	"github.com/google/uuid"
)

type (
	UserRepository interface {
		JoinRoom(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error
		LeaveRoom(ctx context.Context, id uuid.UUID) error
		SendMessage(ctx context.Context, roomID uuid.UUID) error
	}

	User struct {
		ID   uuid.UUID
		Name string
	}
)

func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, nil
	}
	return &User{
		ID:   uuid.New(),
		Name: name,
	}, nil
}

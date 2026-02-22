package room

import (
	"context"

	"github.com/google/uuid"
)

type (
	RoomRepository interface {
		Create(ctx context.Context, room Room) (*Room, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	Room struct {
		ID   uuid.UUID
		Name string
	}
)

func NewRoom(name string) (*Room, error) {
	if name == "" {
		return nil, nil
	}
	return &Room{
		ID:   uuid.New(),
		Name: name,
	}, nil
}

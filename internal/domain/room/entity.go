package room

import (
	"context"
	"strings"
	"wails-chatapp/internal/domain"

	"github.com/google/uuid"
)

type (
	RoomRepository interface {
		Save(ctx context.Context, room *Room) error
		FindByID(ctx context.Context, id uuid.UUID) (*Room, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	Room struct {
		ID   uuid.UUID
		Name string
	}
)

func NewRoom(name string) (*Room, error) {
	if strings.TrimSpace(name) == "" {
		return nil, domain.ErrInvalidRoomName
	}
	return &Room{
		ID:   uuid.New(),
		Name: name,
	}, nil
}

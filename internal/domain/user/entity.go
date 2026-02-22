package user

import (
	"context"
	"strings"
	"wails-chatapp/internal/domain"

	"github.com/google/uuid"
)

type (
	UserRepository interface {
		Save(ctx context.Context, user *User) error
		FindByID(ctx context.Context, id uuid.UUID) (*User, error)
		Delete(ctx context.Context, id uuid.UUID) error
	}

	User struct {
		ID   uuid.UUID
		Name string
	}
)

func NewUser(name string) (*User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, domain.ErrInvalidUserName
	}
	return &User{
		ID:   uuid.New(),
		Name: name,
	}, nil
}

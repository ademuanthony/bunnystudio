package users

import (
	"context"

	"bonnystudio.com/taskmanager/internal"
)

type Service interface {
	Get(ctx context.Context, userID string) (*internal.User, error)
	GetAll(ctx context.Context) ([]internal.User, error)
	Create(ctx context.Context, name string) (*internal.User, error)
	Update(ctx context.Context, user internal.User) error
	Delete(ctx context.Context, userID string) error
}

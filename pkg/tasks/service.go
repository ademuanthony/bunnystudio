package tasks

import (
	"context"

	"bonnystudio.com/taskmanager/internal"
)

type Service interface {
	GetByUserID(ctx context.Context, userID string) ([]internal.Task, error)
	Create(ctx context.Context, task internal.Task) (*internal.Task, error)
	Update(ctx context.Context, task internal.Task) error
}

type store interface {
	Create(user internal.Task) error
	Update(user internal.Task) error
	GetByUserID(userID string) ([]internal.Task, error)
}

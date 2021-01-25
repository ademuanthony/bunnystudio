package tasks

import (
	"context"
	"errors"
	"os"

	"bonnystudio.com/taskmanager/internal"
	"github.com/go-kit/kit/log"
	"github.com/pborman/uuid"
)

type taskService struct {
	store store
}

func NewService(store store) Service {
	return &taskService{store: store}
}

func (s taskService) GetByUserID(_ context.Context, userID string) ([]internal.Task, error) {
	return s.store.GetByUserID(userID)
}

func (s taskService) Create(ctx context.Context, task internal.Task) (*internal.Task, error) {
	if task.Description == "" {
		return nil, errors.New("Cannot create a task with an empty description")
	}
	task.ID = uuid.New()
	if err := s.store.Create(task); err != nil {
		logger.Log("Error in store.Create", err)
		return nil, err
	}
	return &task, nil
}

func (s taskService) Update(_ context.Context, task internal.Task) error {
	if err := s.store.Update(task); err != nil {
		logger.Log("Error in store.Create", err)
		return err
	}
	return nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

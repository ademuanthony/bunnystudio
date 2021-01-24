package users

import (
	"context"
	"os"

	"bonnystudio.com/taskmanager/internal"
	"github.com/go-kit/kit/log"
	"github.com/pborman/uuid"
)

type userService struct {
	store store
}

func NewService(store store) Service {
	return &userService{store: store}
}

func (s userService) Get(_ context.Context, userID string) (*internal.User, error) {
	return s.store.FindByID(userID)
}

func (s userService) GetAll(_ context.Context) ([]internal.User, error) {
	return s.store.GetAll()
}

func (s userService) Create(ctx context.Context, name string) (*internal.User, error) {
	user := internal.User{
		ID:   uuid.New(),
		Name: name,
	}
	if err := s.store.Create(user); err != nil {
		logger.Log("Error in store.Create", err)
		return nil, err
	}
	return &user, nil
}

func (s userService) Update(_ context.Context, user internal.User) error {
	if err := s.store.Update(user); err != nil {
		logger.Log("Error in store.Create", err)
		return err
	}
	return nil
}

func (s userService) Delete(_ context.Context, userID string) error {
	return s.store.Delete(userID)
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

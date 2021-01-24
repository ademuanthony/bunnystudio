package users

import (
	"bonnystudio.com/taskmanager/internal"
	"bonnystudio.com/taskmanager/internal/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const tableName = "Users"

type dynamoStore struct {
	table dynamo.Table
}

func NewDynamoStore() *dynamoStore {
	// endpoint := "http://localhost:8000"
	db := dynamo.New(session.New(), &aws.Config{
		Region:   aws.String("us-west-1"),
		// Endpoint: &endpoint,
	})

	if err := db.CreateTable(tableName, internal.User{}).Run(); err != nil {
		logger.Log("CreateTable", err)
	}

	return &dynamoStore{table: db.Table(tableName)}
}

func (s *dynamoStore) Create(user internal.User) error {
	return s.table.Put(user).Run()
}

func (s *dynamoStore) Update(user internal.User) error {
	err := s.table.Update(user.ID, user).Run()
	if err == dynamo.ErrNotFound {
		err = util.ErrUnknown
	}
	return err
}

func (s *dynamoStore) FindByID(userID string) (*internal.User, error) {
	var user internal.User
	err := s.table.Get("ID", userID).One(&user)
	if err != nil {
		if err == dynamo.ErrNotFound {
			err = util.ErrUnknown
		}
		return nil, err
	}
	return &user, nil
}

func (s *dynamoStore) GetAll() ([]internal.User, error) {
	var users []internal.User
	err := s.table.Scan().All(&users)

	return users, err
}

func (s *dynamoStore) Delete(userID string) error {
	return s.table.Delete("ID", userID).Run()
}

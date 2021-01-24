package tasks

import (
	"bonnystudio.com/taskmanager/internal"
	"bonnystudio.com/taskmanager/internal/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const tableName = "Tasks"

type dynamoStore struct {
	table dynamo.Table
}

func NewDynamoStore() *dynamoStore {
	// endpoint := "http://localhost:8000"
	db := dynamo.New(session.New(), &aws.Config{
		Region:   aws.String("us-west-1"),
		// Endpoint: &endpoint,
	})

	if err := db.CreateTable(tableName, internal.Task{}).Run(); err != nil {
		logger.Log("CreateTable", err)
	}

	return &dynamoStore{table: db.Table(tableName)}
}

func (s *dynamoStore) Create(task internal.Task) error {
	return s.table.Put(task).Run()
}

func (s *dynamoStore) Update(task internal.Task) error {
	err := s.table.Update(task.ID, task).Run()
	if err == dynamo.ErrNotFound {
		err = util.ErrUnknown
	}
	return err
}

func (s *dynamoStore) GetByUserID(userID string) ([]internal.Task, error) {
	var tasks []internal.Task
	err := s.table.Scan().Filter("'UserID' = ?", userID).All(&tasks)

	return tasks, err
}

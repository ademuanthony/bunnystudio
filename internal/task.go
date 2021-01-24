package internal

type Task struct {
	ID          string    `json:"id" dynamo:"ID,hash"`
	Description string    `json:"description"`
	State       TaskState `json:"state"`
	UserID      string    `json:"user_id"`
}

type TaskState string

var (
	Todo TaskState = "to do"
	Done TaskState = "done"
)

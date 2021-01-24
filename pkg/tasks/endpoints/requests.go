package endpoints

import "bonnystudio.com/taskmanager/internal"

// GetByUserID
type GetByUserIDRequest struct {
	UserID string `json:"user_id"`
}

type GetByUserIDResponse struct {
	Tasks []internal.Task `json:"tasks"`
	Err   string          `json:"err,omitempty"`
}

// Create
type CreateRequest struct {
	Description string             `json:"description"`
	State       internal.TaskState `json:"state"`
	UserID      string             `json:"user_id"`
}

type CreateResponse struct {
	Task *internal.Task `json:"task"`
	Err  string         `json:"err,omitempty"`
}

// Update
type UpdateRequest struct {
	Task internal.Task `json:"task"`
}

type UpdateResponse struct {
	Err string `json:"err,omitempty"`
}

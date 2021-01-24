package endpoints

import "bonnystudio.com/taskmanager/internal"

// Get
type GetRequest struct {
	UserID string `json:"user_id,omitempty"`
}

type GetResponse struct {
	User *internal.User `json:"user"`
	Err  string         `json:"err,omitempty"`
}

// GetAll
type GetAllRequest struct {
}

type GetAllResponse struct {
	Users []internal.User `json:"users"`
	Err   string          `json:"err,omitempty"`
}

// Create
type CreateRequest struct {
	Name string `json:"name"`
}

type CreateResponse struct {
	User *internal.User `json:"user"`
	Err  string        `json:"err,omitempty"`
}

// Update
type UpdateRequest struct {
	User internal.User `json:"user"`
}

type UpdateResponse struct {
	Err string `json:"err,omitempty"`
}

// Delete
type DeleteRequest struct {
	UserID string `json:"user_id,omitempty"`
}

type DeleteResponse struct {
	Err string `json:"err,omitempty"`
}

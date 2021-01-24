package internal

// User represent a single user of the application
type User struct {
	ID   string `json:"id" dynamo:"ID,hash"`
	Name string `json:"name"`
}

package internal

// User represent a single user of the application
type User struct {
	ID   string `dynamo:"ID,hash"`
	Name string
}

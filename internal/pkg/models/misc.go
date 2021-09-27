package models

// Credentials Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type GenericResponse struct {
	ID string `json:"user_id"`
}


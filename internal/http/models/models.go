package models

// for now lets keep this simple
type Student struct {
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required" json:"email"`
}

package pgdb

import "time"

// Model is the base model with uuid field
type Model struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// Application is the connection string for database
type Application struct {
	Model
	AppName   string
	IsBlocked bool
}

package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

// Add a new ErrInvalidCredentials error. We'll use this later if a user
// tries to login with an incorrect email address or password.
var ErrInvalidCredentials = errors.New("models: invalid credentials")

// Add a new ErrDuplicateEmail error. We'll use this later if a user
// tries to signup with an email address that's already in use.
var ErrDuplicateEmail = errors.New("models: duplicate email")

type Asset struct {
	Id       string
	Name     string
	Value    float64
	Currency string
	Custody  string
	Created  time.Time
	Finished time.Time
	Active   bool
}

type Balance struct {
	Value float64
}

// Define a new User type. Notice how the field names and types align
// with the columns in the database `users` table?
type User struct {
	Id             string
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

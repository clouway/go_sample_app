package domain

import (
	"time"
)

type SessionStore interface {
	StartSession(user User, time time.Time) (*Session, error)

	FindSessionAvailableAt(ID string, instant time.Time) (*Session, bool, error)
}

type UserStore interface {
	Register(request UserRegistrationRequest) (User, error)

	FindUserByEmail(email string) (User, error)
}

type UserRegistrationRequest struct {
}

type User struct {
}

type Session struct {
	ID        string
	ExpiresAt time.Time
	UserID    string
}

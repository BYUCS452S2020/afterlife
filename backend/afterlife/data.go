package afterlife

import (
	"context"
	"time"
)

type User struct {
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	VerifiedAlive time.Time `json:"verifiedAlive"`
}

type EventType string

const (
	EventTypeEmail EventType = "email"
)

type Event struct {
	Name string    `json:"name"`
	At   time.Time `json:"at"`
	Type EventType `json:"type"`

	Email *EventEmail `json:"email"`
}

type EventEmail struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

type DataService interface {
	Register(context.Context, RegisterRequest) (string, error)
	Login(ctx context.Context, username, password string) (string, error)
	User(context.Context, string) (User, error)
	Timeline(context.Context, string) (Timeline, error)
}

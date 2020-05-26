package afterlife

import (
	"context"
	"time"
)

type UserID string

type User struct {
	ID            UserID    `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	VerifiedAlive time.Time `json:"verifiedAlive"`
}

type EventType string

const (
	EventTypeEmail EventType = "email"
)

type EventID string

type Event struct {
	ID   EventID   `json:"id"`
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
	Register(context.Context, RegisterRequest) error
	Login(context.Context, LoginRequest) (string, error)
	User(context.Context, string) (User, error)

	CreateEvent(context.Context, UserID, Event) error
	DeleteEvent(context.Context, UserID, EventID) error
	UpdateEvent(context.Context, UserID, Event) error
	Timeline(context.Context, UserID) (Timeline, error)
}

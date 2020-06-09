package mongo

import (
	"time"

	"github.com/byuoitav/afterlife"
)

type user struct {
	ID            string    `bson:"_id,omitempty"`
	Email         string    `bson:"email"`
	Password      string    `bson:"password"`
	FirstName     string    `bson:"firstName"`
	LastName      string    `bson:"lastName"`
	VerifiedAlive time.Time `bson:"verifiedAlive"`

	Tokens []token `bson:"tokens"`
	Events []event `bson:"events"`
}

type token struct {
	Token   string    `bson:"token"`
	Created time.Time `bson:"created"`
}

type event struct {
	ID   string    `bson:"id"`
	Name string    `json:"name"`
	At   time.Time `json:"at"`
	Type string    `json:"type"`

	Email *eventEmail `json:"email,omitempty"`
}

type eventEmail struct {
	To      []string `bson:"to"`
	Subject string   `bson:"subject"`
	Body    string   `bson:"body"`
}

func (u user) convert() afterlife.User {
	var user afterlife.User

	user.ID = afterlife.UserID(u.ID)
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.VerifiedAlive = u.VerifiedAlive

	return user
}

func (e event) convert() afterlife.Event {
	var event afterlife.Event

	event.ID = afterlife.EventID(e.ID)
	event.Name = e.Name
	event.At = e.At
	event.Type = afterlife.EventType(e.Type)

	if event.Email != nil {
		event.Email = &afterlife.EventEmail{
			To:      e.Email.To,
			Subject: e.Email.Subject,
			Body:    e.Email.Body,
		}
	}

	return event
}

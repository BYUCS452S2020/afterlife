package mongo

import (
	"time"

	"github.com/byuoitav/afterlife"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Email         string             `bson:"email,omitempty"`
	Password      string             `bson:"password,omitempty"`
	FirstName     string             `bson:"firstName,omitempty"`
	LastName      string             `bson:"lastName,omitempty"`
	VerifiedAlive time.Time          `bson:"verifiedAlive,omitempty"`

	Tokens []token `bson:"tokens,omitempty"`
	Events []event `bson:"events,omitempty"`
}

type token struct {
	Token   string    `bson:"token,omitempty"`
	Created time.Time `bson:"created,omitempty"`
}

type event struct {
	ID   string    `bson:"id,omitempty"`
	Name string    `json:"name,omitempty"`
	At   time.Time `json:"at,omitempty"`
	Type string    `json:"type,omitempty"`

	Email *eventEmail `json:"email,omitempty"`
}

type eventEmail struct {
	To      []string `bson:"to,omitempty"`
	Subject string   `bson:"subject,omitempty"`
	Body    string   `bson:"body,omitempty"`
}

func (u user) convert() afterlife.User {
	var user afterlife.User

	user.ID = afterlife.UserID(u.ID.Hex())
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

	if e.Email != nil {
		event.Email = &afterlife.EventEmail{
			To:      e.Email.To,
			Subject: e.Email.Subject,
			Body:    e.Email.Body,
		}
	}

	return event
}

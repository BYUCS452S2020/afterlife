package pg

import (
	"strconv"
	"time"

	"github.com/byuoitav/afterlife"
	"github.com/lib/pq"
)

type user struct {
	ID            int        `db:"id"`
	Email         string     `db:"email"`
	Password      string     `db:"password"`
	FirstName     string     `db:"first_name"`
	LastName      string     `db:"last_name"`
	CreatedOn     time.Time  `db:"created_on"`
	LastLogin     *time.Time `db:"last_login"`
	VerifiedAlive time.Time  `db:"verified_alive"`
}

type event struct {
	ID     int       `db:"id"`
	UserID int       `db:"user_id"`
	Name   string    `db:"name"`
	At     time.Time `db:"at"`
	Type   string    `db:"event_type"`

	To      *pq.StringArray `db:"email_to"`
	Subject *string         `db:"email_subject"`
	Body    *string         `db:"email_body"`
}

type eventEmail struct {
}

func (u user) convert() afterlife.User {
	var user afterlife.User

	user.ID = afterlife.UserID(strconv.Itoa(u.ID))
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.VerifiedAlive = u.VerifiedAlive

	return user
}

func (e event) convert() afterlife.Event {
	var event afterlife.Event

	event.ID = afterlife.EventID(strconv.Itoa(e.ID))
	event.Name = e.Name
	event.At = e.At
	event.Type = afterlife.EventType(e.Type)

	switch event.Type {
	case afterlife.EventTypeEmail:
		event.Email = &afterlife.EventEmail{}

		if e.To != nil {
			event.Email.To = *e.To
		}

		if e.Subject != nil {
			event.Email.Subject = *e.Subject
		}

		if e.Body != nil {
			event.Email.Body = *e.Body
		}
	}

	return event
}

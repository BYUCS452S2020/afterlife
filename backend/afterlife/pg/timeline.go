package pg

import (
	"context"
	"errors"
	"fmt"

	"github.com/byuoitav/afterlife"
	"github.com/lib/pq"
)

func (d *DataService) CreateEvent(ctx context.Context, id afterlife.UserID, event afterlife.Event) error {
	var stmt string
	var args []interface{}

	switch event.Type {
	case afterlife.EventTypeEmail:
		stmt = "INSERT INTO events (user_id, name, at, event_type, email_to, email_subject, email_body) VALUES ($1, $2, $3, $4, $5, $6, $7)"
		args = []interface{}{
			id,
			event.Name,
			event.At,
			event.Type,
			pq.Array(event.Email.To),
			event.Email.Subject,
			event.Email.Body,
		}
	default:
		return errors.New("invalid event type")
	}

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (d *DataService) UpdateEvent(ctx context.Context, id afterlife.UserID, event afterlife.Event) error {
	var stmt string
	var args []interface{}

	switch event.Type {
	case afterlife.EventTypeEmail:
		stmt = "UPDATE events SET name=$3, at=$4, event_type=$5, email_to=$6, email_subject=$7, email_body=$8 WHERE id=$1 AND user_id=$2"
		args = []interface{}{
			event.ID,
			id,
			event.Name,
			event.At,
			event.Type,
			pq.Array(event.Email.To),
			event.Email.Subject,
			event.Email.Body,
		}
	default:
		return errors.New("invalid event type")
	}

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (d *DataService) DeleteEvent(ctx context.Context, id afterlife.UserID, eventID afterlife.EventID) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt := "DELETE FROM events WHERE id=$1 AND user_id=$2"
	args := []interface{}{
		eventID,
		id,
	}

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (d *DataService) Timeline(ctx context.Context, id afterlife.UserID) (afterlife.Timeline, error) {
	var events []event

	err := d.db.SelectContext(ctx, &events, "SELECT * FROM events WHERE user_id=$1", id)
	if err != nil {
		return afterlife.Timeline{}, err
	}

	fmt.Printf("events: %+v\n", events)

	var ret afterlife.Timeline
	for _, event := range events {
		ret = append(ret, event.convert())
	}

	return ret, nil
}

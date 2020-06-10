package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/byuoitav/afterlife"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *DataService) CreateEvent(ctx context.Context, id afterlife.UserID, e afterlife.Event) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	i, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}

	ksuid, err := ksuid.NewRandom()
	if err != nil {
		return err
	}

	event := event{
		ID:   ksuid.String(),
		Name: e.Name,
		At:   e.At,
		Type: string(e.Type),
	}

	if e.Email != nil {
		event.Email = &eventEmail{
			To:      e.Email.To,
			Subject: e.Email.Subject,
			Body:    e.Email.Body,
		}
	}

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := user{
		ID: i,
	}
	update := bson.M{
		"$push": bson.M{
			"events": event,
		},
	}

	res := d.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (d *DataService) UpdateEvent(ctx context.Context, id afterlife.UserID, e afterlife.Event) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	i, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}

	event := event{
		Name: e.Name,
		At:   e.At,
		Type: string(e.Type),
	}

	if e.Email != nil {
		event.Email = &eventEmail{
			To:      e.Email.To,
			Subject: e.Email.Subject,
			Body:    e.Email.Body,
		}
	}

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := bson.M{
		"_id":       i,
		"events.id": e.ID,
	}
	update := bson.M{
		"$set": bson.M{
			"events.$": event,
		},
	}

	res := d.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (d *DataService) DeleteEvent(ctx context.Context, id afterlife.UserID, eventID afterlife.EventID) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	i, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return err
	}

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := user{
		ID: i,
	}
	update := bson.M{
		"$pull": bson.M{
			"events": event{
				ID: string(eventID),
			},
		},
	}

	res := d.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (d *DataService) Timeline(ctx context.Context, id afterlife.UserID) (afterlife.Timeline, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	i, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return afterlife.Timeline{}, err
	}

	filter := user{
		ID: i,
	}

	res := d.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return afterlife.Timeline{}, res.Err()
	}

	var user user
	if err := res.Decode(&user); err != nil {
		return afterlife.Timeline{}, fmt.Errorf("unable to parse: %w", err)
	}

	var timeline afterlife.Timeline
	for _, event := range user.Events {
		timeline = append(timeline, event.convert())
	}

	return timeline, nil
}

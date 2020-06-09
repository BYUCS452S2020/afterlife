package mongo

import (
	"context"

	"github.com/byuoitav/afterlife"
)

func (d *DataService) CreateEvent(ctx context.Context, id afterlife.UserID, event afterlife.Event) error {
	return nil
}

func (d *DataService) UpdateEvent(ctx context.Context, id afterlife.UserID, event afterlife.Event) error {
	return nil
}

func (d *DataService) DeleteEvent(ctx context.Context, id afterlife.UserID, eventID afterlife.EventID) error {
	return nil
}

func (d *DataService) Timeline(ctx context.Context, id afterlife.UserID) (afterlife.Timeline, error) {
	return afterlife.Timeline{}, nil
}

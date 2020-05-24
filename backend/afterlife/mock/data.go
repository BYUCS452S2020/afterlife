package mock

import (
	"context"
	"errors"
	"time"

	"github.com/byuoitav/afterlife"
)

type DataService struct {
}

func (*DataService) CreateUser(context.Context, afterlife.CreateUserRequest) (string, error) {
	return "", nil
}

func (*DataService) Login(ctx context.Context, username, password string) (string, error) {
	return "12345", nil
}

func (*DataService) User(ctx context.Context, token string) (afterlife.User, error) {
	if token != "12345" {
		return afterlife.User{}, errors.New("invalid token")
	}

	return afterlife.User{
		FirstName:     "Danny",
		LastName:      "Randall",
		VerifiedAlive: time.Now().Add(-2 * 24 * time.Hour),
	}, nil
}

func (*DataService) Timeline(ctx context.Context, token string) (afterlife.Timeline, error) {
	if token != "12345" {
		return afterlife.Timeline{}, errors.New("invalid token")
	}

	return afterlife.Timeline{
		afterlife.Event{
			Name: "Wife Graduation",
			At:   time.Now().Add(30 * 24 * time.Hour),
			Type: afterlife.EventTypeEmail,
			Email: &afterlife.EventEmail{
				To:      []string{"dgrandall@me.com"},
				Subject: "Congrats!",
				Body:    "Yay! You graduated!",
			},
		},
		afterlife.Event{
			Name: "Starting Kindergarten",
			At:   time.Now().Add(120 * 24 * time.Hour),
			Type: afterlife.EventTypeEmail,
			Email: &afterlife.EventEmail{
				To:      []string{"dgrandall@me.com"},
				Subject: "Good luck!",
				Body:    "Good luck at school today!\nLove,\nDanny",
			},
		},
	}, nil
}

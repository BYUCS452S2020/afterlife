package pg

import (
	"context"
	"errors"
	"time"

	"github.com/byuoitav/afterlife"
)

func (d *DataService) Register(ctx context.Context, req afterlife.RegisterRequest) (string, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return "", err
	}

	stmt := "INSERT INTO users (email, password, first_name, last_name, created_on, last_login, verified_alive) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	args := []interface{}{
		req.Email,
		req.Password,
		req.FirstName,
		req.LastName,
		time.Now(),
		time.Now(),
		time.Now(),
	}

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

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


package pg

import (
	"context"
	"time"

	"github.com/byuoitav/afterlife"
	"github.com/segmentio/ksuid"
)

func (d *DataService) Register(ctx context.Context, req afterlife.RegisterRequest) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
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
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (d *DataService) Login(ctx context.Context, req afterlife.LoginRequest) (string, error) {
	var user user

	err := d.db.GetContext(ctx, &user, "SELECT id FROM users WHERE email=$1 AND password=$2", req.Email, req.Password)
	if err != nil {
		return "", err
	}

	// TODO make sure id is set

	// generate a unique token
	ksuid, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}
	tok := ksuid.String()

	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO tokens (user_id, token, created_on) VALUES ($1, $2, $3)", user.ID, tok, time.Now())
	if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return tok, nil
}

func (d *DataService) User(ctx context.Context, token string) (afterlife.User, error) {
	var user user

	err := d.db.GetContext(ctx, &user, "SELECT users.* FROM users JOIN tokens ON tokens.user_id = users.id WHERE tokens.token=$1", token)
	if err != nil {
		return afterlife.User{}, err
	}

	return user.convert(), nil
}

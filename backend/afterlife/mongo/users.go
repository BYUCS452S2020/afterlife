package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/byuoitav/afterlife"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *DataService) Register(ctx context.Context, req afterlife.RegisterRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user := user{
		Email:         req.Email,
		Password:      req.Password,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		VerifiedAlive: time.Now(),
	}

	b, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("unable to marshal user: %w", err)
	}

	res, err := d.collection.InsertOne(ctx, b)
	if err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}

	fmt.Printf("inserted id (%T): %+v\n", res.InsertedID, res.InsertedID)
	return nil
}

func (d *DataService) Login(ctx context.Context, req afterlife.LoginRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ksuid, err := ksuid.NewRandom()
	if err != nil {
		return "", err
	}

	token := token{
		Token:   ksuid.String(),
		Created: time.Now(),
	}

	opts := options.FindOneAndUpdate().SetUpsert(false)
	filter := user{
		Email:    req.Email,
		Password: req.Password,
	}
	update := bson.M{
		"$push": bson.M{
			"tokens": token,
		},
	}

	res := d.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if res.Err() != nil {
		return "", fmt.Errorf("unable to insert token: %w", res.Err())
	}

	return token.Token, nil
}

func (d *DataService) User(ctx context.Context, token string) (afterlife.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{
		"tokens.token": token,
	}

	res := d.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return afterlife.User{}, res.Err()
	}

	var user user
	if err := res.Decode(&user); err != nil {
		return afterlife.User{}, fmt.Errorf("unable to parse: %w", err)
	}

	return user.convert(), nil
}

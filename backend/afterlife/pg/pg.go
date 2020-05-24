package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DataService struct {
	db *sqlx.DB
}

func New(username, dbName string) (*DataService, error) {
	dataSource := fmt.Sprintf("user=%s dbname=%s sslmode=disable", username, dbName)

	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	return &DataService{
		db: db,
	}, nil
}

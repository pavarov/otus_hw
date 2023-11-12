package sqlstorage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pavarov/otus_hw/hw12_13_14_15_calendar/internal/config"
)

type ClientInterface interface {
	Connection() *sqlx.DB
}

type Client struct {
	connection *sqlx.DB
}

func NewClient(cfg config.DBConfig) (ClientInterface, error) {
	conStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.User, cfg.Password, cfg.DBName)

	connection, err := sqlx.Connect(cfg.Driver, conStr)
	if err != nil {
		return nil, err
	}
	connection.SetMaxIdleConns(2)
	connection.SetMaxOpenConns(10)
	return &Client{connection: connection}, nil
}

func (c Client) Connection() *sqlx.DB {
	return c.connection
}

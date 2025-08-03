package database

import (
	"context"
	"fmt"
	"github.com/av-huette/avh-booking-system/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func NewFromConfig() (*DB, error) {
	dbConf := config.LoadDbConfigFromRootEnv()
	return New(dbConf.DbUser, dbConf.DbPassword, dbConf.DbHost, dbConf.DbPort, dbConf.DbName)
}

func New(dbUser string, dbPassword string, dbHost string, dbPort int, dbName string) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s:%d/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	dbPool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s", dsn))
	if err != nil {
		return nil, err
	}

	// test connection
	err = dbPool.Ping(context.Background())
	if err != nil {
		dbPool.Close()

		return nil, err
	}

	return &DB{dbPool}, nil
}

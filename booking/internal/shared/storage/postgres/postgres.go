package postgres

import (
	"context"
	"os"

	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Config struct {
	ConnStr string
}

func GetConfig() (*Config, error) {
	pgConn := os.Getenv("PG_STRING")
	if pgConn == "" {
		return nil, errors.New("not found PG_STRING")
	}
	return &Config{
		ConnStr: pgConn,
	}, nil
}

func InitPostgres(cfg *Config, maxRetries int) (*pgxpool.Pool, error) {
	var retryCount int

	pool, err := pgxpool.New(context.TODO(), cfg.ConnStr)
	for err != nil && retryCount < maxRetries {
		fmt.Println("retryCount", retryCount)
		time.Sleep(time.Second * 6)
		pool, err = pgxpool.New(context.TODO(), cfg.ConnStr)
		retryCount++
	}

	if err != nil {
		return nil, errors.New("Failed to connect to Postgres: " + err.Error())
	}

	retryCount = 0
	err = pool.Ping(context.TODO())

	for err != nil && retryCount < maxRetries {
		fmt.Println(err, "Try to connect")
		time.Sleep(time.Second * 5)
		err = pool.Ping(context.TODO())
		retryCount++
	}

	if err != nil {
		return nil, err
	}

	return pool, nil
}

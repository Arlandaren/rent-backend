package postgres

import (
	"errors"
	"os"
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

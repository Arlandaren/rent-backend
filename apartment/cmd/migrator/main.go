package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"service/internal/shared/storage/postgres"
)

func main() {
	var migrationsPath, migrationsTable string
	var down, rollback bool

	cfg, err := postgres.GetConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	pgString := cfg.ConnStr

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.BoolVar(&down, "down", false, "set this flag to revert all migrations")
	flag.BoolVar(&rollback, "rollback", false, "set this flag to rollback the last migration")
	flag.Parse()

	if migrationsPath == "" {
		log.Fatal("migrations-path is required")
	}

	pgStringWithTable := fmt.Sprintf("%s&x-migrations-table=%s", pgString, migrationsTable)

	m, err := migrate.New(
		"file://"+migrationsPath,
		pgStringWithTable,
	)
	if err != nil {
		log.Fatalf("failed to initialize migrate instance: %v", err)
	}

	if down {
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to revert")
				return
			}
			log.Fatalf("failed to revert all migrations: %v", err)
		}
		fmt.Println("all migrations reverted successfully")
	} else if rollback {
		if err := m.Steps(-1); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migration to rollback")
				return
			}
			log.Fatalf("failed to rollback migration: %v", err)
		}
		fmt.Println("last migration rolled back successfully")
	} else {
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")
				return
			}
			log.Fatalf("failed to apply migrations: %v", err)
		}
		fmt.Println("migrations applied successfully")
	}
}

type Log struct {
	verbose bool
}

func (l *Log) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (l *Log) Verbose() bool {
	return l.verbose
}

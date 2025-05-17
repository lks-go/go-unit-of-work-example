package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-unit-of-work-example/internal/config"
	"log"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse config: %s", err)
	}

	m, err := migrate.New("file://internal/db/migrations", cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

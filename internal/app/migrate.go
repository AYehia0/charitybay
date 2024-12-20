//go:build migrate

package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 5
	_defaultTimeout  = time.Second
)

func init() {
	// database
	dbUrl, ok := os.LookupEnv("POSTGRES_URL")
	if !ok || dbUrl == "" {
		log.Fatal("POSTGRES_URL is not set")
	}

	dbUrl = fmt.Sprintf("%s?sslmode=disable", dbUrl)

	var m *migrate.Migrate
	var err error

	// migrations
	for attempts := _defaultAttempts; attempts > 0; attempts-- {
		m, err = migrate.New("file://migrations", dbUrl)
		if err != nil {
			log.Printf("Migrate: failed to create migration: %s", err)
			break
		}

		log.Printf("Migrate: Trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
	}

	if err != nil {
		log.Fatalf("Migrate: failed to create migration: %s", err)
	}

	err = m.Up()
	defer m.Close()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: failed to apply migration: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("Migrate: No migrations applied")
		return
	}

	log.Println("Migrate: all is good")

}

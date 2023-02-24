package postgres

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"time"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	_defaultAttempts = 10
	_defaultTimeout  = time.Second
)

func Up(migrationsUrl, databaseURL string) {
	if len(databaseURL) == 0 {
		log.Fatalf("Migrate: database url is empty")
	}

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New(migrationsUrl, databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}

func Down(migrationsUrl, databaseURL string) {
	if len(databaseURL) == 0 {
		log.Fatalf("Migrate: database url is empty")
	}

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New(migrationsUrl, databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Down()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: down error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: down success")
}

func LoadFixtures(fixturesPath string) {
	ctx := context.Background()
	q, err := os.ReadFile(fixturesPath)
	if err != nil {
		log.Fatalf("LoadFixtures - os.ReadFile: %s", err)
	}
	pg, err := pgx.Connect(ctx, fixturesPath)
	if err != nil {
		log.Fatalf("LoadFixtures - pgx.Connect: %s", err)
	}
	_, err = pg.Exec(ctx, string(q))
	if err != nil {
		log.Fatalf("LoadFixtures - pg.Exec: %s", err)
	}
	err = pg.Close(ctx)
	if err != nil {
		log.Fatalf("LoadFixtures - pg.Close: %s", err)
	}
}

// CleanDB drops everything inside database.
func CleanDB(migrationsUrl, databaseURL string) {
	if len(databaseURL) == 0 {
		log.Fatalf("CleanDB: database url is empty")
	}

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New(migrationsUrl, databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("CleanDB: postgres connect error: %s", err)
	}

	err = m.Drop()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("CleanDB: drop error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("CleanDB: no change")
		return
	}

	log.Printf("Migrate: drop success")
}

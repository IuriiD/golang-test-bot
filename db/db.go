package db

import (
	"fmt"
	"os"
	"time"

	"github.com/avast/retry-go"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

// DBSession is a wrapper around upper.db's database struct
// this allows for easier custom db methods to be added
type DBSession struct {
	sqlbuilder.Database
}

// New creates a new database session
func New() (*DBSession, error) {
	var dbSession sqlbuilder.Database

	err := retry.Do(
		func() error {
			sess, err := postgresql.Open(getDBSettings())
			if err != nil {
				return err
			}

			dbSession = sess
			return nil
		},
		retry.OnRetry(func(n uint, err error) {
			fmt.Printf("Trying database connection %d\n%v", n, err)
		}),
		retry.Delay(1*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return &DBSession{
		Database: dbSession,
	}, nil
}

func getDBSettings() postgresql.ConnectionURL {
	settings := postgresql.ConnectionURL{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	if os.Getenv("DB_SSL") == "true" {
		settings.Options = map[string]string{
			"sslmode": "require",
		}
	}

	return settings
}

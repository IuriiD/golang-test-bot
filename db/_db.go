package _db

import (
	"fmt"
	"os"

	retry "github.com/avast/retry-go"
	sqlbuilder "upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

// User represents the main db table
type User struct {
	ID       string `db:"id"`
	BotState string `db:"bot_state"`
}

// DBSession is a wrapper around upper.db's database struct
type DBSession struct {
	sqlbuilder.Database
}

// New creates a new database session
func New() (*DBSession, error) {
	var dbSession sqlbuilder.Database

	sess, err := postgresql.Open(GetDBSettings())
	if err != nil {
		fmt.Println("Failed to connect to DB, err: ", err)
		return nil, err
	}
	defer sess.Close()

	dbSession = sess

	return &DBSession{
		Database: dbSession,
	}, nil
}

// GetSchema returns SQL for creating our table
func GetSchema() string {
	return `
        CREATE TABLE IF NOT EXISTS "public"."users" (
            "id" VARCHAR(255) NOT NULL,
            "bot_state" VARCHAR(255),
            CONSTRAINT "users_id_pkey" PRIMARY KEY ("id")
        ) WITH (oids = false);
    `
}

// Migrate will drop and create our "users" table
func (d *DBSession) Migrate() error {
	_, err := d.Exec(GetSchema())
	if err != nil {
		return err
	}

	return nil
}

// GetDBSettings composes the URL for connecting to DB
func GetDBSettings() postgresql.ConnectionURL {
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

// DropSchema returns SQL for dropping the table "users"
func DropSchema() string {
	return `
    DROP TABLE IF EXISTS "users";
    `
}

// DropTables will drop all tables from the database
func (d *DBSession) DropTables() error {
	_, err := d.Exec(DropSchema())
	if err != nil {
		return err
	}

	return nil
}

// GetBotState retrieves current conversation state
// for a chat with user with given PSID
func (d *DBSession) GetBotState(psid string) (string, error) {
	return retry.Do(func() error {
		_, err := d.Find(db.Cond{"id": psid})
		//
	})
}

// UpdateBotState will update bot_state column value (dialog state) in the DB
func (d *DBSession) UpdateBotState(psid, newState string) error {
	return retry.Do(func() error {
		_, err := d.Update("users").Set("bot_state = ?", newState).Where("id = ?", psid).Exec()
		return err
	})
}

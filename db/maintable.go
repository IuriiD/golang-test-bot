package db

import (
	"database/sql"
	"time"

	retry "github.com/avast/retry-go"
	db "upper.io/db.v3"
)

// User represents the main db table
type User struct {
	ID        string    `db:"id,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`

	BotState  *string `db:"bot_state"`
	Templates *string `db:"templates"`
}

// UpdateBotState will update bot_state column value (dialog state) in the DB
func (d *DBSession) UpdateBotState(psid, newState string) error {
	return retry.Do(func() error {
		_, err := d.Update("main").Set("bot_state = ?", newState).Where("id = ?", psid).Exec()
		return err
	})
}

// UpdateUserCountryCode will update the user's country code and original
// country code in the database
func (d *DBSession) UpdateUserCountryCode(psid, cc, cco string) error {
	return retry.Do(func() error {
		_, err := d.Update("users").
			Set("country_code = ?", cc, "country_code_original = ?", cco).
			Where("id = ?", psid).
			Exec()

		return err
	})
}

// UpdateUserEmail will update a user to set the email field
func (d *DBSession) UpdateUserEmail(psid, email string) error {
	return retry.Do(func() error {
		_, err := d.Update("users").
			Set("email = ?", email).
			Where("id = ?", psid).
			Exec()

		return err
	})
}

// UpdateUserNotificationPreference will update the user's notification preference
func (d *DBSession) UpdateUserNotificationPreference(psid string, optIn bool) error {
	return retry.Do(func() error {
		_, err := d.Update("users").
			Set("messenger_notification_optin = ?", optIn).
			Where("id = ?", psid).
			Exec()

		return err
	})
}

// GetTimeSinceLastInteraction will return the time since
// a given user last interacted with the bot. In the case that the user
// has never interacted with the bot, 0 will be returned
func (d *DBSession) GetTimeSinceLastInteraction(psid string) (time.Duration, error) {
	var resRow *sql.Row

	err := retry.Do(func() error {
		row, err := d.QueryRow(`
			SELECT EXTRACT(EPOCH FROM (now() - created_at))
			FROM user_actions
			WHERE user_id = ?
			ORDER BY created_at DESC
			LIMIT 1;
		`, psid)
		resRow = row
		return err
	})
	if err != nil {
		if err == db.ErrNoMoreRows {
			return -1, nil
		}

		return 0, err
	}

	seconds := 0.0
	err = resRow.Scan(&seconds)
	if err != nil {
		return 0, err
	}

	return time.Duration(seconds) * time.Second, nil
}

// GetTimeSinceLastOffsetInteraction will return the time since
// a given user last interacted with the bot. In the case that the user
// has never interacted with the bot, 0 will be returned
// The offset version will offest the events by a given number
func (d *DBSession) GetTimeSinceLastOffsetInteraction(psid string, offset int) (time.Duration, error) {
	var resRow *sql.Row

	err := retry.Do(
		func() error {
			row, err := d.QueryRow(`
			SELECT EXTRACT(EPOCH FROM (now() - created_at))
			FROM user_actions
			WHERE user_id = ?
			ORDER BY created_at DESC
			OFFSET ?
			LIMIT 1;
		`, psid, offset)
			resRow = row
			return err
		},
		retry.RetryIf(func(err error) bool {
			return err != sql.ErrNoRows
		}),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}

		return 0, err
	}

	seconds := 0.0
	err = resRow.Scan(&seconds)
	if err != nil {
		return 0, err
	}

	return time.Duration(seconds) * time.Second, nil
}

package user

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func ByID(db *sql.DB, userID string) User {

	var u User

	err := db.QueryRow(`
		SELECT
			u.firstname,
			u.lastname,
			u.mail
		FROM users AS u
		WHERE u.id = $1`, userID).Scan(
		&u.FirstName,
		&u.LastName,
		&u.MailAddress,
	)
	if err != nil {
		log.Print(err)
	}

	return u
}

package user

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func List(db *sql.DB) []User {

	var ul []User
	var u User

	rows, err := db.Query(`
		SELECT
			u.firstname,
			u.lastname,
			u.mail
		FROM users AS u
		ORDER BY u.id`)
	defer rows.Close()

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		rows.Scan(
			&u.FirstName,
			&u.LastName,
			&u.MailAddress,
		)
		ul = append(ul, u)
	}

	return ul
}

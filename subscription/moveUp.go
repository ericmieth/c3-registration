package subscription

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func moveUp(db *sql.DB) {

	var moveUpID int
	err := db.QueryRow(`
		SELECT id
		FROM subscriptions
		WHERE status_waiting_list = $1
		ORDER BY registration_date ASC`, true).Scan(&moveUpID)

	// return if there is no result
	if err == sql.ErrNoRows {
		return
	}

	_, err = db.Exec(`
		UPDATE subscriptions
		SET status_waiting_list = $1
		WHERE id = $2`, false, moveUpID)
	if err != nil {
		log.Print(err)
	}

	sendMail(db, moveUpID, "moveUp")
}

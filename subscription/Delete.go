package subscription

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func Delete(db *sql.DB, subscriptionID int) {

	var verificationID string
	err := db.QueryRow(`
		DELETE FROM subscriptions
		WHERE id = $1
		RETURNING verification_id`, subscriptionID).Scan(&verificationID)
	if err != nil {
		log.Print(err)
	}

	deleteFile(verificationID)

	// check if someone can move up
	if SlotsAvailable(db) > 0 {
		moveUp(db)
	}
}

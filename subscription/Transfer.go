package subscription

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func Transfer(db *sql.DB, subscriptionID int) {
	_, err := db.Exec(`
		UPDATE subscriptions
		SET status_payment =$1 
		WHERE id = $2`, true, subscriptionID)
	if err != nil {
		log.Print(err)
	}
	sendMail(db, subscriptionID, "payment")
}

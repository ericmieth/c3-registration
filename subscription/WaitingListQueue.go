package subscription

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func WaitingListQueue(db *sql.DB) int {

	var waitingListQueue int
	err := db.QueryRow(`
		SELECT COUNT(*)
		FROM subscriptions
		WHERE status_waiting_list = $1`,
		true).Scan(&waitingListQueue)
	if err != nil {
		log.Print(err)
	}

	return waitingListQueue
}

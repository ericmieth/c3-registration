package subscription

import (
	"github.com/ericmieth/c3-registration/config"

	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func SlotsAvailable(db *sql.DB) int {

	var slotsUsed int
	err := db.QueryRow(`
		SELECT COUNT(*)
		FROM subscriptions
		WHERE status_waiting_list = $1`,
		false).Scan(&slotsUsed)
	if err != nil {
		log.Print(err)
	}

	slotsTotal := config.ReturnSlotsTotal()

	return slotsTotal - slotsUsed
}

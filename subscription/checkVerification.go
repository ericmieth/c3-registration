package subscription

import (
	//"errors"

	"database/sql"
	_ "github.com/lib/pq"
)

func checkVerification(db *sql.DB, verficationID, mailAddress string) error {

	var id int
	err := db.QueryRow(`
		SELECT id
		FROM subscriptions
		WHERE verification_id = $1 AND mail_address = $2`,
		verficationID,
		mailAddress).Scan(&id)

	return err

}

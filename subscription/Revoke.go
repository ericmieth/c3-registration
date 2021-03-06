package subscription

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func Revoke(db *sql.DB, verificationID, mailAddress string) error {

	err := checkVerification(db, verificationID, mailAddress)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		DELETE FROM subscriptions
		WHERE verification_id = $1 AND mail_address = $2`,
		verificationID,
		mailAddress)

	if err != nil {
		log.Print(err)
	}

	deleteFile(verificationID)

	// check if someone can move up
	if SlotsAvailable(db) > 0 {
		moveUp(db)
	}

	return err
}

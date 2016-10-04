package subscription

import (
	//"errors"

	"database/sql"
	_ "github.com/lib/pq"
)

func StatusRequest(db *sql.DB, verificationID, mailAddress string) (Subscription, error) {

	var s Subscription

	err := checkVerification(db, verificationID, mailAddress)
	if err != nil {
		return s, err
	}

	err = db.QueryRow(`
		SELECT
			status_waiting_list,
			status_ticket_uploaded,
			status_payment
		FROM subscriptions
		WHERE verification_id = $1 AND mail_address = $2`,
		verificationID,
		mailAddress).Scan(
		&s.StatusWaitingList,
		&s.StatusTicketUploaded,
		&s.StatusPayment)

	return s, err

}

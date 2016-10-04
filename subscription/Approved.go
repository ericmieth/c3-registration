package subscription

import (
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func Approved(db *sql.DB) []Subscription {

	var subscriptions []Subscription
	var s Subscription

	rows, err := db.Query(`
		SELECT
			id,
			first_name,
			last_name,
			mail_address,
			registration_date,
			verification_id,
			status_waiting_list,
			status_ticket_uploaded,
			status_payment,
			iban
		FROM subscriptions AS s
		WHERE s.status_waiting_list = $1
		ORDER BY s.registration_date ASC`, false)

	defer rows.Close()

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		rows.Scan(
			&s.ID,
			&s.FirstName,
			&s.LastName,
			&s.MailAddress,
			&s.RegistrationDate,
			&s.VerificationID,
			&s.StatusWaitingList,
			&s.StatusTicketUploaded,
			&s.StatusPayment,
			&s.IBAN)

		subscriptions = append(subscriptions, s)
	}

	return subscriptions
}

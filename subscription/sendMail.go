package subscription

import (
	"log"
	"net/smtp"

	"database/sql"
	_ "github.com/lib/pq"
)

func sendMail(db *sql.DB, requestID int, templateName string) {

	var s Subscription

	err := db.QueryRow(`
		SELECT
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
		WHERE s.id = $1`, requestID).Scan(
		&s.FirstName,
		&s.LastName,
		&s.MailAddress,
		&s.RegistrationDate,
		&s.VerificationID,
		&s.StatusWaitingList,
		&s.StatusTicketUploaded,
		&s.StatusPayment,
		&s.IBAN)

	if err != nil {
		log.Print(err)
	}

	auth := smtp.PlainAuth("", "fsinf@fsinf.informatik.uni-leipzig.de", "", "mail.informatik.uni-leipzig.de")

	receiver := []string{s.MailAddress}
	message := renderMailTemplate(s, templateName)

	smtp.SendMail("mail.informatik.uni-leipzig.de:25", auth, "fsinf@fsinf.informatik.uni-leipzig.de", receiver, message)

}

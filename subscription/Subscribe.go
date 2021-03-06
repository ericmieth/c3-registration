package subscription

import (
	"errors"
	"regexp"
	"time"

	"database/sql"
	_ "github.com/lib/pq"
)

// handles a subscription and returns the verification ID an an error
func Subscribe(db *sql.DB, firstName, lastName, mailAddress, iban string) (string, error) {

	// return, when its not the universitary mail address
	isUniversityMail := regexp.MustCompile(".+@studserv.uni-leipzig.de")
	if !isUniversityMail.MatchString(mailAddress) {
		return "", errors.New("This is not an universitary mail address.")
	}

	now := time.Now()

	verificationID := createVerificationID()

	// check if slots available
	var statusWaitingList bool
	if SlotsAvailable(db) < 1 {
		statusWaitingList = true
	} else {
		statusWaitingList = false
	}

	// insert into database

	var insertID int
	err := db.QueryRow(`
		INSERT INTO subscriptions (
			first_name,
			last_name,
			mail_address,
			registration_date,
			verification_id,
			status_waiting_list,
			iban
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`,
		firstName,
		lastName,
		mailAddress,
		now,
		verificationID,
		statusWaitingList,
		iban).Scan(&insertID)

	if err != nil {
		return "", err
	}

	if statusWaitingList {
		sendMail(db, insertID, "waitingList")
	} else {
		sendMail(db, insertID, "approvedRegistration")
	}

	return verificationID, nil
}

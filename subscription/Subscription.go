package subscription

import (
	"time"
)

type Subscription struct {
	ID                   int
	FirstName            string
	LastName             string
	MailAddress          string
	VerificationID       string
	RegistrationDate     time.Time
	StatusWaitingList    bool
	StatusTicketUploaded bool
	StatusPayment        bool
	IBAN                 string
}

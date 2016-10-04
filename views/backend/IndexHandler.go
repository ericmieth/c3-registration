package backend

import (
	"github.com/ericmieth/c3-registration/config"
	"github.com/ericmieth/c3-registration/subscription"

	"net/http"

	"database/sql"
	_ "github.com/lib/pq"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var err error
	var suc string

	data := struct {
		CongressNumber           string
		ErrorMessage             error
		SuccessMessage           string
		SubscriptionsApproved    []subscription.Subscription
		SubscriptionsWaitingList []subscription.Subscription
		SlotsAvailable           int
		WaitingListQueue         int
	}{
		SlotsAvailable:           subscription.SlotsAvailable(db),
		CongressNumber:           config.ReturnCongressNumber(),
		SubscriptionsApproved:    subscription.Approved(db),
		SubscriptionsWaitingList: subscription.WaitingList(db),
		ErrorMessage:             err,
		SuccessMessage:           suc,
		WaitingListQueue:         subscription.WaitingListQueue(db),
	}

	renderView(w, "index", data)
}

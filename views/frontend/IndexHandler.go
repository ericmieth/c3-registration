package frontend

import (
	"github.com/ericmieth/c3-registration/config"
	"github.com/ericmieth/c3-registration/subscription"

	"database/sql"
	"errors"
	"net/http"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var err error
	var suc string
	var s interface{}

	if r.Method == "POST" {
		action := strings.TrimSpace(r.FormValue("action"))

		switch {
		case action == "subscribe":
			firstName := strings.TrimSpace(r.FormValue("first-name"))
			lastName := strings.TrimSpace(r.FormValue("last-name"))
			mailAddress := strings.TrimSpace(r.FormValue("mail-address"))
			iban := strings.TrimSpace(r.FormValue("iban"))

			if len(firstName) == 0 ||
				len(lastName) == 0 ||
				len(mailAddress) == 0 ||
				len(iban) == 0 {
				err = errors.New("no empty fields allowed.")
			} else {

				verificationID, e := subscription.Subscribe(db, firstName, lastName, mailAddress, iban)
				err = e
				if err == nil {
					suc = "Your request was transmitted succesfuly. Your verification ID is: " + verificationID
				}
			}

		case action == "revoke":
			verificationID := strings.TrimSpace(r.FormValue("verification-id"))
			mailAddress := strings.TrimSpace(r.FormValue("mail-address"))

			if len(verificationID) == 0 ||
				len(mailAddress) == 0 {
				err = errors.New("no empty fields allowed.")
			} else {

				err = subscription.Revoke(db, verificationID, mailAddress)
				if err == nil {
					suc = "Your request is revoked. All your data is deleted."
				}
			}

		case action == "status-request":
			verificationID := strings.TrimSpace(r.FormValue("verification-id"))
			mailAddress := strings.TrimSpace(r.FormValue("mail-address"))

			if len(verificationID) == 0 ||
				len(mailAddress) == 0 {
				err = errors.New("no empty fields allowed.")
			} else {
				subscription, errorStatusRequest := subscription.StatusRequest(db, verificationID, mailAddress)
				if errorStatusRequest == nil {
					s = subscription
				} else {
					err = errorStatusRequest
				}
			}

		case action == "upload-ticket":
			verificationID := strings.TrimSpace(r.FormValue("verification-id"))
			mailAddress := strings.TrimSpace(r.FormValue("mail-address"))

			if len(verificationID) == 0 ||
				len(mailAddress) == 0 {
				err = errors.New("no empty fields allowed.")
			} else {

				uploadFile, uploadFileHeader, uploadError := r.FormFile("file")
				defer uploadFile.Close()

				if uploadError != nil {
					err = uploadError
				} else {
					suc = "Your ticket was uploaded succesfuly."
				}

				uploadError = subscription.UploadTicket(w, r, db, verificationID, mailAddress, uploadFile, uploadFileHeader)
				if uploadError != nil {
					err = uploadError
				}
			}
		}
	}

	data := struct {
		CongressNumber   string
		ErrorMessage     error
		SuccessMessage   string
		SlotsAvailable   int
		WaitingListQueue int
		StatusMessage    interface{}
	}{
		SlotsAvailable:   subscription.SlotsAvailable(db),
		CongressNumber:   config.ReturnCongressNumber(),
		ErrorMessage:     err,
		SuccessMessage:   suc,
		WaitingListQueue: subscription.WaitingListQueue(db),
		StatusMessage:    s,
	}

	renderView(w, data)
}

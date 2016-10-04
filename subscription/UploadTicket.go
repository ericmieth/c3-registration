package subscription

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
)

// UploadFile uploads a file from an HTTP POST and stores it in the
// uploads directory. It returns the name of the uploaded file, and an error.
func UploadTicket(w http.ResponseWriter, r *http.Request, db *sql.DB, verificationID, mailAddress string, file multipart.File, fileHeader *multipart.FileHeader) error {

	// check verification

	err := checkVerification(db, verificationID, mailAddress)
	if err != nil {
		return err
	}

	// copy file to server

	pathPrefix := "./uploads/"
	pathSuffix := prepareDocumentLink(verificationID) + ".pdf"

	storeFile, err := os.OpenFile(pathPrefix+pathSuffix, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Print(err)
	}
	defer storeFile.Close()
	io.Copy(storeFile, file)

	// insert file in database

	// TODO: increment number if file exists

	_, err = db.Exec(`
		UPDATE subscriptions
		SET status_ticket_uploaded = $1
		WHERE verification_id = $2`, true, verificationID)

	if err != nil {
		log.Print(err)
	}

	return nil
}

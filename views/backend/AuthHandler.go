package backend

import (
	"github.com/ericmieth/c3-registration/user"

	"net/http"
	"strings"

	"database/sql"
	_ "github.com/lib/pq"
)

func AuthHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	if r.Method != "POST" {
		http.Redirect(w, r, "admin", http.StatusFound)
		return
	}
	mailAddress := strings.ToLower(strings.TrimSpace(r.FormValue("mail-address")))
	passphrase := strings.TrimSpace(r.FormValue("password"))

	// check credentials
	if user.CheckCredentials(mailAddress, passphrase, db) {
		// the credentials are valid, so create a jwt
		user.SetCookieWithToken(w, r, mailAddress, db)
	}
	http.Redirect(w, r, "admin", http.StatusFound)

}

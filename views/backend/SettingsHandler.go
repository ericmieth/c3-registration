package backend

import (
	"github.com/ericmieth/c3-registration/config"
	"github.com/ericmieth/c3-registration/user"

	"net/http"
	"strings"

	"database/sql"
	_ "github.com/lib/pq"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	userID := user.IDFromJWT(r)

	if r.Method == "POST" {
		firstName := strings.TrimSpace(r.FormValue("first-name"))
		lastName := strings.TrimSpace(r.FormValue("last-name"))
		mailAddress := strings.TrimSpace(r.FormValue("mail-address"))
		passphrase1 := strings.TrimSpace(r.FormValue("passphraseFirst"))
		passphrase2 := strings.TrimSpace(r.FormValue("passphraseSecond"))

		user.Update(db, userID, firstName, lastName, mailAddress, passphrase1, passphrase2)
	}

	data := struct {
		CongressNumber string
		UserData       user.User
	}{
		CongressNumber: config.ReturnCongressNumber(),
		UserData:       user.ByID(db, userID),
	}

	renderView(w, "settings", data)
}

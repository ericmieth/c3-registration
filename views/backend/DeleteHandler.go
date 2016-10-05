package backend

import (
	"github.com/ericmieth/c3-registration/subscription"

	"log"
	"net/http"
	"strconv"
	"strings"

	"database/sql"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	URLValues := r.URL.Query()
	subscriptionIDFromURL := strings.TrimSpace(URLValues.Get("id"))
	subscriptionIDFromURLInt, err := strconv.Atoi(subscriptionIDFromURL)

	switch {
	case err != nil:
		log.Print(err)
	case subscriptionIDFromURL == "":
		// return if there is no 'id' argument in url
	default:
		subscription.Delete(db, subscriptionIDFromURLInt)
	}

	http.Redirect(w, r, "/admin", http.StatusFound)
	return
}

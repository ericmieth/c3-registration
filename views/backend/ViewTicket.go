package backend

import (
	"net/http"

	"database/sql"
)

func ViewTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	http.ServeFile(w, r, r.URL.Path[1:])
}

package user

import (
	"net/http"

	"database/sql"
)

type HandlerFuncDB func(http.ResponseWriter, *http.Request, *sql.DB)

// CheckAuth validates if a user is logged in
func CheckAuth(f HandlerFuncDB, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isAuthorized(r) {
			renewJWT(w, r)
			f(w, r, db)
		} else {
			http.Redirect(w, r, "c3//admin/login/", http.StatusFound)
		}
	}
}

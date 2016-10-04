package backend

import (
	"github.com/ericmieth/c3-registration/config"
	"github.com/ericmieth/c3-registration/user"

	"net/http"

	"database/sql"
	_ "github.com/lib/pq"
)

func UserListHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var err error
	var suc string

	data := struct {
		CongressNumber string
		ErrorMessage   error
		SuccessMessage string
		UserList       []user.User
	}{
		CongressNumber: config.ReturnCongressNumber(),
		ErrorMessage:   err,
		SuccessMessage: suc,
		UserList:       user.List(db),
	}

	renderView(w, "userlist", data)
}

package backend

import (
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	renderView(w, "login", nil)

}

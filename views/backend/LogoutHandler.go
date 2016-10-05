package backend

import (
	"github.com/ericmieth/c3-registration/user"

	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	user.InvalidateCookieWithToken(w)
	http.Redirect(w, r, "/c3/admin", http.StatusFound)

}

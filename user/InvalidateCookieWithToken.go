package user

import (
	"net/http"
)

func InvalidateCookieWithToken(w http.ResponseWriter) {

	cookieName := returnCookieName()

	http.SetCookie(w, &http.Cookie{
		Name:       cookieName,
		Value:      "",
		Path:       "/",
		RawExpires: "0",
		Secure:     true,
		HttpOnly:   true,
	})

}

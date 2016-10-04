package user

import (
	"net/http"
)

func getTokenCookie(r *http.Request, tokenName string) (*http.Cookie, error) {
	tokenCookie, err := r.Cookie(tokenName)
	switch {
	case err == http.ErrNoCookie:
		//log.Print("no token")
		return tokenCookie, err
	case err != nil:
		//log.Printf("error parsing cookie: %v\n", err)
		return tokenCookie, err
	}
	return tokenCookie, nil

}

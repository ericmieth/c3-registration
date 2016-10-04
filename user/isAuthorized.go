package user

import (
	jwt "github.com/dgrijalva/jwt-go"
	//"log"
	"net/http"
)

func isAuthorized(r *http.Request) bool {

	cookieName := returnCookieName()
	_, verifyKey := returnRSAKeypair()

	tokenCookie, err := getTokenCookie(r, cookieName)
	if err != nil {
		return false
	}

	// just for the lulz, check if it is empty.. should fail on Parse anyway..
	if tokenCookie.Value == "" {
		//w.WriteHeader(http.StatusUnauthorized)
		//log.Printf("no token")
		return false
	}

	// validate the token
	token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	// branch out into the possible error from signing
	switch err.(type) {

	case nil: // no error

		if !token.Valid { // but may still be invalid
			//log.Printf("invalid token")
			return false
		}

		// see stdout and watch for the CustomUserInfo, nicely unmarshalled
		//log.Printf("access granted with token %+v\n", token)

	case *jwt.ValidationError: // something was wrong during the validation
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			//w.WriteHeader(http.StatusUnauthorized)
			//log.Printf("token expired")
			return false

		default:
			//log.Printf("Error parsing token")
			//log.Printf("ValidationError error: %+v\n", vErr.Errors)
			return false
		}

	default: // something else went wrong
		//log.Printf("Error while Parsing Token!")
		//log.Printf("Token parse error: %v\n", err)
		return false
	}

	if token.Valid {
		return true
	}
	return false
}

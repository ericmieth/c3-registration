package user

import (
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
)

func IDFromJWT(r *http.Request) string {

	tokenCookie, err := getTokenCookie(r, returnCookieName())
	if err != nil {
		log.Print("error getting cookie")
	}

	// just for the lulz, check if it is empty.. should fail on Parse anyway..
	if tokenCookie.Value == "" {
		log.Printf("no token")
		return ""
	}

	// validate the token
	t, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	if err != nil {
		log.Print("error")
	}

	claims := t.Claims.(jwt.MapClaims)
	return strconv.Itoa(int(claims["id"].(float64)))

}

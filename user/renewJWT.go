package user

import (
	"github.com/ericmieth/c3-registration/config"

	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

//  updates the expire date of the jwt inside the token
func renewJWT(w http.ResponseWriter, r *http.Request) {

	cookieName := returnCookieName()
	signKey, _ := returnRSAKeypair()

	//t := jwt.New(jwt.GetSigningMethod("RS256"))
	tokenCookie, err := getTokenCookie(r, returnCookieName())
	if err != nil {
		log.Print("error getting cookie")
	}

	// just for the lulz, check if it is empty.. should fail on Parse anyway..
	if tokenCookie.Value == "" {
		//w.WriteHeader(http.StatusUnauthorized)
		log.Printf("no token")
		return
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
	claims["exp"] = time.Now().Add(time.Minute * config.ReturnLoginExpireMinutes()).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		log.Printf("Token Signing error: %v\n", err)
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:       cookieName,
		Value:      tokenString,
		Path:       "/",
		RawExpires: "0",
		//Secure:     true,
		Secure:   false,
		HttpOnly: true,
	})

}

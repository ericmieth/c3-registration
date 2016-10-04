package user

import (
	"github.com/ericmieth/c3-registration/config"

	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"

	"database/sql"
	_ "github.com/lib/pq"
)

// SetCookieWithToken creates a new token for a user with some information
// stored in a jwt
func SetCookieWithToken(w http.ResponseWriter, r *http.Request, mailAddress string, db *sql.DB) {

	var userID int
	err := db.QueryRow(`
		SELECT id
		FROM users
		WHERE mail = $1`, mailAddress).Scan(&userID)

	cookieName := returnCookieName()
	signKey, _ := returnRSAKeypair()

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * config.ReturnLoginExpireMinutes()).Unix()

	tokenString, err := t.SignedString(signKey)
	if err != nil {
		log.Printf("Token Signing error: %v\n", err)
		return
	}

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

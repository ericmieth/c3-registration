package user

import (
	"encoding/base64"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func CheckCredentials(mailAddress string, passphrase string, db *sql.DB) bool {

	var u User

	err := db.QueryRow(`
		SELECT
			pw_hash,
			salt
		FROM users
		WHERE mail = $1`, mailAddress).Scan(
		&u.PassHash64,
		&u.Salt64,
	)
	if err != nil {
		log.Print(err)
		return false
	}

	u.MailAddress = mailAddress

	// if there is an empty field in database
	if u.PassHash64 == "" || u.Salt64 == "" {
		log.Print(err)
		return false
	}

	u.PassHash, err = base64.StdEncoding.DecodeString(u.PassHash64)
	if err != nil {
		log.Print(err)
	}
	u.Salt, err = base64.StdEncoding.DecodeString(u.Salt64)
	if err != nil {
		log.Print(err)
	}

	generatedHash, err := HashPassphraseWithSalt(passphrase, u.Salt)
	if err != nil {
		log.Print(err)
	}

	valid := true
	for i := range generatedHash {
		if generatedHash[i] != u.PassHash[i] {
			valid = false
		}
	}
	return valid

}

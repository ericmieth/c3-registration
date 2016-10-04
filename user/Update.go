package user

import (
	"encoding/base64"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func Update(db *sql.DB, userID, firstName, lastName, mailAddress, passphrase1, passphrase2 string) {

	// update first name
	if firstName != "" {
		_, err := db.Exec(`
				UPDATE users
				SET firstname = $1
				WHERE id = $2`, firstName, userID)

		if err != nil {
			log.Print(err)
		}
	}

	// update first name
	if lastName != "" {
		_, err := db.Exec(`
				UPDATE users
				SET lastname = $1
				WHERE id = $2`, lastName, userID)

		if err != nil {
			log.Print(err)
		}

	}

	// update mail address
	if mailAddress != "" {
		_, err := db.Exec(`
				UPDATE users
				SET mail = $1
				WHERE id = $2`, mailAddress, userID)

		if err != nil {
			log.Print(err)
		}
	}

	// update passphrase
	if passphrase1 != "" || passphrase2 != "" {
		if passphrase1 != passphrase2 {
			return
		} else {
			newSalt, err := CreateNewSalt()
			if err != nil {
				log.Print(err)
			}
			newSalt64 := base64.StdEncoding.EncodeToString(newSalt)

			hash, err := HashPassphraseWithSalt(passphrase1, newSalt)
			if err != nil {
				log.Print(err)
			}
			hash64 := base64.StdEncoding.EncodeToString(hash)
			_, err = db.Exec(`
					UPDATE users
					SET pw_hash = $1, salt = $2
					WHERE id = $3`, hash64, newSalt64, userID)

			if err != nil {
				log.Print(err)
			}
		}
	}
}

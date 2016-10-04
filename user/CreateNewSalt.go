package user

import (
	"crypto/rand"
	"io"
	"log"
)

const PW_SALT_BYTES = 64

// createNewSaltBase64 creates a new base64 encoded salt and returns
// the salt, the base64 encoded salt and an error
func CreateNewSalt() ([]byte, error) {
	// create a new salt
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Print(err)
	}
	return salt, err
}

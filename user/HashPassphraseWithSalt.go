package user

import (
	"golang.org/x/crypto/scrypt"
)

const PW_HASH_BYTES = 64

// The recommended parameters for interactive logins as of 2009 are N=16384, r=8, p=1.
// https://godoc.org/golang.org/x/crypto/scrypt
func HashPassphraseWithSalt(passphrase string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(passphrase), salt, 16384, 8, 1, PW_HASH_BYTES)
}

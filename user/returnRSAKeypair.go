package user

import (
	"crypto/rsa"
	jwt "github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

const (
	privKeyPath = "static/keys/jwt.rsa"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "static/keys/jwt.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func returnRSAKeypair() (*rsa.PrivateKey, *rsa.PublicKey) {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Print(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Print(err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Print(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Print(err)
	}

	return signKey, verifyKey
}

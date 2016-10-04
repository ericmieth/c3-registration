package subscription

import (
	"github.com/ericmieth/c3-registration/config"

	"strings"
)

// concatenates an event prefix and some random to a unique
// verification ID
func createVerificationID() string {

	prefix := config.ReturnCongressNumber() + "C3"
	randomString := strings.ToUpper(generateRandomString(11))

	var verificationID []string
	verificationID = append(verificationID, prefix)
	verificationID = append(verificationID, randomString)

	return strings.Join(verificationID, "-")
}

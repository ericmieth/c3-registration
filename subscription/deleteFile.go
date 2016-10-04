package subscription

import (
	"log"
	"os"
)

func deleteFile(verificationID string) {

	pathPrefix := "./uploads/"
	err := os.Remove(pathPrefix + verificationID + ".pdf")

	if err != nil {
		log.Print(err)
	}

}

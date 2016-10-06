package subscription

import (
	"regexp"
)

func isUniversityMail(mailAddress string) bool {

	// return, when its not the universitary mail address
	isUniversityMail := regexp.MustCompile(`([[:alpha:]]{3}[[:digit:]]{2}[[:alpha:]]{3}|[[:alpha:]]{2}[[:digit:]]{2}[[:alpha:]]{4})@studserv.uni-leipzig.de`)
	if isUniversityMail.MatchString(mailAddress) {
		return true
	} else {
		return false
	}

}

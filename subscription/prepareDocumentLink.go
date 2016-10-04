package subscription

import (
	"strings"
	"unicode"
)

func prepareDocumentLink(documentName string) string {

	// drop all, thats not a number or a letter
	dropSymobls := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '-' && r != '.' {
			return -1

		}
		return r
	}
	finalString := strings.Replace(documentName, " ", "-", -1)
	finalString = strings.Map(dropSymobls, finalString)
	finalString = strings.Replace(finalString, "ä", "ae", -1)
	finalString = strings.Replace(finalString, "ö", "oe", -1)
	finalString = strings.Replace(finalString, "ü", "ue", -1)
	finalString = strings.Replace(finalString, "ß", "ss", -1)

	return finalString
}

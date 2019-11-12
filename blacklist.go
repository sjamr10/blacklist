package blacklist

import "strings"

// Replace replaces blacklisted words/phrases from text
func Replace(text string, blacklist string) (string, error) {
	text = strings.Replace(text, blacklist, "", -1)
	return text, nil
}

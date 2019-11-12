package blacklist

import "strings"

// Replace replaces blacklisted words/phrases from text
func Replace(text string, blacklist string) (string, error) {
	for _, blacklistWord := range strings.Split(blacklist, "\n") {
		text = strings.Replace(text, blacklistWord, "", -1)
	}

	return text, nil
}

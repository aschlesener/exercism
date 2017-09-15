package acronym

import (
	"bytes"
	"strings"
	"unicode"
)

const testVersion = 2

// Abbreviate makes an acronym out of a phrase
func Abbreviate(phrase string) string {
	if phrase == "" {
		return phrase
	}

	var buffer bytes.Buffer
	words := strings.FieldsFunc(phrase, split)

	for _, word := range words {
		if len(word) > 0 {
			buffer.WriteRune(unicode.ToUpper(rune(word[0])))
			// handle case where there is a lowercase letter followed by an uppercase letter e.g. HyperText Markup Language
			if len(word) > 1 {
				for i := 0; i < len(word)-1; i++ {
					if unicode.IsLetter(rune(word[i])) && unicode.IsLetter(rune(word[i+1])) &&
						rune(word[i]) == unicode.ToLower(rune(word[i])) && rune(word[i+1]) == unicode.ToUpper(rune(word[i+1])) {
						buffer.WriteRune(unicode.ToUpper(rune(word[i+1])))
					}
				}
			}
		}
	}
	return buffer.String()
}

// helper function used to split a string by multiple delimiters
func split(r rune) bool {
	return r == ' ' || r == '-'
}

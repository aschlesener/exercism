package bob

import "unicode"
import "strings"

const testVersion = 2

const (
	fine     = "Fine. Be that way!"
	whatever = "Whatever."
	whoa     = "Whoa, chill out!"
	sure     = "Sure."
)

// Hey represents a teenager's backtalk
// Helper functions are used for readability, if we were running this on a long string/many strings and
// cared about performance we could inspect the input just once to do all the checks
func Hey(input string) string {
	if isAllWhitespace(input) {
		return fine
	}
	input = removeWhitespace(input)
	if lettersAreAllCapital(input) {
		return whoa
	}
	if input[len(input)-1] == '?' {
		return sure
	}

	return whatever
}

// remove all whitespaces from a string
func removeWhitespace(input string) string {
	r := strings.NewReplacer(" ", "")
	input = r.Replace(input)
	return input
}

// determine if a string contains all whitespace
func isAllWhitespace(input string) bool {
	if input == "" {
		return true
	}
	for _, letter := range input {
		if !unicode.IsSpace(letter) {
			return false
		}
	}
	return true
}

// determine if all letters in a string are capital
func lettersAreAllCapital(input string) bool {
	containsLetters := false
	for _, letter := range input {
		if unicode.IsLetter(letter) {
			containsLetters = true
			if letter != unicode.ToUpper(letter) {
				return false
			}
		}
	}

	if containsLetters {
		return true
	}
	return false
}

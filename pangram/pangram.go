package pangram

import "strings"

const testVersion = 1

// IsPangram determines if a string is a pangram (contains all letters of the alphabet)
func IsPangram(input string) bool {
	if input == "" {
		return false
	}

	input = strings.ToLower(input)
	chars := make([]int, 26)

	for i := 0; i < len(input); i++ {
		// check for non-ASCII characters
		if input[i] > 255 {
			return false
		}
		// ignore non-alphabet ASCII characters
		if input[i] < 97 || input[i] > 122 {
			continue
		}
		chars[input[i]-97]++
	}

	// check if all alphabet chars were present in string
	for _, char := range chars {
		if char == 0 {
			return false
		}
	}

	return true
}

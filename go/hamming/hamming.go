package hamming

import "errors"

const testVersion = 5

// Distance calculates the hamming difference between two strings
func Distance(a, b string) (int, error) {
	if a == b {
		return 0, nil
	}
	if len(a) != len(b) {
		return -1, errors.New("string lengths must be the same")
	}

	diffCount := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffCount++
		}
	}
	return diffCount, nil
}

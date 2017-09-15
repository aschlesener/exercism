package raindrops

import "strconv"
import "bytes"

const testVersion = 2

// Convert converts a number to a raindrop using bytes buffer instead of string concat for memory efficiency
func Convert(number int) string {
	var buffer bytes.Buffer

	if number%3 == 0 {
		buffer.WriteString("Pling")
	}
	if number%5 == 0 {
		buffer.WriteString("Plang")
	}
	if number%7 == 0 {
		buffer.WriteString("Plong")
	}

	if buffer.Len() == 0 {
		return strconv.Itoa(number)
	}

	return buffer.String()
}

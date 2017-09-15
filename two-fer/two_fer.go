// Package twofer provides functionality to create a sentence of the "twofer" form
package twofer

// ShareWith returns a string of the form "One for (name), one for you
// If no name is provided, the name in the sentence will default to "you"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}

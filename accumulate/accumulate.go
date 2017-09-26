package accumulate

const testVersion = 1

// Accumulate applies a given operation function to every element of a given collection
func Accumulate(collection []string, operation func(string) string) []string {
	// modify in-place; could instead create new collection
	for i, item := range collection {
		collection[i] = operation(item)
	}
	return collection
}

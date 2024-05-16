package utils

import "math/rand"

const (
	// charSetAlphaNum is the alphanumeric character set for use with
	// RandStringFromCharSet.
	charSetAlphaNum = "abcdefghijklmnopqrstuvwxyz012346789"

	// charSetAlpha is the alphabetical character set for use with
	// RandStringFromCharSet.
	charSetAlpha = "abcdefghijklmnopqrstuvwxyz"

	// Length of the resource name we wish to generate.
	resourceNameLength = 10
)

// GenerateRandomResourceName builds a unique-ish resource identifier to use in
// tests.
func GenerateRandomResourceName() string {
	result := make([]byte, resourceNameLength)
	for i := 0; i < resourceNameLength; i++ {
		result[i] = charSetAlpha[randIntRange(0, len(charSetAlpha))]
	}
	return string(result)
}

// randIntRange returns a random integer between min (inclusive) and max
// (exclusive).
func randIntRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

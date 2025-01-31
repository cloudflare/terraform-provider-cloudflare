package utils

import "math/rand"

const (
	// CharSetAlphaNum is the alphanumeric character set for use with
	// RandStringFromCharSet.
	CharSetAlphaNum = "abcdefghijklmnopqrstuvwxyz012346789"

	// CharSetAlpha is the alphabetical character set for use with
	// RandStringFromCharSet.
	CharSetAlpha = "abcdefghijklmnopqrstuvwxyz"

	// Length of the resource name we wish to generate.
	ResourceNameLength = 10
)

// GenerateRandomResourceName builds a unique-ish resource identifier to use in
// tests.
func GenerateRandomResourceName() string {
	result := make([]byte, ResourceNameLength)
	for i := 0; i < ResourceNameLength; i++ {
		result[i] = CharSetAlpha[randIntRange(0, len(CharSetAlpha))]
	}
	return string(result)
}

// RandStringFromCharSet generates a random string by selecting characters from
// the charset provided
func RandStringFromCharSet(strlen int, charSet string) string {
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSet[RandIntRange(0, len(charSet))]
	}
	return string(result)
}

// RandIntRange returns a random integer between min (inclusive) and max (exclusive)
func RandIntRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

// RandString generates a random alphanumeric string of the length specified
func RandString(strlen int) string {
	return RandStringFromCharSet(strlen, CharSetAlphaNum)
}

// randIntRange returns a random integer between min (inclusive) and max
// (exclusive).
func randIntRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

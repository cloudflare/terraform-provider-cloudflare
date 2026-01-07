package utils

import (
	"os"
	"strings"
)

const (
	// TestResourcePrefix is the standard prefix for all test resources
	// All integration and migration tests must create resources with this prefix
	// Uses underscores instead of dashes for API compatibility (some APIs reject dashes)
	TestResourcePrefix = "cftftest"
)

// IsTestResource checks if a resource name follows the test naming conventions.
// It returns true if the resource name contains the test prefix.
func IsTestResource(name string) bool {
	return strings.Contains(name, TestResourcePrefix)
}

// ShouldSweepResource determines if a resource should be cleaned up by sweepers.
// It checks if the resource name contains the standard test prefix.
//
// DANGER MODE: If SWEEP_DANGEROUSLY_DELETE_ALL environment variable is set to "true",
// this function will return true for ALL resources, bypassing name validation.
// This is extremely dangerous and should only be used with explicit intent.
func ShouldSweepResource(name string) bool {
	// Check for danger mode environment variable
	if os.Getenv("SWEEP_DANGEROUSLY_DELETE_ALL") == "true" {
		return true // DANGER: Delete everything!
	}

	// Normal mode: only delete resources with test prefix
	return IsTestResource(name)
}

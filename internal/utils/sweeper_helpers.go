package utils

import (
	"os"
	"strings"
)

const (
	// TestResourcePrefix is the standard prefix for all test resources
	// All integration and migration tests must create resources with this prefix
	TestResourcePrefix = "cf-tf-test-"

	// LegacyTestResourcePrefix is the old prefix that was used before standardization
	// This is supported temporarily during migration but will be removed in the future
	LegacyTestResourcePrefix = "tf-acctest-"
)

// IsTestResource checks if a resource name follows the test naming conventions.
// It returns true if the resource name starts with the test prefix.
func IsTestResource(name string) bool {
	return strings.HasPrefix(name, TestResourcePrefix)
}

// IsLegacyTestResource checks if a resource name follows the legacy test naming conventions.
// This is supported temporarily during migration.
func IsLegacyTestResource(name string) bool {
	return strings.HasPrefix(name, LegacyTestResourcePrefix)
}

// ShouldSweepResource determines if a resource should be cleaned up by sweepers.
// During the migration period, this supports both new and legacy test prefixes.
// After migration is complete, legacy prefix support should be removed.
//
// DANGER MODE: If SWEEP_DANGEROUSLY_DELETE_ALL environment variable is set to "true",
// this function will return true for ALL resources, bypassing name validation.
// This is extremely dangerous and should only be used with explicit intent.
func ShouldSweepResource(name string) bool {
	// Check for danger mode environment variable
	if os.Getenv("SWEEP_DANGEROUSLY_DELETE_ALL") == "true" {
		return true // DANGER: Delete everything!
	}

	// Normal mode: only delete resources with test prefixes
	return IsTestResource(name) || IsLegacyTestResource(name)
}

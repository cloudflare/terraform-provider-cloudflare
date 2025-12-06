package utils

import (
	"strings"
	"testing"
)

func TestGenerateRandomResourceName(t *testing.T) {
	// Generate a few names and verify they all have the correct format
	for i := 0; i < 10; i++ {
		name := GenerateRandomResourceName()

		// Check it has the correct prefix
		if !strings.HasPrefix(name, TestResourcePrefix) {
			t.Errorf("GenerateRandomResourceName() = %q, expected prefix %q", name, TestResourcePrefix)
		}

		// Check total length (prefix + 10 random chars)
		expectedLength := len(TestResourcePrefix) + ResourceNameLength
		if len(name) != expectedLength {
			t.Errorf("GenerateRandomResourceName() length = %d, want %d", len(name), expectedLength)
		}

		// Check that the suffix after prefix is lowercase letters only
		suffix := strings.TrimPrefix(name, TestResourcePrefix)
		if len(suffix) != ResourceNameLength {
			t.Errorf("Random suffix length = %d, want %d", len(suffix), ResourceNameLength)
		}

		for _, ch := range suffix {
			if ch < 'a' || ch > 'z' {
				t.Errorf("GenerateRandomResourceName() generated non-lowercase char: %c in %q", ch, name)
			}
		}
	}
}

func TestGenerateRandomResourceName_Uniqueness(t *testing.T) {
	// Generate multiple names and verify they're different
	names := make(map[string]bool)
	iterations := 100

	for i := 0; i < iterations; i++ {
		name := GenerateRandomResourceName()
		if names[name] {
			t.Errorf("GenerateRandomResourceName() generated duplicate: %q", name)
		}
		names[name] = true
	}

	// Should have generated unique names (extremely unlikely to have collisions with 26^10 possibilities)
	if len(names) != iterations {
		t.Errorf("Expected %d unique names, got %d", iterations, len(names))
	}
}

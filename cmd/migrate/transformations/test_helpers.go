package transformations

import (
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// normalizeWhitespace normalizes whitespace in a string for comparison
func normalizeWhitespace(s string) string {
	// Replace multiple spaces with single space
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	// Trim leading and trailing whitespace from each line
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	// Remove empty lines
	var nonEmptyLines []string
	for _, line := range lines {
		if line != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return strings.Join(nonEmptyLines, "\n")
}

// compareHCLBlocks compares two HCL configurations semantically, ignoring attribute order
func compareHCLBlocks(t *testing.T, expected, actual string) bool {
	// Parse both configurations
	expectedFile, diags := hclwrite.ParseConfig([]byte(expected), "expected.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("Failed to parse expected config: %v", diags)
	}

	actualFile, diags := hclwrite.ParseConfig([]byte(actual), "actual.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("Failed to parse actual config: %v", diags)
	}

	// Compare all blocks
	expectedBlocks := expectedFile.Body().Blocks()
	actualBlocks := actualFile.Body().Blocks()

	if len(expectedBlocks) != len(actualBlocks) {
		t.Errorf("Block count mismatch: expected %d, got %d", len(expectedBlocks), len(actualBlocks))
		return false
	}

	// Find matching blocks and compare
	for _, expBlock := range expectedBlocks {
		found := false
		for _, actBlock := range actualBlocks {
			if blocksMatch(expBlock, actBlock) {
				if !compareBlockAttributes(t, expBlock, actBlock) {
					return false
				}
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected block not found: %s %v", expBlock.Type(), expBlock.Labels())
			return false
		}
	}

	return true
}

// blocksMatch checks if two blocks have the same type and labels
func blocksMatch(a, b *hclwrite.Block) bool {
	if a.Type() != b.Type() {
		return false
	}

	aLabels := a.Labels()
	bLabels := b.Labels()

	if len(aLabels) != len(bLabels) {
		return false
	}

	for i, label := range aLabels {
		if label != bLabels[i] {
			return false
		}
	}

	return true
}

// compareBlockAttributes compares attributes of two blocks, ignoring order
func compareBlockAttributes(t *testing.T, expected, actual *hclwrite.Block) bool {
	expectedAttrs := expected.Body().Attributes()
	actualAttrs := actual.Body().Attributes()

	// Check all expected attributes exist in actual
	for name, expAttr := range expectedAttrs {
		actAttr, exists := actualAttrs[name]
		if !exists {
			t.Errorf("Missing attribute '%s' in block %s %v", name, expected.Type(), expected.Labels())
			return false
		}

		// Compare attribute values as strings (tokens)
		expTokens := string(expAttr.Expr().BuildTokens(nil).Bytes())
		actTokens := string(actAttr.Expr().BuildTokens(nil).Bytes())

		// Normalize whitespace for comparison
		expNorm := normalizeWhitespace(expTokens)
		actNorm := normalizeWhitespace(actTokens)

		if expNorm != actNorm {
			t.Errorf("Attribute '%s' value mismatch in block %s %v:\n  Expected: %s\n  Got: %s",
				name, expected.Type(), expected.Labels(), expNorm, actNorm)
			return false
		}
	}

	// Check for unexpected attributes in actual
	for name := range actualAttrs {
		if _, exists := expectedAttrs[name]; !exists {
			t.Errorf("Unexpected attribute '%s' in block %s %v", name, actual.Type(), actual.Labels())
			return false
		}
	}

	return true
}
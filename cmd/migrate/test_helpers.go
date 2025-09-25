package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/transformations"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/require"
)

// TestCase represents a single test case for HCL transformation
type TestCase struct {
	Name     string
	Config   string
	State    string
	Expected []string
}

// RunTransformationTests executes a series of test cases for HCL transformation
func RunTransformationTests(t *testing.T, tests []TestCase, transformFunc func([]byte, string) ([]byte, error)) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Parse the input
			file, diags := hclwrite.ParseConfig([]byte(tt.Config), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors(), "Failed to parse input config: %s", diags)

			// Transform the file
			result, err := transformFunc(file.Bytes(), "test.tf")
			require.NoError(t, err, "Transformation failed")

			// Check each expected output fragment
			resultString := string(result)
			for _, expected := range tt.Expected {
				assertHCLContains(t, resultString, expected)
			}
		})
	}
}

// assertHCLContains checks if the expected HCL fragment is present in the actual output
func assertHCLContains(t *testing.T, actual, expected string) {
	t.Helper()

	// Trim leading/trailing whitespace from expected as it's usually not significant
	expected = strings.TrimSpace(expected)

	// First try a smart comparison that handles ordering differences
	if hclFragmentMatches(actual, expected) {
		return
	}

	// If that fails, provide a helpful error message
	actualFormatted := formatHCLForDiff(actual)
	expectedFormatted := formatHCLForDiff(expected)

	t.Errorf("Expected HCL fragment not found in output.\n\nExpected fragment:\n%s\n\nActual output:\n%s\n\nDiff:\n%s",
		expectedFormatted,
		actualFormatted,
		cmp.Diff(expectedFormatted, actualFormatted))
}

// hclFragmentMatches checks if an HCL fragment exists in the actual output
func hclFragmentMatches(actual, expected string) bool {
	// First normalize whitespace aggressively for comparison
	actualNorm := normalizeHCLContent(actual)
	expectedNorm := normalizeHCLContent(expected)

	// Check if the normalized fragment exists
	if strings.Contains(actualNorm, expectedNorm) {
		return true
	}

	// For more complex cases, try semantic matching
	return hclSemanticMatch(actual, expected)
}

// normalizeHCLContent aggressively normalizes HCL for comparison
func normalizeHCLContent(s string) string {
	// Aggressively normalize to handle formatting differences
	// Remove all newlines and extra spaces to compare structure only

	// Replace newlines with spaces
	s = strings.ReplaceAll(s, "\n", " ")

	// Collapse multiple spaces
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}

	// Remove spaces around punctuation
	replacements := []struct{ old, new string }{
		{" = ", "="},
		{" {", "{"},
		{"{ ", "{"},
		{" }", "}"},
		{"} ", "}"},
		{" [", "["},
		{"[ ", "["},
		{" ]", "]"},
		{"] ", "]"},
		{", ", ","},
		{" ,", ","},
	}

	for _, r := range replacements {
		s = strings.ReplaceAll(s, r.old, r.new)
	}

	return strings.TrimSpace(s)
}

// formatHCLForDiff formats HCL nicely for diff output
func formatHCLForDiff(hclContent string) string {
	// Try to parse and reformat for consistent display
	if file, diags := hclwrite.ParseConfig([]byte(hclContent), "", hcl.InitialPos); !diags.HasErrors() {
		return string(hclwrite.Format(file.Bytes()))
	}
	// If it's a fragment, just normalize whitespace
	return normalizeHCLWhitespace(hclContent)
}

// normalizeHCLWhitespace does minimal whitespace normalization
func normalizeHCLWhitespace(s string) string {
	// Trim lines and remove extra spaces
	lines := strings.Split(s, "\n")
	var normalized []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			// Collapse multiple spaces to single space
			normalized = append(normalized, strings.Join(strings.Fields(line), " "))
		}
	}
	return strings.Join(normalized, "\n")
}

// hclSemanticMatch attempts to match HCL semantically, ignoring formatting
func hclSemanticMatch(actual, expected string) bool {
	// Special handling for known patterns that need order-agnostic comparison
	if strings.Contains(expected, "include =") ||
		strings.Contains(expected, "exclude =") ||
		strings.Contains(expected, "require =") {
		return compareAccessPolicyRules(actual, expected)
	}

	// Try to parse both as complete HCL
	actualFile, aDiags := hclwrite.ParseConfig([]byte(actual), "", hcl.InitialPos)
	expectedFile, eDiags := hclwrite.ParseConfig([]byte(expected), "", hcl.InitialPos)

	if !aDiags.HasErrors() && !eDiags.HasErrors() {
		// Both parsed successfully - compare normalized forms
		actualNorm := string(hclwrite.Format(actualFile.Bytes()))
		expectedNorm := string(hclwrite.Format(expectedFile.Bytes()))
		return actualNorm == expectedNorm
	}

	return false
}

// compareAccessPolicyRules handles special comparison for access policy rules
func compareAccessPolicyRules(actual, expected string) bool {
	// This is specifically for access policy include/exclude/require arrays
	// where order doesn't matter semantically

	// Extract all the important tokens/values
	actualTokens := extractPolicyTokens(actual)
	expectedTokens := extractPolicyTokens(expected)

	// Check if all expected tokens are present
	for token := range expectedTokens {
		if _, found := actualTokens[token]; !found {
			return false
		}
	}

	return true
}

// extractPolicyTokens extracts significant tokens for comparison
func extractPolicyTokens(hcl string) map[string]bool {
	tokens := make(map[string]bool)

	// Extract quoted strings (values we care about)
	inQuote := false
	current := ""
	for i := 0; i < len(hcl); i++ {
		ch := hcl[i]
		if !inQuote && ch == '"' {
			inQuote = true
			current = ""
		} else if inQuote && ch == '"' && (i == 0 || hcl[i-1] != '\\') {
			if current != "" {
				tokens[current] = true
			}
			inQuote = false
		} else if inQuote {
			current += string(ch)
		}
	}

	// Also extract key identifiers
	for _, word := range strings.Fields(hcl) {
		if strings.Contains(word, "=") || strings.Contains(word, "{") || strings.Contains(word, "}") {
			continue
		}
		if len(word) > 2 && !strings.HasPrefix(word, "\"") {
			tokens[word] = true
		}
	}

	return tokens
}

// StateTestCase represents a single test case for state transformation
type StateTestCase struct {
	Name     string
	Input    string // Input JSON
	Expected string // Expected JSON output
}

// RunStateTransformationTests executes state transformation tests for a specific transform function
func RunStateTransformationTests(t *testing.T, tests []StateTestCase, transformFunc func(map[string]interface{})) {
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// Parse input JSON
			var inputMap map[string]interface{}
			err := json.Unmarshal([]byte(tc.Input), &inputMap)
			require.NoError(t, err, "Failed to parse input JSON")

			// Apply transformation
			transformFunc(inputMap)

			// Compare with expected
			assertJSONEqual(t, tc.Expected, inputMap)
		})
	}
}

// RunFullStateTransformationTests executes tests for the full state JSON transformation
func RunFullStateTransformationTests(t *testing.T, tests []StateTestCase) {
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// Transform using the JSON-based function
			result, err := transformStateJSON([]byte(tc.Input))
			require.NoError(t, err, "Failed to transform state")

			// Parse result and compare
			var actualData interface{}
			err = json.Unmarshal(result, &actualData)
			require.NoError(t, err, "Failed to parse transformed JSON")

			assertJSONEqual(t, tc.Expected, actualData)
		})
	}
}

// assertJSONEqual compares expected JSON with actual data
func assertJSONEqual(t *testing.T, expectedJSON string, actualData interface{}) {
	t.Helper()

	// Parse expected JSON
	var expectedData interface{}
	err := json.Unmarshal([]byte(expectedJSON), &expectedData)
	require.NoError(t, err, "Failed to parse expected JSON")

	// Normalize both for comparison (handles map ordering)
	expectedNorm := normalizeJSONValue(expectedData)
	actualNorm := normalizeJSONValue(actualData)

	// Compare with helpful diff
	if !reflect.DeepEqual(expectedNorm, actualNorm) {
		// Pretty print for better error messages
		expectedPretty, _ := json.MarshalIndent(expectedNorm, "", "  ")
		actualPretty, _ := json.MarshalIndent(actualNorm, "", "  ")

		diff := cmp.Diff(string(expectedPretty), string(actualPretty))
		t.Errorf("JSON mismatch (-expected +actual):\n%s", diff)
	}
}

// normalizeJSONValue recursively normalizes JSON values for comparison
func normalizeJSONValue(v interface{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		// Sort map keys for consistent comparison
		normalized := make(map[string]interface{})
		for k, v := range val {
			normalized[k] = normalizeJSONValue(v)
		}
		return normalized
	case []interface{}:
		// Normalize array elements
		normalized := make([]interface{}, len(val))
		for i, elem := range val {
			normalized[i] = normalizeJSONValue(elem)
		}
		return normalized
	default:
		return val
	}
}

// normalizeWhitespace normalizes whitespace in a string for simple comparison
// This is kept for backwards compatibility with existing tests
func normalizeWhitespace(s string) string {
	// Collapse multiple spaces to single space
	s = strings.Join(strings.Fields(s), " ")
	// Remove spaces around brackets and braces
	replacements := []struct{ old, new string }{
		{" {", "{"}, {"{ ", "{"},
		{" }", "}"}, {"} ", "}"},
		{" [", "["}, {"[ ", "["},
		{" ]", "]"}, {"] ", "]"},
	}
	for _, r := range replacements {
		s = strings.ReplaceAll(s, r.old, r.new)
	}
	return strings.TrimSpace(s)
}

// transformFileWithYAML applies both YAML and Go transformations for testing
func transformFileWithYAML(content []byte, filename string) ([]byte, error) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "migrate-test-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	// Write the content to a temporary file
	tempFile := filepath.Join(tempDir, filename)
	if err := os.WriteFile(tempFile, content, 0644); err != nil {
		return nil, err
	}

	// Apply YAML transformations using the actual transformer
	if err := applyYAMLTransformationsToFile(tempFile); err != nil {
		return nil, err
	}

	// Read the transformed content
	transformed, err := os.ReadFile(tempFile)
	if err != nil {
		return nil, err
	}

	// Then apply Go transformations
	return transformFileDefault(transformed, filename)
}

// applyYAMLTransformationsToFile applies all YAML transformations to a file
func applyYAMLTransformationsToFile(filePath string) error {
	// Get the path to the YAML config files
	configDir := filepath.Join(filepath.Dir(filePath), "..", "transformations", "config")
	
	// If the config directory doesn't exist relative to the test, try finding it from the module root
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		// Try to find the config directory from the current working directory
		cwd, _ := os.Getwd()
		configDir = filepath.Join(cwd, "transformations", "config")
	}

	// List of transformation configs to apply in order
	transformationConfigs := []string{
		"cloudflare_terraform_v5_resource_renames_configuration.yaml",
		"cloudflare_terraform_v5_block_to_attribute_configuration.yaml",
		"cloudflare_terraform_v5_attribute_renames_configuration.yaml",
	}

	for _, configFile := range transformationConfigs {
		configPath := filepath.Join(configDir, configFile)
		
		// Check if the config file exists
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			// Skip if config doesn't exist (tests may not need all transformations)
			continue
		}

		// Create the appropriate transformer
		var transformer interface {
			TransformFile(input, output string) error
		}

		if strings.Contains(configFile, "resource_renames") {
			t, err := transformations.NewResourceRenameTransformer(configPath)
			if err != nil {
				return err
			}
			transformer = t
		} else if strings.Contains(configFile, "block_to_attribute") || strings.Contains(configFile, "attribute_renames") {
			t, err := transformations.NewHCLTransformer(configPath)
			if err != nil {
				return err
			}
			transformer = t
		}

		if transformer != nil {
			// Apply the transformation in-place
			if err := transformer.TransformFile(filePath, filePath); err != nil {
				return err
			}
		}
	}

	return nil
}

// RunFullTransformationTests executes test cases with both YAML and Go transformations
func RunFullTransformationTests(t *testing.T, tests []TestCase) {
	RunTransformationTests(t, tests, transformFileWithYAML)
}

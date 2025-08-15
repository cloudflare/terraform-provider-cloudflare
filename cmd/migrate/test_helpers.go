package main

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
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
			assert.False(t, diags.HasErrors())

			// Transform the file
			result, err := transformFunc(file.Bytes(), "test.tf")
			assert.NoError(t, err)
			resultString := string(result)

			// Check each expected output
			for _, expected := range tt.Expected {
				normalizedExpected := string(hclwrite.Format([]byte(expected)))
				assert.Contains(t, resultString, normalizedExpected, "== Actual output ==\n%s\n== Expected to contain ==\n%s\n", resultString, normalizedExpected)
			}
		})
	}
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
			// Parse input JSON to map
			var inputMap map[string]interface{}
			if err := json.Unmarshal([]byte(tc.Input), &inputMap); err != nil {
				t.Fatalf("Failed to parse input JSON: %v", err)
			}

			// Apply transformation
			transformFunc(inputMap)

			// Convert back to JSON for comparison
			actualJSON, err := json.Marshal(inputMap)
			if err != nil {
				t.Fatalf("Failed to marshal result: %v", err)
			}

			// Parse expected JSON for normalized comparison
			var expectedMap map[string]interface{}
			if err := json.Unmarshal([]byte(tc.Expected), &expectedMap); err != nil {
				t.Fatalf("Failed to parse expected JSON: %v", err)
			}
			expectedJSON, err := json.Marshal(expectedMap)
			if err != nil {
				t.Fatalf("Failed to marshal expected: %v", err)
			}

			// Compare normalized JSON
			if string(actualJSON) != string(expectedJSON) {
				// For better error output, pretty print both
				actualPretty, _ := json.MarshalIndent(inputMap, "", "  ")
				expectedPretty, _ := json.MarshalIndent(expectedMap, "", "  ")
				t.Errorf("Transformation failed\nExpected:\n%s\n\nGot:\n%s", expectedPretty, actualPretty)
			}
		})
	}
}

// RunFullStateTransformationTests executes tests for the full state JSON transformation
func RunFullStateTransformationTests(t *testing.T, tests []StateTestCase) {
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// Transform using the JSON-based function
			result, err := transformStateJSON([]byte(tc.Input))
			if err != nil {
				t.Fatalf("Failed to transform state: %v", err)
			}

			// Compare the JSON structures semantically (ignoring formatting and key order)
			if !compareJSON(t, string(result), tc.Expected) {
				// For better error output, pretty print both
				var resultMap, expectedMap interface{}
				json.Unmarshal(result, &resultMap)
				json.Unmarshal([]byte(tc.Expected), &expectedMap)

				resultPretty, _ := json.MarshalIndent(resultMap, "", "  ")
				expectedPretty, _ := json.MarshalIndent(expectedMap, "", "  ")

				t.Errorf("Transformation failed\nExpected:\n%s\n\nGot:\n%s", expectedPretty, resultPretty)
			}
		})
	}
}

// compareJSON compares two JSON strings semantically, ignoring formatting and key order
func compareJSON(t *testing.T, actual, expected string) bool {
	// Parse both JSON strings into generic interfaces
	var actualData, expectedData interface{}

	if err := json.Unmarshal([]byte(actual), &actualData); err != nil {
		t.Fatalf("Failed to parse actual JSON: %v", err)
		return false
	}

	if err := json.Unmarshal([]byte(expected), &expectedData); err != nil {
		t.Fatalf("Failed to parse expected JSON: %v", err)
		return false
	}

	// Marshal both to normalize formatting (this handles key ordering)
	actualNorm, _ := json.Marshal(actualData)
	expectedNorm, _ := json.Marshal(expectedData)

	return string(actualNorm) == string(expectedNorm)
}

// normalizeWhitespace normalizes whitespace in a string for comparison
func normalizeWhitespace(s string) string {
	// Replace multiple spaces with single space
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")

	// Trim leading and trailing whitespace
	s = strings.TrimSpace(s)

	// Remove spaces around certain punctuation
	s = strings.ReplaceAll(s, " {", "{")
	s = strings.ReplaceAll(s, " }", "}")
	s = strings.ReplaceAll(s, " [", "[")
	s = strings.ReplaceAll(s, " ]", "]")
	s = strings.ReplaceAll(s, "{ ", "{")
	s = strings.ReplaceAll(s, "} ", "}")
	s = strings.ReplaceAll(s, "[ ", "[")
	s = strings.ReplaceAll(s, "] ", "]")

	return s
}

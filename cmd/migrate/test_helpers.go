package main

import (
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

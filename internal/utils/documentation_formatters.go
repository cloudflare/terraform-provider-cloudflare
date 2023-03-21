package utils

import (
	"fmt"
	"strings"
)

// RenderAvailableDocumentationValuesStringSlice takes a slice of strings and
// formats it for documentation output use.
//
// Example: ["foo", "bar", "baz"] -> `foo`, `bar`, `baz`.
func RenderAvailableDocumentationValuesStringSlice(s []string) string {
	output := ""
	if s != nil && len(s) > 0 {
		values := make([]string, len(s))
		for i, c := range s {
			values[i] = fmt.Sprintf("`%s`", c)
		}
		output = fmt.Sprintf("Available values: %s", strings.Join(values, ", "))
	}
	return output
}

// RenderAvailableDocumentationValuesIntSlice takes a slice of ints and
// formats it for documentation output use.
//
// Example: [1, 2, 3] -> `1`, `2`, `3`.
func RenderAvailableDocumentationValuesIntSlice(s []int) string {
	output := ""
	if s != nil && len(s) > 0 {
		values := make([]string, len(s))
		for i, c := range s {
			values[i] = fmt.Sprintf("`%d`", c)
		}
		output = fmt.Sprintf("Available values: %s", strings.Join(values, ", "))
	}
	return output
}

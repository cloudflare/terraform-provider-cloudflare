package transformations

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// CompareHCL compares two HCL strings semantically, ignoring attribute ordering
func CompareHCL(expected, actual string) (bool, string, error) {
	// Parse both HCL strings
	expectedFile, diags := hclwrite.ParseConfig([]byte(expected), "expected.hcl", hcl.InitialPos)
	if diags.HasErrors() {
		return false, "", fmt.Errorf("failed to parse expected HCL: %s", diags.Error())
	}

	actualFile, diags := hclwrite.ParseConfig([]byte(actual), "actual.hcl", hcl.InitialPos)
	if diags.HasErrors() {
		return false, "", fmt.Errorf("failed to parse actual HCL: %s", diags.Error())
	}

	// Compare the bodies
	diff := compareBody(expectedFile.Body(), actualFile.Body(), "")
	if diff != "" {
		return false, diff, nil
	}

	return true, "", nil
}

// compareBody compares two HCL bodies semantically
func compareBody(expected, actual *hclwrite.Body, path string) string {
	var diffs []string

	// Compare attributes
	expectedAttrs := getAttributeMap(expected)
	actualAttrs := getAttributeMap(actual)

	// Check for missing or extra attributes
	for name := range expectedAttrs {
		if _, exists := actualAttrs[name]; !exists {
			diffs = append(diffs, fmt.Sprintf("%s: missing attribute '%s'", path, name))
		}
	}

	for name := range actualAttrs {
		if _, exists := expectedAttrs[name]; !exists {
			diffs = append(diffs, fmt.Sprintf("%s: unexpected attribute '%s'", path, name))
		}
	}

	// Compare attribute values
	for name, expectedAttr := range expectedAttrs {
		if actualAttr, exists := actualAttrs[name]; exists {
			attrPath := fmt.Sprintf("%s.%s", path, name)
			if path == "" {
				attrPath = name
			}
			
			if diff := compareAttributeValues(expectedAttr, actualAttr, attrPath); diff != "" {
				diffs = append(diffs, diff)
			}
		}
	}

	// Compare blocks
	expectedBlocks := getBlocksByType(expected)
	actualBlocks := getBlocksByType(actual)

	// Check for missing or extra block types
	for blockType := range expectedBlocks {
		if _, exists := actualBlocks[blockType]; !exists {
			diffs = append(diffs, fmt.Sprintf("%s: missing block type '%s'", path, blockType))
		}
	}

	for blockType := range actualBlocks {
		if _, exists := expectedBlocks[blockType]; !exists {
			diffs = append(diffs, fmt.Sprintf("%s: unexpected block type '%s'", path, blockType))
		}
	}

	// Compare blocks of each type
	for blockType, expectedBlockList := range expectedBlocks {
		if actualBlockList, exists := actualBlocks[blockType]; exists {
			blockPath := fmt.Sprintf("%s.%s", path, blockType)
			if path == "" {
				blockPath = blockType
			}

			if diff := compareBlockLists(expectedBlockList, actualBlockList, blockPath); diff != "" {
				diffs = append(diffs, diff)
			}
		}
	}

	if len(diffs) > 0 {
		return strings.Join(diffs, "\n")
	}
	return ""
}

// getAttributeMap returns a map of attribute names to attributes
func getAttributeMap(body *hclwrite.Body) map[string]*hclwrite.Attribute {
	attrs := make(map[string]*hclwrite.Attribute)
	for name, attr := range body.Attributes() {
		attrs[name] = attr
	}
	return attrs
}

// getBlocksByType groups blocks by their type
func getBlocksByType(body *hclwrite.Body) map[string][]*hclwrite.Block {
	blocks := make(map[string][]*hclwrite.Block)
	for _, block := range body.Blocks() {
		blockType := block.Type()
		blocks[blockType] = append(blocks[blockType], block)
	}
	return blocks
}

// compareAttributeValues compares two attribute values as strings
func compareAttributeValues(expected, actual *hclwrite.Attribute, path string) string {
	// Get the tokens for comparison
	expectedTokens := string(expected.Expr().BuildTokens(nil).Bytes())
	actualTokens := string(actual.Expr().BuildTokens(nil).Bytes())

	// Normalize whitespace for comparison
	expectedNorm := normalizeWhitespace(expectedTokens)
	actualNorm := normalizeWhitespace(actualTokens)

	if expectedNorm != actualNorm {
		return fmt.Sprintf("%s: value mismatch\n  expected: %s\n  actual:   %s", path, expectedNorm, actualNorm)
	}
	return ""
}

// compareBlockLists compares two lists of blocks
func compareBlockLists(expected, actual []*hclwrite.Block, path string) string {
	if len(expected) != len(actual) {
		return fmt.Sprintf("%s: block count mismatch (expected %d, got %d)", path, len(expected), len(actual))
	}

	// For simplicity, compare blocks in order
	// This could be enhanced to match blocks by labels or content
	for i := 0; i < len(expected); i++ {
		blockPath := fmt.Sprintf("%s[%d]", path, i)
		
		// Compare labels
		expectedLabels := expected[i].Labels()
		actualLabels := actual[i].Labels()
		
		if len(expectedLabels) != len(actualLabels) {
			return fmt.Sprintf("%s: label count mismatch", blockPath)
		}
		
		for j := range expectedLabels {
			if expectedLabels[j] != actualLabels[j] {
				return fmt.Sprintf("%s: label mismatch at position %d", blockPath, j)
			}
		}
		
		// Compare block bodies
		if diff := compareBody(expected[i].Body(), actual[i].Body(), blockPath); diff != "" {
			return diff
		}
	}

	return ""
}

// CompareHCLResources compares two HCL resources, focusing on resource type and name
func CompareHCLResources(expected, actual string) (bool, string, error) {
	// Parse both HCL strings
	expectedFile, diags := hclwrite.ParseConfig([]byte(expected), "expected.hcl", hcl.InitialPos)
	if diags.HasErrors() {
		return false, "", fmt.Errorf("failed to parse expected HCL: %s", diags.Error())
	}

	actualFile, diags := hclwrite.ParseConfig([]byte(actual), "actual.hcl", hcl.InitialPos)
	if diags.HasErrors() {
		return false, "", fmt.Errorf("failed to parse actual HCL: %s", diags.Error())
	}

	// Get resource blocks
	expectedResources := findResourceBlocks(expectedFile.Body())
	actualResources := findResourceBlocks(actualFile.Body())

	if len(expectedResources) != len(actualResources) {
		return false, fmt.Sprintf("resource count mismatch: expected %d, got %d", len(expectedResources), len(actualResources)), nil
	}

	// Compare each resource
	for i, expectedRes := range expectedResources {
		actualRes := actualResources[i]
		
		// Compare labels (resource type and name)
		expectedLabels := expectedRes.Labels()
		actualLabels := actualRes.Labels()
		
		if len(expectedLabels) < 2 || len(actualLabels) < 2 {
			return false, "invalid resource block: missing labels", nil
		}
		
		if expectedLabels[0] != actualLabels[0] {
			return false, fmt.Sprintf("resource type mismatch: expected %s, got %s", expectedLabels[0], actualLabels[0]), nil
		}
		
		if expectedLabels[1] != actualLabels[1] {
			return false, fmt.Sprintf("resource name mismatch: expected %s, got %s", expectedLabels[1], actualLabels[1]), nil
		}
		
		// Compare resource body semantically
		if diff := compareResourceBody(expectedRes.Body(), actualRes.Body()); diff != "" {
			return false, fmt.Sprintf("resource %s.%s: %s", expectedLabels[0], expectedLabels[1], diff), nil
		}
	}

	return true, "", nil
}

// findResourceBlocks finds all resource blocks in a body
func findResourceBlocks(body *hclwrite.Body) []*hclwrite.Block {
	var resources []*hclwrite.Block
	for _, block := range body.Blocks() {
		if block.Type() == "resource" {
			resources = append(resources, block)
		}
	}
	return resources
}

// compareResourceBody compares resource bodies with special handling for common patterns
func compareResourceBody(expected, actual *hclwrite.Body) string {
	// Get all attributes from both bodies
	expectedAttrs := make(map[string]string)
	actualAttrs := make(map[string]string)
	
	for name, attr := range expected.Attributes() {
		expectedAttrs[name] = normalizeAttributeValue(attr)
	}
	
	for name, attr := range actual.Attributes() {
		actualAttrs[name] = normalizeAttributeValue(attr)
	}
	
	// Check for missing attributes
	var missing []string
	for name := range expectedAttrs {
		if _, exists := actualAttrs[name]; !exists {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		return fmt.Sprintf("missing attributes: %s", strings.Join(missing, ", "))
	}
	
	// Check for extra attributes
	var extra []string
	for name := range actualAttrs {
		if _, exists := expectedAttrs[name]; !exists {
			extra = append(extra, name)
		}
	}
	if len(extra) > 0 {
		sort.Strings(extra)
		return fmt.Sprintf("unexpected attributes: %s", strings.Join(extra, ", "))
	}
	
	// Compare attribute values
	for name, expectedVal := range expectedAttrs {
		actualVal := actualAttrs[name]
		if expectedVal != actualVal {
			return fmt.Sprintf("attribute '%s' mismatch:\n  expected: %s\n  actual:   %s", name, expectedVal, actualVal)
		}
	}
	
	// Compare nested blocks
	expectedBlocks := getBlocksByType(expected)
	actualBlocks := getBlocksByType(actual)
	
	// Compare block types
	for blockType, expectedBlockList := range expectedBlocks {
		actualBlockList, exists := actualBlocks[blockType]
		if !exists {
			return fmt.Sprintf("missing block type: %s", blockType)
		}
		
		if len(expectedBlockList) != len(actualBlockList) {
			return fmt.Sprintf("block '%s' count mismatch: expected %d, got %d", 
				blockType, len(expectedBlockList), len(actualBlockList))
		}
		
		// For each block, recursively compare
		for i := 0; i < len(expectedBlockList); i++ {
			if diff := compareResourceBody(expectedBlockList[i].Body(), actualBlockList[i].Body()); diff != "" {
				return fmt.Sprintf("in block '%s[%d]': %s", blockType, i, diff)
			}
		}
	}
	
	// Check for extra block types
	for blockType := range actualBlocks {
		if _, exists := expectedBlocks[blockType]; !exists {
			return fmt.Sprintf("unexpected block type: %s", blockType)
		}
	}
	
	return ""
}

// normalizeAttributeValue normalizes an attribute value for comparison
func normalizeAttributeValue(attr *hclwrite.Attribute) string {
	tokens := attr.Expr().BuildTokens(nil).Bytes()
	return normalizeWhitespace(string(tokens))
}
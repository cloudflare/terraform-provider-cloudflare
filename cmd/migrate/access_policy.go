package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isAccessPolicyResource checks if a block is a cloudflare_zero_trust_access_policy resource
// (grit has already renamed from cloudflare_access_policy)
func isAccessPolicyResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		block.Labels()[0] == "cloudflare_zero_trust_access_policy"
}

// transformAccessPolicyBlock transforms include/exclude/require attributes from v4 to v5 format
// After Grit transforms, we still need to handle complex value transformations
func transformAccessPolicyBlock(block *hclwrite.Block) {
	// Process include, exclude, and require attributes (grit has already converted them to lists)
	conditionAttributes := []string{"include", "exclude", "require"}

	for _, attrName := range conditionAttributes {
		attr := block.Body().GetAttribute(attrName)
		if attr == nil {
			continue // Attribute doesn't exist
		}

		// Get the current tokens and transform them
		tokens := attr.Expr().BuildTokens(nil)
		transformedTokens := transformConditionListTokens(tokens)
		if transformedTokens != nil {
			block.Body().SetAttributeRaw(attrName, transformedTokens)
		}
	}
}

// transformConditionListTokens transforms the condition list tokens using string manipulation
// This is simpler and more reliable than complex token parsing
func transformConditionListTokens(tokens hclwrite.Tokens) hclwrite.Tokens {
	// Convert tokens to string for easier manipulation
	var content strings.Builder
	for _, token := range tokens {
		content.Write(token.Bytes)
	}
	
	originalContent := content.String()
	transformedContent := transformConditionString(originalContent)
	
	// If no transformation was applied, return nil to keep original
	if originalContent == transformedContent {
		return nil
	}
	
	// Parse the transformed content back into tokens
	// This is a simplified approach - create a single token with the transformed content
	return hclwrite.Tokens{
		&hclwrite.Token{
			Type:  hclsyntax.TokenEOF, // Using a generic token type
			Bytes: []byte(transformedContent),
		},
	}
}

// transformConditionString applies string-based transformations to condition content
// Converts v4 condition blocks to v5 condition objects, creating separate objects for each condition type
func transformConditionString(content string) string {
	// Check if this looks like a condition list (starts with [ and contains condition blocks)
	if !strings.Contains(content, "{") {
		return content
	}
	
	var conditionObjects []string
	
	// Define all condition mappings
	listAttrs := map[string]string{
		"email":         "email",
		"email_domain":  "domain", 
		"email_list":    "id",
		"ip":            "ip",
		"ip_list":       "id",
		"service_token": "token_id",
		"group":         "id",
		"geo":           "country_code",
		"login_method":  "id",
		"device_posture": "integration_uid",
	}
	
	// Preserve the original order from the v4 configuration by finding the actual position of each attribute
	
	// Find all attributes with their positions in the source
	type attrPosition struct {
		name string
		pos  int
	}
	
	var attrPositions []attrPosition
	allPossibleAttrs := []string{
		"email", "email_domain", "ip", "everyone", "any_valid_service_token", "geo", 
		"certificate", "email_list", "ip_list", "service_token", "group", "login_method", 
		"device_posture", "common_name", "auth_method",
	}
	
	// Find the position of each attribute in the source content
	for _, attr := range allPossibleAttrs {
		var pattern *regexp.Regexp
		var found bool
		
		if attr == "everyone" || attr == "certificate" || attr == "any_valid_service_token" {
			pattern = regexp.MustCompile(`\s*` + attr + `\s*=\s*true`)
		} else if _, exists := listAttrs[attr]; exists {
			// Try list pattern first
			listPattern := regexp.MustCompile(`\s*` + attr + `\s*=\s*\[\s*([^\]]+)\s*\]`)
			if listPattern.MatchString(content) {
				pattern = listPattern
				found = true
			} else {
				// Try single pattern
				pattern = regexp.MustCompile(`\s*` + attr + `\s*=\s*"([^"]+)"`)
			}
		} else if attr == "common_name" || attr == "auth_method" {
			pattern = regexp.MustCompile(`\s*` + attr + `\s*=\s*"([^"]+)"`)
		}
		
		if pattern != nil && (found || pattern.MatchString(content)) {
			loc := pattern.FindStringIndex(content)
			if loc != nil {
				attrPositions = append(attrPositions, attrPosition{name: attr, pos: loc[0]})
			}
		}
	}
	
	// Sort by position in source to preserve original order
	sort.Slice(attrPositions, func(i, j int) bool {
		return attrPositions[i].pos < attrPositions[j].pos
	})
	
	// Process attributes in the order they appear in the source
	for _, attrPos := range attrPositions {
		attr := attrPos.name
		// Handle boolean attributes
		if attr == "everyone" || attr == "certificate" || attr == "any_valid_service_token" {
			pattern := regexp.MustCompile(`\s*` + attr + `\s*=\s*true`)
			if pattern.MatchString(content) {
				conditionObjects = append(conditionObjects, "    { "+attr+" = {} }")
			}
			continue
		}
		
		// Handle list attributes
		if nestedKey, exists := listAttrs[attr]; exists {
			// Pattern to match: attr = ["value1", "value2", ...]
			listPattern := regexp.MustCompile(`\s*` + attr + `\s*=\s*\[\s*([^\]]+)\s*\]`)
			matches := listPattern.FindStringSubmatch(content)
			if len(matches) > 1 {
				// Parse the list values
				valueString := matches[1]
				// Match quoted strings
				valuePattern := regexp.MustCompile(`"([^"]+)"`)
				values := valuePattern.FindAllStringSubmatch(valueString, -1)
				for _, valueMatch := range values {
					if len(valueMatch) > 1 {
						value := valueMatch[1]
						conditionObjects = append(conditionObjects, fmt.Sprintf("    { %s = { %s = \"%s\" } }", attr, nestedKey, value))
					}
				}
			}
			
			// Also handle single values: attr = "value"
			singlePattern := regexp.MustCompile(`\s*` + attr + `\s*=\s*"([^"]+)"`)
			singleMatches := singlePattern.FindStringSubmatch(content)
			if len(singleMatches) > 1 {
				value := singleMatches[1]
				conditionObjects = append(conditionObjects, fmt.Sprintf("    { %s = { %s = \"%s\" } }", attr, nestedKey, value))
			}
			continue
		}
		
		// Handle string attributes
		if attr == "common_name" || attr == "auth_method" {
			pattern := regexp.MustCompile(`\s*` + attr + `\s*=\s*"([^"]+)"`)
			matches := pattern.FindStringSubmatch(content)
			if len(matches) > 1 && matches[1] != "" {
				value := matches[1]
				conditionObjects = append(conditionObjects, fmt.Sprintf("    { %s = { %s = \"%s\" } }", attr, attr, value))
			}
		}
	}
	
	// If no conditions found, return original
	if len(conditionObjects) == 0 {
		return content
	}
	
	// Build the final result as a list of condition objects
	result := "[\n" + strings.Join(conditionObjects, ",\n") + "\n  ]"
	return result
}
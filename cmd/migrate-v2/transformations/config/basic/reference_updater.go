package basic

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// ReferenceUpdater creates a transformer that updates resource references
//
// Example YAML configuration:
//   reference_mappings:
//     cloudflare_firewall_rule: cloudflare_ruleset
//     cloudflare_access_application: cloudflare_zero_trust_access_application
//
// Transforms:
//   resource "example" "test" {
//     firewall_id = cloudflare_firewall_rule.main.id
//     app_id = "${cloudflare_access_application.webapp.id}"
//     depends_on = [cloudflare_firewall_rule.main]
//   }
//
// Into:
//   resource "example" "test" {
//     firewall_id = cloudflare_ruleset.main.id
//     app_id = "${cloudflare_zero_trust_access_application.webapp.id}"
//     depends_on = [cloudflare_ruleset.main]
//   }
func ReferenceUpdater(mappings map[string]string) TransformFunc {
	if len(mappings) == 0 {
		return func(block *hclwrite.Block, ctx *TransformContext) error {
			return nil
		}
	}

	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		
		// Process all attributes
		for name, attr := range body.Attributes() {
			tokens := attr.Expr().BuildTokens(nil)
			value := string(tokens.Bytes())
			
			// Check if value contains references
			if containsReference(value) {
				newValue := updateReferences(value, mappings)
				if newValue != value {
					// Parse and set the updated expression
					setUpdatedExpression(body, name, newValue)
				}
			}
		}
		
		// Process nested blocks
		for _, nestedBlock := range body.Blocks() {
			if err := updateBlockReferences(nestedBlock, mappings); err != nil {
				return err
			}
		}
		
		return nil
	}
}

// updateBlockReferences recursively updates references in nested blocks
func updateBlockReferences(block *hclwrite.Block, mappings map[string]string) error {
	body := block.Body()
	
	// Process attributes
	for name, attr := range body.Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		value := string(tokens.Bytes())
		
		if containsReference(value) {
			newValue := updateReferences(value, mappings)
			if newValue != value {
				setUpdatedExpression(body, name, newValue)
			}
		}
	}
	
	// Process nested blocks recursively
	for _, nestedBlock := range body.Blocks() {
		if err := updateBlockReferences(nestedBlock, mappings); err != nil {
			return err
		}
	}
	
	return nil
}

// containsReference checks if a value contains Terraform references
func containsReference(value string) bool {
	// Check for common reference patterns
	patterns := []string{
		`\.id\b`,           // .id references
		`\.arn\b`,          // .arn references
		`data\.`,           // data source references
		`var\.`,            // variable references
		`local\.`,          // local references
		`module\.`,         // module references
		`\$\{.*\}`,         // interpolation syntax
		`cloudflare_[a-z_]+\.`, // cloudflare resource references
	}
	
	for _, pattern := range patterns {
		if matched, _ := regexp.MatchString(pattern, value); matched {
			return true
		}
	}
	
	return false
}

// updateReferences updates all references in a value based on mappings
func updateReferences(value string, mappings map[string]string) string {
	result := value
	
	// Sort mappings by length (longest first) to avoid partial replacements
	var sortedMappings []struct{ from, to string }
	for from, to := range mappings {
		sortedMappings = append(sortedMappings, struct{ from, to string }{from, to})
	}
	// Sort by length descending
	for i := 0; i < len(sortedMappings); i++ {
		for j := i + 1; j < len(sortedMappings); j++ {
			if len(sortedMappings[j].from) > len(sortedMappings[i].from) {
				sortedMappings[i], sortedMappings[j] = sortedMappings[j], sortedMappings[i]
			}
		}
	}
	
	// Apply replacements
	for _, mapping := range sortedMappings {
		// Handle different reference formats
		patterns := []struct {
			pattern string
			replace string
		}{
			// Direct resource type replacement
			{
				pattern: fmt.Sprintf(`\b%s\.`, regexp.QuoteMeta(mapping.from)),
				replace: fmt.Sprintf(`%s.`, mapping.to),
			},
			// In interpolations
			{
				pattern: fmt.Sprintf(`\$\{%s\.`, regexp.QuoteMeta(mapping.from)),
				replace: fmt.Sprintf(`${%s.`, mapping.to),
			},
			// In for_each references
			{
				pattern: fmt.Sprintf(`each\.value\.%s`, regexp.QuoteMeta(mapping.from)),
				replace: fmt.Sprintf(`each.value.%s`, mapping.to),
			},
		}
		
		for _, p := range patterns {
			re := regexp.MustCompile(p.pattern)
			result = re.ReplaceAllString(result, p.replace)
		}
	}
	
	return result
}

// setUpdatedExpression sets an updated expression preserving its type
func setUpdatedExpression(body *hclwrite.Body, name, newValue string) {
	// Try to determine if it's a string literal or expression
	if strings.HasPrefix(newValue, `"`) && strings.HasSuffix(newValue, `"`) {
		// String literal
		body.SetAttributeValue(name, cty.StringVal(strings.Trim(newValue, `"`)))
	} else if strings.Contains(newValue, "${") || strings.Contains(newValue, ".") {
		// Expression or interpolation - parse it
		tempConfig := fmt.Sprintf("x = %s", newValue)
		tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
		if tempFile != nil && tempFile.Body() != nil {
			if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
				body.SetAttributeRaw(name, tempAttr.Expr().BuildTokens(nil))
			}
		}
	} else {
		// Try to parse as traversal
		if traversal, diags := hclsyntax.ParseTraversalAbs([]byte(newValue), "", hcl.InitialPos); !diags.HasErrors() {
			body.SetAttributeTraversal(name, traversal)
		} else {
			// Fall back to string value
			body.SetAttributeValue(name, cty.StringVal(newValue))
		}
	}
}

// ChainedReferenceUpdater creates a transformer that applies multiple reference updates in sequence
func ChainedReferenceUpdater(updates []ReferenceUpdate) TransformFunc {
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		for _, update := range updates {
			if update.Condition != nil && !update.Condition(block) {
				continue
			}
			
			updater := ReferenceUpdater(update.Mappings)
			if err := updater(block, ctx); err != nil {
				return err
			}
		}
		
		return nil
	}
}

// ReferenceUpdate defines a conditional reference update
type ReferenceUpdate struct {
	Condition func(*hclwrite.Block) bool `yaml:"-"`
	Mappings  map[string]string           `yaml:"mappings"`
}

// GlobalReferenceUpdater updates references across all blocks in a file
func GlobalReferenceUpdater(mappings map[string]string) func(*hclwrite.File) error {
	return func(file *hclwrite.File) error {
		if file == nil || file.Body() == nil {
			return nil
		}
		
		// Process all top-level blocks
		for _, block := range file.Body().Blocks() {
			ctx := &TransformContext{}
			updater := ReferenceUpdater(mappings)
			if err := updater(block, ctx); err != nil {
				return err
			}
		}
		
		// Process any root-level attributes (rare but possible)
		body := file.Body()
		for name, attr := range body.Attributes() {
			tokens := attr.Expr().BuildTokens(nil)
			value := string(tokens.Bytes())
			
			if containsReference(value) {
				newValue := updateReferences(value, mappings)
				if newValue != value {
					setUpdatedExpression(body, name, newValue)
				}
			}
		}
		
		return nil
	}
}

// ReferenceUpdaterForState updates references in state
func ReferenceUpdaterForState(mappings map[string]string) func(map[string]interface{}) error {
	return func(state map[string]interface{}) error {
		if state == nil {
			return nil
		}
		
		attributes, ok := state["attributes"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("state does not contain attributes map")
		}
		
		// Process all attributes recursively
		updateStateReferences(attributes, mappings)
		
		// Process dependencies if present
		if deps, ok := state["dependencies"].([]interface{}); ok {
			for i, dep := range deps {
				if depStr, ok := dep.(string); ok {
					deps[i] = updateReferences(depStr, mappings)
				}
			}
		}
		
		return nil
	}
}

// updateStateReferences recursively updates references in state attributes
func updateStateReferences(data interface{}, mappings map[string]string) {
	switch v := data.(type) {
	case map[string]interface{}:
		// Recursively update map values
		for key, value := range v {
			switch val := value.(type) {
			case string:
				if containsReference(val) {
					v[key] = updateReferences(val, mappings)
				}
			case map[string]interface{}, []interface{}:
				updateStateReferences(val, mappings)
			}
		}
		
	case []interface{}:
		// Update array elements
		for i, item := range v {
			switch val := item.(type) {
			case string:
				if containsReference(val) {
					v[i] = updateReferences(val, mappings)
				}
			case map[string]interface{}, []interface{}:
				updateStateReferences(val, mappings)
			}
		}
	}
}

// SpecificAttributeReferenceUpdater updates references only in specific attributes
func SpecificAttributeReferenceUpdater(attributes []string, mappings map[string]string) TransformFunc {
	attrSet := make(map[string]bool)
	for _, attr := range attributes {
		attrSet[attr] = true
	}
	
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		
		for name, attr := range body.Attributes() {
			if !attrSet[name] {
				continue
			}
			
			tokens := attr.Expr().BuildTokens(nil)
			value := string(tokens.Bytes())
			
			if containsReference(value) {
				newValue := updateReferences(value, mappings)
				if newValue != value {
					setUpdatedExpression(body, name, newValue)
				}
			}
		}
		
		return nil
	}
}

// PatternBasedReferenceUpdater updates references based on regex patterns
func PatternBasedReferenceUpdater(patterns []ReferencePattern) TransformFunc {
	return func(block *hclwrite.Block, ctx *TransformContext) error {
		body := block.Body()
		
		for name, attr := range body.Attributes() {
			tokens := attr.Expr().BuildTokens(nil)
			value := string(tokens.Bytes())
			updated := false
			
			for _, pattern := range patterns {
				re, err := regexp.Compile(pattern.Pattern)
				if err != nil {
					continue
				}
				
				if re.MatchString(value) {
					value = re.ReplaceAllString(value, pattern.Replacement)
					updated = true
				}
			}
			
			if updated {
				setUpdatedExpression(body, name, value)
			}
		}
		
		// Process nested blocks
		for _, nestedBlock := range body.Blocks() {
			if err := applyPatternUpdates(nestedBlock, patterns); err != nil {
				return err
			}
		}
		
		return nil
	}
}

// ReferencePattern defines a regex-based reference update
type ReferencePattern struct {
	Pattern     string `yaml:"pattern"`
	Replacement string `yaml:"replacement"`
}

// applyPatternUpdates recursively applies pattern-based updates
func applyPatternUpdates(block *hclwrite.Block, patterns []ReferencePattern) error {
	body := block.Body()
	
	for name, attr := range body.Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		value := string(tokens.Bytes())
		updated := false
		
		for _, pattern := range patterns {
			re, err := regexp.Compile(pattern.Pattern)
			if err != nil {
				continue
			}
			
			if re.MatchString(value) {
				value = re.ReplaceAllString(value, pattern.Replacement)
				updated = true
			}
		}
		
		if updated {
			setUpdatedExpression(body, name, value)
		}
	}
	
	// Recurse into nested blocks
	for _, nestedBlock := range body.Blocks() {
		if err := applyPatternUpdates(nestedBlock, patterns); err != nil {
			return err
		}
	}
	
	return nil
}
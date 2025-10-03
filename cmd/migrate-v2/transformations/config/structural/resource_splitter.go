package structural

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

// ResourceSplitter creates a transformer that splits resources based on rules
//
// Example YAML configuration:
//   resource_splits:
//     source_resource: cloudflare_combined_resource
//     splits:
//       - when_attribute_exists: waf_settings
//         create_resource: cloudflare_waf_rule
//         attribute_mappings:
//           waf_settings: configuration
//         copy_attributes:
//           - zone_id
//         resource_name_suffix: waf
//       - when_attribute_exists: firewall_rules
//         create_resource: cloudflare_ruleset
//         resource_name_suffix: firewall
//     fallback:
//       change_resource_type: cloudflare_base_resource
//
// Transforms:
//   resource "cloudflare_combined_resource" "example" {
//     zone_id = "abc123"
//     name = "test"
//     waf_settings {
//       enabled = true
//     }
//   }
//
// Into two resources:
//   resource "cloudflare_base_resource" "example" {
//     zone_id = "abc123"
//     name = "test"
//   }
//
//   resource "cloudflare_waf_rule" "example_waf" {
//     zone_id = "abc123"
//     configuration {
//       enabled = true
//     }
//   }
func ResourceSplitter(split basic.ResourceSplit) basic.TransformFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		// Only process resource blocks
		if block.Type() != "resource" {
			return nil
		}
		
		labels := block.Labels()
		if len(labels) < 2 {
			return nil
		}
		
		resourceType := labels[0]
		resourceName := labels[1]
		
		// Check if this is the resource type we want to split
		if split.SourceResource != "" && resourceType != split.SourceResource {
			return nil
		}
		
		body := block.Body()
		
		// Process each split rule
		for _, rule := range split.Splits {
			if shouldApplySplitRule(body, rule) {
				newBlock := createSplitResource(block, rule, resourceName)
				if newBlock != nil {
					// Add to context for later processing
					if ctx.NewBlocks == nil {
						ctx.NewBlocks = []*hclwrite.Block{}
					}
					ctx.NewBlocks = append(ctx.NewBlocks, newBlock)
					
					// Apply attribute mappings to original block
					applyAttributeMappingsToOriginal(body, rule)
					
					// Generate moved block if requested
					if split.GenerateMovedBlocks && rule.CreateResource != "" {
						generateMovedBlock(ctx, resourceType, resourceName, rule)
					}
				}
			}
		}
		
		// Apply fallback rule if specified
		if split.Fallback != nil {
			if !anyRuleApplied(body, split.Splits) {
				applyFallbackRule(block, split.Fallback, ctx)
			}
		}
		
		// Mark original for removal if requested
		if split.RemoveOriginal {
			if ctx.BlocksToRemove == nil {
				ctx.BlocksToRemove = make(map[*hclwrite.Block]bool)
			}
			ctx.BlocksToRemove[block] = true
		}
		
		return nil
	}
}

// shouldApplySplitRule checks if a split rule should be applied
func shouldApplySplitRule(body *hclwrite.Body, rule basic.SplitRule) bool {
	if rule.WhenAttributeExists != "" {
		attr := body.GetAttribute(rule.WhenAttributeExists)
		return attr != nil
	}
	// If no condition specified, always apply
	return true
}

// createSplitResource creates a new resource based on split rule
func createSplitResource(originalBlock *hclwrite.Block, rule basic.SplitRule, originalName string) *hclwrite.Block {
	labels := originalBlock.Labels()
	if len(labels) < 2 {
		return nil
	}
	
	// Determine new resource type
	var newResourceType string
	if rule.ChangeResourceType != "" {
		newResourceType = rule.ChangeResourceType
	} else if rule.CreateResource != "" {
		newResourceType = rule.CreateResource
	} else {
		newResourceType = labels[0] // Keep original type
	}
	
	// Determine new resource name
	newResourceName := originalName
	if rule.ResourceNameSuffix != "" {
		newResourceName = fmt.Sprintf("%s_%s", originalName, rule.ResourceNameSuffix)
	}
	
	// Create new block
	newBlock := hclwrite.NewBlock("resource", []string{newResourceType, newResourceName})
	newBody := newBlock.Body()
	originalBody := originalBlock.Body()
	
	// Copy attributes based on rules
	if len(rule.CopyAttributes) > 0 {
		// Copy specific attributes
		for _, attrName := range rule.CopyAttributes {
			if attr := originalBody.GetAttribute(attrName); attr != nil {
				newBody.SetAttributeRaw(attrName, attr.Expr().BuildTokens(nil))
			}
		}
	}
	
	// Apply attribute mappings
	for fromAttr, toAttr := range rule.AttributeMappings {
		if attr := originalBody.GetAttribute(fromAttr); attr != nil {
			newBody.SetAttributeRaw(toAttr, attr.Expr().BuildTokens(nil))
		}
	}
	
	// Set new attributes
	for attrName, value := range rule.SetAttributes {
		setAttributeValue(newBody, attrName, value)
	}
	
	return newBlock
}

// applyAttributeMappingsToOriginal removes or renames attributes in original
func applyAttributeMappingsToOriginal(body *hclwrite.Body, rule basic.SplitRule) {
	// Remove attributes that were moved to new resource
	for fromAttr := range rule.AttributeMappings {
		body.RemoveAttribute(fromAttr)
	}
	
	// Remove copied attributes if they should only exist in new resource
	if rule.WhenAttributeExists != "" {
		// If splitting based on attribute existence, remove it from original
		body.RemoveAttribute(rule.WhenAttributeExists)
	}
}

// anyRuleApplied checks if any rule was applied
func anyRuleApplied(body *hclwrite.Body, rules []basic.SplitRule) bool {
	for _, rule := range rules {
		if shouldApplySplitRule(body, rule) {
			return true
		}
	}
	return false
}

// applyFallbackRule applies the fallback rule
func applyFallbackRule(block *hclwrite.Block, fallback *basic.SplitRule, ctx *basic.TransformContext) {
	labels := block.Labels()
	if len(labels) < 2 {
		return
	}
	
	// Change resource type if specified
	if fallback.ChangeResourceType != "" {
		if ctx.ResourceTypeChanges == nil {
			ctx.ResourceTypeChanges = make(map[*hclwrite.Block]string)
		}
		ctx.ResourceTypeChanges[block] = fallback.ChangeResourceType
	}
	
	body := block.Body()
	
	// Apply attribute mappings
	for fromAttr, toAttr := range fallback.AttributeMappings {
		if attr := body.GetAttribute(fromAttr); attr != nil {
			body.SetAttributeRaw(toAttr, attr.Expr().BuildTokens(nil))
			if fromAttr != toAttr {
				body.RemoveAttribute(fromAttr)
			}
		}
	}
	
	// Set attributes
	for attrName, value := range fallback.SetAttributes {
		setAttributeValue(body, attrName, value)
	}
}

// generateMovedBlock creates a moved block for resource migration
func generateMovedBlock(ctx *basic.TransformContext, resourceType, resourceName string, rule basic.SplitRule) {
	fromRef := fmt.Sprintf("%s.%s", resourceType, resourceName)
	
	var toRef string
	if rule.CreateResource != "" {
		suffix := resourceName
		if rule.ResourceNameSuffix != "" {
			suffix = fmt.Sprintf("%s_%s", resourceName, rule.ResourceNameSuffix)
		}
		toRef = fmt.Sprintf("%s.%s", rule.CreateResource, suffix)
	} else if rule.ChangeResourceType != "" {
		toRef = fmt.Sprintf("%s.%s", rule.ChangeResourceType, resourceName)
	}
	
	if toRef != "" {
		if ctx.MovedBlocks == nil {
			ctx.MovedBlocks = make(map[string]string)
		}
		ctx.MovedBlocks[fromRef] = toRef
	}
}

// setAttributeValue sets an attribute value handling different types
func setAttributeValue(body *hclwrite.Body, name string, value interface{}) {
	switch v := value.(type) {
	case string:
		// Check if it's a reference (starts with var., local., etc.)
		if strings.HasPrefix(v, "var.") || strings.HasPrefix(v, "local.") || 
		   strings.HasPrefix(v, "data.") || strings.HasPrefix(v, "module.") {
			// Set as traversal
			parts := strings.Split(v, ".")
			traversal := hcl.Traversal{
				hcl.TraverseRoot{Name: parts[0]},
			}
			for i := 1; i < len(parts); i++ {
				traversal = append(traversal, hcl.TraverseAttr{Name: parts[i]})
			}
			body.SetAttributeTraversal(name, traversal)
		} else {
			// Set as string value
			body.SetAttributeValue(name, cty.StringVal(v))
		}
	case bool:
		body.SetAttributeValue(name, cty.BoolVal(v))
	case int:
		body.SetAttributeValue(name, cty.NumberIntVal(int64(v)))
	case int64:
		body.SetAttributeValue(name, cty.NumberIntVal(v))
	case float64:
		body.SetAttributeValue(name, cty.NumberFloatVal(v))
	case []interface{}:
		// Convert to HCL tokens
		var listStr strings.Builder
		listStr.WriteString("[")
		for i, item := range v {
			if i > 0 {
				listStr.WriteString(", ")
			}
			switch itemVal := item.(type) {
			case string:
				listStr.WriteString(fmt.Sprintf(`"%s"`, itemVal))
			default:
				listStr.WriteString(fmt.Sprintf("%v", itemVal))
			}
		}
		listStr.WriteString("]")
		
		// Parse and set
		tempConfig := fmt.Sprintf("x = %s", listStr.String())
		tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
		if tempFile != nil && tempFile.Body() != nil {
			if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
				body.SetAttributeRaw(name, tempAttr.Expr().BuildTokens(nil))
			}
		}
	case map[string]interface{}:
		// Convert map to HCL object
		var objStr strings.Builder
		objStr.WriteString("{")
		first := true
		for k, val := range v {
			if !first {
				objStr.WriteString(", ")
			}
			first = false
			switch v := val.(type) {
			case string:
				objStr.WriteString(fmt.Sprintf(`%s = "%s"`, k, v))
			default:
				objStr.WriteString(fmt.Sprintf("%s = %v", k, val))
			}
		}
		objStr.WriteString("}")
		
		// Parse and set
		tempConfig := fmt.Sprintf("x = %s", objStr.String())
		tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
		if tempFile != nil && tempFile.Body() != nil {
			if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
				body.SetAttributeRaw(name, tempAttr.Expr().BuildTokens(nil))
			}
		}
	default:
		// Default to string representation
		body.SetAttributeValue(name, cty.StringVal(fmt.Sprintf("%v", value)))
	}
}

// ResourceSplitterForState applies resource splitting to state
func ResourceSplitterForState(split basic.ResourceSplit) func(map[string]interface{}) error {
	return func(state map[string]interface{}) error {
		if state == nil {
			return nil
		}
		
		attributes, ok := state["attributes"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("state does not contain attributes map")
		}
		
		// Check each split rule
		for _, rule := range split.Splits {
			if shouldApplySplitRuleInState(attributes, rule) {
				// Create new state for split resource
				newState := createSplitStateResource(attributes, rule)
				
				// Store in metadata for later processing
				if metadata, ok := state["_split_resources"].([]map[string]interface{}); ok {
					state["_split_resources"] = append(metadata, newState)
				} else {
					state["_split_resources"] = []map[string]interface{}{newState}
				}
				
				// Remove moved attributes from original
				removeMovedAttributesFromState(attributes, rule)
			}
		}
		
		// Apply fallback if no rules matched
		if split.Fallback != nil {
			if !anyRuleAppliedToState(attributes, split.Splits) {
				applyFallbackToState(attributes, split.Fallback)
			}
		}
		
		return nil
	}
}

// shouldApplySplitRuleInState checks if rule applies to state
func shouldApplySplitRuleInState(attributes map[string]interface{}, rule basic.SplitRule) bool {
	if rule.WhenAttributeExists != "" {
		_, exists := attributes[rule.WhenAttributeExists]
		return exists
	}
	return true
}

// createSplitStateResource creates new state for split resource
func createSplitStateResource(originalAttrs map[string]interface{}, rule basic.SplitRule) map[string]interface{} {
	newAttrs := make(map[string]interface{})
	
	// Copy specified attributes
	for _, attrName := range rule.CopyAttributes {
		if val, exists := originalAttrs[attrName]; exists {
			newAttrs[attrName] = val
		}
	}
	
	// Apply mappings
	for fromAttr, toAttr := range rule.AttributeMappings {
		if val, exists := originalAttrs[fromAttr]; exists {
			newAttrs[toAttr] = val
		}
	}
	
	// Set new attributes
	for attrName, value := range rule.SetAttributes {
		newAttrs[attrName] = value
	}
	
	return map[string]interface{}{
		"attributes": newAttrs,
		"_resource_type": rule.CreateResource,
		"_resource_name_suffix": rule.ResourceNameSuffix,
	}
}

// removeMovedAttributesFromState removes attributes that were split
func removeMovedAttributesFromState(attributes map[string]interface{}, rule basic.SplitRule) {
	// Remove mapped attributes
	for fromAttr := range rule.AttributeMappings {
		delete(attributes, fromAttr)
	}
	
	// Remove trigger attribute if specified
	if rule.WhenAttributeExists != "" {
		delete(attributes, rule.WhenAttributeExists)
	}
}

// anyRuleAppliedToState checks if any rule applies to state
func anyRuleAppliedToState(attributes map[string]interface{}, rules []basic.SplitRule) bool {
	for _, rule := range rules {
		if shouldApplySplitRuleInState(attributes, rule) {
			return true
		}
	}
	return false
}

// applyFallbackToState applies fallback rule to state
func applyFallbackToState(attributes map[string]interface{}, fallback *basic.SplitRule) {
	// Apply attribute mappings
	for fromAttr, toAttr := range fallback.AttributeMappings {
		if val, exists := attributes[fromAttr]; exists {
			attributes[toAttr] = val
			if fromAttr != toAttr {
				delete(attributes, fromAttr)
			}
		}
	}
	
	// Set attributes
	for attrName, value := range fallback.SetAttributes {
		attributes[attrName] = value
	}
}
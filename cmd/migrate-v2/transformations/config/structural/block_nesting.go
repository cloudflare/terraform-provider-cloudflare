package structural

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

// BlockNester creates a transformer that nests attributes into blocks
//
// Example YAML configuration:
//   structural_changes:
//     - type: nest_block
//       parameters:
//         block_name: network_settings
//         attributes:
//           - ip_address
//           - port
//           - protocol
//         remove_originals: true
//
// Transforms:
//   resource "example" "test" {
//     name = "test"
//     ip_address = "192.168.1.1"
//     port = 8080
//     protocol = "tcp"
//   }
//
// Into:
//   resource "example" "test" {
//     name = "test"
//     
//     network_settings {
//       ip_address = "192.168.1.1"
//       port = 8080
//       protocol = "tcp"
//     }
//   }
func BlockNester(blockName string, attributes []string, removeOriginals bool) basic.TransformFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		
		// Collect values for attributes to nest
		attrValues := make(map[string]hclwrite.Tokens)
		foundAny := false
		
		for _, attrName := range attributes {
			if attr := body.GetAttribute(attrName); attr != nil {
				attrValues[attrName] = attr.Expr().BuildTokens(nil)
				foundAny = true
			}
		}
		
		// If no attributes found, nothing to do
		if !foundAny {
			return nil
		}
		
		// Create new nested block
		nestedBlock := hclwrite.NewBlock(blockName, []string{})
		nestedBody := nestedBlock.Body()
		
		// Add attributes to nested block
		for attrName, tokens := range attrValues {
			nestedBody.SetAttributeRaw(attrName, tokens)
			
			// Remove from parent if requested
			if removeOriginals {
				body.RemoveAttribute(attrName)
			}
		}
		
		// Add the nested block to parent
		body.AppendBlock(nestedBlock)
		
		return nil
	}
}

// BlockUnnester creates a transformer that flattens nested blocks to parent level
//
// Example YAML configuration:
//   structural_changes:
//     - type: unnest_block
//       parameters:
//         block_name: network_settings
//         promote_attributes: []  # Empty means promote all
//
// Transforms:
//   resource "example" "test" {
//     name = "test"
//     
//     network_settings {
//       ip_address = "192.168.1.1"
//       port = 8080
//       protocol = "tcp"
//     }
//   }
//
// Into:
//   resource "example" "test" {
//     name = "test"
//     ip_address = "192.168.1.1"
//     port = 8080
//     protocol = "tcp"
//   }
func BlockUnnester(blockName string, promoteAttributes []string) basic.TransformFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		
		// Find the nested block
		var targetBlock *hclwrite.Block
		var blocksToRemove []*hclwrite.Block
		
		for _, nestedBlock := range body.Blocks() {
			if nestedBlock.Type() == blockName {
				targetBlock = nestedBlock
				blocksToRemove = append(blocksToRemove, nestedBlock)
				break // Only process first matching block
			}
		}
		
		if targetBlock == nil {
			return nil // Block not found
		}
		
		nestedBody := targetBlock.Body()
		
		// Promote specified attributes (or all if empty list)
		if len(promoteAttributes) == 0 {
			// Promote all attributes
			for name, attr := range nestedBody.Attributes() {
				body.SetAttributeRaw(name, attr.Expr().BuildTokens(nil))
			}
		} else {
			// Promote specific attributes
			for _, attrName := range promoteAttributes {
				if attr := nestedBody.GetAttribute(attrName); attr != nil {
					body.SetAttributeRaw(attrName, attr.Expr().BuildTokens(nil))
				}
			}
		}
		
		// Remove the nested block
		for _, blockToRemove := range blocksToRemove {
			body.RemoveBlock(blockToRemove)
		}
		
		return nil
	}
}

// MultiLevelNester creates nested blocks multiple levels deep
func MultiLevelNester(nestingSpec []NestingLevel) basic.TransformFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		
		// Process each level of nesting independently
		// Each level creates a new block at the root level
		for _, level := range nestingSpec {
			// Collect attributes for this level from the original body
			attrValues := make(map[string]hclwrite.Tokens)
			foundAny := false
			
			for _, attrName := range level.Attributes {
				if attr := body.GetAttribute(attrName); attr != nil {
					attrValues[attrName] = attr.Expr().BuildTokens(nil)
					foundAny = true
				}
			}
			
			if !foundAny && !level.CreateEmpty {
				// No attributes found and not creating empty block
				continue
			}
			
			// Create nested block for this level
			nestedBlock := hclwrite.NewBlock(level.BlockName, level.Labels)
			nestedBody := nestedBlock.Body()
			
			// Add attributes
			for attrName, tokens := range attrValues {
				nestedBody.SetAttributeRaw(attrName, tokens)
				
				// Remove from original body if requested
				if level.RemoveOriginals {
					body.RemoveAttribute(attrName)
				}
			}
			
			// Add to main body
			body.AppendBlock(nestedBlock)
		}
		
		return nil
	}
}

// NestingLevel defines a level in multi-level nesting
type NestingLevel struct {
	BlockName       string   `yaml:"block_name"`
	Labels          []string `yaml:"labels,omitempty"`
	Attributes      []string `yaml:"attributes"`
	RemoveOriginals bool     `yaml:"remove_originals,omitempty"`
	CreateEmpty     bool     `yaml:"create_empty,omitempty"`
}

// ConditionalNester nests attributes based on conditions
func ConditionalNester(blockName string, condition func(*hclwrite.Block) bool, attributes []string) basic.TransformFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		// Check condition
		if !condition(block) {
			return nil
		}
		
		// Apply nesting
		nester := BlockNester(blockName, attributes, true)
		return nester(block, ctx)
	}
}

// DynamicBlockConverter converts repeated blocks to dynamic blocks
func DynamicBlockConverter(blockType string, iteratorName string) basic.TransformFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		
		// Collect all blocks of the specified type
		var targetBlocks []*hclwrite.Block
		for _, nestedBlock := range body.Blocks() {
			if nestedBlock.Type() == blockType {
				targetBlocks = append(targetBlocks, nestedBlock)
			}
		}
		
		if len(targetBlocks) <= 1 {
			// No need for dynamic block if 0 or 1 instance
			return nil
		}
		
		// Create list of configurations
		var configs []string
		for _, targetBlock := range targetBlocks {
			configStr := extractBlockConfig(targetBlock)
			configs = append(configs, configStr)
		}
		
		// Create dynamic block
		dynamicBlock := hclwrite.NewBlock("dynamic", []string{blockType})
		dynamicBody := dynamicBlock.Body()
		
		// Set for_each
		forEachValue := fmt.Sprintf("[%s]", strings.Join(configs, ", "))
		tempConfig := fmt.Sprintf("x = %s", forEachValue)
		tempFile, _ := hclwrite.ParseConfig([]byte(tempConfig), "", hcl.InitialPos)
		if tempFile != nil && tempFile.Body() != nil {
			if tempAttr := tempFile.Body().GetAttribute("x"); tempAttr != nil {
				dynamicBody.SetAttributeRaw("for_each", tempAttr.Expr().BuildTokens(nil))
			}
		}
		
		// Set iterator if custom name specified
		if iteratorName != "" && iteratorName != blockType {
			dynamicBody.SetAttributeValue("iterator", cty.StringVal(iteratorName))
		}
		
		// Create content block
		contentBlock := hclwrite.NewBlock("content", []string{})
		contentBody := contentBlock.Body()
		
		// Use first block as template for content
		if len(targetBlocks) > 0 {
			templateBlock := targetBlocks[0]
			templateBody := templateBlock.Body()
			
			// Copy attributes with iterator references
			for name, attr := range templateBody.Attributes() {
				// Convert to use iterator
				tokens := attr.Expr().BuildTokens(nil)
				tokenStr := string(tokens.Bytes())
				
				// Simple reference conversion
				iterRef := iteratorName
				if iterRef == "" {
					iterRef = blockType
				}
				
				// Check if value needs iterator reference
				if strings.Contains(tokenStr, "${") || isSimpleValue(tokenStr) {
					contentBody.SetAttributeRaw(name, tokens)
				} else {
					// Create iterator reference
					traversal := hcl.Traversal{
						hcl.TraverseRoot{Name: iterRef},
						hcl.TraverseAttr{Name: "value"},
						hcl.TraverseAttr{Name: name},
					}
					contentBody.SetAttributeTraversal(name, traversal)
				}
			}
		}
		
		dynamicBody.AppendBlock(contentBlock)
		
		// Remove original blocks
		for _, targetBlock := range targetBlocks {
			body.RemoveBlock(targetBlock)
		}
		
		// Add dynamic block
		body.AppendBlock(dynamicBlock)
		
		return nil
	}
}

// extractBlockConfig extracts configuration from a block as a map string
func extractBlockConfig(block *hclwrite.Block) string {
	body := block.Body()
	var parts []string
	
	for name, attr := range body.Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		value := strings.TrimSpace(string(tokens.Bytes()))
		parts = append(parts, fmt.Sprintf("%s = %s", name, value))
	}
	
	return fmt.Sprintf("{\n      %s\n    }", strings.Join(parts, "\n      "))
}

// isSimpleValue checks if a value is a simple literal
func isSimpleValue(value string) bool {
	// Check for string literals
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		return true
	}
	// Check for numbers
	if _, err := fmt.Sscanf(value, "%f", new(float64)); err == nil {
		return true
	}
	// Check for booleans
	return value == "true" || value == "false"
}

// BlockNesterForState nests attributes in state
func BlockNesterForState(blockName string, attributes []string) func(map[string]interface{}) error {
	return func(state map[string]interface{}) error {
		if state == nil {
			return nil
		}
		
		attrs, ok := state["attributes"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("state does not contain attributes map")
		}
		
		// Collect values to nest
		nestedAttrs := make(map[string]interface{})
		foundAny := false
		
		for _, attrName := range attributes {
			if val, exists := attrs[attrName]; exists {
				nestedAttrs[attrName] = val
				foundAny = true
			}
		}
		
		if !foundAny {
			return nil
		}
		
		// Create nested structure
		attrs[blockName] = []interface{}{nestedAttrs}
		
		// Remove original attributes
		for _, attrName := range attributes {
			delete(attrs, attrName)
		}
		
		return nil
	}
}

// BlockUnnesterForState unnests blocks in state
func BlockUnnesterForState(blockName string) func(map[string]interface{}) error {
	return func(state map[string]interface{}) error {
		if state == nil {
			return nil
		}
		
		attrs, ok := state["attributes"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("state does not contain attributes map")
		}
		
		// Check for nested block
		nestedVal, exists := attrs[blockName]
		if !exists {
			return nil
		}
		
		// Handle different nested formats
		var nestedAttrs map[string]interface{}
		
		switch v := nestedVal.(type) {
		case []interface{}:
			// List of blocks - take first
			if len(v) > 0 {
				if m, ok := v[0].(map[string]interface{}); ok {
					nestedAttrs = m
				}
			}
		case map[string]interface{}:
			// Single block
			nestedAttrs = v
		default:
			return nil
		}
		
		// Promote attributes
		for k, v := range nestedAttrs {
			attrs[k] = v
		}
		
		// Remove nested block
		delete(attrs, blockName)
		
		return nil
	}
}
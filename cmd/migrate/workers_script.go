package main

import (
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

// isWorkersScriptResource checks if a block is a cloudflare_workers_script resource
func isWorkersScriptResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" &&
		len(block.Labels()) >= 2 &&
		(block.Labels()[0] == "cloudflare_worker_script" || block.Labels()[0] == "cloudflare_workers_script")
}

// transformWorkersScriptBlock transforms v4 workers_script to v5 format
//
// Handles:
// 1. Resource rename: cloudflare_worker_script → cloudflare_workers_script
// 2. Attribute rename: name → script_name
// 3. Binding transformation: all binding blocks → unified bindings array
// 4. Remove module attribute
// 5. Transform placement block → placement object
func transformWorkersScriptBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	// 1. Rename resource type (cloudflare_worker_script → cloudflare_workers_script)
	if block.Type() == "resource" && len(block.Labels()) >= 2 && block.Labels()[0] == "cloudflare_worker_script" {
		// Update the resource type label
		block.SetLabels([]string{"cloudflare_workers_script", block.Labels()[1]})
	}
	
	body := block.Body()
	
	// 2. Rename top-level name attribute to script_name
	if nameAttr := body.GetAttribute("name"); nameAttr != nil {
		// Get the current value and set it as script_name
		body.SetAttributeRaw("script_name", nameAttr.Expr().BuildTokens(nil))
		body.RemoveAttribute("name")
	}
	
	// Collect all binding blocks and remove them
	var bindings []cty.Value
	var blocksToRemove []*hclwrite.Block
	
	// Process each block in the resource
	for _, childBlock := range body.Blocks() {
		bindingType := childBlock.Type()
		
		switch bindingType {
		case "analytics_engine_binding":
			binding := transformAnalyticsEngineBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "d1_database_binding":
			binding := transformD1DatabaseBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "kv_namespace_binding":
			binding := transformKVNamespaceBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "plain_text_binding":
			binding := transformPlainTextBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "secret_text_binding":
			binding := transformSecretTextBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "queue_binding":
			binding := transformQueueBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "r2_bucket_binding":
			binding := transformR2BucketBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "service_binding":
			binding := transformServiceBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "hyperdrive_config_binding":
			binding := transformHyperdriveBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "webassembly_binding":
			binding := transformWebAssemblyBinding(childBlock, diags)
			if !binding.IsNull() {
				bindings = append(bindings, binding)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
			
		case "placement":
			// Handle placement block - convert to placement object
			placement := transformPlacementBlock(childBlock, diags)
			if !placement.IsNull() {
				body.SetAttributeValue("placement", placement)
			}
			blocksToRemove = append(blocksToRemove, childBlock)
		}
	}
	
	// Remove all binding blocks
	for _, block := range blocksToRemove {
		body.RemoveBlock(block)
	}
	
	// Remove the 'module' attribute if it exists (doesn't exist in v5)
	if body.GetAttribute("module") != nil {
		body.RemoveAttribute("module")
	}
	
	// Add bindings array if we have any bindings
	if len(bindings) > 0 {
		bindingsArray := cty.ListVal(bindings)
		body.SetAttributeValue("bindings", bindingsArray)
	}
}

// Helper function to get string attribute value from block
func getStringAttr(block *hclwrite.Block, attrName string) (string, bool) {
	attr := block.Body().GetAttribute(attrName)
	if attr == nil {
		return "", false
	}
	
	// Get the attribute value by cloning and parsing
	tokens := attr.Expr().BuildTokens(nil)
	if len(tokens) == 0 {
		return "", false
	}
	
	// Join all tokens to get the complete value
	var fullValue string
	for _, token := range tokens {
		fullValue += string(token.Bytes)
	}
	
	// Remove surrounding quotes if present
	fullValue = strings.TrimSpace(fullValue)
	if len(fullValue) >= 2 && fullValue[0] == '"' && fullValue[len(fullValue)-1] == '"' {
		return fullValue[1 : len(fullValue)-1], true
	}
	
	// Return as-is for unquoted values
	return fullValue, true
}

// Transform analytics engine binding
func transformAnalyticsEngineBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	dataset, hasDataset := getStringAttr(block, "dataset")
	
	if !hasName || !hasDataset {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name":    cty.StringVal(name),
		"type":    cty.StringVal("analytics_engine"),
		"dataset": cty.StringVal(dataset),
	})
}

// Transform D1 database binding
func transformD1DatabaseBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	databaseId, hasId := getStringAttr(block, "database_id")
	
	if !hasName || !hasId {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name": cty.StringVal(name),
		"type": cty.StringVal("d1"),
		"id":   cty.StringVal(databaseId),
	})
}

// Transform KV namespace binding
func transformKVNamespaceBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	namespaceId, hasId := getStringAttr(block, "namespace_id")
	
	if !hasName || !hasId {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name":         cty.StringVal(name),
		"type":         cty.StringVal("kv_namespace"),
		"namespace_id": cty.StringVal(namespaceId),
	})
}

// Transform plain text binding
func transformPlainTextBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	text, hasText := getStringAttr(block, "text")
	
	if !hasName || !hasText {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name": cty.StringVal(name),
		"type": cty.StringVal("plain_text"),
		"text": cty.StringVal(text),
	})
}

// Transform secret text binding
func transformSecretTextBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	text, hasText := getStringAttr(block, "text")
	
	if !hasName || !hasText {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name": cty.StringVal(name),
		"type": cty.StringVal("secret_text"),
		"text": cty.StringVal(text),
	})
}

// Transform queue binding
func transformQueueBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	queue, hasQueue := getStringAttr(block, "queue")
	
	if !hasName || !hasQueue {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name":       cty.StringVal(name),
		"type":       cty.StringVal("queue"),
		"queue_name": cty.StringVal(queue),
	})
}

// Transform R2 bucket binding
func transformR2BucketBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	bucketName, hasBucket := getStringAttr(block, "bucket_name")
	
	if !hasName || !hasBucket {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name":        cty.StringVal(name),
		"type":        cty.StringVal("r2_bucket"),
		"bucket_name": cty.StringVal(bucketName),
	})
}

// Transform service binding
func transformServiceBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	service, hasService := getStringAttr(block, "service")
	
	if !hasName || !hasService {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	binding := map[string]cty.Value{
		"name":    cty.StringVal(name),
		"type":    cty.StringVal("service"),
		"service": cty.StringVal(service),
	}
	
	// Add environment if present
	if env, hasEnv := getStringAttr(block, "environment"); hasEnv {
		binding["environment"] = cty.StringVal(env)
	}
	
	return cty.ObjectVal(binding)
}

// Transform hyperdrive binding
func transformHyperdriveBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	bindingId, hasId := getStringAttr(block, "binding_id")
	
	if !hasName || !hasId {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name": cty.StringVal(name),
		"type": cty.StringVal("hyperdrive"),
		"id":   cty.StringVal(bindingId),
	})
}

// Transform WebAssembly binding
func transformWebAssemblyBinding(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	name, hasName := getStringAttr(block, "name")
	module, hasModule := getStringAttr(block, "module")
	
	if !hasName || !hasModule {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"name": cty.StringVal(name),
		"type": cty.StringVal("wasm"),
		"part": cty.StringVal(module),
	})
}

// Transform placement block to placement object
func transformPlacementBlock(block *hclwrite.Block, diags ast.Diagnostics) cty.Value {
	mode, hasMode := getStringAttr(block, "mode")
	
	if !hasMode {
		// Skip transformation if required attributes are missing
		return cty.NullVal(cty.Object(map[string]cty.Type{}))
	}
	
	return cty.ObjectVal(map[string]cty.Value{
		"mode": cty.StringVal(mode),
	})
}
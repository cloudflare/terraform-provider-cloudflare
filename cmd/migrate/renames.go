package main

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

var resourceRenames = map[string][]struct {
	old string
	new string
}{
	"cloudflare_custom_pages": {
		{
			old: "type",
			new: "identifier",
		},
	},
	"cloudflare_zero_trust_access_policy": {
		{
			old: "approval_group",
			new: "approval_groups",
		},
	},
}

func applyRenames(block *hclwrite.Block) {
	if len(block.Labels()) == 0 {
		return
	}
	if rename, ok := resourceRenames[block.Labels()[0]]; ok {
		for _, rename := range rename {
			block.Body().RenameAttribute(rename.old, rename.new)
		}
	}

	// Also rename resource references in attribute values
	renameResourceReferences(block.Body())
}

// renameResourceReferences renames resource type references in expressions
// e.g., cloudflare_access_policy.foo.id -> cloudflare_zero_trust_access_policy.foo.id
// and attribute references like cloudflare_workers_script.foo.name -> cloudflare_workers_script.foo.script_name
func renameResourceReferences(body *hclwrite.Body) {
	// Map of old resource names to new ones
	resourceTypeRenames := map[string]string{
		"cloudflare_access_policy":      "cloudflare_zero_trust_access_policy",
		"cloudflare_access_application": "cloudflare_zero_trust_access_application",
		"cloudflare_access_group":       "cloudflare_zero_trust_access_group",
		// Add more renames as needed
	}

	// Map of resource types to their attribute renames
	// Maps: resource_type -> old_attribute -> new_attribute
	attributeRenames := map[string]map[string]string{
		"cloudflare_workers_script": {
			"name": "script_name",
		},
		"cloudflare_worker_script": { // also handle singular version
			"name": "script_name",
		},
	}

	// Process all attributes
	for name, attr := range body.Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		transformedTokens := renameResourceReferencesInTokens(tokens, resourceTypeRenames, attributeRenames)
		if transformedTokens != nil {
			body.SetAttributeRaw(name, transformedTokens)
		}
	}
}

// renameResourceReferencesInTokens renames resource references in token stream
func renameResourceReferencesInTokens(tokens hclwrite.Tokens, renames map[string]string, attributeRenames map[string]map[string]string) hclwrite.Tokens {
	var result hclwrite.Tokens
	changed := false
	i := 0

	for i < len(tokens) {
		token := tokens[i]

		// Look for identifiers that match resource names
		if token.Type == hclsyntax.TokenIdent {
			resourceType := string(token.Bytes)

			// Check for resource type renames (e.g., cloudflare_access_policy -> cloudflare_zero_trust_access_policy)
			if newType, ok := renames[resourceType]; ok {
				// Check if the next token is a dot (indicating a resource reference)
				if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenDot {
					// Replace the old resource type with the new one
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte(newType),
					})
					changed = true
					i++
					continue
				}
			}

			// Check for attribute renames in resource references
			// Pattern: cloudflare_workers_script.example.name -> cloudflare_workers_script.example.script_name
			if attrRenames, ok := attributeRenames[resourceType]; ok {
				// Look for pattern: resourceType.instanceName.attributeName
				if i+3 < len(tokens) &&
					tokens[i+1].Type == hclsyntax.TokenDot &&
					tokens[i+2].Type == hclsyntax.TokenIdent &&
					tokens[i+3].Type == hclsyntax.TokenDot &&
					i+4 < len(tokens) && tokens[i+4].Type == hclsyntax.TokenIdent {

					attributeName := string(tokens[i+4].Bytes)
					if newAttr, ok := attrRenames[attributeName]; ok {
						// Copy the resource type, dot, instance name, dot
						result = append(result, token)       // resource type
						result = append(result, tokens[i+1]) // first dot
						result = append(result, tokens[i+2]) // instance name
						result = append(result, tokens[i+3]) // second dot

						// Replace the attribute name
						result = append(result, &hclwrite.Token{
							Type:  hclsyntax.TokenIdent,
							Bytes: []byte(newAttr),
						})
						changed = true
						i += 5 // Skip the 5 tokens we just processed
						continue
					}
				}
			}
		}

		result = append(result, token)
		i++
	}

	if changed {
		return result
	}
	return nil // Return nil if no changes were made
}

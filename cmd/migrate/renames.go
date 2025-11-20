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
func renameResourceReferences(body *hclwrite.Body) {
	// Map of old resource names to new ones, and new names to themselves (for attribute renaming)
	resourceTypeRenames := map[string]string{
		"cloudflare_access_policy":      "cloudflare_zero_trust_access_policy",
		"cloudflare_access_application": "cloudflare_zero_trust_access_application",
		"cloudflare_access_group":       "cloudflare_zero_trust_access_group",
		// Workers resource renames - old to new
		"cloudflare_worker_route":        "cloudflare_workers_route",
		"cloudflare_worker_script":       "cloudflare_workers_script",
		"cloudflare_worker_cron_trigger": "cloudflare_workers_cron_trigger",
		"cloudflare_worker_domain":       "cloudflare_workers_custom_domain",
		// Workers resource renames - new to new (for attribute renaming on already-migrated references)
		"cloudflare_workers_script": "cloudflare_workers_script",
		"cloudflare_workers_route":  "cloudflare_workers_route",
		// Note: cloudflare_worker_secret is migrated to secret_text bindings, not renamed
	}

	// Process all attributes
	for name, attr := range body.Attributes() {
		tokens := attr.Expr().BuildTokens(nil)
		transformedTokens := renameResourceReferencesInTokens(tokens, resourceTypeRenames)
		if transformedTokens != nil {
			body.SetAttributeRaw(name, transformedTokens)
		}
	}
}

// renameResourceReferencesInTokens renames resource references in token stream
// Handles both resource type renames and attribute renames within references
func renameResourceReferencesInTokens(tokens hclwrite.Tokens, renames map[string]string) hclwrite.Tokens {
	var result hclwrite.Tokens
	changed := false
	i := 0

	for i < len(tokens) {
		token := tokens[i]

		// Look for identifiers that match resource names
		if token.Type == hclsyntax.TokenIdent {
			resourceType := string(token.Bytes)
			if newType, ok := renames[resourceType]; ok {
				// Check if the next token is a dot (indicating a resource reference)
				if i+1 < len(tokens) && tokens[i+1].Type == hclsyntax.TokenDot {
					// Replace the old resource type with the new one
					result = append(result, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte(newType),
					})
					changed = true

					// Add the dot
					result = append(result, tokens[i+1])
					i += 2

					// Skip resource name (e.g., "my_worker")
					if i < len(tokens) && tokens[i].Type == hclsyntax.TokenIdent {
						result = append(result, tokens[i])
						i++

						// Check for another dot (attribute access)
						if i < len(tokens) && tokens[i].Type == hclsyntax.TokenDot {
							result = append(result, tokens[i]) // Add dot
							i++

							// Check for attribute rename
							if i < len(tokens) && tokens[i].Type == hclsyntax.TokenIdent {
								attrName := string(tokens[i].Bytes)
								newAttrName := renameWorkerAttribute(resourceType, newType, attrName)
								if newAttrName != attrName {
									result = append(result, &hclwrite.Token{
										Type:  hclsyntax.TokenIdent,
										Bytes: []byte(newAttrName),
									})
									changed = true
									i++
									continue
								}
							}
						}
					}
					continue
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

// renameWorkerAttribute handles attribute renames for worker resource references
func renameWorkerAttribute(oldResourceType, newResourceType, attrName string) string {
	// Handle workers_script attribute renames - both old->new and new->new (already migrated resource type)
	if (oldResourceType == "cloudflare_worker_script" && newResourceType == "cloudflare_workers_script") ||
		(oldResourceType == "cloudflare_workers_script" && newResourceType == "cloudflare_workers_script") {
		if attrName == "name" {
			return "script_name"
		}
	}

	// Handle workers_route attribute renames
	if oldResourceType == "cloudflare_worker_route" && newResourceType == "cloudflare_workers_route" {
		if attrName == "script_name" {
			return "script"
		}
	}

	return attrName // No change
}

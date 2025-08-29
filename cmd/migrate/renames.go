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
	// Map of old resource names to new ones
	resourceTypeRenames := map[string]string{
		"cloudflare_access_policy": "cloudflare_zero_trust_access_policy",
		"cloudflare_access_application": "cloudflare_zero_trust_access_application",
		"cloudflare_access_group": "cloudflare_zero_trust_access_group",
		// Add more renames as needed
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
					i++
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

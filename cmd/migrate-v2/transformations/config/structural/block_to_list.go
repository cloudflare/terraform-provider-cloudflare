package structural

import (
	"sort"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// BlocksToListConverter converts multiple blocks of a type to a list attribute
func BlocksToListConverter(blockType string) basic.TransformerFunc {
	return func(block *hclwrite.Block, ctx *basic.TransformContext) error {
		body := block.Body()
		blocks := body.Blocks()
		
		var items []hclwrite.Tokens
		blocksToRemove := []*hclwrite.Block{}
		
		for _, b := range blocks {
			if b.Type() == blockType {
				tokens := blockToObjectTokens(b)
				items = append(items, tokens)
				blocksToRemove = append(blocksToRemove, b)
			}
		}
		
		if len(items) > 0 {
			// Remove all blocks of this type
			for _, b := range blocksToRemove {
				body.RemoveBlock(b)
			}
			
			// Create list attribute
			listTokens := createListTokens(items)
			body.SetAttributeRaw(blockType, listTokens)
		}
		
		return nil
	}
}

func blockToObjectTokens(block *hclwrite.Block) hclwrite.Tokens {
	tokens := hclwrite.Tokens{}
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte{'{'}})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
	
	// Get all attributes from the block
	attrs := block.Body().Attributes()
	
	// Collect attribute names and sort them alphabetically
	var attrNames []string
	for name := range attrs {
		attrNames = append(attrNames, name)
	}
	sort.Strings(attrNames)
	
	// Add attributes in alphabetical order
	for _, name := range attrNames {
		attr := attrs[name]
		
		// Add indentation (6 spaces for proper nesting)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})
		
		// Add attribute name
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)})
		
		// Add padding to align equals signs (optional, for consistent formatting)
		// Calculate padding based on longest attribute name for alignment
		padding := ""
		maxLen := 0
		for _, n := range attrNames {
			if len(n) > maxLen {
				maxLen = len(n)
			}
		}
		for i := len(name); i < maxLen; i++ {
			padding += " "
		}
		if padding != "" {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(padding)})
		}
		
		// Add equals sign with spaces
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
		
		// Add the attribute value
		tokens = append(tokens, attr.Expr().BuildTokens(nil)...)
		
		// Add newline
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
	}
	
	// Process nested blocks (like uri blocks inside destinations)
	nestedBlocks := block.Body().Blocks()
	for _, nestedBlock := range nestedBlocks {
		// Add indentation (6 spaces for proper nesting)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})
		
		// Add the nested block type as an attribute with object value
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(nestedBlock.Type())})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
		
		// Recursively convert the nested block to object
		tokens = append(tokens, nestedBlockToObjectTokens(nestedBlock)...)
		
		// Add newline
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
	}
	
	// Use 4 spaces for the closing brace
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte{'}'}})
	return tokens
}

func nestedBlockToObjectTokens(block *hclwrite.Block) hclwrite.Tokens {
	tokens := hclwrite.Tokens{}
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrace, Bytes: []byte{'{'}})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
	
	// Get all attributes from the nested block
	attrs := block.Body().Attributes()
	for name, attr := range attrs {
		// Add indentation (8 spaces for nested attributes)
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("        ")})
		
		// Add attribute name
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(name)})
		
		// Add equals sign
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
		
		// Add the attribute value
		tokens = append(tokens, attr.Expr().BuildTokens(nil)...)
		
		// Add newline
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
	}
	
	// Handle further nested blocks if any
	for _, nestedBlock := range block.Body().Blocks() {
		// Add indentation
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("        ")})
		
		// Add the nested block type as an attribute
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(nestedBlock.Type())})
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenEqual, Bytes: []byte(" = ")})
		
		// Recursively convert
		tokens = append(tokens, nestedBlockToObjectTokens(nestedBlock)...)
		
		// Add newline
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
	}
	
	// Closing brace with proper indentation
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("      ")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrace, Bytes: []byte{'}'}})
	return tokens
}

func createListTokens(items []hclwrite.Tokens) hclwrite.Tokens {
	tokens := hclwrite.Tokens{}
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenOBrack, Bytes: []byte{'['}})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
	
	for i, item := range items {
		// Use 4 spaces for indentation of list items
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("    ")})
		tokens = append(tokens, item...)
		if i < len(items)-1 {
			tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte{','}})
		}
		tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte{'\n'}})
	}
	
	// Use 2 spaces for the closing bracket
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("  ")})
	tokens = append(tokens, &hclwrite.Token{Type: hclsyntax.TokenCBrack, Bytes: []byte{']'}})
	return tokens
}
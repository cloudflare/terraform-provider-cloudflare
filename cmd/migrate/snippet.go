package main

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// isSnippetResource checks if a block is a cloudflare_snippet resource
func isSnippetResource(block *hclwrite.Block) bool {
	return block.Type() == "resource" && len(block.Labels()) >= 1 && block.Labels()[0] == "cloudflare_snippet"
}

// transformSnippetBlock handles migration of cloudflare_snippet resources in config
func transformSnippetBlock(block *hclwrite.Block, diags ast.Diagnostics) {
	body := block.Body()

	// Rename "name" to "snippet_name"
	if nameAttr := body.GetAttribute("name"); nameAttr != nil {
		tokens := nameAttr.Expr().BuildTokens(nil)
		body.SetAttributeRaw("snippet_name", tokens)
		body.RemoveAttribute("name")
	}

	// Move "main_module" to nested "metadata" object
	if mainModuleAttr := body.GetAttribute("main_module"); mainModuleAttr != nil {
		// Get the raw tokens from the original expression to preserve variable references
		mainModuleTokens := mainModuleAttr.Expr().BuildTokens(nil)

		// Build the metadata object with raw tokens to preserve the original expression
		var metadataTokens hclwrite.Tokens
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenOBrace,
			Bytes: []byte("{"),
		})
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte("\n"),
		})
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte("    main_module"),
		})
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(" "),
		})
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte("="),
		})
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(" "),
		})
		// Append the original expression tokens (preserves variables, literals, etc.)
		metadataTokens = append(metadataTokens, mainModuleTokens...)
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte("\n"),
		})
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte("  "),
		})
		metadataTokens = append(metadataTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenCBrace,
			Bytes: []byte("}"),
		})

		// Set metadata attribute using raw tokens
		body.SetAttributeRaw("metadata", metadataTokens)
		body.RemoveAttribute("main_module")
	}

	// Convert "files" blocks to list attribute
	filesBlocks := body.Blocks()
	var blocksToRemove []*hclwrite.Block
	var hasFiles bool

	// First pass: check if we have files blocks
	for _, filesBlock := range filesBlocks {
		if filesBlock.Type() == "files" {
			hasFiles = true
			blocksToRemove = append(blocksToRemove, filesBlock)
		}
	}

	// Remove all files blocks
	for _, filesBlock := range blocksToRemove {
		body.RemoveBlock(filesBlock)
	}

	// Add files as a list attribute if we have any files
	// We'll preserve the raw tokens to maintain heredoc format
	if hasFiles {
		// Build the files list manually using raw tokens to preserve heredoc strings
		var filesTokens hclwrite.Tokens
		filesTokens = append(filesTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenOBrack,
			Bytes: []byte("["),
		})
		filesTokens = append(filesTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenNewline,
			Bytes: []byte("\n"),
		})

		// Process each files block
		for i, filesBlock := range blocksToRemove {
			if filesBlock.Type() == "files" {
				// Add indentation
				filesTokens = append(filesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte("    "),
				})
				filesTokens = append(filesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenOBrace,
					Bytes: []byte("{"),
				})
				filesTokens = append(filesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenNewline,
					Bytes: []byte("\n"),
				})

				fileBody := filesBlock.Body()

				// Add name attribute
				if nameAttr := fileBody.GetAttribute("name"); nameAttr != nil {
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("      name"),
					})
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("    "),
					})
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenEqual,
						Bytes: []byte("="),
					})
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte(" "),
					})
					// Append the raw tokens from the name expression
					filesTokens = append(filesTokens, nameAttr.Expr().BuildTokens(nil)...)
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenNewline,
						Bytes: []byte("\n"),
					})
				}

				// Add content attribute - preserve heredoc format
				if contentAttr := fileBody.GetAttribute("content"); contentAttr != nil {
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte("      content"),
					})
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte(" "),
					})
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenEqual,
						Bytes: []byte("="),
					})
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenIdent,
						Bytes: []byte(" "),
					})
					// Append the raw tokens from the content expression (preserves heredoc)
					filesTokens = append(filesTokens, contentAttr.Expr().BuildTokens(nil)...)
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenNewline,
						Bytes: []byte("\n"),
					})
				}

				// Close object
				filesTokens = append(filesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte("    "),
				})
				filesTokens = append(filesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenCBrace,
					Bytes: []byte("}"),
				})

				// Add comma if not the last item
				if i < len(blocksToRemove)-1 {
					filesTokens = append(filesTokens, &hclwrite.Token{
						Type:  hclsyntax.TokenComma,
						Bytes: []byte(","),
					})
				}
				filesTokens = append(filesTokens, &hclwrite.Token{
					Type:  hclsyntax.TokenNewline,
					Bytes: []byte("\n"),
				})
			}
		}

		// Close list
		filesTokens = append(filesTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte("  "),
		})
		filesTokens = append(filesTokens, &hclwrite.Token{
			Type:  hclsyntax.TokenCBrack,
			Bytes: []byte("]"),
		})

		// Set the files attribute using raw tokens
		body.SetAttributeRaw("files", filesTokens)
	}
}


package main

import (
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
}

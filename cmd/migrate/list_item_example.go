package main

import (
	"slices"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func transformListItem(config []byte, filename string) ([]byte, error) {
	diags := ast.NewDiagnostics()
	body := ast.ParseIntoSyntaxBody(config, filename, diags)

	newBlocks := hclsyntax.Blocks{}
	for _, block := range body.Blocks {
		// we want to translate existing blocks, and add new resource blocks for "cloudflare_list_item"
		newBlocks = append(newBlocks, block)

		if block.Type == "resource" && len(block.Labels) == 2 && block.Labels[0] == "cloudflare_list" {
			expr, ok := ast.Body2Expr(*block.Body, diags).(*hclsyntax.ObjectConsExpr)
			if !ok {
				continue
			}

			i := slices.IndexFunc(expr.Items, func(item hclsyntax.ObjectConsItem) bool {
				key := ast.Expr2S(item.KeyExpr, diags)
				return key == "item"
			})
			if i == -1 {
				continue
			}
			item := expr.Items[i]

			// we want to extract out the list items
			block.Body.Blocks = slices.DeleteFunc(block.Body.Blocks, func(b *hclsyntax.Block) bool {
				label := strings.Join(b.Labels, "")
				return b.Type == "item" || (b.Type == "dynamic" && label == "item")
			})

			// construct new resource based on transformed AST
			labels := []string{"cloudflare_list_item", block.Labels[1]}
			syn := &hclsyntax.Block{
				Type:   "resource",
				Labels: labels,
				Body: &hclsyntax.Body{
					Attributes: hclsyntax.Attributes{
						"for_each": &hclsyntax.Attribute{
							Name: "for_each",
							Expr: item.ValueExpr,
						},
						"account_id": &hclsyntax.Attribute{
							Name: "account_id",
							Expr: ast.NewScopeTraversal(append(slices.Clone(labels), "account_id")...),
						},
						"list_id": &hclsyntax.Attribute{
							Name: "list_id",
							Expr: ast.NewScopeTraversal(append(slices.Clone(labels), "id")...),
						},
						"comment": &hclsyntax.Attribute{Name: "comment",
							Expr: ast.NewScopeTraversal("each", "value", "content", "comment"),
						},
						"ip": &hclsyntax.Attribute{Name: "ip",
							Expr: ast.NewScopeTraversal("each", "value", "content", "ip", "value"),
						},
					},
				},
			}

			newBlocks = append(newBlocks, syn)
		}
	}

	out := ast.Blocks2S(newBlocks, diags)
	return []byte(out), nil
}

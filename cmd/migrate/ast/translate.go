package ast

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func dynBlock2Expr(block hclsyntax.Block, diags Diagnostics) hclsyntax.Expression {
	if block.Type != "dynamic" {
		return &hclsyntax.TupleConsExpr{Exprs: []hclsyntax.Expression{Body2Expr(*block.Body, diags)}}
	}

	valVar := strings.Join(block.Labels, "")
	forEach, ok := block.Body.Attributes["for_each"]
	delete(block.Body.Attributes, "for_each")
	if !ok {
		return &hclsyntax.TupleConsExpr{Exprs: []hclsyntax.Expression{Body2Expr(*block.Body, diags)}}
	}

	if varVal, ok := block.Body.Attributes["iterator"]; ok {
		delete(block.Body.Attributes, "iterator")
		valVar = Expr2S(varVal.Expr, diags)
	}

	return &hclsyntax.ForExpr{
		ValVar:  valVar,
		ValExpr: Body2Expr(*block.Body, diags),
		CollExpr: &hclsyntax.ForExpr{
			ValVar: "value",
			ValExpr: &hclsyntax.ObjectConsExpr{
				Items: []hclsyntax.ObjectConsItem{
					{
						KeyExpr: NewKeyExpr("key"),
						ValueExpr: &hclsyntax.ScopeTraversalExpr{
							Traversal: hcl.Traversal{
								hcl.TraverseRoot{Name: "value"},
							},
						},
					},
					{
						KeyExpr: NewKeyExpr("value"),
						ValueExpr: &hclsyntax.ScopeTraversalExpr{
							Traversal: hcl.Traversal{
								hcl.TraverseRoot{Name: "value"},
							},
						},
					},
				},
			},
			CollExpr: forEach.Expr,
		},
	}
}

// translate a dynamic "<block-name>" block into an equivalent [for xyz in ... : ...] list comprehension of an attribute that is the same name
//
// should probably also take a callback to modify the nodes, pending on specific requirements
func Body2Expr(body hclsyntax.Body, diags Diagnostics) hclsyntax.Expression {
	items := []hclsyntax.ObjectConsItem{}

	for _, attr := range OrderedAttributes(body.Attributes) {
		items = append(items, hclsyntax.ObjectConsItem{
			KeyExpr:   NewKeyExpr(attr.Name),
			ValueExpr: attr.Expr,
		})
	}

	dynBlocks := map[string]hclsyntax.Blocks{}
	for _, block := range body.Blocks {
		if block.Type == "dynamic" {
			label := strings.Join(block.Labels, "")
			dynBlocks[label] = append(dynBlocks[label], block)
		}
	}

	for _, block := range body.Blocks {
		if block.Type == "dynamic" {
			continue
		}
		if blocks, ok := dynBlocks[block.Type]; ok {
			dynBlocks[block.Type] = append(blocks, block)
			continue
		}

		items = append(items, hclsyntax.ObjectConsItem{
			KeyExpr:   NewKeyExpr(block.Type),
			ValueExpr: Body2Expr(*block.Body, diags),
		})
	}

	for key, blocks := range dynBlocks {
		switch len(blocks) {
		case 0:
			continue
		case 1:
			expr := dynBlock2Expr(*blocks[0], diags)
			items = append(items, hclsyntax.ObjectConsItem{KeyExpr: NewKeyExpr(key), ValueExpr: expr})
		default:
			acc := []hclsyntax.Expression{}
			for _, block := range blocks {
				expr := dynBlock2Expr(*block, diags)
				acc = append(acc, expr)
			}
			expr := hclsyntax.FunctionCallExpr{Name: "concat", Args: acc}
			items = append(items, hclsyntax.ObjectConsItem{KeyExpr: NewKeyExpr(key), ValueExpr: &expr})
		}
	}

	return &hclsyntax.ObjectConsExpr{Items: items}
}

func Block2AttrsRoot(body hclsyntax.Body, diags Diagnostics) hclsyntax.Body {
	for _, block := range body.Blocks {
		expr, ok := Body2Expr(*block.Body, diags).(*hclsyntax.ObjectConsExpr)
		if !ok {
			continue
		}

		for _, item := range expr.Items {
			key := Expr2S(item.KeyExpr, diags)
			block.Body.Attributes[key] = &hclsyntax.Attribute{
				Name: key,
				Expr: item.ValueExpr,
			}
		}

		block.Body.Blocks = nil
	}

	return body
}

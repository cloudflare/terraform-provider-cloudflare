package ast

import (
	"slices"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Embed block & object so we can make them satisfy a common interface
type Block struct {
	*hclwrite.Block
}

type Object struct {
	*hclsyntax.ObjectConsExpr
	// map for more convenient read/update access to object attributes
	attrs map[string]*hclsyntax.Expression
}

func NewObject(obj *hclsyntax.ObjectConsExpr, diags Diagnostics) Object {
	attrs := map[string]*hclsyntax.Expression{}
	for i := range obj.Items {
		key := Expr2S(obj.Items[i].KeyExpr, diags)
		attrs[key] = &obj.Items[i].ValueExpr
	}
	return Object{obj, attrs}
}

type HasAttributes interface {
	// Get list of attributes
	Attributes(diags Diagnostics) map[string]*hclsyntax.Expression
	SetAttribute(key string, val hclsyntax.Expression, diags Diagnostics)
	RemoveAttribute(key string, diags Diagnostics)
}

func (b Block) Attributes(diags Diagnostics) map[string]*hclsyntax.Expression {
	attrs := map[string]*hclsyntax.Expression{}
	for k, v := range b.Body().Attributes() {
		expr := WriteExpr2Expr(v.Expr(), diags)
		attrs[k] = &expr
	}
	return attrs
}

func (b Block) SetAttribute(key string, val hclsyntax.Expression, diags Diagnostics) {
	e := Expr2WriteExpr(val, diags)
	eTokens := e.BuildTokens(nil)
	b.Body().SetAttributeRaw(key, eTokens)
}

func (b Block) RemoveAttribute(key string, diag Diagnostics) {
	b.Body().RemoveAttribute(key)
}

func (o Object) Attributes(diags Diagnostics) map[string]*hclsyntax.Expression {
	return o.attrs
}

func (o Object) SetAttribute(key string, val hclsyntax.Expression, diags Diagnostics) {
	if o.attrs[key] == nil {
		// object doesn't have this attribute yet, create it & add to map
		attr := hclsyntax.ObjectConsItem{
			KeyExpr:   NewKeyExpr(key),
			ValueExpr: val,
		}
		o.Items = append(o.Items, attr)
		o.attrs[key] = &attr.ValueExpr
	} else {
		// already exists, just update it
		*(o.attrs[key]) = val
	}
}

func (o Object) RemoveAttribute(key string, diags Diagnostics) {
	// Remove attribute from attrs map & underlying list of items
	delete(o.attrs, key)
	o.Items = slices.DeleteFunc(o.Items, func(item hclsyntax.ObjectConsItem) bool {
		return Expr2S(item.KeyExpr, diags) == key
	})
}

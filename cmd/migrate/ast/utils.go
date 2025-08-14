package ast

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/hashicorp/hcl/v2"
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
		expr := WriteExpr2Expr(*v.Expr(), diags)
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
		o.attrs[key] = &o.Items[len(o.Items)-1].ValueExpr
	} else {
		// already exists, just update it
		*(o.attrs[key]) = val
	}
}

func (o Object) RemoveAttribute(key string, diags Diagnostics) {
	// Remove attribute from attrs map & underlying list of items
	delete(o.attrs, key)
	acc := []hclsyntax.ObjectConsItem{}
	for _, item := range o.Items {
		k := Expr2S(item.KeyExpr, diags)
		if k != key {
			acc = append(acc, item)
		}
	}
	o.Items = acc
}

// ExprTransformer mutates an attribute value in place
// or sets to nil if the attribute should be removed
type ExprTransformer func(*hclsyntax.Expression, Diagnostics)

func NewKeyExpr(key string) hclsyntax.Expression {
	return &hclsyntax.ObjectConsKeyExpr{
		Wrapped: &hclsyntax.ScopeTraversalExpr{
			Traversal: hcl.Traversal{
				hcl.TraverseRoot{Name: key},
			},
		},
	}
}

func ApplyTransformToAttributes(objOrBlock HasAttributes, transforms map[string]ExprTransformer, diags Diagnostics) {

	unusedTransforms := maps.Clone(transforms)

	for key, v := range objOrBlock.Attributes(diags) {
		if transform, ok := transforms[key]; ok {

			// transform it in place
			transform(v, diags)

			// if result of transform is nil, remove the attribute
			// (nil indicates we should remove it)
			if *v == nil {
				objOrBlock.RemoveAttribute(key, diags)
			} else {
				// For Block types, we need to set the attribute back since
				// Attributes() returns copies, not references to the originals
				objOrBlock.SetAttribute(key, *v, diags)
			}

			// remove key from unusedTransforms b/c we just used it
			delete(unusedTransforms, key)

		}
	}

	// unusedTransforms apply to attributes that weren't present in the original object
	// we call each of them with nil, which gives them the opportunity to add the attribute
	for key, transform := range unusedTransforms {
		var newVal hclsyntax.Expression = nil
		transform(&newVal, diags)
		if newVal != nil {
			objOrBlock.SetAttribute(key, newVal, diags)
		}
	}
}

func RawWarning(msg string) *hclwrite.Expression {
	eot := "WARNING"
	heredoc := hclwrite.Tokens{&hclwrite.Token{Type: hclsyntax.TokenOHeredoc, Bytes: []byte("<<-" + eot + "\n")}}
	for _, m := range strings.Split(msg, "\n") {
		str := strings.TrimRight(m, " \t") + "\n"
		if str != "\n" {
			str = "  " + str
		}
		heredoc = append(heredoc, &hclwrite.Token{Type: hclsyntax.TokenStringLit, Bytes: []byte(str)})
	}
	heredoc = append(heredoc, &hclwrite.Token{Type: hclsyntax.TokenCHeredoc, Bytes: []byte(eot)})

	tokens := slices.Concat(
		hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte("regex")},
			&hclwrite.Token{Type: hclsyntax.TokenOParen, Bytes: []byte("(")},
		},
		heredoc,
		hclwrite.Tokens{
			&hclwrite.Token{Type: hclsyntax.TokenNewline, Bytes: []byte("\n")},
			&hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(", ")},

			&hclwrite.Token{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
			&hclwrite.Token{Type: hclsyntax.TokenStringLit, Bytes: []byte("")},
			&hclwrite.Token{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},

			&hclwrite.Token{Type: hclsyntax.TokenCParen, Bytes: []byte(")")},
		},
	)
	return hclwrite.NewExpressionRaw(tokens)
}

func warnManualMessage(path string, previous string) string {
	msg := fmt.Sprintf(`
> We were unable to automatically migrate this resource to the new provider.
> Please refer to %q for manual migration.
		`,
		"https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/"+path,
	)

	if previous != "" {
		msg = msg + "\n> Your previous configuration was:\n" + previous
	}
	return msg
}

func WarnManualMigration4Attr(path string, previous *hclwrite.Attribute) *hclwrite.Expression {
	str := ""
	if previous != nil {
		str = string(previous.BuildTokens(nil).Bytes())
	}

	msg := warnManualMessage(path, str)
	return RawWarning(msg)
}

func WarnManualMigration4Expr(path string, previous *hclwrite.Expression) *hclwrite.Expression {
	str := ""
	if previous != nil {
		str = string(previous.BuildTokens(nil).Bytes())
	}

	msg := warnManualMessage(path, str)
	return RawWarning(msg)
}

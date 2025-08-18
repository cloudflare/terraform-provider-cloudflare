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

func rawWarning(msg string) *hclwrite.Expression {
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

func warning(msg string, diags Diagnostics) *hclsyntax.Expression {
	raw := rawWarning(msg)
	e := WriteExpr2Expr(raw, diags)
	return &e
}

func WarnManualMigration4AttrWrite(path string, previous *hclwrite.Attribute) *hclwrite.Expression {
	str := ""
	if previous != nil {
		str = string(previous.BuildTokens(nil).Bytes())
	}

	msg := warnManualMessage(path, str)
	return rawWarning(msg)
}

func WarnManualMigration4ExprWrite(path string, previous *hclwrite.Expression) *hclwrite.Expression {
	str := ""
	if previous != nil {
		str = string(previous.BuildTokens(nil).Bytes())
	}

	msg := warnManualMessage(path, str)
	return rawWarning(msg)
}

func WarnManualMigration4Attr(path string, previous *hclsyntax.Attribute, diags Diagnostics) *hclsyntax.Expression {
	str := ""
	if previous != nil {
		str = Attr2S(*previous, diags)
	}

	msg := warnManualMessage(path, str)
	return warning(msg, diags)
}

func WarnManualMigration4Expr(path string, previous *hclsyntax.Expression, diags Diagnostics) *hclsyntax.Expression {
	str := ""
	if previous != nil {
		str = Expr2S(*previous, diags)
	}

	msg := warnManualMessage(path, str)
	return warning(msg, diags)
}

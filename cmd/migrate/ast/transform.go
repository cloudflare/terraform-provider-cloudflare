package ast

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func Str2Expr(str string, diags Diagnostics) hcl.Expression {
	expr, d := hclsyntax.ParseExpression([]byte(str), "dog.hcl", hcl.InitialPos)
	diags.HclDiagnostics.Extend(d)
	return expr
}

func WriteExpr2Expr(write hclwrite.Expression, diags Diagnostics) hclsyntax.Expression {
	raw := write.BuildTokens(nil).Bytes()
	expr, d := hclsyntax.ParseExpression(raw, "dog.hcl", hcl.InitialPos)
	diags.HclDiagnostics.Extend(d)
	return expr
}

func op2S(op hclsyntax.Operation) string {
	switch op.Impl {
	case hclsyntax.OpLogicalOr.Impl:
		return "||"
	case hclsyntax.OpLogicalAnd.Impl:
		return "&&"
	case hclsyntax.OpLogicalNot.Impl:
		return "!"
	case hclsyntax.OpEqual.Impl:
		return "=="
	case hclsyntax.OpNotEqual.Impl:
		return "!="
	case hclsyntax.OpGreaterThan.Impl:
		return ">"
	case hclsyntax.OpGreaterThanOrEqual.Impl:
		return ">="
	case hclsyntax.OpLessThan.Impl:
		return "<"
	case hclsyntax.OpLessThanOrEqual.Impl:
		return "<="
	case hclsyntax.OpAdd.Impl:
		return "+"
	case hclsyntax.OpSubtract.Impl:
		return "-"
	case hclsyntax.OpMultiply.Impl:
		return "*"
	case hclsyntax.OpDivide.Impl:
		return "/"
	case hclsyntax.OpModulo.Impl:
		return "-"
	case hclsyntax.OpNegate.Impl:
		return "-"
	default:
		return ""
	}
}

func traversal2S(tr hcl.Traversal) string {
	if tr.IsRelative() {
		tr = hcl.TraversalJoin(hcl.Traversal{hcl.TraverseRoot{}}, tr)
	}
	raw := (hclwrite.NewExpressionAbsTraversal(tr).BuildTokens(nil).Bytes())
	str, _ := strings.CutPrefix(string(raw), ".")
	return str
}

func Expr2S(expr hcl.Expression, diag Diagnostics) (token string) {
	switch e := expr.(type) {

	case *hclsyntax.AnonSymbolExpr:
		diag.ComplicatedHCL.Append(e)

	case *hclsyntax.BinaryOpExpr:
		op := op2S(*e.Op)
		if op == "" {
			diag.ComplicatedHCL.Append(e)
		}
		lhs := Expr2S(e.LHS, diag)
		rhs := Expr2S(e.RHS, diag)

		token = lhs + " " + op + " " + rhs

	case *hclsyntax.ConditionalExpr:
		cond := Expr2S(e.Condition, diag)
		ye := Expr2S(e.TrueResult, diag)
		na := Expr2S(e.FalseResult, diag)
		token = cond + " ? " + ye + " : " + na

	/*
		List comprehension:
		[ for val_var in coll_expr : val_expr if cond_expr]
		Dict comprehension:
		{ for key_var, val_var in coll_expr : key_expr => val_expr if cond_expr }
	*/
	case *hclsyntax.ForExpr:
		coll := Expr2S(e.CollExpr, diag)
		valLine := e.ValVar
		val := Expr2S(e.ValExpr, diag)
		// If it's a list comprehension, KeyVar will be empty and KeyExpr will be nil
		// If it's a dict comprehension both should be non-empty/non-nil
		if e.KeyVar != "" {
			valLine = e.KeyVar + ", " + valLine
		}
		if e.KeyExpr != nil {
			key := Expr2S(e.KeyExpr, diag)
			val = key + " => " + val
		}
		forExp := "for " + valLine + " in " + coll + " : " + val
		if e.CondExpr != nil {
			cond := Expr2S(e.CondExpr, diag)
			forExp = token + " if " + cond
		}
		if e.KeyVar != "" {
			token = "{" + forExp + "}"
		} else {
			token = "[" + forExp + "]"
		}

	case *hclsyntax.FunctionCallExpr:
		args := []string{}
		for _, a := range e.Args {
			arg := Expr2S(a, diag)
			args = append(args, arg)
		}
		token = e.Name + "(" + strings.Join(args, ",") + ")"

	case *hclsyntax.IndexExpr:
		lhs := Expr2S(e.Collection, diag)
		rhs := Expr2S(e.Key, diag)
		token = lhs + "[" + rhs + "]"

	case *hclsyntax.LiteralValueExpr:
		raw := hclwrite.TokensForValue(e.Val).Bytes()
		token = string(raw)

	case *hclsyntax.ObjectConsExpr:
		lines := []string{}
		for _, item := range e.Items {
			key := Expr2S(item.KeyExpr, diag)
			val := Expr2S(item.ValueExpr, diag)
			lines = append(lines, "  "+key+" = "+val)
		}
		token = "{" + strings.Join(lines, "\n") + "}"

	// TODO might need to wrap inner key in extra syntax,
	// e.g. ["my syntactically weird key!?"]
	// but unlikely - would only come up in dynamic blocks if users are
	// making bad choices about key names
	case *hclsyntax.ObjectConsKeyExpr:
		return Expr2S(e.Wrapped, diag)

	case *hclsyntax.ParenthesesExpr:
		inner := Expr2S(e.Expression, diag)
		token = "(" + inner + ")"

	case *hclsyntax.RelativeTraversalExpr:
		token = traversal2S(e.Traversal)
	case *hclsyntax.ScopeTraversalExpr:
		token = traversal2S(e.Traversal)

	case *hclsyntax.SplatExpr:
		src := Expr2S(e.Source, diag)
		token = src + "[*]"

	case *hclsyntax.TemplateExpr:

		if e.IsStringLiteral() {
			// TODO may need to be quoted
			return Expr2S(e.Parts[0], diag)
		}
		parts := []string{}
		for _, p := range e.Parts {
			if lit, ok := p.(*hclsyntax.LiteralValueExpr); ok {
				if lit.Val.Type().Equals(cty.Bool) || lit.Val.Type().Equals(cty.Number) || lit.Val.Type().Equals(cty.String) {
					parts = append(parts, lit.Val.AsString())
					continue
				}
			}

			part := Expr2S(p, diag)
			parts = append(parts, "${"+part+"}")
		}
		token = `"` + strings.Join(parts, "") + `"`

	// TODO not sure if this is right
	case *hclsyntax.TemplateJoinExpr:
		return Expr2S(e.Tuple, diag)

	case *hclsyntax.TemplateWrapExpr:
		part := Expr2S(e.Wrapped, diag)
		token = `"${` + part + `}"`

	case *hclsyntax.TupleConsExpr:
		items := []string{}
		for _, exp := range e.Exprs {
			arg := Expr2S(exp, diag)
			items = append(items, arg)
		}
		token = "[" + strings.Join(items, ",") + "]"

	case *hclsyntax.UnaryOpExpr:
		op := op2S(*e.Op)
		if op == "" {
			// We couldn't figure out what unary operation this is
			diag.ComplicatedHCL.Append((e))
		}
		val := Expr2S(e.Val, diag)
		token = op + val

	default:
		// What expression is this??
		diag.ComplicatedHCL.Append((e))
	}

	return
}

func Expr2WriteExpr(expr hcl.Expression, diag Diagnostics) hclwrite.Expression {
	str := Expr2S(expr, diag)

	file := fmt.Sprintf(`
resource "dog" "dog" {
	dog = %s
}
`, str)

	parsed, d := hclwrite.ParseConfig([]byte(file), "dog.hcl", hcl.InitialPos)
	diag.HclDiagnostics.Extend(d)
	// TODO use hclformat to format parsed config before returning it?
	attr := parsed.Body().Blocks()[0].Body().GetAttribute("dog")
	return *attr.Expr()
}

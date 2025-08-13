package ast

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func Str2Expr(str string) (hcl.Expression, hcl.Diagnostics) {
	expr, d := hclsyntax.ParseExpression([]byte(str), "dog.hcl", hcl.InitialPos)
	return expr, d
}

func WriteExpr2Expr(write hclwrite.Expression) (hcl.Expression, hcl.Diagnostics) {
	raw := write.BuildTokens(nil).Bytes()
	expr, d := hclsyntax.ParseExpression(raw, "dog.hcl", hcl.InitialPos)
	return expr, d
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

func Expr2S(expr hcl.Expression) (token string, interesting []hcl.Expression) {
	switch e := expr.(type) {
	case *hclsyntax.AnonSymbolExpr:
		interesting = append(interesting, e)

	case *hclsyntax.BinaryOpExpr:
		op := op2S(*e.Op)
		if op == "" {
			interesting = append(interesting, e)
		}
		lhs, i := Expr2S(e.LHS)
		interesting = append(interesting, i...)
		rhs, i := Expr2S(e.RHS)
		interesting = append(interesting, i...)
		token = lhs + " " + op + " " + rhs

	case *hclsyntax.ConditionalExpr:
		cond, i := Expr2S(e.Condition)
		interesting = append(interesting, i...)
		ye, i := Expr2S(e.TrueResult)
		interesting = append(interesting, i...)
		na, i := Expr2S(e.FalseResult)
		interesting = append(interesting, i...)
		token = cond + " ? " + ye + " : " + na

	case *hclsyntax.ForExpr:
		coll, i := Expr2S(e.CollExpr)
		interesting = append(interesting, i...)
		valLine := e.ValVar
		val, i := Expr2S(e.ValExpr)
		interesting = append(interesting, i...)
		if e.KeyVar != "" {
			valLine = e.KeyVar + ", " + valLine
		}
		if e.KeyExpr != nil {
			key, i := Expr2S(e.KeyExpr)
			interesting = append(interesting, i...)
			val = key + " => " + val + val
		}
		forExp := "for " + valLine + " in " + coll + " : " + val
		if e.CondExpr != nil {
			cond, i := Expr2S(e.CondExpr)
			interesting = append(interesting, i...)
			forExp = token + " if " + cond
		}
		if e.KeyVar != "" {
			token = "{" + forExp + "}"
		} else {
			token = "(" + forExp + ")"
		}

	case *hclsyntax.FunctionCallExpr:
		args := []string{}
		for _, a := range e.Args {
			arg, i := Expr2S(a)
			interesting = append(interesting, i...)
			args = append(args, arg)
		}
		token = e.Name + "(" + strings.Join(args, ",") + ")"

	case *hclsyntax.IndexExpr:
		lhs, i := Expr2S(e.Collection)
		interesting = append(interesting, i...)
		rhs, i := Expr2S(e.Key)
		interesting = append(interesting, i...)
		token = lhs + "[" + rhs + "]"

	case *hclsyntax.LiteralValueExpr:
		raw := hclwrite.TokensForValue(e.Val).Bytes()
		token = string(raw)

	case *hclsyntax.ObjectConsExpr:
		lines := []string{}
		for _, item := range e.Items {
			key, i := Expr2S(item.KeyExpr)
			interesting = append(interesting, i...)
			val, i := Expr2S(item.ValueExpr)
			interesting = append(interesting, i...)
			lines = append(lines, "  "+key+" = "+val)
		}
		token = "{" + strings.Join(lines, "\n") + "}"

	case *hclsyntax.ObjectConsKeyExpr:
		return Expr2S(e.Wrapped)

	case *hclsyntax.ParenthesesExpr:
		inner, i := Expr2S(e.Expression)
		interesting = append(interesting, i...)
		token = "(" + inner + ")"

	case *hclsyntax.RelativeTraversalExpr:
		token = traversal2S(e.Traversal)
	case *hclsyntax.ScopeTraversalExpr:
		token = traversal2S(e.Traversal)

	case *hclsyntax.SplatExpr:
		src, i := Expr2S(e.Source)
		interesting = append(interesting, i...)
		token = src + "[*]"

	case *hclsyntax.TemplateExpr:
		if e.IsStringLiteral() {
			return Expr2S(e.Parts[0])
		}
		parts := []string{}
		for _, p := range e.Parts {
			if lit, ok := p.(*hclsyntax.LiteralValueExpr); ok {
				if lit.Val.Type().Equals(cty.Bool) || lit.Val.Type().Equals(cty.Number) || lit.Val.Type().Equals(cty.String) {
					parts = append(parts, lit.Val.AsString())
					continue
				}
			}

			part, i := Expr2S(p)
			interesting = append(interesting, i...)
			parts = append(parts, "${"+part+"}")
		}
		token = `"` + strings.Join(parts, "") + `"`

	case *hclsyntax.TemplateJoinExpr:
		return Expr2S(e.Tuple)

	case *hclsyntax.TemplateWrapExpr:
		part, i := Expr2S(e.Wrapped)
		interesting = append(interesting, i...)
		token = `"${` + part + `}"`

	case *hclsyntax.TupleConsExpr:
		items := []string{}
		for _, exp := range e.Exprs {
			arg, i := Expr2S(exp)
			interesting = append(interesting, i...)
			items = append(items, arg)
		}
		token = "[" + strings.Join(items, ",") + "]"

	case *hclsyntax.UnaryOpExpr:
		op := op2S(*e.Op)
		if op == "" {
			interesting = append(interesting, e)
		}
		val, i := Expr2S(e.Val)
		interesting = append(interesting, i...)
		token = op + val

	default:
		interesting = append(interesting, e)
	}

	return
}

func Expr2WriteExpr(expr hcl.Expression) (hclwrite.Expression, hcl.Diagnostics) {
	str, interesting := Expr2S(expr)
	if len(interesting) != 0 {
		spew.Dump(interesting)
	}

	file := fmt.Sprintf(`
resource "dog" "dog" {
	dog = %s
}
`, str)

	parsed, d := hclwrite.ParseConfig([]byte(file), "dog.hcl", hcl.InitialPos)
	attr := parsed.Body().Blocks()[0].Body().GetAttribute("dog")
	return *attr.Expr(), d
}

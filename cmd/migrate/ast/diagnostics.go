package ast

import "github.com/hashicorp/hcl/v2"

type Expressions []hcl.Expression

func (exps *Expressions) Append(exp hcl.Expression) {
	*exps = append(*exps, exp)
}

// Collection of diagnostics we'll accumulate during parsing & stringifying
type Diagnostics struct {
	// diagnostics from the hcl library
	HclDiagnostics hcl.Diagnostics

	// expressions we ran into and didn't know how to deal with
	// (ideally we'll handle these gracefully but we collect them here for debugging)
	ComplicatedHCL Expressions
}

func NewDiagnostics() Diagnostics {
	return Diagnostics{}
}

package certificate_pack

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"runtime"
	"testing"
)

func TestIsDeletedCertificatePackStatus(t *testing.T) {
	t.Parallel()

	tests := map[string]bool{
		"deleted":              true,
		"pending_deletion":     true,
		"deletion_timed_out":   true,
		"active":               false,
		"initializing":         false,
		"pending_validation":   false,
		"validation_timed_out": false,
		"":                     false,
	}

	for status, want := range tests {
		t.Run(status, func(t *testing.T) {
			t.Parallel()
			if got := isDeletedCertificatePackStatus(status); got != want {
				t.Fatalf("isDeletedCertificatePackStatus(%q) = %t, want %t", status, got, want)
			}
		})
	}
}

// Read is generated code, so this guards against a regeneration dropping the
// soft-deletion check: Read must test the status, remove the resource and
// return before it writes state.
func TestCertificatePackReadChecksSoftDeletion(t *testing.T) {
	t.Parallel()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("could not locate test source")
	}

	file, err := parser.ParseFile(token.NewFileSet(), filepath.Join(filepath.Dir(filename), "resource.go"), nil, 0)
	if err != nil {
		t.Fatalf("parse resource.go: %v", err)
	}

	var read *ast.FuncDecl
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Recv != nil && fn.Name.Name == "Read" {
			read = fn
		}
	}
	if read == nil {
		t.Fatal("CertificatePackResource.Read not found")
	}

	check, set := -1, -1
	for i, stmt := range read.Body.List {
		if ifStmt, ok := stmt.(*ast.IfStmt); ok && check < 0 && callsFunc(ifStmt.Cond, "isDeletedCertificatePackStatus") {
			if !callsFunc(ifStmt.Body, "RemoveResource") || !endsWithReturn(ifStmt.Body) {
				t.Fatal("the soft-deletion check in Read must remove the resource and return")
			}
			check = i
		}
		if set < 0 && callsSelector(stmt, "State", "Set") {
			set = i
		}
	}

	if check < 0 {
		t.Fatal("CertificatePackResource.Read must remove soft-deleted certificate packs from state")
	}
	if set >= 0 && set < check {
		t.Fatal("the soft-deletion check in Read must come before State.Set")
	}
}

func callsFunc(node ast.Node, name string) bool {
	found := false
	ast.Inspect(node, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return !found
		}
		switch fn := call.Fun.(type) {
		case *ast.Ident:
			found = found || fn.Name == name
		case *ast.SelectorExpr:
			found = found || fn.Sel.Name == name
		}
		return !found
	})
	return found
}

func callsSelector(node ast.Node, receiver, method string) bool {
	found := false
	ast.Inspect(node, func(n ast.Node) bool {
		sel, ok := n.(*ast.SelectorExpr)
		if ok && sel.Sel.Name == method {
			if inner, ok := sel.X.(*ast.SelectorExpr); ok && inner.Sel.Name == receiver {
				found = true
			}
		}
		return !found
	})
	return found
}

func endsWithReturn(block *ast.BlockStmt) bool {
	if len(block.List) == 0 {
		return false
	}
	_, ok := block.List[len(block.List)-1].(*ast.ReturnStmt)
	return ok
}

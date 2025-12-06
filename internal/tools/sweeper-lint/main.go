package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type Violation struct {
	File    string
	Line    int
	Message string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sweeper-lint <directory>")
		fmt.Println("Example: sweeper-lint ./internal/services")
		os.Exit(1)
	}

	directory := os.Args[1]
	violations := []Violation{}
	filesScanned := 0

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process resource_test.go files
		if !info.IsDir() && strings.HasSuffix(path, "resource_test.go") {
			fileViolations := lintFile(path)
			violations = append(violations, fileViolations...)
			filesScanned++
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// Print results
	if len(violations) == 0 {
		fmt.Printf("✓ No violations found in %d files\n", filesScanned)
		os.Exit(0)
	}

	fmt.Printf("Found %d violation(s) in %d file(s):\n\n", len(violations), filesScanned)

	for _, v := range violations {
		fmt.Printf("⚠ %s:%d: %s\n", v.File, v.Line, v.Message)
	}

	fmt.Printf("\nSummary: %d warning(s) in %d file(s)\n", len(violations), filesScanned)
	os.Exit(0) // Exit 0 for now (warnings only - change to os.Exit(1) when ready to enforce)
}

func lintFile(filePath string) []Violation {
	violations := []Violation{}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return violations
	}

	hasSweeperRegistration := false
	sweeperFuncName := ""

	// Find sweeper registration and function name
	ast.Inspect(node, func(n ast.Node) bool {
		if callExpr, ok := n.(*ast.CallExpr); ok {
			if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				if ident, ok := selExpr.X.(*ast.Ident); ok && ident.Name == "resource" {
					if selExpr.Sel.Name == "AddTestSweepers" {
						hasSweeperRegistration = true

						// Extract sweeper function name
						if len(callExpr.Args) >= 2 {
							if unary, ok := callExpr.Args[1].(*ast.UnaryExpr); ok {
								if compositeLit, ok := unary.X.(*ast.CompositeLit); ok {
									for _, elt := range compositeLit.Elts {
										if kv, ok := elt.(*ast.KeyValueExpr); ok {
											if key, ok := kv.Key.(*ast.Ident); ok && key.Name == "F" {
												if fn, ok := kv.Value.(*ast.Ident); ok {
													sweeperFuncName = fn.Name
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	if !hasSweeperRegistration {
		violations = append(violations, Violation{
			File:    filePath,
			Line:    1,
			Message: "Missing sweeper registration - should have resource.AddTestSweepers() in init()",
		})
		return violations
	}

	// Check sweeper function implementation
	if sweeperFuncName != "" {
		hasShouldSweepCheck := false
		isSingletonSweeper := false

		ast.Inspect(node, func(n ast.Node) bool {
			if funcDecl, ok := n.(*ast.FuncDecl); ok {
				if funcDecl.Name.Name == sweeperFuncName {
					// Check if this is a singleton sweeper (no for loops over resources)
					hasForLoop := false
					ast.Inspect(funcDecl.Body, func(inner ast.Node) bool {
						if _, ok := inner.(*ast.RangeStmt); ok {
							hasForLoop = true
						}
						if _, ok := inner.(*ast.ForStmt); ok {
							hasForLoop = true
						}
						return true
					})

					// If no for loops, it's likely a singleton sweeper that doesn't need filtering
					if !hasForLoop {
						isSingletonSweeper = true
					}

					// Check if function uses utils.ShouldSweepResource
					ast.Inspect(funcDecl.Body, func(inner ast.Node) bool {
						if callExpr, ok := inner.(*ast.CallExpr); ok {
							if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
								if ident, ok := selExpr.X.(*ast.Ident); ok {
									if ident.Name == "utils" && selExpr.Sel.Name == "ShouldSweepResource" {
										hasShouldSweepCheck = true
									}
								}
							}
						}
						return true
					})

					// Only report violation if it's not a singleton and doesn't have filtering
					if !hasShouldSweepCheck && !isSingletonSweeper {
						violations = append(violations, Violation{
							File:    filePath,
							Line:    fset.Position(funcDecl.Pos()).Line,
							Message: "Sweeper should use utils.ShouldSweepResource() for filtering",
						})
					}
				}
			}
			return true
		})
	}

	return violations
}

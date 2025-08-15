# Writing Source Migrations with AST Helpers

## Overview

When writing source migrations for Terraform configurations, use the AST helper functions in `ast/utils.go` to manipulate HCL code safely and reliably. This guide covers best practices and common patterns.

## Core Principles

### 1. Use AST Helpers, Not Regex
Never use regular expressions to modify HCL. The AST helpers preserve formatting, comments, and handle edge cases correctly.

### 2. Add Helpers to utils.go as Needed
If you need a common transformation pattern that doesn't exist, add it to `ast/utils.go` rather than implementing it inline. This promotes reusability and consistency.

### 3. Keep Transformations Focused
When writing transformers, avoid including special-case logic related to one attribute when editing another attribute. Each transformer should have a single, clear responsibility.

### 4. Handle Unparseable Content Gracefully
Always use `WarnManualMigration4Expr` or `WarnManualMigration4Attr` when encountering content that cannot be parsed or transformed automatically. This inserts a helpful warning message in the source file.

## Common Patterns

### Pattern 1: Simple Attribute Transformations

Use `ExprTransformer` functions for in-place modifications:

```go
func myTransformer(expr *hclsyntax.Expression, diags ast.Diagnostics) {
    if *expr == nil {
        return
    }
    
    // Transform the expression
    *expr = &hclsyntax.ObjectConsExpr{...}
}

// Apply to attributes
transforms := map[string]ast.ExprTransformer{
    "attribute_name": myTransformer,
}
ast.ApplyTransformToAttributes(obj, transforms, diags)
```

### Pattern 2: Complex Structural Changes

For transformations that change the structure (e.g., splitting objects, expanding arrays):

```go
func transformComplexStructure(expr *hclsyntax.Expression, diags ast.Diagnostics) {
    tup, ok := (*expr).(*hclsyntax.TupleConsExpr)
    if !ok {
        // Can't parse - add warning
        *expr = *ast.WarnManualMigration4Expr("resources/my_resource", expr, diags)
        return
    }
    
    // Perform transformation...
}
```

### Pattern 3: Working with Blocks vs Objects

Remember that `Block` and `Object` have different behaviors:
- `Block.Attributes()` returns copies - always call `SetAttribute` after transforming
- `Object.Attributes()` returns references - in-place modifications work directly

```go
// For blocks, you must set the attribute back after transformation
ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)

// For objects, modifications happen in-place
ast.ApplyTransformToAttributes(ast.NewObject(obj, diags), transforms, diags)
```

## Helper Functions Reference

### Core Types
- `Block` - Wrapper around `hclwrite.Block` for uniform interface
- `Object` - Wrapper around `hclsyntax.ObjectConsExpr` for uniform interface
- `ExprTransformer` - Function type for transforming expressions

### Key Functions
- `ApplyTransformToAttributes()` - Apply transformers to multiple attributes
- `NewKeyExpr()` - Create a key expression for object attributes
- `Expr2S()` - Convert expression to string
- `WriteExpr2Expr()` - Convert between write and syntax expressions

### Warning Functions
- `WarnManualMigration4Expr()` - Add warning for expressions that can't be migrated
- `WarnManualMigration4Attr()` - Add warning for attributes that can't be migrated
- `WarnManualMigration4ExprWrite()` - Add warning in write context
- `WarnManualMigration4AttrWrite()` - Add warning for attributes in write context

## Example: Complete Migration

```go
func transformMyResource(block *hclwrite.Block, diags ast.Diagnostics) {
    transforms := map[string]ast.ExprTransformer{
        "old_attr": renameToNewAttr,
        "complex_attr": handleComplexTransform,
    }
    
    ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)
}

func handleComplexTransform(expr *hclsyntax.Expression, diags ast.Diagnostics) {
    if *expr == nil {
        return
    }
    
    // Try to parse as expected type
    obj, ok := (*expr).(*hclsyntax.ObjectConsExpr)
    if !ok {
        // Can't parse - add warning with helpful message
        *expr = *ast.WarnManualMigration4Expr(
            "resources/my_resource#complex_attr",
            expr,
            diags,
        )
        return
    }
    
    // Perform transformation...
}
```

## Testing

Always write tests using the `RunTransformationTests` helper from `test_helpers.go`:

```go
func TestMyTransformation(t *testing.T) {
    tests := []TestCase{
        {
            Name: "transform complex attribute",
            Config: `resource "cloudflare_my_resource" "test" {
                old_attr = "value"
            }`,
            Expected: []string{`resource "cloudflare_my_resource" "test" {
                new_attr = "value"
            }`},
        },
    }
    
    RunTransformationTests(t, tests, transformFile)
}
```

## Best Practices Summary

1. **Always handle nil expressions** - Check for nil before dereferencing
2. **Use type assertions carefully** - Always check the `ok` value
3. **Add warnings for unparseable content** - Never silently skip transformations
4. **Test edge cases** - Empty values, missing attributes, complex nesting
5. **Document transformations** - Add clear comments with before/after examples
6. **Keep transformers pure** - Avoid side effects outside the expression being transformed
7. **Reuse existing helpers** - Check `ast/utils.go` before writing new logic

## Common Pitfalls to Avoid

### Don't Mix Concerns
```go
// BAD: Mixing email and group logic in one transformer
func transformEmail(expr *hclsyntax.Expression, diags ast.Diagnostics) {
    // ... email transformation ...
    
    // Don't do this - group logic doesn't belong here!
    if hasGroupAttribute {
        transformGroup(...)
    }
}

// GOOD: Separate transformers for separate concerns
transforms := map[string]ast.ExprTransformer{
    "email": transformEmail,
    "group": transformGroup,
}
```

### Always Add Warnings for Unparseable Content
```go
// BAD: Silently skipping
if !ok {
    return
}

// GOOD: Add warning to inform user
if !ok {
    *expr = *ast.WarnManualMigration4Expr("resources/my_resource", expr, diags)
    return
}
```

### Remember Block vs Object Semantics
```go
// For Block types, ApplyTransformToAttributes will call SetAttribute
// after transformation to persist changes
ast.ApplyTransformToAttributes(ast.Block{Block: block}, transforms, diags)

// For Object types, changes are made in-place through pointers
ast.ApplyTransformToAttributes(ast.NewObject(obj, diags), transforms, diags)
```
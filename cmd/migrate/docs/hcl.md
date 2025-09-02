# HashiCorp HCL v2 Library Documentation

## Overview

HCL (HashiCorp Configuration Language) is a toolkit for creating structured configuration languages that are both human- and machine-friendly. The HCL v2 library provides a set of Go packages for parsing, writing, and manipulating HCL configuration files.

## Core Packages

### 1. `github.com/hashicorp/hcl/v2`
The main HCL package providing fundamental types and interfaces.

**Key Types:**
- `Expression`: Represents configuration values that can be evaluated
- `Traversal`: Describes how to navigate through values (attribute lookup, indexing)
- `Diagnostic`: Structured error/warning messages with source location
- `Body`: Container for attributes and blocks
- `File`: Represents a parsed HCL file

**Use Case:** Core data structures and interfaces used by all other HCL packages.

### 2. `github.com/hashicorp/hcl/v2/hclparse`
Main API entry point for parsing both HCL native syntax and HCL JSON.

**Key Functions:**
- `ParseHCL()`: Parse HCL from a buffer
- `ParseHCLFile()`: Parse HCL from a file
- `ParseJSON()`: Parse JSON-formatted HCL
- `ParseJSONFile()`: Parse JSON HCL from a file
- `Files()`: Get map of parsed files for diagnostics

**Example:**
```go
parser := hclparse.NewParser()
file, diags := parser.ParseHCLFile("config.hcl")
if diags.HasErrors() {
    log.Fatal(diags.Error())
}
```

### 3. `github.com/hashicorp/hcl/v2/hclwrite`
Package for generating HCL and making surgical changes to existing configurations while preserving formatting, comments, and structure.

**Key Types:**
- `File`: Represents an HCL file for writing
- `Body`: Container for attributes and blocks
- `Block`: Represents an HCL block
- `Expression`: Represents HCL expressions

**Key Functions:**
- `NewEmptyFile()`: Create an empty HCL file
- `ParseConfig()`: Parse existing HCL for modification
- `AppendBlock()`: Add existing blocks
- `AppendNewBlock()`: Create and add new blocks
- `SetAttributeValue()`: Set attribute values with cty.Value
- `SetAttributeTraversal()`: Set attribute to reference another value
- `SetAttributeRaw()`: Set attribute with raw tokens

**Example:**
```go
f := hclwrite.NewEmptyFile()
rootBody := f.Body()
rootBody.SetAttributeValue("string_attr", cty.StringVal("value"))
block := rootBody.AppendNewBlock("resource", []string{"aws_instance", "example"})
blockBody := block.Body()
blockBody.SetAttributeValue("ami", cty.StringVal("ami-12345"))
```

### 4. `github.com/hashicorp/hcl/v2/hclsyntax`
Native HCL language parser and Abstract Syntax Tree (AST). Lower-level than hclparse.

**Use Case:** Direct access to HCL syntax parsing when you need fine-grained control over the parsing process.

### 5. `github.com/hashicorp/hcl/v2/hclsimple`
High-level entry point for loading configuration files directly into Go structs.

**Example:**
```go
type Config struct {
    IOMode  string        `hcl:"io_mode"`
    Service ServiceConfig `hcl:"service,block"`
}

var config Config
err := hclsimple.DecodeFile("config.hcl", nil, &config)
```

### 6. `github.com/hashicorp/hcl/v2/gohcl`
Decoding HCL bodies into Go structs with struct tags.

**Use Case:** When you want to decode HCL configurations into strongly-typed Go structures.

### 7. `github.com/hashicorp/hcl/v2/hcldec`
Schema-based decoding of HCL bodies.

**Use Case:** When you need to decode HCL based on a dynamic schema rather than static struct tags.

## Working with hclwrite

### Creating New Files
```go
import (
    "github.com/hashicorp/hcl/v2/hclwrite"
    "github.com/zclconf/go-cty/cty"
)

// Create empty file
f := hclwrite.NewEmptyFile()
rootBody := f.Body()

// Add attributes
rootBody.SetAttributeValue("name", cty.StringVal("example"))
rootBody.SetAttributeValue("count", cty.NumberIntVal(3))
rootBody.SetAttributeValue("enabled", cty.BoolVal(true))

// Add a block
block := rootBody.AppendNewBlock("resource", []string{"aws_instance", "web"})
blockBody := block.Body()
blockBody.SetAttributeValue("instance_type", cty.StringVal("t2.micro"))

// Write to file
bytes := f.Bytes()
```

### Modifying Existing Files
```go
// Parse existing file
file, diags := hclwrite.ParseConfig([]byte(content), "config.hcl", hcl.InitialPos)
if diags.HasErrors() {
    return diags
}

// Find and modify blocks
for _, block := range file.Body().Blocks() {
    if block.Type() == "resource" && len(block.Labels()) >= 2 {
        if block.Labels()[0] == "aws_instance" {
            // Modify attributes
            block.Body().SetAttributeValue("instance_type", cty.StringVal("t3.micro"))
        }
    }
}

// Get modified content
modifiedBytes := file.Bytes()
```

### Working with References and Traversals
```go
import "github.com/hashicorp/hcl/v2"

// Create a reference to another resource
traversal := hcl.Traversal{
    hcl.TraverseRoot{Name: "aws_vpc"},
    hcl.TraverseAttr{Name: "main"},
    hcl.TraverseAttr{Name: "id"},
}
body.SetAttributeTraversal("vpc_id", traversal)
// Results in: vpc_id = aws_vpc.main.id
```

### Working with Complex Types
```go
// Lists
listVal := cty.ListVal([]cty.Value{
    cty.StringVal("item1"),
    cty.StringVal("item2"),
})
body.SetAttributeValue("items", listVal)

// Maps/Objects
mapVal := cty.ObjectVal(map[string]cty.Value{
    "key1": cty.StringVal("value1"),
    "key2": cty.NumberIntVal(42),
})
body.SetAttributeValue("config", mapVal)

// Tuples (for mixed-type lists)
tupleVal := cty.TupleVal([]cty.Value{
    cty.StringVal("string"),
    cty.NumberIntVal(123),
    cty.BoolVal(true),
})
body.SetAttributeValue("mixed", tupleVal)
```

## Best Practices

### 1. Use the Right Package for the Task
- **Parsing**: Use `hclparse` for general parsing
- **Writing/Modifying**: Use `hclwrite` to preserve formatting
- **Decoding to structs**: Use `hclsimple` or `gohcl`
- **Dynamic schemas**: Use `hcldec`

### 2. Handle Diagnostics Properly
```go
file, diags := parser.ParseHCLFile("config.hcl")
if diags.HasErrors() {
    // Print detailed error information
    for _, diag := range diags {
        fmt.Printf("Error: %s\n", diag.Summary)
        if diag.Detail != "" {
            fmt.Printf("  Detail: %s\n", diag.Detail)
        }
    }
    return diags
}
```

### 3. Preserve Formatting with hclwrite
When modifying existing files, `hclwrite` preserves:
- Comments
- Whitespace and indentation
- Original formatting style

### 4. Avoid Regex for HCL Manipulation
Using regular expressions to modify HCL is brittle and error-prone because:
- HCL syntax is context-sensitive
- Comments and formatting can break regex patterns
- Expressions and traversals have complex syntax
- Multi-line values are difficult to handle

Instead, use `hclwrite` for reliable modifications.

## Limitations and Workarounds

### 1. Nested Attribute Modification
Currently, hclwrite doesn't support surgical editing of nested attributes within complex objects. You must replace the entire attribute value.

**Workaround:**
```go
// Read the current value (if possible)
attr := block.Body().GetAttribute("complex_attr")
// Reconstruct with modifications
newValue := reconstructWithChanges(attr)
// Replace entire value
block.Body().SetAttributeValue("complex_attr", newValue)
```

### 2. Preserving Complex Expressions
When modifying attributes that contain complex expressions (functions, conditionals), you may need to work with raw tokens:

```go
// Get existing attribute tokens
attr := body.GetAttribute("expression_attr")
tokens := attr.Expr().BuildTokens(nil)
// Modify tokens as needed
// Set back using SetAttributeRaw
body.SetAttributeRaw("expression_attr", tokens)
```

### 3. Working with Dynamic Blocks
Dynamic blocks require special handling as they're not regular blocks:

```go
// Dynamic blocks are treated as attributes with complex expressions
// You typically need to reconstruct the entire dynamic block
```

## Example: Proper HCL Transformation

Instead of using regex (brittle approach):
```go
// DON'T DO THIS
result := regexp.MustCompile(`header\s*{\s*header\s*=\s*"Host"\s*values\s*=\s*(\[[^\]]+\])\s*}`).
    ReplaceAllString(content, `header = { host = $1 }`)
```

Use hclwrite (robust approach):
```go
// DO THIS
file, diags := hclwrite.ParseConfig([]byte(content), "config.tf", hcl.InitialPos)
if diags.HasErrors() {
    return diags
}

for _, block := range file.Body().Blocks() {
    if block.Type() == "resource" && block.Labels()[0] == "cloudflare_load_balancer_pool" {
        // Get origins attribute
        originsAttr := block.Body().GetAttribute("origins")
        if originsAttr != nil {
            // Parse and reconstruct origins with transformed headers
            transformedOrigins := transformOrigins(originsAttr)
            block.Body().SetAttributeValue("origins", transformedOrigins)
        }
    }
}

result := file.Bytes()
```

## Resources

- [Official HCL GitHub Repository](https://github.com/hashicorp/hcl)
- [HCL v2 Go Package Documentation](https://pkg.go.dev/github.com/hashicorp/hcl/v2)
- [hclwrite Package Documentation](https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclwrite)
- [hclparse Package Documentation](https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclparse)
- [go-cty Package (for values)](https://pkg.go.dev/github.com/zclconf/go-cty/cty)
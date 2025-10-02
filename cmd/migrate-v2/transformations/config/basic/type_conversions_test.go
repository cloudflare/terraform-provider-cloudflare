package basic

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetToListConverter_ComplexCases(t *testing.T) {
	tests := []struct {
		name       string
		hcl        string
		attributes []string
		validate   func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "convert multiple toset wrappers",
			hcl: `resource "test" "example" {
  set1 = toset(["a", "b", "c"])
  set2 = toset([1, 2, 3])
  set3 = toset(var.my_list)
  normal_list = ["x", "y", "z"]
  not_a_set = "string_value"
}`,
			attributes: []string{"set1", "set2", "set3", "normal_list", "not_a_set"},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Check toset wrappers are removed
				set1 := string(body.GetAttribute("set1").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, set1, "toset")
				assert.Contains(t, set1, `["a", "b", "c"]`)
				
				set2 := string(body.GetAttribute("set2").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, set2, "toset")
				assert.Contains(t, set2, "[1, 2, 3]")
				
				set3 := string(body.GetAttribute("set3").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, set3, "toset")
				assert.Contains(t, set3, "var.my_list")
				
				// Normal list should remain unchanged
				normalList := string(body.GetAttribute("normal_list").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, normalList, `["x", "y", "z"]`)
				
				// Non-list should remain unchanged
				notASet := string(body.GetAttribute("not_a_set").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, notASet, "string_value")
			},
		},
		{
			name: "nested toset calls",
			hcl: `resource "test" "example" {
  nested = toset(toset(["a", "b"]))
  double_wrapped = toset(concat(["a"], toset(["b", "c"])))
}`,
			attributes: []string{"nested", "double_wrapped"},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Only removes the outermost toset
				nested := string(body.GetAttribute("nested").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, nested, "toset(toset")
				assert.Contains(t, nested, "toset(") // Inner one remains
				
				doubleWrapped := string(body.GetAttribute("double_wrapped").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, doubleWrapped, "toset(concat")
				assert.Contains(t, doubleWrapped, "concat")
			},
		},
		{
			name: "toset with complex expressions",
			hcl: `resource "test" "example" {
  complex1 = toset(split(",", var.csv_string))
  complex2 = toset(concat(var.list1, var.list2))
  complex3 = toset([for s in var.list : upper(s)])
  complex4 = toset(distinct(var.duplicates))
}`,
			attributes: []string{"complex1", "complex2", "complex3", "complex4"},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Check that toset is removed but expressions remain
				complex1 := string(body.GetAttribute("complex1").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, complex1, "toset(")
				assert.Contains(t, complex1, "split")
				
				complex2 := string(body.GetAttribute("complex2").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, complex2, "toset(")
				assert.Contains(t, complex2, "concat")
				
				complex3 := string(body.GetAttribute("complex3").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, complex3, "toset(")
				assert.Contains(t, complex3, "for s in var.list")
				
				complex4 := string(body.GetAttribute("complex4").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, complex4, "toset(")
				assert.Contains(t, complex4, "distinct")
			},
		},
		{
			name: "toset with whitespace variations",
			hcl: `resource "test" "example" {
  no_space = toset(["a"])
  extra_spaces = toset(  ["b"]  )
  newlines = toset(
    ["c"]
  )
  tabs = toset(	["d"]	)
}`,
			attributes: []string{"no_space", "extra_spaces", "newlines", "tabs"},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// All should have toset removed
				for _, attr := range []string{"no_space", "extra_spaces", "newlines", "tabs"} {
					value := string(body.GetAttribute(attr).Expr().BuildTokens(nil).Bytes())
					assert.NotContains(t, value, "toset", "Attribute %s should not contain toset", attr)
				}
			},
		},
		{
			name: "empty attributes list",
			hcl: `resource "test" "example" {
  set1 = toset(["a"])
  set2 = toset(["b"])
}`,
			attributes: []string{},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Nothing should change
				set1 := string(body.GetAttribute("set1").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, set1, "toset")
				
				set2 := string(body.GetAttribute("set2").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, set2, "toset")
			},
		},
		{
			name: "toset with special characters",
			hcl: `resource "test" "example" {
  special = toset(["a-b", "c_d", "e.f", "g/h"])
  unicode = toset(["hello", "世界", "🚀"])
}`,
			attributes: []string{"special", "unicode"},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				special := string(body.GetAttribute("special").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, special, "toset")
				assert.Contains(t, special, "a-b")
				assert.Contains(t, special, "c_d")
				
				unicode := string(body.GetAttribute("unicode").Expr().BuildTokens(nil).Bytes())
				assert.NotContains(t, unicode, "toset")
				assert.Contains(t, unicode, "世界")
				assert.Contains(t, unicode, "🚀")
			},
		},
		{
			name: "not toset patterns",
			hcl: `resource "test" "example" {
  not_toset1 = "toset([\"a\"])"
  not_toset2 = tolist(["b"])
  not_toset3 = tosetting(["c"])
  not_toset4 = toset_custom(["d"])
  comment = "# toset([\"e\"])"
}`,
			attributes: []string{"not_toset1", "not_toset2", "not_toset3", "not_toset4", "comment"},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// These should remain unchanged
				assert.Contains(t, string(body.GetAttribute("not_toset1").Expr().BuildTokens(nil).Bytes()), "toset")
				assert.Contains(t, string(body.GetAttribute("not_toset2").Expr().BuildTokens(nil).Bytes()), "tolist")
				assert.Contains(t, string(body.GetAttribute("not_toset3").Expr().BuildTokens(nil).Bytes()), "tosetting")
				assert.Contains(t, string(body.GetAttribute("not_toset4").Expr().BuildTokens(nil).Bytes()), "toset_custom")
				assert.Contains(t, string(body.GetAttribute("comment").Expr().BuildTokens(nil).Bytes()), "toset")
			},
		},
		{
			name: "malformed toset",
			hcl: `resource "test" "example" {
  incomplete1 = "toset("
  incomplete2 = toset()
  incomplete3 = "toset([\"a\""
  extra_paren = toset(["a"])
}`,
			attributes: []string{"incomplete1", "incomplete2", "incomplete3", "extra_paren"},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// These are string literals, not actual toset calls
				// So they should remain unchanged
				if attr := body.GetAttribute("incomplete1"); attr != nil {
					// String literal, not a toset call
					assert.Contains(t, string(attr.Expr().BuildTokens(nil).Bytes()), "toset(")
				}
				// incomplete2 has toset() which might result in empty value after processing
				attr := body.GetAttribute("incomplete2")
				assert.NotNil(t, attr, "incomplete2 attribute should exist")
				if attr := body.GetAttribute("incomplete3"); attr != nil {
					// String literal
					assert.Contains(t, string(attr.Expr().BuildTokens(nil).Bytes()), "toset")
				}
				if attr := body.GetAttribute("extra_paren"); attr != nil {
					// This one is well-formed toset, should have it removed
					value := string(attr.Expr().BuildTokens(nil).Bytes())
					assert.NotContains(t, value, "toset(")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &TransformContext{}

			transformer := SetToListConverter(tt.attributes...)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestSetToListConverter_EdgeCases(t *testing.T) {
	t.Run("nil attributes", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  attr = toset(["a"])
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		// Call with nil (through varargs)
		transformer := SetToListConverter()
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		// Attribute should remain unchanged
		value := string(block.Body().GetAttribute("attr").Expr().BuildTokens(nil).Bytes())
		assert.Contains(t, value, "toset")
	})
	
	t.Run("non-existent attributes", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  real_attr = toset(["a"])
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		// Try to convert non-existent attributes
		transformer := SetToListConverter("non_existent1", "non_existent2", "real_attr")
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		// Only real_attr should be processed
		value := string(block.Body().GetAttribute("real_attr").Expr().BuildTokens(nil).Bytes())
		assert.NotContains(t, value, "toset")
	})
	
	t.Run("very long attribute value", func(t *testing.T) {
		// Create a very long list
		var items []string
		for i := 0; i < 100; i++ {
			items = append(items, fmt.Sprintf(`"item%d"`, i))
		}
		longList := strings.Join(items, ", ")
		
		hclContent := fmt.Sprintf(`resource "test" "example" {
  long_set = toset([%s])
}`, longList)
		
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())
		
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := SetToListConverter("long_set")
		err := transformer(block, ctx)
		require.NoError(t, err)
		
		// Should process correctly
		value := string(block.Body().GetAttribute("long_set").Expr().BuildTokens(nil).Bytes())
		assert.NotContains(t, value, "toset")
		assert.Contains(t, value, "item0")
		assert.Contains(t, value, "item99")
	})
}

func TestSetToListConverter_RealWorldScenarios(t *testing.T) {
	tests := []struct {
		name string
		hcl  string
	}{
		{
			name: "AWS security group rules",
			hcl: `resource "test" "example" {
  ingress_rules = toset([
    {
      from_port   = 80
      to_port     = 80
      protocol    = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    },
    {
      from_port   = 443
      to_port     = 443
      protocol    = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  ])
}`,
		},
		{
			name: "Kubernetes labels",
			hcl: `resource "test" "example" {
  labels = toset([
    "app=frontend",
    "env=production",
    "version=1.0.0"
  ])
}`,
		},
		{
			name: "Terraform for_each conversion",
			hcl: `resource "test" "example" {
  instance_ids = toset(aws_instance.example[*].id)
  subnet_ids   = toset(data.aws_subnets.example.ids)
  zone_names   = toset(keys(var.availability_zones))
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &TransformContext{}

			// Get all attribute names
			var attrs []string
			// In real code, we'd iterate through attributes
			// For testing, we'll manually specify them
			switch tt.name {
			case "AWS security group rules":
				attrs = []string{"ingress_rules"}
			case "Kubernetes labels":
				attrs = []string{"labels"}
			case "Terraform for_each conversion":
				attrs = []string{"instance_ids", "subnet_ids", "zone_names"}
			}

			transformer := SetToListConverter(attrs...)
			err := transformer(block, ctx)
			require.NoError(t, err)

			// All toset wrappers should be removed
			body := block.Body()
			for _, attr := range attrs {
				if a := body.GetAttribute(attr); a != nil {
					value := string(a.Expr().BuildTokens(nil).Bytes())
					assert.NotContains(t, value, "toset(", "Attribute %s should not contain toset", attr)
				}
			}
		})
	}
}

func BenchmarkSetToListConverter(b *testing.B) {
	hclContent := `resource "test" "example" {
  set1 = toset(["a", "b", "c"])
  set2 = toset([1, 2, 3])
  set3 = toset(var.my_list)
  normal = ["x", "y", "z"]
}`
	
	attrs := []string{"set1", "set2", "set3", "normal"}
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := SetToListConverter(attrs...)
		_ = transformer(block, ctx)
	}
}

func BenchmarkSetToListConverter_LongExpression(b *testing.B) {
	// Create a long expression
	var items []string
	for i := 0; i < 50; i++ {
		items = append(items, fmt.Sprintf(`"item%d"`, i))
	}
	longList := strings.Join(items, ", ")
	
	hclContent := fmt.Sprintf(`resource "test" "example" {
  long_set = toset([%s])
}`, longList)
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := SetToListConverter("long_set")
		_ = transformer(block, ctx)
	}
}

func BenchmarkSetToListConverter_ManyAttributes(b *testing.B) {
	// Create HCL with many toset attributes
	var hclBuilder strings.Builder
	hclBuilder.WriteString(`resource "test" "example" {`)
	
	var attrs []string
	for i := 0; i < 20; i++ {
		hclBuilder.WriteString(fmt.Sprintf("\n  attr%d = toset([\"value%d\"])", i, i))
		attrs = append(attrs, fmt.Sprintf("attr%d", i))
	}
	hclBuilder.WriteString("\n}")
	
	hclContent := hclBuilder.String()
	
	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &TransformContext{}
		
		transformer := SetToListConverter(attrs...)
		_ = transformer(block, ctx)
	}
}
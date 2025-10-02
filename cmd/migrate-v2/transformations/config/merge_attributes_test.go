package config

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/common"
)

func TestMergeAttributesTransformer(t *testing.T) {
	tests := []struct {
		name        string
		hcl         string
		target      string
		sourceAttrs []string
		format      string
		validate    func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "merge into object format",
			hcl: `resource "test" "example" {
  host = "localhost"
  port = 5432
  database = "mydb"
}`,
			target:      "connection",
			sourceAttrs: []string{"host", "port", "database"},
			format:      "object",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Source attributes should be removed
				assert.Nil(t, body.GetAttribute("host"))
				assert.Nil(t, body.GetAttribute("port"))
				assert.Nil(t, body.GetAttribute("database"))
				
				// Target object should exist
				connAttr := body.GetAttribute("connection")
				assert.NotNil(t, connAttr)
				
				// Check that the object contains the merged fields
				value := string(connAttr.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "host")
				assert.Contains(t, value, "port")
				assert.Contains(t, value, "database")
				assert.Contains(t, value, "localhost")
				assert.Contains(t, value, "5432")
				assert.Contains(t, value, "mydb")
			},
		},
		{
			name: "merge into list format",
			hcl: `resource "test" "example" {
  tag1 = "production"
  tag2 = "web"
  tag3 = "public"
}`,
			target:      "tags",
			sourceAttrs: []string{"tag1", "tag2", "tag3"},
			format:      "list",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Source attributes should be removed
				assert.Nil(t, body.GetAttribute("tag1"))
				assert.Nil(t, body.GetAttribute("tag2"))
				assert.Nil(t, body.GetAttribute("tag3"))
				
				// Target list should exist
				tagsAttr := body.GetAttribute("tags")
				assert.NotNil(t, tagsAttr)
				
				// Check that the list contains the values
				value := string(tagsAttr.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "[")
				assert.Contains(t, value, "]")
				assert.Contains(t, value, "production")
				assert.Contains(t, value, "web")
				assert.Contains(t, value, "public")
			},
		},
		{
			name: "merge with missing source attributes",
			hcl: `resource "test" "example" {
  field1 = "value1"
  other = "keep"
}`,
			target:      "merged",
			sourceAttrs: []string{"field1", "field2", "field3"},
			format:      "object",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Existing source should be removed
				assert.Nil(t, body.GetAttribute("field1"))
				
				// Non-existent sources don't cause error
				assert.Nil(t, body.GetAttribute("field2"))
				assert.Nil(t, body.GetAttribute("field3"))
				
				// Other attributes preserved
				assert.NotNil(t, body.GetAttribute("other"))
				
				// Target created with available fields
				assert.NotNil(t, body.GetAttribute("merged"))
			},
		},
		{
			name: "empty source attributes list",
			hcl: `resource "test" "example" {
  keep = "value"
}`,
			target:      "empty_merge",
			sourceAttrs: []string{},
			format:      "object",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should create empty target
				assert.NotNil(t, body.GetAttribute("empty_merge"))
				
				// Other attributes unchanged
				assert.NotNil(t, body.GetAttribute("keep"))
			},
		},
		{
			name: "merge with references",
			hcl: `resource "test" "example" {
  vpc_id = aws_vpc.main.id
  subnet_id = aws_subnet.private.id
  security_group = aws_security_group.web.id
}`,
			target:      "network_config",
			sourceAttrs: []string{"vpc_id", "subnet_id", "security_group"},
			format:      "object",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Sources removed
				assert.Nil(t, body.GetAttribute("vpc_id"))
				assert.Nil(t, body.GetAttribute("subnet_id"))
				assert.Nil(t, body.GetAttribute("security_group"))
				
				// Target created with references preserved
				netConfig := body.GetAttribute("network_config")
				assert.NotNil(t, netConfig)
				value := string(netConfig.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "aws_vpc.main.id")
				assert.Contains(t, value, "aws_subnet.private.id")
				assert.Contains(t, value, "aws_security_group.web.id")
			},
		},
		{
			name: "merge different data types",
			hcl: `resource "test" "example" {
  enabled = true
  count = 42
  name = "test"
  ratio = 0.75
}`,
			target:      "settings",
			sourceAttrs: []string{"enabled", "count", "name", "ratio"},
			format:      "object",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// All sources removed
				assert.Nil(t, body.GetAttribute("enabled"))
				assert.Nil(t, body.GetAttribute("count"))
				assert.Nil(t, body.GetAttribute("name"))
				assert.Nil(t, body.GetAttribute("ratio"))
				
				// Target contains all types
				settings := body.GetAttribute("settings")
				assert.NotNil(t, settings)
				value := string(settings.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "true")
				assert.Contains(t, value, "42")
				assert.Contains(t, value, "test")
				assert.Contains(t, value, "0.75")
			},
		},
		{
			name: "invalid format defaults to object",
			hcl: `resource "test" "example" {
  a = "1"
  b = "2"
}`,
			target:      "result",
			sourceAttrs: []string{"a", "b"},
			format:      "invalid",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should default to object format
				assert.Nil(t, body.GetAttribute("a"))
				assert.Nil(t, body.GetAttribute("b"))
				assert.NotNil(t, body.GetAttribute("result"))
			},
		},
		{
			name: "target already exists",
			hcl: `resource "test" "example" {
  existing_target = "old_value"
  field1 = "new1"
  field2 = "new2"
}`,
			target:      "existing_target",
			sourceAttrs: []string{"field1", "field2"},
			format:      "object",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Sources removed
				assert.Nil(t, body.GetAttribute("field1"))
				assert.Nil(t, body.GetAttribute("field2"))
				
				// Target should be overwritten
				target := body.GetAttribute("existing_target")
				assert.NotNil(t, target)
				value := string(target.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "field1")
				assert.Contains(t, value, "field2")
				assert.NotContains(t, value, "old_value")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &common.TransformContext{}

			transformer := MergeAttributesTransformer(tt.target, tt.sourceAttrs, tt.format)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestMergeAttributesWithMapping(t *testing.T) {
	tests := []struct {
		name         string
		hcl          string
		target       string
		attributeMap map[string]string
		validate     func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "basic attribute mapping",
			hcl: `resource "test" "example" {
  db_host = "localhost"
  db_port = 5432
  db_name = "mydb"
}`,
			target: "database",
			attributeMap: map[string]string{
				"db_host": "host",
				"db_port": "port",
				"db_name": "name",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original attributes removed
				assert.Nil(t, body.GetAttribute("db_host"))
				assert.Nil(t, body.GetAttribute("db_port"))
				assert.Nil(t, body.GetAttribute("db_name"))
				
				// Target with mapped names
				db := body.GetAttribute("database")
				assert.NotNil(t, db)
				value := string(db.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "host")
				assert.Contains(t, value, "port")
				assert.Contains(t, value, "name")
				assert.NotContains(t, value, "db_host")
			},
		},
		{
			name: "partial mapping",
			hcl: `resource "test" "example" {
  old_field1 = "value1"
  old_field2 = "value2"
  other = "keep"
}`,
			target: "config",
			attributeMap: map[string]string{
				"old_field1": "new_field1",
				"old_field2": "new_field2",
				"missing":    "ignored",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Mapped sources removed
				assert.Nil(t, body.GetAttribute("old_field1"))
				assert.Nil(t, body.GetAttribute("old_field2"))
				
				// Other preserved
				assert.NotNil(t, body.GetAttribute("other"))
				
				// Target created with mapped names
				config := body.GetAttribute("config")
				assert.NotNil(t, config)
				value := string(config.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "new_field1")
				assert.Contains(t, value, "new_field2")
				assert.NotContains(t, value, "missing")
			},
		},
		{
			name: "empty mapping",
			hcl: `resource "test" "example" {
  field = "value"
}`,
			target:       "result",
			attributeMap: map[string]string{},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original unchanged
				assert.NotNil(t, body.GetAttribute("field"))
				
				// Empty object created
				result := body.GetAttribute("result")
				assert.NotNil(t, result)
			},
		},
		{
			name: "nil mapping",
			hcl: `resource "test" "example" {
  attr = "val"
}`,
			target:       "output",
			attributeMap: nil,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original unchanged
				assert.NotNil(t, body.GetAttribute("attr"))
				
				// Empty object created
				output := body.GetAttribute("output")
				assert.NotNil(t, output)
			},
		},
		{
			name: "rename to same name",
			hcl: `resource "test" "example" {
  field1 = "value1"
  field2 = "value2"
}`,
			target: "merged",
			attributeMap: map[string]string{
				"field1": "field1", // Same name
				"field2": "renamed",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Sources removed
				assert.Nil(t, body.GetAttribute("field1"))
				assert.Nil(t, body.GetAttribute("field2"))
				
				// Target with both names
				merged := body.GetAttribute("merged")
				assert.NotNil(t, merged)
				value := string(merged.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "field1")
				assert.Contains(t, value, "renamed")
			},
		},
		{
			name: "complex value mapping",
			hcl: `resource "test" "example" {
  list_attr = ["a", "b", "c"]
  obj_attr = {
    nested = "value"
  }
  ref_attr = var.some_var
}`,
			target: "complex",
			attributeMap: map[string]string{
				"list_attr": "items",
				"obj_attr":  "config",
				"ref_attr":  "reference",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// All sources removed
				assert.Nil(t, body.GetAttribute("list_attr"))
				assert.Nil(t, body.GetAttribute("obj_attr"))
				assert.Nil(t, body.GetAttribute("ref_attr"))
				
				// Target with mapped complex values
				complex := body.GetAttribute("complex")
				assert.NotNil(t, complex)
				value := string(complex.Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, value, "items")
				assert.Contains(t, value, "config")
				assert.Contains(t, value, "reference")
				assert.Contains(t, value, "var.some_var")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &common.TransformContext{}

			transformer := MergeAttributesWithMapping(tt.target, tt.attributeMap)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestMergeAttributesTransformer_EdgeCases(t *testing.T) {
	t.Run("empty target name", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  field = "value"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())

		block := file.Body().Blocks()[0]
		ctx := &common.TransformContext{}

		transformer := MergeAttributesTransformer("", []string{"field"}, "object")
		err := transformer(block, ctx)
		require.NoError(t, err)

		// Should not crash, but field should remain
		assert.NotNil(t, block.Body().GetAttribute("field"))
	})

	t.Run("duplicate source attributes", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  field1 = "value1"
  field2 = "value2"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())

		block := file.Body().Blocks()[0]
		ctx := &common.TransformContext{}

		// Duplicate field1 in source list
		transformer := MergeAttributesTransformer("merged", []string{"field1", "field1", "field2"}, "object")
		err := transformer(block, ctx)
		require.NoError(t, err)

		// Should handle gracefully
		assert.Nil(t, block.Body().GetAttribute("field1"))
		assert.Nil(t, block.Body().GetAttribute("field2"))
		assert.NotNil(t, block.Body().GetAttribute("merged"))
	})
}

func BenchmarkMergeAttributesTransformer(b *testing.B) {
	hclContent := `resource "test" "example" {
  host = "localhost"
  port = 5432
  database = "mydb"
  user = "admin"
  password = "secret"
}`

	sourceAttrs := []string{"host", "port", "database", "user", "password"}

	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &common.TransformContext{}
		
		transformer := MergeAttributesTransformer("connection", sourceAttrs, "object")
		_ = transformer(block, ctx)
	}
}

func BenchmarkMergeAttributesWithMapping(b *testing.B) {
	hclContent := `resource "test" "example" {
  field1 = "value1"
  field2 = "value2"
  field3 = "value3"
  field4 = "value4"
  field5 = "value5"
}`

	attributeMap := map[string]string{
		"field1": "renamed1",
		"field2": "renamed2",
		"field3": "renamed3",
		"field4": "renamed4",
		"field5": "renamed5",
	}

	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &common.TransformContext{}
		
		transformer := MergeAttributesWithMapping("merged", attributeMap)
		_ = transformer(block, ctx)
	}
}
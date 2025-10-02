package config

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/common"
)

func TestSplitObjectTransformer(t *testing.T) {
	tests := []struct {
		name       string
		hcl        string
		source     string
		attributes []string
		prefix     string
		validate   func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "basic object splitting",
			hcl: `resource "test" "example" {
  connection = {
    host = "localhost"
    port = 5432
    database = "mydb"
  }
}`,
			source:     "connection",
			attributes: []string{"host", "port", "database"},
			prefix:     "db_",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original object should be removed
				assert.Nil(t, body.GetAttribute("connection"))
				
				// Split attributes should exist with prefix
				assert.NotNil(t, body.GetAttribute("db_host"))
				assert.NotNil(t, body.GetAttribute("db_port"))
				assert.NotNil(t, body.GetAttribute("db_database"))
				
				// Check values
				host := string(body.GetAttribute("db_host").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, host, "localhost")
				port := string(body.GetAttribute("db_port").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, port, "5432")
				db := string(body.GetAttribute("db_database").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, db, "mydb")
			},
		},
		{
			name: "splitting without prefix",
			hcl: `resource "test" "example" {
  settings = {
    enabled = true
    timeout = 30
    retries = 3
  }
}`,
			source:     "settings",
			attributes: []string{"enabled", "timeout", "retries"},
			prefix:     "",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("settings"))
				
				// Attributes without prefix
				assert.NotNil(t, body.GetAttribute("enabled"))
				assert.NotNil(t, body.GetAttribute("timeout"))
				assert.NotNil(t, body.GetAttribute("retries"))
			},
		},
		{
			name: "partial extraction",
			hcl: `resource "test" "example" {
  config = {
    public_field1 = "value1"
    public_field2 = "value2"
    private_field = "secret"
    internal_data = "hidden"
  }
}`,
			source:     "config",
			attributes: []string{"public_field1", "public_field2"},
			prefix:     "",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original still exists (has remaining fields)
				assert.NotNil(t, body.GetAttribute("config"))
				
				// Extracted fields
				assert.NotNil(t, body.GetAttribute("public_field1"))
				assert.NotNil(t, body.GetAttribute("public_field2"))
				
				// Check original still has other fields
				configValue := string(body.GetAttribute("config").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, configValue, "private_field")
				assert.Contains(t, configValue, "internal_data")
				assert.NotContains(t, configValue, "public_field1")
				assert.NotContains(t, configValue, "public_field2")
			},
		},
		{
			name: "empty attributes list extracts all",
			hcl: `resource "test" "example" {
  data = {
    field1 = "value1"
    field2 = "value2"
    field3 = "value3"
  }
}`,
			source:     "data",
			attributes: []string{},
			prefix:     "extracted_",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original removed
				assert.Nil(t, body.GetAttribute("data"))
				
				// All fields extracted
				assert.NotNil(t, body.GetAttribute("extracted_field1"))
				assert.NotNil(t, body.GetAttribute("extracted_field2"))
				assert.NotNil(t, body.GetAttribute("extracted_field3"))
			},
		},
		{
			name: "non-existent source",
			hcl: `resource "test" "example" {
  other = "value"
}`,
			source:     "missing",
			attributes: []string{"field"},
			prefix:     "pre_",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should not affect existing attributes
				assert.NotNil(t, body.GetAttribute("other"))
				
				// No new attributes created
				assert.Nil(t, body.GetAttribute("pre_field"))
			},
		},
		{
			name: "nested objects in source",
			hcl: `resource "test" "example" {
  complex = {
    simple = "value"
    nested = {
      inner = "deep"
    }
    list = ["a", "b"]
  }
}`,
			source:     "complex",
			attributes: []string{"simple", "nested", "list"},
			prefix:     "flat_",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("complex"))
				
				// All types extracted
				assert.NotNil(t, body.GetAttribute("flat_simple"))
				assert.NotNil(t, body.GetAttribute("flat_nested"))
				assert.NotNil(t, body.GetAttribute("flat_list"))
				
				// Nested structure preserved
				nestedValue := string(body.GetAttribute("flat_nested").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, nestedValue, "inner")
			},
		},
		{
			name: "source with references",
			hcl: `resource "test" "example" {
  network = {
    vpc_id = aws_vpc.main.id
    subnet_id = aws_subnet.private.id
    security_group = aws_security_group.web.id
  }
}`,
			source:     "network",
			attributes: []string{"vpc_id", "subnet_id", "security_group"},
			prefix:     "net_",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("network"))
				
				// References preserved
				vpcValue := string(body.GetAttribute("net_vpc_id").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, vpcValue, "aws_vpc.main.id")
				
				subnetValue := string(body.GetAttribute("net_subnet_id").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, subnetValue, "aws_subnet.private.id")
			},
		},
		{
			name: "extract non-existent fields",
			hcl: `resource "test" "example" {
  obj = {
    real_field = "value"
  }
}`,
			source:     "obj",
			attributes: []string{"real_field", "fake_field"},
			prefix:     "",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original removed (all specified fields attempted)
				assert.Nil(t, body.GetAttribute("obj"))
				
				// Only real field extracted
				assert.NotNil(t, body.GetAttribute("real_field"))
				assert.Nil(t, body.GetAttribute("fake_field"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &common.TransformContext{}

			transformer := SplitObjectTransformer(tt.source, tt.attributes, tt.prefix)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestSplitObjectWithMapping(t *testing.T) {
	tests := []struct {
		name         string
		hcl          string
		source       string
		attributeMap map[string]string
		validate     func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "basic mapping",
			hcl: `resource "test" "example" {
  database = {
    host = "localhost"
    port = 5432
    name = "mydb"
  }
}`,
			source: "database",
			attributeMap: map[string]string{
				"host": "db_host",
				"port": "db_port",
				"name": "db_name",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("database"))
				
				// Mapped names
				assert.NotNil(t, body.GetAttribute("db_host"))
				assert.NotNil(t, body.GetAttribute("db_port"))
				assert.NotNil(t, body.GetAttribute("db_name"))
				
				// Values preserved
				host := string(body.GetAttribute("db_host").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, host, "localhost")
			},
		},
		{
			name: "partial mapping",
			hcl: `resource "test" "example" {
  config = {
    field1 = "value1"
    field2 = "value2"
    field3 = "value3"
  }
}`,
			source: "config",
			attributeMap: map[string]string{
				"field1": "renamed1",
				"field3": "renamed3",
				// field2 not mapped - uses original name
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Source removed
				assert.Nil(t, body.GetAttribute("config"))
				
				// Mapped fields
				assert.NotNil(t, body.GetAttribute("renamed1"))
				assert.NotNil(t, body.GetAttribute("renamed3"))
				
				// Unmapped field uses original name
				assert.NotNil(t, body.GetAttribute("field2"))
			},
		},
		{
			name: "empty mapping extracts all with original names",
			hcl: `resource "test" "example" {
  obj = {
    attr1 = "val1"
    attr2 = "val2"
  }
}`,
			source:       "obj",
			attributeMap: map[string]string{},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("obj"))
				assert.NotNil(t, body.GetAttribute("attr1"))
				assert.NotNil(t, body.GetAttribute("attr2"))
			},
		},
		{
			name: "nil mapping extracts all",
			hcl: `resource "test" "example" {
  data = {
    x = "1"
    y = "2"
  }
}`,
			source:       "data",
			attributeMap: nil,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("data"))
				assert.NotNil(t, body.GetAttribute("x"))
				assert.NotNil(t, body.GetAttribute("y"))
			},
		},
		{
			name: "mapping to same name",
			hcl: `resource "test" "example" {
  source = {
    keep_name = "value"
    change_name = "value"
  }
}`,
			source: "source",
			attributeMap: map[string]string{
				"keep_name":   "keep_name", // Same name
				"change_name": "new_name",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("source"))
				assert.NotNil(t, body.GetAttribute("keep_name"))
				assert.NotNil(t, body.GetAttribute("new_name"))
				assert.Nil(t, body.GetAttribute("change_name"))
			},
		},
		{
			name: "mapping non-existent fields",
			hcl: `resource "test" "example" {
  obj = {
    real = "value"
  }
}`,
			source: "obj",
			attributeMap: map[string]string{
				"real": "extracted_real",
				"fake": "extracted_fake", // Doesn't exist
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Only real field extracted
				assert.NotNil(t, body.GetAttribute("extracted_real"))
				assert.Nil(t, body.GetAttribute("extracted_fake"))
			},
		},
		{
			name: "complex value types with mapping",
			hcl: `resource "test" "example" {
  mixed = {
    string_val = "text"
    number_val = 42
    bool_val = true
    list_val = ["a", "b"]
    object_val = {
      nested = "value"
    }
  }
}`,
			source: "mixed",
			attributeMap: map[string]string{
				"string_val": "str",
				"number_val": "num",
				"bool_val":   "bool",
				"list_val":   "list",
				"object_val": "obj",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("mixed"))
				
				// All types extracted with mapped names
				assert.NotNil(t, body.GetAttribute("str"))
				assert.NotNil(t, body.GetAttribute("num"))
				assert.NotNil(t, body.GetAttribute("bool"))
				assert.NotNil(t, body.GetAttribute("list"))
				assert.NotNil(t, body.GetAttribute("obj"))
				
				// Check types preserved
				strVal := string(body.GetAttribute("str").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, strVal, "text")
				
				numVal := string(body.GetAttribute("num").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, numVal, "42")
				
				boolVal := string(body.GetAttribute("bool").Expr().BuildTokens(nil).Bytes())
				assert.Contains(t, boolVal, "true")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &common.TransformContext{}

			transformer := SplitObjectWithMapping(tt.source, tt.attributeMap)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestSplitObjectTransformer_EdgeCases(t *testing.T) {
	t.Run("source is not an object", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  scalar = "not an object"
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())

		block := file.Body().Blocks()[0]
		ctx := &common.TransformContext{}

		transformer := SplitObjectTransformer("scalar", []string{}, "")
		err := transformer(block, ctx)
		require.NoError(t, err)

		// Should not crash, scalar should remain
		assert.NotNil(t, block.Body().GetAttribute("scalar"))
	})

	t.Run("empty source name", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  obj = {
    field = "value"
  }
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())

		block := file.Body().Blocks()[0]
		ctx := &common.TransformContext{}

		transformer := SplitObjectTransformer("", []string{"field"}, "pre_")
		err := transformer(block, ctx)
		require.NoError(t, err)

		// Original unchanged
		assert.NotNil(t, block.Body().GetAttribute("obj"))
		assert.Nil(t, block.Body().GetAttribute("pre_field"))
	})

	t.Run("attribute exists at top level", func(t *testing.T) {
		hclContent := `resource "test" "example" {
  existing_field = "original"
  obj = {
    existing_field = "from_object"
  }
}`
		file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		require.False(t, diags.HasErrors())

		block := file.Body().Blocks()[0]
		ctx := &common.TransformContext{}

		transformer := SplitObjectTransformer("obj", []string{"existing_field"}, "")
		err := transformer(block, ctx)
		require.NoError(t, err)

		// Should overwrite existing
		field := block.Body().GetAttribute("existing_field")
		assert.NotNil(t, field)
		value := string(field.Expr().BuildTokens(nil).Bytes())
		assert.Contains(t, value, "from_object")
		assert.NotContains(t, value, "original")
	})
}

func BenchmarkSplitObjectTransformer(b *testing.B) {
	hclContent := `resource "test" "example" {
  config = {
    host = "localhost"
    port = 5432
    database = "mydb"
    user = "admin"
    password = "secret"
    timeout = 30
    retries = 3
  }
}`

	attributes := []string{"host", "port", "database", "user", "password", "timeout", "retries"}

	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &common.TransformContext{}
		
		transformer := SplitObjectTransformer("config", attributes, "db_")
		_ = transformer(block, ctx)
	}
}

func BenchmarkSplitObjectWithMapping(b *testing.B) {
	hclContent := `resource "test" "example" {
  data = {
    field1 = "value1"
    field2 = "value2"
    field3 = "value3"
    field4 = "value4"
    field5 = "value5"
  }
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
		
		transformer := SplitObjectWithMapping("data", attributeMap)
		_ = transformer(block, ctx)
	}
}
package structural

import (
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

func TestFlattenNestedTransformer(t *testing.T) {
	tests := []struct {
		name      string
		hcl       string
		source    string
		separator string
		maxDepth  int
		validate  func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "basic nested object flattening",
			hcl: `resource "test" "example" {
  nested = {
    level1 = {
      level2 = "value"
      other = "data"
    }
    simple = "direct"
  }
}`,
			source:    "nested",
			separator: "_",
			maxDepth:  2,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original should be removed
				assert.Nil(t, body.GetAttribute("nested"))
				
				// Flattened attributes should exist
				assert.NotNil(t, body.GetAttribute("nested_level1_level2"))
				assert.NotNil(t, body.GetAttribute("nested_level1_other"))
				assert.NotNil(t, body.GetAttribute("nested_simple"))
			},
		},
		{
			name: "single level flattening",
			hcl: `resource "test" "example" {
  config = {
    timeout = 30
    retries = 3
    enabled = true
  }
}`,
			source:    "config",
			separator: ".",
			maxDepth:  1,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("config"))
				assert.NotNil(t, body.GetAttribute("config.timeout"))
				assert.NotNil(t, body.GetAttribute("config.retries"))
				assert.NotNil(t, body.GetAttribute("config.enabled"))
			},
		},
		{
			name: "deeply nested structure",
			hcl: `resource "test" "example" {
  settings = {
    network = {
      firewall = {
        rules = {
          allow = "all"
        }
      }
    }
  }
}`,
			source:    "settings",
			separator: "_",
			maxDepth:  4,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("settings"))
				assert.NotNil(t, body.GetAttribute("settings_network_firewall_rules_allow"))
			},
		},
		{
			name: "max depth limitation",
			hcl: `resource "test" "example" {
  deep = {
    level1 = {
      level2 = {
        level3 = "too deep"
      }
    }
  }
}`,
			source:    "deep",
			separator: "_",
			maxDepth:  2,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should flatten up to maxDepth
				assert.Nil(t, body.GetAttribute("deep"))
				assert.NotNil(t, body.GetAttribute("deep_level1_level2"))
				// Should not go deeper
				assert.Nil(t, body.GetAttribute("deep_level1_level2_level3"))
			},
		},
		{
			name: "non-existent source",
			hcl: `resource "test" "example" {
  other = "value"
}`,
			source:    "missing",
			separator: "_",
			maxDepth:  2,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should not modify anything
				assert.NotNil(t, body.GetAttribute("other"))
				assert.Nil(t, body.GetAttribute("missing"))
			},
		},
		{
			name: "empty separator",
			hcl: `resource "test" "example" {
  obj = {
    field1 = "value1"
    field2 = "value2"
  }
}`,
			source:    "obj",
			separator: "",
			maxDepth:  1,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("obj"))
				assert.NotNil(t, body.GetAttribute("objfield1"))
				assert.NotNil(t, body.GetAttribute("objfield2"))
			},
		},
		{
			name: "zero max depth",
			hcl: `resource "test" "example" {
  data = {
    field = "value"
  }
}`,
			source:    "data",
			separator: "_",
			maxDepth:  0,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should not flatten anything with maxDepth 0
				assert.NotNil(t, body.GetAttribute("data"))
			},
		},
		{
			name: "mixed types in nested structure",
			hcl: `resource "test" "example" {
  mixed = {
    string = "text"
    number = 42
    boolean = true
    list = ["a", "b"]
    nested = {
      inner = "value"
    }
  }
}`,
			source:    "mixed",
			separator: "_",
			maxDepth:  2,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("mixed"))
				assert.NotNil(t, body.GetAttribute("mixed_string"))
				assert.NotNil(t, body.GetAttribute("mixed_number"))
				assert.NotNil(t, body.GetAttribute("mixed_boolean"))
				assert.NotNil(t, body.GetAttribute("mixed_list"))
				assert.NotNil(t, body.GetAttribute("mixed_nested_inner"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &basic.TransformContext{}

			transformer := FlattenNestedTransformer(tt.source, tt.separator, tt.maxDepth)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestFlattenWithPrefix(t *testing.T) {
	tests := []struct {
		name      string
		hcl       string
		source    string
		prefix    string
		separator string
		validate  func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "basic prefix addition",
			hcl: `resource "test" "example" {
  data = {
    field1 = "value1"
    field2 = "value2"
  }
}`,
			source:    "data",
			prefix:    "prefixed",
			separator: "_",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("data"))
				assert.NotNil(t, body.GetAttribute("prefixed_field1"))
				assert.NotNil(t, body.GetAttribute("prefixed_field2"))
			},
		},
		{
			name: "empty prefix",
			hcl: `resource "test" "example" {
  obj = {
    attr = "value"
  }
}`,
			source:    "obj",
			prefix:    "",
			separator: "_",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("obj"))
				assert.NotNil(t, body.GetAttribute("attr"))
			},
		},
		{
			name: "prefix with custom separator",
			hcl: `resource "test" "example" {
  config = {
    timeout = 30
    retries = 3
  }
}`,
			source:    "config",
			prefix:    "app",
			separator: ".",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("config"))
				assert.NotNil(t, body.GetAttribute("app.timeout"))
				assert.NotNil(t, body.GetAttribute("app.retries"))
			},
		},
		{
			name: "nested structure with prefix",
			hcl: `resource "test" "example" {
  settings = {
    advanced = {
      option = "enabled"
    }
  }
}`,
			source:    "settings",
			prefix:    "global",
			separator: "_",
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("settings"))
				assert.NotNil(t, body.GetAttribute("global_advanced_option"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &basic.TransformContext{}

			transformer := FlattenWithPrefix(tt.source, tt.prefix, tt.separator)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func TestNestedRestructureTransformer(t *testing.T) {
	tests := []struct {
		name     string
		hcl      string
		source   string
		paths    map[string]string
		validate func(*testing.T, *hclwrite.Block)
	}{
		{
			name: "basic path mapping",
			hcl: `resource "test" "example" {
  config = {
    database = {
      host = "localhost"
      port = 5432
    }
    cache = {
      ttl = 300
    }
  }
}`,
			source: "config",
			paths: map[string]string{
				"database.host": "db_host",
				"database.port": "db_port",
				"cache.ttl":     "cache_timeout",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("config"))
				assert.NotNil(t, body.GetAttribute("db_host"))
				assert.NotNil(t, body.GetAttribute("db_port"))
				assert.NotNil(t, body.GetAttribute("cache_timeout"))
			},
		},
		{
			name: "partial path mapping",
			hcl: `resource "test" "example" {
  settings = {
    feature1 = {
      enabled = true
      value = "a"
    }
    feature2 = {
      enabled = false
      value = "b"
    }
  }
}`,
			source: "settings",
			paths: map[string]string{
				"feature1.enabled": "feature1_on",
				"feature2.value":   "feature2_data",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Original removed
				assert.Nil(t, body.GetAttribute("settings"))
				
				// Mapped paths exist
				assert.NotNil(t, body.GetAttribute("feature1_on"))
				assert.NotNil(t, body.GetAttribute("feature2_data"))
				
				// Unmapped paths also flattened with default naming
				assert.NotNil(t, body.GetAttribute("settings_feature1_value"))
				assert.NotNil(t, body.GetAttribute("settings_feature2_enabled"))
			},
		},
		{
			name: "deep path mapping",
			hcl: `resource "test" "example" {
  app = {
    server = {
      http = {
        port = 8080
        timeout = 30
      }
    }
  }
}`,
			source: "app",
			paths: map[string]string{
				"server.http.port":    "http_port",
				"server.http.timeout": "http_timeout_seconds",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				assert.Nil(t, body.GetAttribute("app"))
				assert.NotNil(t, body.GetAttribute("http_port"))
				assert.NotNil(t, body.GetAttribute("http_timeout_seconds"))
			},
		},
		{
			name: "empty paths map",
			hcl: `resource "test" "example" {
  data = {
    field = "value"
  }
}`,
			source: "data",
			paths:  map[string]string{},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should flatten with default naming
				assert.Nil(t, body.GetAttribute("data"))
				assert.NotNil(t, body.GetAttribute("data_field"))
			},
		},
		{
			name: "nil paths map",
			hcl: `resource "test" "example" {
  obj = {
    attr = "val"
  }
}`,
			source: "obj",
			paths:  nil,
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Should flatten with default naming
				assert.Nil(t, body.GetAttribute("obj"))
				assert.NotNil(t, body.GetAttribute("obj_attr"))
			},
		},
		{
			name: "non-matching paths",
			hcl: `resource "test" "example" {
  config = {
    real_field = "value"
  }
}`,
			source: "config",
			paths: map[string]string{
				"non_existent.field": "mapped_name",
			},
			validate: func(t *testing.T, block *hclwrite.Block) {
				body := block.Body()
				
				// Non-matching path ignored
				assert.Nil(t, body.GetAttribute("mapped_name"))
				
				// Real field flattened with default
				assert.NotNil(t, body.GetAttribute("config_real_field"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.hcl), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			block := file.Body().Blocks()[0]
			ctx := &basic.TransformContext{}

			transformer := NestedRestructureTransformer(tt.source, tt.paths)
			err := transformer(block, ctx)
			require.NoError(t, err)

			tt.validate(t, block)
		})
	}
}

func BenchmarkFlattenNestedTransformer(b *testing.B) {
	hclContent := `resource "test" "example" {
  nested = {
    level1 = {
      field1 = "value1"
      field2 = "value2"
      level2 = {
        field3 = "value3"
        field4 = "value4"
      }
    }
    other = {
      data = "test"
    }
  }
}`

	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &basic.TransformContext{}
		
		transformer := FlattenNestedTransformer("nested", "_", 3)
		_ = transformer(block, ctx)
	}
}

func BenchmarkNestedRestructureTransformer(b *testing.B) {
	hclContent := `resource "test" "example" {
  config = {
    database = {
      host = "localhost"
      port = 5432
      user = "admin"
    }
    cache = {
      ttl = 300
      size = 1000
    }
  }
}`

	paths := map[string]string{
		"database.host": "db_host",
		"database.port": "db_port",
		"database.user": "db_user",
		"cache.ttl":     "cache_timeout",
		"cache.size":    "cache_max_size",
	}

	for i := 0; i < b.N; i++ {
		file, _ := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
		block := file.Body().Blocks()[0]
		ctx := &basic.TransformContext{}
		
		transformer := NestedRestructureTransformer("config", paths)
		_ = transformer(block, ctx)
	}
}
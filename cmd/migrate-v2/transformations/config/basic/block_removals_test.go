package basic

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	
)

func TestBlockRemover(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		blockTypes     []string
		expectedOutput string
	}{
		{
			name: "remove single block type",
			input: `resource "cloudflare_regional_hostname" "example" {
  zone_id  = "abc123"
  hostname = "example.com"
  
  timeouts {
    create = "30m"
    update = "30m"
  }
}`,
			blockTypes: []string{"timeouts"},
			expectedOutput: `resource "cloudflare_regional_hostname" "example" {
  zone_id  = "abc123"
  hostname = "example.com"
}`,
		},
		{
			name: "remove multiple block types",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  deprecated_config {
    old_setting = "value"
  }
  
  lifecycle {
    create_before_destroy = true
  }
  
  timeouts {
    create = "30m"
  }
}`,
			blockTypes: []string{"timeouts", "deprecated_config"},
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  lifecycle {
    create_before_destroy = true
  }
}`,
		},
		{
			name: "remove multiple blocks of same type",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  header {
    name  = "X-Custom-1"
    value = "value1"
  }
  
  header {
    name  = "X-Custom-2"
    value = "value2"
  }
  
  config {
    setting = "keep"
  }
}`,
			blockTypes: []string{"header"},
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  config {
    setting = "keep"
  }
}`,
		},
		{
			name: "remove nested blocks",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  configuration {
    setting = "value"
    
    deprecated {
      old = "remove_me"
    }
    
    keep {
      important = "data"
    }
  }
}`,
			blockTypes: []string{"deprecated"},
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  configuration {
    setting = "value"
    
    keep {
      important = "data"
    }
  }
}`,
		},
		{
			name: "no removal when block type not found",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  lifecycle {
    create_before_destroy = true
  }
}`,
			blockTypes: []string{"timeouts"},
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  lifecycle {
    create_before_destroy = true
  }
}`,
		},
		{
			name: "empty block types list",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  timeouts {
    create = "30m"
  }
}`,
			blockTypes: []string{},
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  timeouts {
    create = "30m"
  }
}`,
		},
		{
			name: "preserve other content when removing blocks",
			input: `resource "cloudflare_regional_hostname" "example" {
  zone_id  = "abc123"
  hostname = "example.com"
  region   = "us"
  
  timeouts {
    create = "30m"
    update = "30m"
    delete = "30m"
  }
  
  lifecycle {
    prevent_destroy = true
  }
}`,
			blockTypes: []string{"timeouts"},
			expectedOutput: `resource "cloudflare_regional_hostname" "example" {
  zone_id  = "abc123"
  hostname = "example.com"
  region   = "us"
  
  lifecycle {
    prevent_destroy = true
  }
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			// Get the first block
			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			block := blocks[0]

			// Create context and apply transformation
			ctx := &TransformContext{}
			transformer := BlockRemover(tt.blockTypes...)
			
			err := transformer(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Normalize and compare
			expectedNorm := normalizeHCLOutput(tt.expectedOutput)
			actualNorm := normalizeHCLOutput(output)

			assert.Equal(t, expectedNorm, actualNorm)
		})
	}
}

// normalizeHCLOutput normalizes HCL for comparison
func normalizeHCLOutput(s string) string {
	// Remove extra whitespace and normalize line endings
	lines := strings.Split(strings.TrimSpace(s), "\n")
	var normalized []string
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			// Also normalize spaces around = to handle HCL formatting variations
			trimmed = strings.ReplaceAll(trimmed, "  =", " =")
			trimmed = strings.ReplaceAll(trimmed, "=  ", "= ")
			normalized = append(normalized, trimmed)
		}
	}
	
	return strings.Join(normalized, "\n")
}

func TestSelectiveBlockRemover(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		criteria       func(*hclwrite.Block) bool
		expectedOutput string
	}{
		{
			name: "remove blocks with specific labels",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  rule "allow" {
    action = "allow"
  }
  
  rule "deny" {
    action = "deny"
  }
  
  rule "log" {
    action = "log"
  }
}`,
			criteria: func(block *hclwrite.Block) bool {
				if block.Type() == "rule" {
					labels := block.Labels()
					return len(labels) > 0 && labels[0] == "deny"
				}
				return false
			},
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  rule "allow" {
    action = "allow"
  }
  
  rule "log" {
    action = "log"
  }
}`,
		},
		{
			name: "remove blocks based on attribute content",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  origin {
    name    = "primary"
    address = "192.0.2.1"
    enabled = true
  }
  
  origin {
    name    = "secondary"
    address = "192.0.2.2"
    enabled = false
  }
  
  origin {
    name    = "tertiary"
    address = "192.0.2.3"
    enabled = true
  }
}`,
			criteria: func(block *hclwrite.Block) bool {
				if block.Type() == "origin" {
					// Check if enabled = false
					attr := block.Body().GetAttribute("enabled")
					if attr != nil {
						tokens := attr.Expr().BuildTokens(nil)
						value := strings.TrimSpace(string(tokens.Bytes()))
						return value == "false"
					}
				}
				return false
			},
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  origin {
    name    = "primary"
    address = "192.0.2.1"
    enabled = true
  }
  
  origin {
    name    = "tertiary"
    address = "192.0.2.3"
    enabled = true
  }
}`,
		},
		{
			name: "nil criteria returns unchanged",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  timeouts {
    create = "30m"
  }
}`,
			criteria: nil,
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  timeouts {
    create = "30m"
  }
}`,
		},
		{
			name: "complex nested removal",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  parent {
    name = "parent1"
    
    child {
      type = "remove"
      value = "data"
    }
    
    child {
      type = "keep"
      value = "important"
    }
  }
  
  parent {
    name = "parent2"
    
    child {
      type = "keep"
      value = "also_important"
    }
  }
}`,
			criteria: func(block *hclwrite.Block) bool {
				if block.Type() == "child" {
					attr := block.Body().GetAttribute("type")
					if attr != nil {
						tokens := attr.Expr().BuildTokens(nil)
						value := strings.TrimSpace(string(tokens.Bytes()))
						return value == `"remove"`
					}
				}
				return false
			},
			expectedOutput: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  parent {
    name = "parent1"
    
    child {
      type = "keep"
      value = "important"
    }
  }
  
  parent {
    name = "parent2"
    
    child {
      type = "keep"
      value = "also_important"
    }
  }
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			// Get the first block
			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			block := blocks[0]

			// Create context and apply transformation
			ctx := &TransformContext{}
			transformer := SelectiveBlockRemover(tt.criteria)
			
			err := transformer(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Normalize and compare
			expectedNorm := normalizeHCLOutput(tt.expectedOutput)
			actualNorm := normalizeHCLOutput(output)

			assert.Equal(t, expectedNorm, actualNorm)
		})
	}
}
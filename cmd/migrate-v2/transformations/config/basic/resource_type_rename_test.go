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

func TestResourceTypeChanger(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		change         *ResourceTypeChange
		expectedOutput string
		expectMoved    bool
	}{
		{
			name: "simple resource type change",
			input: `resource "cloudflare_managed_headers" "example" {
  zone_id = "abc123"
}`,
			change: &ResourceTypeChange{
				From:               "cloudflare_managed_headers",
				To:                 "cloudflare_managed_transforms",
				GenerateMovedBlock: false,
			},
			expectedOutput: `resource "cloudflare_managed_transforms" "example" {
  zone_id = "abc123"
}`,
			expectMoved: false,
		},
		{
			name: "resource type change with moved block",
			input: `resource "cloudflare_teams_list" "test" {
  account_id = "xyz789"
  name       = "test-list"
}`,
			change: &ResourceTypeChange{
				From:               "cloudflare_teams_list",
				To:                 "cloudflare_zero_trust_list",
				GenerateMovedBlock: true,
			},
			expectedOutput: `resource "cloudflare_zero_trust_list" "test" {
  account_id = "xyz789"
  name       = "test-list"
}

moved {
  from = cloudflare_teams_list.test
  to   = cloudflare_zero_trust_list.test
}`,
			expectMoved: true,
		},
		{
			name: "no change when resource type doesn't match",
			input: `resource "cloudflare_other_resource" "example" {
  zone_id = "abc123"
}`,
			change: &ResourceTypeChange{
				From:               "cloudflare_managed_headers",
				To:                 "cloudflare_managed_transforms",
				GenerateMovedBlock: false,
			},
			expectedOutput: `resource "cloudflare_other_resource" "example" {
  zone_id = "abc123"
}`,
			expectMoved: false,
		},
		{
			name: "multiple resources with only matching ones changed",
			input: `resource "cloudflare_managed_headers" "first" {
  zone_id = "abc123"
}

resource "cloudflare_other" "second" {
  zone_id = "def456"
}

resource "cloudflare_managed_headers" "third" {
  zone_id = "ghi789"
}`,
			change: &ResourceTypeChange{
				From:               "cloudflare_managed_headers",
				To:                 "cloudflare_managed_transforms",
				GenerateMovedBlock: true,
			},
			expectedOutput: `resource "cloudflare_managed_transforms" "first" {
  zone_id = "abc123"
}

resource "cloudflare_other" "second" {
  zone_id = "def456"
}

resource "cloudflare_managed_transforms" "third" {
  zone_id = "ghi789"
}

moved {
  from = cloudflare_managed_headers.first
  to   = cloudflare_managed_transforms.first
}

moved {
  from = cloudflare_managed_headers.third
  to   = cloudflare_managed_transforms.third
}`,
			expectMoved: true,
		},
		{
			name: "preserve nested blocks during type change",
			input: `resource "cloudflare_access_application" "app" {
  zone_id = "abc123"
  name    = "test-app"
  
  cors_headers {
    allowed_methods = ["GET", "POST"]
    allowed_origins = ["https://example.com"]
  }
  
  lifecycle {
    create_before_destroy = true
  }
}`,
			change: &ResourceTypeChange{
				From:               "cloudflare_access_application",
				To:                 "cloudflare_zero_trust_access_application",
				GenerateMovedBlock: false,
			},
			expectedOutput: `resource "cloudflare_zero_trust_access_application" "app" {
  zone_id = "abc123"
  name    = "test-app"
  
  cors_headers {
    allowed_methods = ["GET", "POST"]
    allowed_origins = ["https://example.com"]
  }
  
  lifecycle {
    create_before_destroy = true
  }
}`,
			expectMoved: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			// Get blocks
			blocks := file.Body().Blocks()
			
			// Create context and apply transformation
			ctx := &TransformContext{
				ResourceTypeChanges: make(map[*hclwrite.Block]string),
				MovedBlocks:        make(map[string]string),
			}
			
			transformer := ResourceTypeChanger(tt.change)
			for _, block := range blocks {
				err := transformer(block, ctx)
				require.NoError(t, err)
			}
			
			// Apply the resource type changes
			resultBlocks := ApplyResourceTypeChanges(blocks, ctx)
			
			// Reconstruct the file with new blocks
			newFile := hclwrite.NewFile()
			newBody := newFile.Body()
			
			for _, block := range resultBlocks {
				newBody.AppendBlock(block)
			}
			
			// Get output
			output := string(hclwrite.Format(newFile.Bytes()))
			
			// Check that the resource type was changed (only if it should have been)
			if tt.change != nil && tt.change.From != "" && tt.change.To != "" && strings.Contains(tt.input, tt.change.From) {
				assert.Contains(t, output, fmt.Sprintf(`resource "%s"`, tt.change.To))
				assert.NotContains(t, output, fmt.Sprintf(`resource "%s"`, tt.change.From))
			}
			
			// Check specific content exists (order-independent)
			for _, block := range resultBlocks {
				if block.Type() == "resource" {
					// Verify resource type change
					labels := block.Labels()
					if len(labels) > 1 {
						resourceType := labels[0]
						resourceName := labels[1]
						
						// Find the corresponding input block
						inputHasType := strings.Contains(tt.input, fmt.Sprintf(`resource "%s" "%s"`, tt.change.From, resourceName))
						if tt.change != nil && inputHasType {
							// This specific resource should have been changed
							assert.Equal(t, tt.change.To, resourceType)
						}
					}
				}
			}
			
			// Verify moved blocks were generated if expected
			if tt.expectMoved {
				assert.NotEmpty(t, ctx.MovedBlocks)
				// Check that moved blocks were actually added
				assert.Contains(t, output, "moved {")
			} else {
				assert.Empty(t, ctx.MovedBlocks)
			}
		})
	}
}

func TestResourceTypeChangerNilHandling(t *testing.T) {
	// Test with nil change
	transformer := ResourceTypeChanger(nil)
	
	input := `resource "cloudflare_managed_headers" "example" {
  zone_id = "abc123"
}`
	
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())
	
	blocks := file.Body().Blocks()
	ctx := &TransformContext{}
	
	// Should not error with nil change
	for _, block := range blocks {
		err := transformer(block, ctx)
		assert.NoError(t, err)
	}
	
	// No changes should be made
	assert.Empty(t, ctx.ResourceTypeChanges)
	assert.Empty(t, ctx.MovedBlocks)
}

// Helper function to normalize HCL for comparison
func normalizeHCL(s string) string {
	// Remove extra whitespace and normalize line endings
	lines := strings.Split(strings.TrimSpace(s), "\n")
	var normalized []string
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			normalized = append(normalized, trimmed)
		}
	}
	
	return strings.Join(normalized, "\n")
}
package structural

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/transformations/config/basic"
)

func TestBlockNester(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		blockName      string
		attributes     []string
		removeOriginals bool
		checkFunc      func(t *testing.T, output string)
	}{
		{
			name: "nest simple attributes into block",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  address = "192.168.1.1"
  port = 8080
  protocol = "tcp"
}`,
			blockName:      "network",
			attributes:     []string{"address", "port", "protocol"},
			removeOriginals: true,
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "network {")
				assert.Contains(t, output, `address  = "192.168.1.1"`)
				assert.Contains(t, output, "port     = 8080")
				assert.Contains(t, output, `protocol = "tcp"`)
				// Original attributes should be removed from top level
				assert.Equal(t, 1, countOccurrences(output, "address"))
				assert.Equal(t, 1, countOccurrences(output, "port"))
			},
		},
		{
			name: "nest attributes without removing originals",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  ssl = true
  min_tls = "1.2"
}`,
			blockName:      "tls_config",
			attributes:     []string{"ssl", "min_tls"},
			removeOriginals: false,
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "tls_config {")
				// Should have attributes in both places
				assert.Equal(t, 2, countOccurrences(output, "ssl"))
				assert.Equal(t, 2, countOccurrences(output, "min_tls"))
			},
		},
		{
			name: "nest partial attributes when some missing",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  enabled = true
}`,
			blockName:      "settings",
			attributes:     []string{"enabled", "timeout", "retry_count"},
			removeOriginals: true,
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "settings {")
				assert.Contains(t, output, "enabled = true")
				assert.NotContains(t, output, "timeout")
				assert.NotContains(t, output, "retry_count")
			},
		},
		{
			name: "no nesting when no attributes found",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
}`,
			blockName:      "missing",
			attributes:     []string{"attr1", "attr2"},
			removeOriginals: true,
			checkFunc: func(t *testing.T, output string) {
				assert.NotContains(t, output, "missing {")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			block := blocks[0]

			// Apply transformation
			ctx := &basic.TransformContext{}
			nester := BlockNester(tt.blockName, tt.attributes, tt.removeOriginals)
			
			err := nester(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Run custom checks
			if tt.checkFunc != nil {
				tt.checkFunc(t, output)
			}
		})
	}
}

func TestBlockUnnester(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		blockName         string
		promoteAttributes []string
		checkFunc         func(t *testing.T, output string)
	}{
		{
			name: "unnest all attributes from block",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  network {
    address = "192.168.1.1"
    port = 8080
    protocol = "tcp"
  }
}`,
			blockName:         "network",
			promoteAttributes: []string{},
			checkFunc: func(t *testing.T, output string) {
				assert.NotContains(t, output, "network {")
				assert.Contains(t, output, `address  = "192.168.1.1"`)
				assert.Contains(t, output, "port     = 8080")
				assert.Contains(t, output, `protocol = "tcp"`)
			},
		},
		{
			name: "unnest specific attributes",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  settings {
    enabled = true
    timeout = 30
    debug = false
  }
}`,
			blockName:         "settings",
			promoteAttributes: []string{"enabled", "timeout"},
			checkFunc: func(t *testing.T, output string) {
				assert.NotContains(t, output, "settings {")
				assert.Contains(t, output, "enabled = true")
				assert.Contains(t, output, "timeout = 30")
				// debug should not be promoted
				assert.NotContains(t, output, "debug")
			},
		},
		{
			name: "no unnesting when block not found",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  name = "example"
}`,
			blockName:         "nonexistent",
			promoteAttributes: []string{},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "zone_id")
				assert.Contains(t, output, "name")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			block := blocks[0]

			// Apply transformation
			ctx := &basic.TransformContext{}
			unnester := BlockUnnester(tt.blockName, tt.promoteAttributes)
			
			err := unnester(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Run custom checks
			if tt.checkFunc != nil {
				tt.checkFunc(t, output)
			}
		})
	}
}

func TestMultiLevelNester(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		nestingSpec []NestingLevel
		checkFunc   func(t *testing.T, output string)
	}{
		{
			name: "two-level nesting",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  host = "example.com"
  port = 443
  ssl = true
  cipher = "AES256"
}`,
			nestingSpec: []NestingLevel{
				{
					BlockName:       "connection",
					Attributes:      []string{"host", "port"},
					RemoveOriginals: true,
				},
				{
					BlockName:       "tls",
					Attributes:      []string{"ssl", "cipher"},
					RemoveOriginals: false,
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "connection {")
				assert.Contains(t, output, "tls {")
				// Check nesting structure
				assert.Contains(t, output, "host")
				assert.Contains(t, output, "port")
				assert.Contains(t, output, "ssl")
				assert.Contains(t, output, "cipher")
			},
		},
		{
			name: "nesting with labels",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  rule = "allow"
  priority = 1
}`,
			nestingSpec: []NestingLevel{
				{
					BlockName:       "policy",
					Labels:          []string{"firewall"},
					Attributes:      []string{"rule", "priority"},
					RemoveOriginals: true,
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `policy "firewall" {`)
				assert.Contains(t, output, "rule")
				assert.Contains(t, output, "priority")
			},
		},
		{
			name: "create empty blocks",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
}`,
			nestingSpec: []NestingLevel{
				{
					BlockName:   "defaults",
					Attributes:  []string{"nonexistent"},
					CreateEmpty: true,
				},
			},
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, "defaults {")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			block := blocks[0]

			// Apply transformation
			ctx := &basic.TransformContext{}
			nester := MultiLevelNester(tt.nestingSpec)
			
			err := nester(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Run custom checks
			if tt.checkFunc != nil {
				tt.checkFunc(t, output)
			}
		})
	}
}

func TestDynamicBlockConverter(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		blockType    string
		iteratorName string
		checkFunc    func(t *testing.T, output string)
	}{
		{
			name: "convert repeated blocks to dynamic",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  rule {
    expression = "ip.src eq 1.2.3.4"
    action = "allow"
  }
  
  rule {
    expression = "ip.src eq 5.6.7.8"
    action = "deny"
  }
  
  rule {
    expression = "ip.src eq 9.10.11.12"
    action = "challenge"
  }
}`,
			blockType:    "rule",
			iteratorName: "",
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `dynamic "rule" {`)
				assert.Contains(t, output, "for_each")
				assert.Contains(t, output, "content {")
				// Original rule blocks should be removed
				assert.Equal(t, 1, countOccurrences(output, `dynamic "rule"`))
			},
		},
		{
			name: "dynamic block with custom iterator",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  origin {
    address = "192.168.1.1"
    enabled = true
  }
  
  origin {
    address = "192.168.1.2"
    enabled = false
  }
}`,
			blockType:    "origin",
			iteratorName: "each_origin",
			checkFunc: func(t *testing.T, output string) {
				assert.Contains(t, output, `dynamic "origin" {`)
				assert.Contains(t, output, `iterator = "each_origin"`)
				assert.Contains(t, output, "content {")
			},
		},
		{
			name: "no conversion for single block",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  
  rule {
    expression = "ip.src eq 1.2.3.4"
    action = "allow"
  }
}`,
			blockType:    "rule",
			iteratorName: "",
			checkFunc: func(t *testing.T, output string) {
				assert.NotContains(t, output, "dynamic")
				assert.Contains(t, output, "rule {")
			},
		},
		{
			name: "no conversion when no blocks found",
			input: `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  name = "example"
}`,
			blockType:    "nonexistent",
			iteratorName: "",
			checkFunc: func(t *testing.T, output string) {
				assert.NotContains(t, output, "dynamic")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			blocks := file.Body().Blocks()
			require.Len(t, blocks, 1)
			block := blocks[0]

			// Apply transformation
			ctx := &basic.TransformContext{}
			converter := DynamicBlockConverter(tt.blockType, tt.iteratorName)
			
			err := converter(block, ctx)
			require.NoError(t, err)

			// Get output
			output := string(hclwrite.Format(file.Bytes()))

			// Run custom checks
			if tt.checkFunc != nil {
				tt.checkFunc(t, output)
			}
		})
	}
}

func TestBlockNesterForState(t *testing.T) {
	tests := []struct {
		name          string
		state         map[string]interface{}
		blockName     string
		attributes    []string
		expectedState map[string]interface{}
	}{
		{
			name: "nest attributes in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":  "abc123",
					"address":  "192.168.1.1",
					"port":     8080,
					"protocol": "tcp",
				},
			},
			blockName:  "network",
			attributes: []string{"address", "port", "protocol"},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"network": []interface{}{
						map[string]interface{}{
							"address":  "192.168.1.1",
							"port":     8080,
							"protocol": "tcp",
						},
					},
				},
			},
		},
		{
			name: "partial nesting when some attributes missing",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"enabled": true,
				},
			},
			blockName:  "settings",
			attributes: []string{"enabled", "timeout"},
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"settings": []interface{}{
						map[string]interface{}{
							"enabled": true,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply state transformation
			nester := BlockNesterForState(tt.blockName, tt.attributes)
			err := nester(tt.state)
			require.NoError(t, err)

			// Compare states
			assert.Equal(t, tt.expectedState, tt.state)
		})
	}
}

func TestBlockUnnesterForState(t *testing.T) {
	tests := []struct {
		name          string
		state         map[string]interface{}
		blockName     string
		expectedState map[string]interface{}
	}{
		{
			name: "unnest block in state",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"network": []interface{}{
						map[string]interface{}{
							"address":  "192.168.1.1",
							"port":     8080,
							"protocol": "tcp",
						},
					},
				},
			},
			blockName: "network",
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id":  "abc123",
					"address":  "192.168.1.1",
					"port":     8080,
					"protocol": "tcp",
				},
			},
		},
		{
			name: "unnest single block format",
			state: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"settings": map[string]interface{}{
						"enabled": true,
						"timeout": 30,
					},
				},
			},
			blockName: "settings",
			expectedState: map[string]interface{}{
				"attributes": map[string]interface{}{
					"zone_id": "abc123",
					"enabled": true,
					"timeout": 30,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Apply state transformation
			unnester := BlockUnnesterForState(tt.blockName)
			err := unnester(tt.state)
			require.NoError(t, err)

			// Compare states
			assert.Equal(t, tt.expectedState, tt.state)
		})
	}
}

func TestConditionalNester(t *testing.T) {
	input := `resource "cloudflare_example" "test" {
  zone_id = "abc123"
  type = "premium"
  setting1 = "value1"
  setting2 = "value2"
}`

	// Parse input
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	blocks := file.Body().Blocks()
	require.Len(t, blocks, 1)
	block := blocks[0]

	// Create conditional nester
	condition := func(b *hclwrite.Block) bool {
		body := b.Body()
		if attr := body.GetAttribute("type"); attr != nil {
			tokens := attr.Expr().BuildTokens(nil)
			value := strings.TrimSpace(string(tokens.Bytes()))
			return value == `"premium"`
		}
		return false
	}

	ctx := &basic.TransformContext{}
	nester := ConditionalNester("premium_settings", condition, []string{"setting1", "setting2"})
	
	err := nester(block, ctx)
	require.NoError(t, err)

	// Get output
	output := string(hclwrite.Format(file.Bytes()))

	// Should have nested the settings
	assert.Contains(t, output, "premium_settings {")
	assert.Contains(t, output, "setting1")
	assert.Contains(t, output, "setting2")
}

// Helper function to count occurrences of a string
func countOccurrences(s, substr string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if i+len(substr) <= len(s) && s[i:i+len(substr)] == substr {
			count++
		}
	}
	return count
}
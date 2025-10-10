package struct_transform

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStructTransformHandler(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []string // Strings that should appear in output
		notExpected []string // Strings that should NOT appear
	}{
		{
			name: "transform A record",
			input: `
resource "cloudflare_record" "test" {
  zone_id = "zone123"
  name    = "example.com"
  type    = "A"
  value   = "192.0.2.1"
  ttl     = 300
  proxied = true
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "test"`, // Resource type renamed
				`content = "192.0.2.1"`,                   // value -> content
				`zone_id = "zone123"`,
				`name    = "example.com"`,
				`type    = "A"`,
				`ttl     = 300`,
				`proxied = true`,
			},
			notExpected: []string{
				`value =`,               // Should be renamed to content
				`cloudflare_record`,     // Should be renamed to cloudflare_dns_record
			},
		},
		{
			name: "transform with deprecated fields",
			input: `
resource "cloudflare_record" "deprecated" {
  zone_id         = "zone456"
  hostname        = "old.example.com"
  type            = "CNAME"
  value           = "target.example.com"
  allow_overwrite = true
  ttl             = 3600
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "deprecated"`,
				`content = "target.example.com"`,
				`zone_id = "zone456"`,
				`type    = "CNAME"`,
				`ttl     = 3600`,
			},
			notExpected: []string{
				`hostname`,        // Deprecated field
				`allow_overwrite`, // Deprecated field
				`value =`,
			},
		},
		{
			name: "transform CAA record with data block",
			input: `
resource "cloudflare_record" "caa" {
  zone_id = "zone789"
  name    = "example.com"
  type    = "CAA"

  data {
    flags   = "0"
    tag     = "issue"
    content = "letsencrypt.org"
  }
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "caa"`,
				`zone_id = "zone789"`,
				`name    = "example.com"`,
				`type    = "CAA"`,
				`data {`,
				`flags = 0`,                      // String "0" -> number 0
				`tag   = "issue"`,
				`value = "letsencrypt.org"`,      // content -> value in data block
			},
			notExpected: []string{
				`content =`, // In data block, should be 'value'
				`flags   = "0"`, // Should be numeric
			},
		},
		{
			name: "transform SRV record with data block",
			input: `
resource "cloudflare_record" "srv" {
  zone_id = "zone_srv"
  name    = "_sip._tcp"
  type    = "SRV"

  data {
    priority = 10
    weight   = 60
    port     = 5060
    target   = "sip.example.com"
    service  = "_sip"
    proto    = "_tcp"
    name     = "example.com"
  }
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "srv"`,
				`data {`,
				`priority = 10`,
				`weight   = 60`,
				`port     = 5060`,
				`target   = "sip.example.com"`,
				`service  = "_sip"`,
			},
			notExpected: []string{
				`proto =`, // Removed in v5
				`name  =`, // Removed in v5 data block
			},
		},
		{
			name: "add default TTL when missing",
			input: `
resource "cloudflare_record" "no_ttl" {
  zone_id = "zone_ttl"
  name    = "auto-ttl"
  type    = "A"
  value   = "192.0.2.2"
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "no_ttl"`,
				`ttl     = 1`, // Default TTL added
			},
		},
		{
			name: "skip transformation for non-cloudflare resource",
			input: `
resource "aws_instance" "skip" {
  ami           = "ami-12345"
  instance_type = "t2.micro"
}`,
			expected: []string{
				`aws_instance`,    // Should remain unchanged
				`ami           =`, // Should remain unchanged
			},
			notExpected: []string{
				`cloudflare_dns_record`, // Should NOT appear
			},
		},
		{
			name: "handle multiple resources",
			input: `
resource "cloudflare_record" "first" {
  zone_id = "zone1"
  name    = "first"
  type    = "A"
  value   = "192.0.2.10"
  ttl     = 300
}

resource "aws_instance" "other" {
  ami           = "ami-12345"
  instance_type = "t2.micro"
}

resource "cloudflare_dns_record" "second" {
  zone_id = "zone2"
  name    = "second"
  type    = "CNAME"
  content = "target.example.com"
  ttl     = 600
}`,
			expected: []string{
				`resource "cloudflare_dns_record" "first"`,
				`content = "192.0.2.10"`,
				`resource "aws_instance" "other"`,  // Non-CF resource unchanged
				`ami           = "ami-12345"`,
				`resource "cloudflare_dns_record" "second"`, // Already v5, unchanged
				`content = "target.example.com"`,
			},
			notExpected: []string{
				`cloudflare_record`,
				`value = "192.0.2.10"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			// Create context
			ctx := &interfaces.TransformContext{
				Content:     []byte(tt.input),
				Filename:    "test.tf",
				AST:         file,
				Diagnostics: nil,
				Metadata:    make(map[string]interface{}),
			}

			// Create handler with struct-based registry
			handler := NewStructTransformHandler(nil) // Will use default struct registry

			// Process
			result, err := handler.Handle(ctx)
			assert.NoError(t, err)

			// Get output
			output := string(result.AST.Bytes())

			// Check expected strings
			for _, expected := range tt.expected {
				assert.Contains(t, output, expected, "Output should contain: %s", expected)
			}

			// Check strings that should NOT appear
			for _, notExpected := range tt.notExpected {
				assert.NotContains(t, output, notExpected, "Output should NOT contain: %s", notExpected)
			}

			// Check metadata was added for transformed resources
			if _, ok := result.Metadata["struct_mode_used"]; ok {
				assert.True(t, result.Metadata["struct_mode_used"].(bool))

				// Check transformation count
				if _, ok := result.Metadata["struct_transformed_cloudflare_record"]; ok {
					count := result.Metadata["struct_transformed_cloudflare_record"].(int)
					assert.Greater(t, count, 0, "Should have transformed at least one record")
				}
			}
		})
	}
}

func TestStructTransformHandlerErrors(t *testing.T) {
	tests := []struct {
		name        string
		ctx         *interfaces.TransformContext
		expectError bool
		errorMsg    string
	}{
		{
			name: "nil AST",
			ctx: &interfaces.TransformContext{
				Content:  []byte("test"),
				Filename: "test.tf",
				AST:      nil,
			},
			expectError: true,
			errorMsg:    "AST is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := NewStructTransformHandler(nil)
			result, err := handler.Handle(tt.ctx)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestStructTransformIntegration(t *testing.T) {
	// Test the complete transformation pipeline
	input := `
# Test file with various DNS records
resource "cloudflare_record" "a_record" {
  zone_id = var.zone_id
  name    = "test-a"
  type    = "A"
  value   = "192.0.2.100"
  ttl     = 300
  proxied = true
}

resource "cloudflare_record" "mx_record" {
  zone_id  = var.zone_id
  name     = "@"
  type     = "MX"
  value    = "mail.example.com"
  priority = 10
  ttl      = 3600
}

resource "cloudflare_record" "caa_record" {
  zone_id = var.zone_id
  name    = "@"
  type    = "CAA"

  data {
    flags   = "0"
    tag     = "issue"
    content = "letsencrypt.org"
  }

  data {
    flags   = "128"
    tag     = "issuewild"
    content = ";"
  }
}

resource "cloudflare_record" "srv_record" {
  zone_id = var.zone_id
  name    = "_sip._tcp"
  type    = "SRV"
  ttl     = 86400

  data {
    priority = 10
    weight   = 60
    port     = 5060
    target   = "sip.example.com"
    service  = "_sip"
    proto    = "_tcp"
    name     = "example.com"
  }
}`

	// Parse input
	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	// Create context
	ctx := &interfaces.TransformContext{
		Content:     []byte(input),
		Filename:    "test.tf",
		AST:         file,
		Diagnostics: nil,
		Metadata:    make(map[string]interface{}),
	}

	// Create and run handler
	handler := NewStructTransformHandler(nil) // Uses default struct registry
	result, err := handler.Handle(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	output := string(result.AST.Bytes())

	// Verify all resources were transformed
	assert.Contains(t, output, `resource "cloudflare_dns_record" "a_record"`)
	assert.Contains(t, output, `resource "cloudflare_dns_record" "mx_record"`)
	assert.Contains(t, output, `resource "cloudflare_dns_record" "caa_record"`)
	assert.Contains(t, output, `resource "cloudflare_dns_record" "srv_record"`)

	// Verify value -> content transformation
	assert.NotContains(t, output, `value   = "192.0.2.100"`)
	assert.Contains(t, output, `content = "192.0.2.100"`)
	assert.NotContains(t, output, `value    = "mail.example.com"`)
	assert.Contains(t, output, `content  = "mail.example.com"`)

	// Verify CAA flags conversion (only first data block is transformed)
	assert.NotContains(t, output, `flags   = "0"`)
	assert.Contains(t, output, `flags = 0`)
	// Note: Second data block is not transformed in current implementation
	// as v5 schema uses single data object, not array

	// Verify CAA content -> value in data block
	assert.Contains(t, output, `value = "letsencrypt.org"`)

	// Verify SRV proto/name removal
	assert.NotContains(t, output, `proto =`)
	assert.NotContains(t, output, `name     = "example.com"`)

	// Verify priority handling
	assert.Contains(t, output, `priority = 10`)

	// Check metadata
	assert.True(t, result.Metadata["struct_mode_used"].(bool))
	assert.Equal(t, 4, result.Metadata["struct_transformed_cloudflare_record"].(int))
}

func TestChainOfResponsibility(t *testing.T) {
	// Test that the handler properly chains to the next handler
	input := `
resource "cloudflare_record" "test" {
  zone_id = "zone123"
  name    = "test"
  type    = "A"
  value   = "192.0.2.1"
}`

	file, diags := hclwrite.ParseConfig([]byte(input), "test.tf", hcl.InitialPos)
	require.False(t, diags.HasErrors())

	ctx := &interfaces.TransformContext{
		Content:  []byte(input),
		Filename: "test.tf",
		AST:      file,
		Metadata: make(map[string]interface{}),
	}

	// Create a mock next handler
	nextHandlerCalled := false
	nextHandler := &mockHandler{
		handleFunc: func(ctx *interfaces.TransformContext) (*interfaces.TransformContext, error) {
			nextHandlerCalled = true
			ctx.Metadata["next_handler_called"] = true
			return ctx, nil
		},
	}

	// Test with struct handler
	handler := NewStructTransformHandler(nil)
	handler.SetNext(nextHandler)

	result, err := handler.Handle(ctx)
	assert.NoError(t, err)
	assert.True(t, nextHandlerCalled, "Next handler should be called")
	assert.True(t, result.Metadata["next_handler_called"].(bool))
	assert.True(t, result.Metadata["struct_mode_used"].(bool))
}

// mockHandler implements TransformationHandler for testing
type mockHandler struct {
	interfaces.BaseHandler
	handleFunc func(*interfaces.TransformContext) (*interfaces.TransformContext, error)
}

func (m *mockHandler) Handle(ctx *interfaces.TransformContext) (*interfaces.TransformContext, error) {
	if m.handleFunc != nil {
		return m.handleFunc(ctx)
	}
	return ctx, nil
}
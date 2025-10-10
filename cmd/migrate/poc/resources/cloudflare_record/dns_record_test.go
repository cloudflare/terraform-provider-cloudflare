package cloudflare_record_test

import (
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/resources/cloudflare_record"
	"github.com/tidwall/gjson"
)

// TestDNSRecordTransformation tests the DNS record wrapper
func TestDNSRecordTransformation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string // Multiple expected strings to check for
	}{
		{
			name: "Basic DNS record transformation",
			input: `resource "cloudflare_record" "example" {
  zone_id = "abc123"
  name    = "example"
  type    = "A"
  value   = "192.0.2.1"
}`,
			expected: []string{
				"cloudflare_dns_record", // Resource type renamed
				"ttl",                   // TTL added
				"content",               // value renamed to content
				"!value",                // value should be removed
			},
		},
		{
			name: "DNS record with deprecated attributes",
			input: `resource "cloudflare_record" "example" {
  zone_id         = "abc123"
  name            = "example"
  type            = "A"
  value           = "192.0.2.1"
  allow_overwrite = true
  hostname        = "old.example.com"
}`,
			expected: []string{
				"cloudflare_dns_record",
				"ttl",
				"content",
				"!allow_overwrite", // Should be removed
				"!hostname",        // Should be removed
			},
		},
		{
			name: "CAA record with data block",
			input: `resource "cloudflare_record" "caa" {
  zone_id = "abc123"
  name    = "example"
  type    = "CAA"
  data {
    flags = "0"
    tag   = "issue"
    content = "ca.example.com"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"ttl",
				"data =", // data block converted to attribute
			},
		},
		{
			name: "SRV record with priority",
			input: `resource "cloudflare_record" "srv" {
  zone_id = "abc123"
  name    = "_service._proto.example"
  type    = "SRV"
  data {
    priority = 10
    weight   = 60
    port     = 5060
    target   = "srv.example.com"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"ttl",
				"priority", // Priority should be at root level
				"data =",   // data block converted to attribute
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a production pipeline with the wrapper
			reg := registry.NewStrategyRegistry()
			reg.Register(cloudflare_record.NewDNSRecord())

			// Build pipeline
			pipeline := poc.NewPipelineBuilder().
				WithPreprocessing(reg).
				WithParsing().
				WithResourceTransformation(reg).
				WithFormatting().
				Build()

			// TransformConfig
			result, err := pipeline.Transform([]byte(tt.input), "test.tf")
			if err != nil {
				t.Fatalf("Transformation failed: %v", err)
			}

			resultStr := string(result)

			// Check expected patterns
			for _, expected := range tt.expected {
				if strings.HasPrefix(expected, "!") {
					// Should NOT contain
					pattern := expected[1:]
					if strings.Contains(resultStr, pattern) {
						t.Errorf("Result should not contain '%s', but it does:\n%s", pattern, resultStr)
					}
				} else {
					// Should contain
					if !strings.Contains(resultStr, expected) {
						t.Errorf("Result should contain '%s', but it doesn't:\n%s", expected, resultStr)
					}
				}
			}
		})
	}
}

// TestDNSRecordStateTransformation tests state file transformations
func TestDNSRecordStateTransformation(t *testing.T) {
	stateJSON := `{
		"resources": [{
			"type": "cloudflare_record",
			"name": "example",
			"instances": [{
				"attributes": {
					"zone_id": "abc123",
					"name": "example.com",
					"type": "A",
					"value": "192.0.2.1"
				}
			}]
		}]
	}`

	wrapper := cloudflare_record.NewDNSRecord()

	json := gjson.Parse(stateJSON)
	resourcePath := "resources.0"

	result, err := wrapper.TransformState(json, resourcePath)
	if err != nil {
		t.Fatalf("State transformation failed: %v", err)
	}

	// Check that resource type was renamed
	newType := gjson.Get(result, resourcePath+".type").String()
	if newType != "cloudflare_dns_record" {
		t.Errorf("Resource type not transformed. Got: %s", newType)
	}
}

// TestPipelineWithMultipleHandlers tests the full pipeline
func TestPipelineWithMultipleHandlers(t *testing.T) {
	// This test demonstrates how the pipeline would work
	// with multiple handlers processing different aspects

	input := `resource "cloudflare_record" "test1" {
  zone_id = "zone1"
  name    = "test1"
  type    = "A"
  value   = "1.2.3.4"
}

resource "cloudflare_load_balancer" "lb" {
  zone_id = "zone1"
  name    = "example-lb"
}

resource "cloudflare_record" "test2" {
  zone_id = "zone1"
  name    = "test2"
  type    = "CNAME"
  value   = "example.com"
}`

	// Build pipeline
	reg := registry.NewStrategyRegistry()
	reg.Register(cloudflare_record.NewDNSRecord())
	// In full implementation: reg.Register(poc.NewLoadBalancerWrapper())

	pipeline := poc.NewPipelineBuilder().
		WithPreprocessing(reg).
		WithParsing().
		WithResourceTransformation(reg).
		WithFormatting().
		Build()

	result, err := pipeline.Transform([]byte(input), "multi.tf")
	if err != nil {
		t.Fatalf("Transformation failed: %v", err)
	}

	resultStr := string(result)

	// Check DNS records were transformed
	count := strings.Count(resultStr, "cloudflare_dns_record")
	if count != 2 {
		t.Errorf("Expected 2 DNS records to be transformed, got %d. Output:\n%s", count, resultStr)
	}

	// Check load balancer was NOT transformed (no strategy registered)
	if !strings.Contains(resultStr, "cloudflare_load_balancer") {
		t.Error("Load balancer should remain unchanged")
	}

	// Check TTL was added to DNS records
	ttlCount := strings.Count(resultStr, "ttl")
	if ttlCount < 2 {
		t.Error("Expected TTL to be added to DNS records")
	}
}

// TestComplexDataTransformation tests complex record types
func TestComplexDataTransformation(t *testing.T) {
	input := `resource "cloudflare_record" "caa" {
  zone_id = "abc123"
  name    = "example.com"
  type    = "CAA"
  data {
    flags   = "0"
    tag     = "issue"
    content = "letsencrypt.org"
  }
}`

	reg := registry.NewStrategyRegistry()
	reg.Register(cloudflare_record.NewDNSRecord())

	pipeline := poc.NewPipelineBuilder().
		WithPreprocessing(reg).
		WithParsing().
		WithResourceTransformation(reg).
		WithFormatting().
		Build()

	result, err := pipeline.Transform([]byte(input), "caa.tf")
	if err != nil {
		t.Fatalf("Transformation failed: %v", err)
	}

	resultStr := string(result)

	// Check transformations
	if !strings.Contains(resultStr, "cloudflare_dns_record") {
		t.Error("Resource type not transformed")
	}

	if !strings.Contains(resultStr, "ttl") {
		t.Error("TTL not added")
	}

	// Data should be converted from block to attribute
	if strings.Contains(resultStr, "data {") {
		t.Error("Data block should be converted to attribute")
	}

	if !strings.Contains(resultStr, "data =") {
		t.Error("Data attribute not found")
	}
}

// TestIntegrationWithExistingCode shows how this would integrate with existing code
func TestIntegrationWithExistingCode(t *testing.T) {
	// This test demonstrates the integration pattern
	t.Log(`
Integration Pattern:

1. Create wrapper strategies for each resource type:
   - DNSRecordWrapper (wraps ProcessDNSRecordConfig)
   - LoadBalancerWrapper (wraps transformLoadBalancerPoolBlock)
   - ZoneSettingsWrapper (wraps zone settings logic)

2. Register all strategies in the registry:
   registry.Register(NewDNSRecordWrapper())
   registry.Register(NewLoadBalancerWrapper())
   ... (30+ strategies)

3. Build the pipeline with all handlers:
   - GritHandler (if --grit flag set)
   - StringPreprocessHandler (all string transformations)
   - ParseHandler (HCL parsing)
   - ResourceTransformHandler (apply strategies)
   - CrossResourceHandler (workers secrets, etc.)
   - ImportGeneratorHandler (for split resources)
   - ValidationHandler (check completeness)
   - FormatterHandler (final output)

4. Use feature flag for gradual rollout:
   if *useNewPipeline {
       return pipeline.TransformConfig(content, filename)
   }
   return legacyTransformFile(content, filename)

This allows:
- Zero risk migration (feature flag protection)
- Reuse of ALL existing transformation logic
- Gradual refactoring of internals
- Clear separation of concerns
- Better testability and maintainability
`)
}

// MockComplexHandler demonstrates a custom handler for testing
type MockComplexHandler struct {
	interfaces.BaseHandler
	processed bool
}

func (h *MockComplexHandler) Handle(ctx *interfaces.TransformContext) (*interfaces.TransformContext, error) {
	h.processed = true
	// Add metadata to track processing
	if ctx.Metadata == nil {
		ctx.Metadata = make(map[string]interface{})
	}
	ctx.Metadata["mock_processed"] = true
	return h.CallNext(ctx)
}

// TestCustomHandlerIntegration tests adding custom handlers to the pipeline
func TestCustomHandlerIntegration(t *testing.T) {
	mockHandler := &MockComplexHandler{}

	// Create an empty registry for this test (no resources needed)
	reg := registry.NewStrategyRegistry()

	// Build pipeline with custom handler
	pipeline := poc.NewPipelineBuilder().
		WithPreprocessing(reg).
		WithParsing().
		WithHandler(mockHandler). // Custom handler
		WithFormatting().
		Build()

	input := `resource "cloudflare_record" "test" {
  zone_id = "test"
  name    = "test"
  type    = "A"
  value   = "1.2.3.4"
}`

	result, err := pipeline.Transform([]byte(input), "test.tf")
	if err != nil {
		t.Fatalf("Transformation failed: %v", err)
	}

	if !mockHandler.processed {
		t.Error("Custom handler was not processed")
	}

	// Result should still be valid HCL
	if len(result) == 0 {
		t.Error("Empty result")
	}
}

// TestCompleteCAA tests the complete CAA record transformation including data attribute handling
func TestCompleteCAATransformation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "CAA with data block",
			input: `resource "cloudflare_record" "caa" {
  zone_id = "abc123"
  name    = "example.com"
  type    = "CAA"
  data {
    flags   = "0"
    tag     = "issue"
    content = "ca.example.com"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"ttl",
				`data =`,   // Converted to attribute
				"!data {",  // Block should be removed
				"value",    // content renamed to value in data
				"!content", // content should not be in data
			},
		},
		{
			name: "CAA with data attribute and flags conversion",
			input: `resource "cloudflare_record" "caa" {
  zone_id = "abc123"
  name    = "example.com"
  type    = "CAA"
  data = {
    flags   = "0"
    tag     = "issue"
    content = "ca.example.com"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"ttl",
				`data =`,
				"value",    // content renamed to value
				"!content", // content removed from data
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := registry.NewStrategyRegistry()
			reg.Register(cloudflare_record.NewDNSRecord())

			pipeline := poc.NewPipelineBuilder().
				WithPreprocessing(reg).
				WithParsing().
				WithResourceTransformation(reg).
				WithFormatting().
				Build()

			result, err := pipeline.Transform([]byte(tt.input), "test.tf")
			if err != nil {
				t.Fatalf("Transformation failed: %v", err)
			}

			resultStr := string(result)

			for _, expected := range tt.expected {
				if strings.HasPrefix(expected, "!") {
					pattern := expected[1:]
					if strings.Contains(resultStr, pattern) {
						t.Errorf("Result should not contain '%s', but it does:\n%s", pattern, resultStr)
					}
				} else {
					if !strings.Contains(resultStr, expected) {
						t.Errorf("Result should contain '%s', but it doesn't:\n%s", expected, resultStr)
					}
				}
			}
		})
	}
}

// TestCompleteStateTransformation tests the complete state transformation logic
func TestCompleteStateTransformation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, result string)
	}{
		{
			name: "Simple record state transformation",
			input: `{
				"resources": [{
					"type": "cloudflare_record",
					"name": "example",
					"instances": [{
						"attributes": {
							"zone_id": "abc123",
							"name": "example.com",
							"type": "A",
							"value": "192.0.2.1",
							"allow_overwrite": true,
							"hostname": "old.example.com"
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check resource type
				if !gjson.Get(result, "resources.0.type").Exists() ||
					gjson.Get(result, "resources.0.type").String() != "cloudflare_dns_record" {
					t.Error("Resource type not transformed")
				}

				// Check content field
				if !gjson.Get(result, "resources.0.instances.0.attributes.content").Exists() {
					t.Error("Value not renamed to content")
				}

				// Check TTL added
				if !gjson.Get(result, "resources.0.instances.0.attributes.ttl").Exists() {
					t.Error("TTL not added")
				}

				// Check deprecated fields removed
				if gjson.Get(result, "resources.0.instances.0.attributes.allow_overwrite").Exists() {
					t.Error("allow_overwrite not removed")
				}
				if gjson.Get(result, "resources.0.instances.0.attributes.hostname").Exists() {
					t.Error("hostname not removed")
				}
			},
		},
		{
			name: "CAA record state transformation",
			input: `{
				"resources": [{
					"type": "cloudflare_record",
					"name": "caa",
					"instances": [{
						"attributes": {
							"zone_id": "abc123",
							"name": "example.com",
							"type": "CAA",
							"data": [{
								"flags": "0",
								"tag": "issue",
								"content": "ca.example.com"
							}]
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check data is object not array
				data := gjson.Get(result, "resources.0.instances.0.attributes.data")
				if !data.IsObject() {
					t.Error("Data should be object, not array")
				}

				// Check flags wrapping
				flags := gjson.Get(result, "resources.0.instances.0.attributes.data.flags")
				if !flags.IsObject() || !flags.Get("value").Exists() {
					t.Error("Flags should be wrapped object")
				}

				// Check content -> value rename
				if !gjson.Get(result, "resources.0.instances.0.attributes.data.value").Exists() {
					t.Error("Content not renamed to value in data")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapper := cloudflare_record.NewDNSRecord()
			json := gjson.Parse(tt.input)
			result, err := wrapper.TransformState(json, "resources.0")
			if err != nil {
				t.Fatalf("State transformation failed: %v", err)
			}

			tt.check(t, result)
		})
	}
}

// TestDNSRecordWrapper verifies the complete functionality
func TestDNSRecordWrapper(t *testing.T) {
	wrapper := cloudflare_record.NewDNSRecord()

	// Test CanHandle
	if !wrapper.CanHandle("cloudflare_record") {
		t.Error("Should handle cloudflare_record")
	}
	if !wrapper.CanHandle("cloudflare_dns_record") {
		t.Error("Should handle cloudflare_dns_record")
	}
	if wrapper.CanHandle("cloudflare_other") {
		t.Error("Should not handle other resource types")
	}

	// Test GetResourceType
	if wrapper.GetResourceType() != "cloudflare_record" {
		t.Error("Unexpected resource type")
	}
}

// TestAllRecordTypes tests various DNS record types
func TestAllRecordTypes(t *testing.T) {
	recordTypes := []struct {
		name       string
		recordType string
		hasData    bool
	}{
		{"A record", "A", false},
		{"AAAA record", "AAAA", false},
		{"CNAME record", "CNAME", false},
		{"TXT record", "TXT", false},
		{"MX record", "MX", false},
		{"NS record", "NS", false},
		{"CAA record", "CAA", true},
		{"SRV record", "SRV", true},
		{"URI record", "URI", true},
	}

	for _, rt := range recordTypes {
		t.Run(rt.name, func(t *testing.T) {
			var input string
			if rt.hasData {
				input = `resource "cloudflare_record" "test" {
  zone_id = "abc123"
  name    = "example"
  type    = "` + rt.recordType + `"
  data {
    flags = "0"
    tag = "test"
  }
}`
			} else {
				input = `resource "cloudflare_record" "test" {
  zone_id = "abc123"
  name    = "example"
  type    = "` + rt.recordType + `"
  value   = "test.example.com"
}`
			}

			reg := registry.NewStrategyRegistry()
			reg.Register(cloudflare_record.NewDNSRecord())

			pipeline := poc.NewPipelineBuilder().
				WithPreprocessing(reg).
				WithParsing().
				WithResourceTransformation(reg).
				WithFormatting().
				Build()

			result, err := pipeline.Transform([]byte(input), "test.tf")
			if err != nil {
				t.Fatalf("Transformation failed for %s: %v", rt.name, err)
			}

			resultStr := string(result)

			// All should be transformed to cloudflare_dns_record
			if !strings.Contains(resultStr, "cloudflare_dns_record") {
				t.Errorf("%s: Resource type not transformed", rt.name)
			}

			// All should have TTL
			if !strings.Contains(resultStr, "ttl") {
				t.Errorf("%s: TTL not added", rt.name)
			}

			// Simple types should have content, not value
			if !rt.hasData {
				if !strings.Contains(resultStr, "content") {
					t.Errorf("%s: Value not renamed to content", rt.name)
				}
				if strings.Contains(resultStr, "value") && !strings.Contains(resultStr, "# value") {
					t.Errorf("%s: Value attribute should be removed", rt.name)
				}
			}
		})
	}
}

// TestPortedDNSRecordCAATransformation ports the CAA transformation tests from cmd/migrate
func TestPortedDNSRecordCAATransformation(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  []string
		notExpect []string
	}{
		{
			name: "CAA record with numeric flags in data block - content renamed to value",
			input: `
resource "cloudflare_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data {
    flags   = 0
    tag     = "issue"
    content = "letsencrypt.org"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"zone_id = \"0da42c8d2132a9ddaf714f9e7c920711\"",
				"name    = \"test.example.com\"",
				"type    = \"CAA\"",
				"ttl     = 3600",
				"data =",
				"flags = 0",
				"tag   = \"issue\"",
				"value = \"letsencrypt.org\"",
			},
			notExpect: []string{
				"content = \"letsencrypt.org\"", // content should be renamed to value
				"data {",                        // data should be attribute, not block
			},
		},
		{
			name: "CAA record with numeric flags in data attribute map - content renamed to value",
			input: `
resource "cloudflare_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600
  data    {
    flags   = 0
    tag     = "issue"
    content = "letsencrypt.org"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"zone_id = \"0da42c8d2132a9ddaf714f9e7c920711\"",
				"name    = \"test.example.com\"",
				"type    = \"CAA\"",
				"ttl     = 3600",
				"data =",
				"flags = 0",
				"tag   = \"issue\"",
				"value = \"letsencrypt.org\"",
			},
			notExpect: []string{
				"content = \"letsencrypt.org\"",
			},
		},
		{
			name: "CAA record with flags already as string - content still renamed to value",
			input: `
resource "cloudflare_record" "caa_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data {
    flags   = "0"
    tag     = "issue"
    content = "letsencrypt.org"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"zone_id = \"0da42c8d2132a9ddaf714f9e7c920711\"",
				"name    = \"test.example.com\"",
				"type    = \"CAA\"",
				"ttl     = 3600",
				"data =",
				"flags = \"0\"",
				"tag   = \"issue\"",
				"value = \"letsencrypt.org\"",
			},
			notExpect: []string{
				"content = \"letsencrypt.org\"",
			},
		},
		{
			name: "Non-CAA record should not be modified",
			input: `
resource "cloudflare_record" "a_test" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "A"
  ttl     = 3600
  content = "192.168.1.1"
}`,
			expected: []string{
				"cloudflare_dns_record",
				"zone_id = \"0da42c8d2132a9ddaf714f9e7c920711\"",
				"name    = \"test.example.com\"",
				"type    = \"A\"",
				"ttl     = 3600",
				"content = \"192.168.1.1\"",
			},
			notExpect: []string{
				"value", // A records use content, not value
			},
		},
		{
			name: "cloudflare_record (legacy) with CAA type - content renamed to value",
			input: `
resource "cloudflare_record" "caa_legacy" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "CAA"
  ttl     = 3600

  data {
    flags   = 128
    tag     = "issuewild"
    content = "pki.goog"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"zone_id = \"0da42c8d2132a9ddaf714f9e7c920711\"",
				"name    = \"test.example.com\"",
				"type    = \"CAA\"",
				"ttl     = 3600",
				"data =",
				"flags = 128",
				"tag   = \"issuewild\"",
				"value = \"pki.goog\"",
			},
			notExpect: []string{
				"content = \"pki.goog\"",
			},
		},
		{
			name: "DNS record without TTL - should add TTL with default value",
			input: `
resource "cloudflare_record" "mx_test" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  name     = "test.example.com"
  type     = "MX"
  content  = "mx.sendgrid.net"
  priority = 10
}`,
			expected: []string{
				"cloudflare_dns_record",
				"zone_id  = \"0da42c8d2132a9ddaf714f9e7c920711\"",
				"name     = \"test.example.com\"",
				"type     = \"MX\"",
				"content  = \"mx.sendgrid.net\"",
				"priority = 10",
				"ttl      = 1",
			},
		},
		{
			name: "DNS record with existing TTL - should keep existing value",
			input: `
resource "cloudflare_record" "a_test_ttl" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test.example.com"
  type    = "A"
  ttl     = 3600
  content = "192.168.1.1"
}`,
			expected: []string{
				"cloudflare_dns_record",
				"zone_id = \"0da42c8d2132a9ddaf714f9e7c920711\"",
				"name    = \"test.example.com\"",
				"type    = \"A\"",
				"ttl     = 3600",
				"content = \"192.168.1.1\"",
			},
			notExpect: []string{
				"value", // A records keep content
			},
		},
		{
			name: "Multiple CAA records in same file - content renamed to value and TTL added",
			input: `
resource "cloudflare_record" "caa_test1" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test1.example.com"
  type    = "CAA"
  data {
    flags   = 0
    tag     = "issue"
    content = "letsencrypt.org"
  }
}

resource "cloudflare_record" "caa_test2" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "test2.example.com"
  type    = "CAA"
  data {
    flags   = 128
    tag     = "issuewild"
    content = "pki.goog"
  }
}`,
			expected: []string{
				"cloudflare_dns_record",
				"caa_test1",
				"test1.example.com",
				"ttl     = 1",
				"flags = 0",
				"value = \"letsencrypt.org\"",
				"caa_test2",
				"test2.example.com",
				"flags = 128",
				"value = \"pki.goog\"",
			},
			notExpect: []string{
				"content = \"letsencrypt.org\"",
				"content = \"pki.goog\"",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create pipeline with DNS record wrapper
			reg := registry.NewStrategyRegistry()
			reg.Register(cloudflare_record.NewDNSRecord())

			pipeline := poc.NewPipelineBuilder().
				WithPreprocessing(reg).
				WithParsing().
				WithResourceTransformation(reg).
				WithFormatting().
				Build()

			// TransformConfig
			result, err := pipeline.Transform([]byte(tt.input), "test.tf")
			if err != nil {
				t.Fatalf("Transformation failed: %v", err)
			}

			resultStr := string(result)

			// Check expected patterns
			for _, expected := range tt.expected {
				if !strings.Contains(resultStr, expected) {
					t.Errorf("Result should contain '%s', but it doesn't:\n%s", expected, resultStr)
				}
			}

			// Check patterns that should NOT be present
			for _, notExpected := range tt.notExpect {
				if strings.Contains(resultStr, notExpected) {
					t.Errorf("Result should NOT contain '%s', but it does:\n%s", notExpected, resultStr)
				}
			}
		})
	}
}

// TestPortedDNSRecordStateTransformation ports the state transformation tests from cmd/migrate
func TestPortedDNSRecordStateTransformation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
		check    func(t *testing.T, result string)
	}{
		{
			name: "CAA record v4 format with array data and numeric flags",
			input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"content": "0 issue letsencrypt.org",
							"data": [{
								"flags": 0,
								"tag": "issue",
								"content": "letsencrypt.org"
							}]
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check data is now an object
				data := gjson.Get(result, "resources.0.instances.0.attributes.data")
				if !data.IsObject() {
					t.Error("Data should be an object, not an array")
				}

				// Check flags is wrapped
				flags := gjson.Get(result, "resources.0.instances.0.attributes.data.flags")
				if !flags.IsObject() {
					t.Error("Flags should be a wrapped object")
				}

				flagsValue := gjson.Get(result, "resources.0.instances.0.attributes.data.flags.value")
				if flagsValue.Float() != 0 {
					t.Errorf("Expected flags value to be 0, got %v", flagsValue.Float())
				}

				// Check content renamed to value in data
				value := gjson.Get(result, "resources.0.instances.0.attributes.data.value")
				if value.String() != "letsencrypt.org" {
					t.Errorf("Expected value to be 'letsencrypt.org', got '%s'", value.String())
				}

				// Check TTL added
				ttl := gjson.Get(result, "resources.0.instances.0.attributes.ttl")
				if !ttl.Exists() || ttl.Float() != 1 {
					t.Error("TTL should be added with value 1")
				}

				// Check computed fields added
				createdOn := gjson.Get(result, "resources.0.instances.0.attributes.created_on")
				if !createdOn.Exists() {
					t.Error("created_on should be added")
				}
			},
		},
		{
			name: "CAA record with numeric flags 128",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"content": "128 issuewild pki.goog",
							"data": [{
								"flags": 128,
								"tag": "issuewild",
								"content": "pki.goog"
							}]
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check flags value is 128
				flagsValue := gjson.Get(result, "resources.0.instances.0.attributes.data.flags.value")
				if flagsValue.Float() != 128 {
					t.Errorf("Expected flags value to be 128, got %v", flagsValue.Float())
				}

				// Check tag is issuewild
				tag := gjson.Get(result, "resources.0.instances.0.attributes.data.tag")
				if tag.String() != "issuewild" {
					t.Errorf("Expected tag to be 'issuewild', got '%s'", tag.String())
				}

				// Check value is pki.goog
				value := gjson.Get(result, "resources.0.instances.0.attributes.data.value")
				if value.String() != "pki.goog" {
					t.Errorf("Expected value to be 'pki.goog', got '%s'", value.String())
				}
			},
		},
		{
			name: "CAA record already migrated to object format",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"data": {
								"flags": {
									"value": 0,
									"type": "number"
								},
								"tag": "issue",
								"value": "letsencrypt.org"
							}
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check data is still an object
				data := gjson.Get(result, "resources.0.instances.0.attributes.data")
				if !data.IsObject() {
					t.Error("Data should remain an object")
				}

				// Check flags structure is preserved
				flagsType := gjson.Get(result, "resources.0.instances.0.attributes.data.flags.type")
				if flagsType.String() != "number" {
					t.Error("Flags type should be preserved")
				}

				// Check value field exists
				value := gjson.Get(result, "resources.0.instances.0.attributes.data.value")
				if value.String() != "letsencrypt.org" {
					t.Errorf("Value should be 'letsencrypt.org', got '%s'", value.String())
				}

				// Check TTL added
				ttl := gjson.Get(result, "resources.0.instances.0.attributes.ttl")
				if !ttl.Exists() || ttl.Float() != 1 {
					t.Error("TTL should be added with value 1")
				}
			},
		},
		{
			name: "Simple A record should set data field to null",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "a_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "A",
							"content": "192.168.1.1",
							"data": []
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check data field is removed/null for A record
				data := gjson.Get(result, "resources.0.instances.0.attributes.data")
				if data.Exists() {
					t.Error("Data field should not exist for A record")
				}

				// Check content field exists
				content := gjson.Get(result, "resources.0.instances.0.attributes.content")
				if content.String() != "192.168.1.1" {
					t.Errorf("Content should be '192.168.1.1', got '%s'", content.String())
				}

				// Check TTL added
				ttl := gjson.Get(result, "resources.0.instances.0.attributes.ttl")
				if !ttl.Exists() || ttl.Float() != 1 {
					t.Error("TTL should be added with value 1")
				}
			},
		},
		{
			name: "cloudflare_record (legacy) renamed to cloudflare_dns_record",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_record",
					"name": "caa_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "CAA",
							"content": "0 issue letsencrypt.org",
							"data": [{
								"flags": 0,
								"tag": "issue",
								"content": "letsencrypt.org"
							}]
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check resource type renamed
				resourceType := gjson.Get(result, "resources.0.type")
				if resourceType.String() != "cloudflare_dns_record" {
					t.Errorf("Resource type should be 'cloudflare_dns_record', got '%s'", resourceType.String())
				}

				// Check data is object
				data := gjson.Get(result, "resources.0.instances.0.attributes.data")
				if !data.IsObject() {
					t.Error("Data should be an object")
				}

				// Check content renamed to value in data
				value := gjson.Get(result, "resources.0.instances.0.attributes.data.value")
				if value.String() != "letsencrypt.org" {
					t.Errorf("Value should be 'letsencrypt.org', got '%s'", value.String())
				}
			},
		},
		{
			name: "SRV record with array data migration",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "srv_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "_sip._tcp.example.com",
							"type": "SRV",
							"data": [{
								"priority": 10,
								"weight": 60,
								"port": 5060,
								"target": "sipserver.example.com"
							}]
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check data is now an object
				data := gjson.Get(result, "resources.0.instances.0.attributes.data")
				if !data.IsObject() {
					t.Error("Data should be an object, not an array")
				}

				// Check priority is also at root level
				rootPriority := gjson.Get(result, "resources.0.instances.0.attributes.priority")
				if rootPriority.Float() != 10 {
					t.Errorf("Priority should be 10 at root level, got %v", rootPriority.Float())
				}

				// Check data fields
				dataPriority := gjson.Get(result, "resources.0.instances.0.attributes.data.priority")
				if dataPriority.Float() != 10 {
					t.Errorf("Data priority should be 10, got %v", dataPriority.Float())
				}

				weight := gjson.Get(result, "resources.0.instances.0.attributes.data.weight")
				if weight.Float() != 60 {
					t.Errorf("Weight should be 60, got %v", weight.Float())
				}

				port := gjson.Get(result, "resources.0.instances.0.attributes.data.port")
				if port.Float() != 5060 {
					t.Errorf("Port should be 5060, got %v", port.Float())
				}

				target := gjson.Get(result, "resources.0.instances.0.attributes.data.target")
				if target.String() != "sipserver.example.com" {
					t.Errorf("Target should be 'sipserver.example.com', got '%s'", target.String())
				}
			},
		},
		{
			name: "Record with value field renamed to content",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_record",
					"name": "a_test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"name": "test.example.com",
							"type": "A",
							"value": "192.168.1.1",
							"hostname": "test.example.com",
							"allow_overwrite": true
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check resource type renamed
				resourceType := gjson.Get(result, "resources.0.type")
				if resourceType.String() != "cloudflare_dns_record" {
					t.Errorf("Resource type should be 'cloudflare_dns_record', got '%s'", resourceType.String())
				}

				// Check value renamed to content
				content := gjson.Get(result, "resources.0.instances.0.attributes.content")
				if content.String() != "192.168.1.1" {
					t.Errorf("Content should be '192.168.1.1', got '%s'", content.String())
				}

				// Check value field removed
				value := gjson.Get(result, "resources.0.instances.0.attributes.value")
				if value.Exists() {
					t.Error("Value field should be removed")
				}

				// Check deprecated fields removed
				hostname := gjson.Get(result, "resources.0.instances.0.attributes.hostname")
				if hostname.Exists() {
					t.Error("Hostname field should be removed")
				}

				allowOverwrite := gjson.Get(result, "resources.0.instances.0.attributes.allow_overwrite")
				if allowOverwrite.Exists() {
					t.Error("allow_overwrite field should be removed")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapper := cloudflare_record.NewDNSRecord()
			json := gjson.Parse(tt.input)

			// Find the resource path
			resourcePath := "resources.0"
			result, err := wrapper.TransformState(json, resourcePath)
			if err != nil {
				t.Fatalf("State transformation failed: %v", err)
			}

			// Run custom checks
			if tt.check != nil {
				tt.check(t, result)
			}
		})
	}
}

// TestPortedDNSRecordStateTransformationWithComputedFields ports tests for computed fields
func TestPortedDNSRecordStateTransformationWithComputedFields(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, result string)
	}{
		{
			name: "TXT record with all computed fields from v4",
			input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [{
					"type": "cloudflare_record",
					"name": "txt_spf",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "7c2d6320347f97de16dd2569e1fcd6b5",
							"name": "static.example.com.terraform.cfapi.net",
							"type": "TXT",
							"value": "v=spf1 include:_spf.mx.cloudflare.net ~all",
							"ttl": 1,
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"created_on": "2025-08-26T05:22:13.523335Z",
							"modified_on": "2025-08-26T05:22:13.523335Z",
							"proxied": false,
							"meta": "{}",
							"tags": ["tf-applied"],
							"tags_modified_on": "2025-08-26T05:22:13Z",
							"settings": {
								"flatten_cname": null,
								"ipv4_only": null,
								"ipv6_only": null
							}
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check resource type renamed
				resourceType := gjson.Get(result, "resources.0.type")
				if resourceType.String() != "cloudflare_dns_record" {
					t.Errorf("Resource type should be 'cloudflare_dns_record', got '%s'", resourceType.String())
				}

				// Check value renamed to content
				content := gjson.Get(result, "resources.0.instances.0.attributes.content")
				if content.String() != "v=spf1 include:_spf.mx.cloudflare.net ~all" {
					t.Errorf("Content should be 'v=spf1 include:_spf.mx.cloudflare.net ~all', got '%s'", content.String())
				}

				// Check deprecated fields removed
				meta := gjson.Get(result, "resources.0.instances.0.attributes.meta")
				if meta.Exists() {
					t.Error("meta field should be removed")
				}

				settings := gjson.Get(result, "resources.0.instances.0.attributes.settings")
				if settings.Exists() {
					t.Error("settings field should be removed")
				}

				// Check tags preserved
				tags := gjson.Get(result, "resources.0.instances.0.attributes.tags")
				if !tags.IsArray() {
					t.Error("tags should be preserved as array")
				}
			},
		},
		{
			name: "A record with missing computed fields - should add them",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_record",
					"name": "a_test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "test-id",
							"name": "test.example.com",
							"type": "A",
							"value": "192.168.1.1",
							"ttl": 3600,
							"zone_id": "test-zone"
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check created_on added
				createdOn := gjson.Get(result, "resources.0.instances.0.attributes.created_on")
				if !createdOn.Exists() {
					t.Error("created_on should be added")
				}

				// Check modified_on added
				modifiedOn := gjson.Get(result, "resources.0.instances.0.attributes.modified_on")
				if !modifiedOn.Exists() {
					t.Error("modified_on should be added")
				}

				// Check value renamed to content
				content := gjson.Get(result, "resources.0.instances.0.attributes.content")
				if content.String() != "192.168.1.1" {
					t.Errorf("Content should be '192.168.1.1', got '%s'", content.String())
				}
			},
		},
		{
			name: "CAA record with all computed fields",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_dns_record",
					"name": "caa_google",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"id": "47480c33c49b0240b17dc9168d4442dd",
							"name": "static.example.com.terraform.cfapi.net",
							"type": "CAA",
							"content": "0 issue \"pki.goog\"",
							"ttl": 1,
							"zone_id": "0da42c8d2132a9ddaf714f9e7c920711",
							"data": [{
								"flags": 0,
								"tag": "issue",
								"value": "pki.goog"
							}],
							"created_on": "2025-08-26T05:21:58Z",
							"modified_on": "2025-08-26T05:21:58Z",
							"proxied": false,
							"meta": "{}",
							"comment": null,
							"comment_modified_on": null,
							"tags": ["tf-applied"],
							"tags_modified_on": "2025-08-26T05:21:58Z",
							"settings": {
								"flatten_cname": null,
								"ipv4_only": null,
								"ipv6_only": null
							}
						}
					}]
				}]
			}`,
			check: func(t *testing.T, result string) {
				// Check data transformed correctly
				data := gjson.Get(result, "resources.0.instances.0.attributes.data")
				if !data.IsObject() {
					t.Error("Data should be an object")
				}

				// Check flags wrapped
				flagsValue := gjson.Get(result, "resources.0.instances.0.attributes.data.flags.value")
				if flagsValue.Float() != 0 {
					t.Errorf("Flags value should be 0, got %v", flagsValue.Float())
				}

				// Check value field in data
				value := gjson.Get(result, "resources.0.instances.0.attributes.data.value")
				if value.String() != "pki.goog" {
					t.Errorf("Value should be 'pki.goog', got '%s'", value.String())
				}

				// Check deprecated fields removed
				meta := gjson.Get(result, "resources.0.instances.0.attributes.meta")
				if meta.Exists() {
					t.Error("meta field should be removed")
				}

				settings := gjson.Get(result, "resources.0.instances.0.attributes.settings")
				if settings.Exists() {
					t.Error("settings field should be removed")
				}

				// Check comment fields preserved as they're not deprecated
				comment := gjson.Get(result, "resources.0.instances.0.attributes.comment")
				if !comment.Exists() {
					t.Error("comment field should be preserved")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapper := cloudflare_record.NewDNSRecord()
			json := gjson.Parse(tt.input)

			// Find the resource path
			resourcePath := "resources.0"
			result, err := wrapper.TransformState(json, resourcePath)
			if err != nil {
				t.Fatalf("State transformation failed: %v", err)
			}

			// Run custom checks
			if tt.check != nil {
				tt.check(t, result)
			}
		})
	}
}

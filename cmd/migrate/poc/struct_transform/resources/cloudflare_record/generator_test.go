package cloudflare_record

import (
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/v5_models"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDNSRecord(t *testing.T) {
	tests := []struct {
		name           string
		v5Model        *v5_models.DNSRecordV5Model
		labels         []string
		expectedBlocks []string // Strings that should appear in the output
		notExpected    []string // Strings that should NOT appear
	}{
		{
			name: "generate basic A record",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID:  types.StringValue("zone123"),
				Name:    types.StringValue("example.com"),
				Type:    types.StringValue("A"),
				Content: types.StringValue("192.0.2.1"),
				TTL:     types.Float64Value(300),
				Proxied: types.BoolValue(true),
			},
			labels: []string{"cloudflare_record", "test"},
			expectedBlocks: []string{
				`resource "cloudflare_dns_record" "test"`,
				`zone_id = "zone123"`,
				`name    = "example.com"`,
				`type    = "A"`,
				`content = "192.0.2.1"`,
				`ttl     = 300`,
				`proxied = true`,
			},
			notExpected: []string{
				`value =`, // Should use 'content' not 'value'
			},
		},
		{
			name: "generate MX record with priority",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID:   types.StringValue("zone456"),
				Name:     types.StringValue("mail"),
				Type:     types.StringValue("MX"),
				Content:  types.StringValue("mail.example.com"),
				TTL:      types.Float64Value(3600),
				Priority: types.Float64Value(10),
			},
			labels: []string{"cloudflare_record", "mx"},
			expectedBlocks: []string{
				`resource "cloudflare_dns_record" "mx"`,
				`zone_id  = "zone456"`,
				`name     = "mail"`,
				`type     = "MX"`,
				`content  = "mail.example.com"`,
				`ttl      = 3600`,
				`priority = 10`,
			},
		},
		{
			name: "generate with tags",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID:  types.StringValue("zone789"),
				Name:    types.StringValue("tagged"),
				Type:    types.StringValue("A"),
				Content: types.StringValue("192.0.2.2"),
				TTL:     types.Float64Value(1),
				Tags: []types.String{
					types.StringValue("production"),
					types.StringValue("primary"),
				},
			},
			labels: []string{"cloudflare_record", "tagged"},
			expectedBlocks: []string{
				`tags    = ["production", "primary"]`,
			},
		},
		{
			name: "generate with comment",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID:  types.StringValue("zone_comment"),
				Name:    types.StringValue("commented"),
				Type:    types.StringValue("TXT"),
				Content: types.StringValue("test value"),
				TTL:     types.Float64Value(300),
				Comment: types.StringValue("This is a test record"),
			},
			labels: []string{"cloudflare_record", "commented"},
			expectedBlocks: []string{
				`comment = "This is a test record"`,
			},
		},
		{
			name: "skip null/unknown values",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID:   types.StringValue("zone_minimal"),
				Name:     types.StringValue("minimal"),
				Type:     types.StringValue("A"),
				Content:  types.StringValue("192.0.2.3"),
				TTL:      types.Float64Value(1),
				Proxied:  types.BoolNull(),    // Null value
				Priority: types.Float64Null(), // Null value
				Comment:  types.StringNull(),   // Null value
			},
			labels: []string{"cloudflare_record", "minimal"},
			expectedBlocks: []string{
				`zone_id = "zone_minimal"`,
				`name    = "minimal"`,
				`type    = "A"`,
				`content = "192.0.2.3"`,
				`ttl     = 1`,
			},
			notExpected: []string{
				`proxied =`,
				`priority =`,
				`comment =`,
			},
		},
	}

	// Use the DNS-specific generator
	generator := NewDNSRecordGenerator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate as interface{} to match ModelGenerator interface
			block := generator.Generate(tt.v5Model, tt.labels)

			// Convert to string
			file := hclwrite.NewEmptyFile()
			file.Body().AppendBlock(block)
			output := string(file.Bytes())

			// Check expected strings
			for _, expected := range tt.expectedBlocks {
				assert.Contains(t, output, expected, "Output should contain: %s", expected)
			}

			// Check strings that should NOT appear
			for _, notExpected := range tt.notExpected {
				assert.NotContains(t, output, notExpected, "Output should NOT contain: %s", notExpected)
			}

			// Verify it's valid HCL
			_, diags := hclwrite.ParseConfig(file.Bytes(), "test.tf", hcl.InitialPos)
			assert.False(t, diags.HasErrors(), "Generated HCL should be valid")
		})
	}
}

func TestGenerateDataBlock(t *testing.T) {
	tests := []struct {
		name           string
		v5Model        *v5_models.DNSRecordV5Model
		expectedBlocks []string
		notExpected    []string
	}{
		{
			name: "generate CAA with data block",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID:  types.StringValue("zone_caa"),
				Name:    types.StringValue("example.com"),
				Type:    types.StringValue("CAA"),
				TTL:     types.Float64Value(300),
				Data: &v5_models.DNSRecordDataV5Model{
					Flags: types.Float64Value(0),
					Tag:   types.StringValue("issue"),
					Value: types.StringValue("letsencrypt.org"),
				},
			},
			expectedBlocks: []string{
				`data {`,
				`flags = 0`,
				`tag   = "issue"`,
				`value = "letsencrypt.org"`, // v5 uses 'value'
			},
			notExpected: []string{
				`content =`, // CAA data uses 'value', not 'content'
			},
		},
		{
			name: "generate SRV with data block",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID: types.StringValue("zone_srv"),
				Name:   types.StringValue("_sip._tcp"),
				Type:   types.StringValue("SRV"),
				TTL:    types.Float64Value(300),
				Data: &v5_models.DNSRecordDataV5Model{
					Priority: types.Float64Value(10),
					Weight:   types.Float64Value(60),
					Port:     types.Float64Value(5060),
					Target:   types.StringValue("sip.example.com"),
					Service:  types.StringValue("_sip"),
				},
			},
			expectedBlocks: []string{
				`data {`,
				`priority = 10`,
				`weight   = 60`,
				`port     = 5060`,
				`target   = "sip.example.com"`,
				`service  = "_sip"`,
			},
			notExpected: []string{
				`proto =`, // v5 doesn't have proto
				`name =`,  // v5 doesn't have name in data
			},
		},
		{
			name: "generate LOC with data block",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID: types.StringValue("zone_loc"),
				Name:   types.StringValue("location"),
				Type:   types.StringValue("LOC"),
				TTL:    types.Float64Value(300),
				Data: &v5_models.DNSRecordDataV5Model{
					Altitude:      types.Float64Value(100),
					LatDegrees:    types.Float64Value(37),
					LatDirection:  types.StringValue("N"),
					LatMinutes:    types.Float64Value(46),
					LatSeconds:    types.Float64Value(46),
					LongDegrees:   types.Float64Value(122),
					LongDirection: types.StringValue("W"),
					LongMinutes:   types.Float64Value(23),
					LongSeconds:   types.Float64Value(35),
					PrecisionHorz: types.Float64Value(10),
					PrecisionVert: types.Float64Value(2),
					Size:          types.Float64Value(1),
				},
			},
			expectedBlocks: []string{
				`data {`,
				`altitude       = 100`,
				`lat_degrees    = 37`,
				`lat_direction  = "N"`,
				`lat_minutes    = 46`,
				`lat_seconds    = 46`,
				`long_degrees   = 122`,
				`long_direction = "W"`,
				`long_minutes   = 23`,
				`long_seconds   = 35`,
				`precision_horz = 10`,
				`precision_vert = 2`,
				`size           = 1`,
			},
		},
		{
			name: "generate URI with data block",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID: types.StringValue("zone_uri"),
				Name:   types.StringValue("_http._tcp"),
				Type:   types.StringValue("URI"),
				TTL:    types.Float64Value(300),
				Data: &v5_models.DNSRecordDataV5Model{
					Priority: types.Float64Value(10),
					Weight:   types.Float64Value(1),
					Target:   types.StringValue("http://example.com/path"),
				},
			},
			expectedBlocks: []string{
				`data {`,
				`priority = 10`,
				`weight   = 1`,
				`target   = "http://example.com/path"`,
			},
		},
		{
			name: "skip empty data block",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID:  types.StringValue("zone_no_data"),
				Name:    types.StringValue("no-data"),
				Type:    types.StringValue("A"),
				Content: types.StringValue("192.0.2.4"),
				TTL:     types.Float64Value(300),
				Data: &v5_models.DNSRecordDataV5Model{
					// All fields are null
				},
			},
			expectedBlocks: []string{
				`zone_id = "zone_no_data"`,
				`name    = "no-data"`,
			},
			notExpected: []string{
				`data {`, // Empty data block should not be generated
			},
		},
		{
			name: "data block with partial values",
			v5Model: &v5_models.DNSRecordV5Model{
				ZoneID: types.StringValue("zone_partial"),
				Name:   types.StringValue("partial"),
				Type:   types.StringValue("CAA"),
				TTL:    types.Float64Value(300),
				Data: &v5_models.DNSRecordDataV5Model{
					Flags: types.Float64Value(128),
					Tag:   types.StringValue("issuewild"),
					// Value is null
				},
			},
			expectedBlocks: []string{
				`data {`,
				`flags = 128`,
				`tag   = "issuewild"`,
			},
			notExpected: []string{
				`value =`, // Null value should not appear
			},
		},
	}

	generator := NewDNSRecordGenerator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			labels := []string{"cloudflare_record", "test"}
			block := generator.Generate(tt.v5Model, labels)

			file := hclwrite.NewEmptyFile()
			file.Body().AppendBlock(block)
			output := string(file.Bytes())

			// Check expected strings
			for _, expected := range tt.expectedBlocks {
				assert.Contains(t, output, expected, "Output should contain: %s", expected)
			}

			// Check strings that should NOT appear
			for _, notExpected := range tt.notExpected {
				assert.NotContains(t, output, notExpected, "Output should NOT contain: %s", notExpected)
			}

			// Verify valid HCL
			_, diags := hclwrite.ParseConfig(file.Bytes(), "test.tf", hcl.InitialPos)
			assert.False(t, diags.HasErrors(), "Generated HCL should be valid")
		})
	}
}

func TestResourceTypeTransformation(t *testing.T) {
	// Test that v4 cloudflare_record becomes v5 cloudflare_dns_record
	v5Model := &v5_models.DNSRecordV5Model{
		ZoneID:  types.StringValue("zone_rename"),
		Name:    types.StringValue("renamed"),
		Type:    types.StringValue("A"),
		Content: types.StringValue("192.0.2.5"),
		TTL:     types.Float64Value(300),
	}

	generator := NewDNSRecordGenerator()

	// Test with v4 labels
	v4Labels := []string{"cloudflare_record", "myrecord"}
	block := generator.Generate(v5Model, v4Labels)

	file := hclwrite.NewEmptyFile()
	file.Body().AppendBlock(block)
	output := string(file.Bytes())

	// Should generate cloudflare_dns_record, not cloudflare_record
	assert.Contains(t, output, `resource "cloudflare_dns_record" "myrecord"`)
	assert.NotContains(t, output, `resource "cloudflare_record"`)
}

func TestGeneratorOptions(t *testing.T) {
	v5Model := &v5_models.DNSRecordV5Model{
		ZoneID:  types.StringValue("zone_options"),
		Name:    types.StringValue("test"),
		Type:    types.StringValue("A"),
		Content: types.StringValue("192.0.2.6"),
		TTL:     types.Float64Value(300),
	}

	// Test with different options
	tests := []struct {
		name      string
		generator *DNSRecordGenerator
		validate  func(*testing.T, string)
	}{
		{
			name:      "default options",
			generator: NewDNSRecordGenerator(),
			validate: func(t *testing.T, output string) {
				// Should be formatted normally
				assert.Contains(t, output, `zone_id = "zone_options"`)
			},
		},
		// Additional option tests could be added here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			block := tt.generator.Generate(v5Model, []string{"cloudflare_record", "test"})
			file := hclwrite.NewEmptyFile()
			file.Body().AppendBlock(block)
			output := string(file.Bytes())
			tt.validate(t, output)
		})
	}
}

func TestAttributeOrdering(t *testing.T) {
	// Test that attributes appear in logical order
	v5Model := &v5_models.DNSRecordV5Model{
		Comment: types.StringValue("last attribute"),
		Type:    types.StringValue("A"),
		Name:    types.StringValue("ordered"),
		Content: types.StringValue("192.0.2.7"),
		ZoneID:  types.StringValue("zone_order"),
		TTL:     types.Float64Value(300),
		Proxied: types.BoolValue(false),
	}

	generator := NewDNSRecordGenerator()
	block := generator.Generate(v5Model, []string{"cloudflare_record", "test"})

	file := hclwrite.NewEmptyFile()
	file.Body().AppendBlock(block)
	output := string(file.Bytes())

	// Find positions of attributes in output
	zonePos := strings.Index(output, "zone_id")
	namePos := strings.Index(output, "name")
	typePos := strings.Index(output, "type")
	contentPos := strings.Index(output, "content")
	ttlPos := strings.Index(output, "ttl")
	proxiedPos := strings.Index(output, "proxied")
	commentPos := strings.Index(output, "comment")

	// Verify logical ordering: zone_id, name, type, content, ttl, optional fields
	assert.Less(t, zonePos, namePos, "zone_id should come before name")
	assert.Less(t, namePos, typePos, "name should come before type")
	assert.Less(t, typePos, contentPos, "type should come before content")
	assert.Less(t, contentPos, ttlPos, "content should come before ttl")
	assert.Less(t, ttlPos, proxiedPos, "ttl should come before proxied")
	assert.Less(t, proxiedPos, commentPos, "proxied should come before comment")
}
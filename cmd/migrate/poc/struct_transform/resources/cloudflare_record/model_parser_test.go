package cloudflare_record

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/v4_models"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
)

func TestParseDNSRecord(t *testing.T) {
	tests := []struct {
		name        string
		hclContent  string
		expectError bool
		validate    func(*testing.T, *v4_models.DNSRecordV4Model)
	}{
		{
			name: "parse basic A record",
			hclContent: `
resource "cloudflare_record" "test" {
  zone_id = "abc123"
  name    = "example.com"
  type    = "A"
  value   = "192.0.2.1"
  ttl     = 300
  proxied = true
}`,
			expectError: false,
			validate: func(t *testing.T, model *v4_models.DNSRecordV4Model) {
				assert.Equal(t, "abc123", model.ZoneID.ValueString())
				assert.Equal(t, "example.com", model.Name.ValueString())
				assert.Equal(t, "A", model.Type.ValueString())
				assert.Equal(t, "192.0.2.1", model.Value.ValueString())
				assert.Equal(t, float64(300), model.TTL.ValueFloat64())
				assert.True(t, model.Proxied.ValueBool())
			},
		},
		{
			name: "parse MX record with priority",
			hclContent: `
resource "cloudflare_record" "mx" {
  zone_id  = "zone456"
  name     = "mail"
  type     = "MX"
  value    = "mail.example.com"
  priority = 10
  ttl      = 3600
}`,
			expectError: false,
			validate: func(t *testing.T, model *v4_models.DNSRecordV4Model) {
				assert.Equal(t, "zone456", model.ZoneID.ValueString())
				assert.Equal(t, "mail", model.Name.ValueString())
				assert.Equal(t, "MX", model.Type.ValueString())
				assert.Equal(t, "mail.example.com", model.Value.ValueString())
				assert.Equal(t, float64(10), model.Priority.ValueFloat64())
				assert.Equal(t, float64(3600), model.TTL.ValueFloat64())
			},
		},
		{
			name: "parse CAA record with data block",
			hclContent: `
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
			expectError: false,
			validate: func(t *testing.T, model *v4_models.DNSRecordV4Model) {
				assert.Equal(t, "zone789", model.ZoneID.ValueString())
				assert.Equal(t, "example.com", model.Name.ValueString())
				assert.Equal(t, "CAA", model.Type.ValueString())
				assert.Len(t, model.Data, 1)

				data := model.Data[0]
				assert.Equal(t, "0", data.Flags.ValueString())
				assert.Equal(t, "issue", data.Tag.ValueString())
				assert.Equal(t, "letsencrypt.org", data.Content.ValueString())
			},
		},
		{
			name: "parse SRV record with data block",
			hclContent: `
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
			expectError: false,
			validate: func(t *testing.T, model *v4_models.DNSRecordV4Model) {
				assert.Equal(t, "zone_srv", model.ZoneID.ValueString())
				assert.Equal(t, "_sip._tcp", model.Name.ValueString())
				assert.Equal(t, "SRV", model.Type.ValueString())
				assert.Len(t, model.Data, 1)

				data := model.Data[0]
				assert.Equal(t, float64(10), data.Priority.ValueFloat64())
				assert.Equal(t, float64(60), data.Weight.ValueFloat64())
				assert.Equal(t, float64(5060), data.Port.ValueFloat64())
				assert.Equal(t, "sip.example.com", data.Target.ValueString())
				assert.Equal(t, "_sip", data.Service.ValueString())
				assert.Equal(t, "_tcp", data.Proto.ValueString())
				assert.Equal(t, "example.com", data.Name.ValueString())
			},
		},
		{
			name: "parse with variable references",
			hclContent: `
resource "cloudflare_record" "var_test" {
  zone_id = var.zone_id
  name    = local.hostname
  type    = "CNAME"
  value   = aws_instance.main.public_dns
  ttl     = 300
}`,
			expectError: false,
			validate: func(t *testing.T, model *v4_models.DNSRecordV4Model) {
				// Variable references are parsed as identifiers
				assert.NotNil(t, model.ZoneID)
				assert.NotNil(t, model.Name)
				assert.Equal(t, "CNAME", model.Type.ValueString())
				assert.NotNil(t, model.Value)
			},
		},
		{
			name: "parse with deprecated fields",
			hclContent: `
resource "cloudflare_record" "deprecated" {
  zone_id         = "zone_dep"
  hostname        = "old.example.com"
  type            = "A"
  value           = "192.0.2.10"
  allow_overwrite = true
}`,
			expectError: false,
			validate: func(t *testing.T, model *v4_models.DNSRecordV4Model) {
				assert.Equal(t, "zone_dep", model.ZoneID.ValueString())
				assert.Equal(t, "old.example.com", model.Hostname.ValueString())
				assert.Equal(t, "A", model.Type.ValueString())
				assert.Equal(t, "192.0.2.10", model.Value.ValueString())
				assert.True(t, model.AllowOverwrite.ValueBool())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse HCL
			file, diags := hclwrite.ParseConfig([]byte(tt.hclContent), "test.tf", hcl.InitialPos)
			assert.False(t, diags.HasErrors())

			blocks := file.Body().Blocks()
			assert.Len(t, blocks, 1)

			// Parse to model
			parser := NewHCLParser()
			model, err := parser.ParseDNSRecord(blocks[0])

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tt.validate(t, model)
			}
		})
	}
}

func TestParseDataBlock(t *testing.T) {
	tests := []struct {
		name        string
		hclContent  string
		expectError bool
		validate    func(*testing.T, *v4_models.DNSRecordDataV4Model)
	}{
		{
			name: "parse LOC data fields",
			hclContent: `
data {
  altitude       = 100
  lat_degrees    = 37
  lat_direction  = "N"
  lat_minutes    = 46
  lat_seconds    = 46
  long_degrees   = 122
  long_direction = "W"
  long_minutes   = 23
  long_seconds   = 35
  precision_horz = 10
  precision_vert = 2
  size           = 1
}`,
			expectError: false,
			validate: func(t *testing.T, data *v4_models.DNSRecordDataV4Model) {
				assert.Equal(t, float64(100), data.Altitude.ValueFloat64())
				assert.Equal(t, float64(37), data.LatDegrees.ValueFloat64())
				assert.Equal(t, "N", data.LatDirection.ValueString())
				assert.Equal(t, float64(46), data.LatMinutes.ValueFloat64())
				assert.Equal(t, float64(46), data.LatSeconds.ValueFloat64())
				assert.Equal(t, float64(122), data.LongDegrees.ValueFloat64())
				assert.Equal(t, "W", data.LongDirection.ValueString())
				assert.Equal(t, float64(23), data.LongMinutes.ValueFloat64())
				assert.Equal(t, float64(35), data.LongSeconds.ValueFloat64())
				assert.Equal(t, float64(10), data.PrecisionHorz.ValueFloat64())
				assert.Equal(t, float64(2), data.PrecisionVert.ValueFloat64())
				assert.Equal(t, float64(1), data.Size.ValueFloat64())
			},
		},
		{
			name: "parse with unknown attributes in non-strict mode",
			hclContent: `
data {
  flags         = "0"
  tag           = "issue"
  content       = "ca.example.com"
  unknown_field = "ignored"
}`,
			expectError: false,
			validate: func(t *testing.T, data *v4_models.DNSRecordDataV4Model) {
				assert.Equal(t, "0", data.Flags.ValueString())
				assert.Equal(t, "issue", data.Tag.ValueString())
				assert.Equal(t, "ca.example.com", data.Content.ValueString())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse HCL
			file, diags := hclwrite.ParseConfig([]byte(tt.hclContent), "test.tf", hcl.InitialPos)
			assert.False(t, diags.HasErrors())

			blocks := file.Body().Blocks()
			assert.Len(t, blocks, 1)

			// Parse data block
			parser := NewHCLParser()
			model, err := parser.parseDataBlock(blocks[0])

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tt.validate(t, model)
			}
		})
	}
}

func TestStrictMode(t *testing.T) {
	hclContent := `
resource "cloudflare_record" "test" {
  zone_id       = "zone123"
  name          = "test"
  type          = "A"
  value         = "192.0.2.1"
  unknown_field = "should_fail_in_strict"
}`

	file, diags := hclwrite.ParseConfig([]byte(hclContent), "test.tf", hcl.InitialPos)
	assert.False(t, diags.HasErrors())

	blocks := file.Body().Blocks()
	assert.Len(t, blocks, 1)

	// Test non-strict mode (default)
	parser := NewHCLParser()
	model, err := parser.ParseDNSRecord(blocks[0])
	assert.NoError(t, err)
	assert.NotNil(t, model)

	// Test strict mode
	strictParser := &HCLParser{StrictMode: true}
	_, err = strictParser.ParseDNSRecord(blocks[0])
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown attribute")
}
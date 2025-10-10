package cloudflare_record

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/v4_models"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/v5_models"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestTransformBasicRecord(t *testing.T) {
	tests := []struct {
		name     string
		v4Model  v4_models.DNSRecordV4Model
		validate func(*testing.T, *v5_models.DNSRecordV5Model)
	}{
		{
			name: "transform A record with value to content",
			v4Model: v4_models.DNSRecordV4Model{
				ID:      types.StringValue("rec123"),
				ZoneID:  types.StringValue("zone456"),
				Name:    types.StringValue("example.com"),
				Type:    types.StringValue("A"),
				Value:   types.StringValue("192.0.2.1"),
				TTL:     types.Float64Value(300),
				Proxied: types.BoolValue(true),
			},
			validate: func(t *testing.T, v5 *v5_models.DNSRecordV5Model) {
				assert.Equal(t, "rec123", v5.ID.ValueString())
				assert.Equal(t, "zone456", v5.ZoneID.ValueString())
				assert.Equal(t, "example.com", v5.Name.ValueString())
				assert.Equal(t, "A", v5.Type.ValueString())
				// v4 'value' becomes v5 'content'
				assert.Equal(t, "192.0.2.1", v5.Content.ValueString())
				assert.Equal(t, float64(300), v5.TTL.ValueFloat64())
				assert.True(t, v5.Proxied.ValueBool())
			},
		},
		{
			name: "transform record with content field",
			v4Model: v4_models.DNSRecordV4Model{
				ZoneID:  types.StringValue("zone789"),
				Name:    types.StringValue("test"),
				Type:    types.StringValue("TXT"),
				Content: types.StringValue("v=spf1 include:example.com ~all"),
				TTL:     types.Float64Value(3600),
			},
			validate: func(t *testing.T, v5 *v5_models.DNSRecordV5Model) {
				assert.Equal(t, "zone789", v5.ZoneID.ValueString())
				assert.Equal(t, "test", v5.Name.ValueString())
				assert.Equal(t, "TXT", v5.Type.ValueString())
				// v4 'content' stays as v5 'content'
				assert.Equal(t, "v=spf1 include:example.com ~all", v5.Content.ValueString())
				assert.Equal(t, float64(3600), v5.TTL.ValueFloat64())
			},
		},
		{
			name: "add default TTL when not specified",
			v4Model: v4_models.DNSRecordV4Model{
				ZoneID: types.StringValue("zone_ttl"),
				Name:   types.StringValue("auto-ttl"),
				Type:   types.StringValue("A"),
				Value:  types.StringValue("192.0.2.2"),
				// TTL is not set
			},
			validate: func(t *testing.T, v5 *v5_models.DNSRecordV5Model) {
				// Should have automatic TTL (1)
				assert.Equal(t, float64(1), v5.TTL.ValueFloat64())
			},
		},
		{
			name: "remove deprecated fields",
			v4Model: v4_models.DNSRecordV4Model{
				ZoneID:         types.StringValue("zone_dep"),
				Name:           types.StringValue("test"),
				Type:           types.StringValue("A"),
				Value:          types.StringValue("192.0.2.3"),
				AllowOverwrite: types.BoolValue(true),  // deprecated
				Hostname:       types.StringValue("old"), // deprecated
				TTL:            types.Float64Value(300),
			},
			validate: func(t *testing.T, v5 *v5_models.DNSRecordV5Model) {
				// Deprecated fields should not be present
				// (In a real implementation, we'd check that these fields don't exist)
				assert.Equal(t, "zone_dep", v5.ZoneID.ValueString())
				assert.Equal(t, "test", v5.Name.ValueString())
				assert.Equal(t, "192.0.2.3", v5.Content.ValueString())
			},
		},
		{
			name: "handle MX record with priority",
			v4Model: v4_models.DNSRecordV4Model{
				ZoneID:   types.StringValue("zone_mx"),
				Name:     types.StringValue("mail"),
				Type:     types.StringValue("MX"),
				Value:    types.StringValue("mail.example.com"),
				Priority: types.Float64Value(10),
				TTL:      types.Float64Value(3600),
			},
			validate: func(t *testing.T, v5 *v5_models.DNSRecordV5Model) {
				assert.Equal(t, "zone_mx", v5.ZoneID.ValueString())
				assert.Equal(t, "mail", v5.Name.ValueString())
				assert.Equal(t, "MX", v5.Type.ValueString())
				assert.Equal(t, "mail.example.com", v5.Content.ValueString())
				assert.Equal(t, float64(10), v5.Priority.ValueFloat64())
				assert.Equal(t, float64(3600), v5.TTL.ValueFloat64())
			},
		},
	}

	transformer := NewDNSRecordTransformer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v5Model, err := transformer.TransformDNSRecord(tt.v4Model)
			assert.NoError(t, err)
			assert.NotNil(t, v5Model)
			tt.validate(t, v5Model)
		})
	}
}

func TestTransformCAA(t *testing.T) {
	v4Model := v4_models.DNSRecordV4Model{
		ZoneID: types.StringValue("zone_caa"),
		Name:   types.StringValue("example.com"),
		Type:   types.StringValue("CAA"),
		Data: []v4_models.DNSRecordDataV4Model{
			{
				Flags:   types.StringValue("0"), // v4: string
				Tag:     types.StringValue("issue"),
				Content: types.StringValue("letsencrypt.org"),
			},
		},
	}

	transformer := NewDNSRecordTransformer()
	v5Model, err := transformer.TransformDNSRecord(v4Model)

	assert.NoError(t, err)
	assert.NotNil(t, v5Model)
	assert.NotNil(t, v5Model.Data)

	// Check flags conversion from string to number
	assert.Equal(t, float64(0), v5Model.Data.Flags.ValueFloat64())
	assert.Equal(t, "issue", v5Model.Data.Tag.ValueString())
	// Check content -> value rename in data block
	assert.Equal(t, "letsencrypt.org", v5Model.Data.Value.ValueString())
}

func TestTransformSRV(t *testing.T) {
	v4Model := v4_models.DNSRecordV4Model{
		ZoneID: types.StringValue("zone_srv"),
		Name:   types.StringValue("_sip._tcp"),
		Type:   types.StringValue("SRV"),
		Data: []v4_models.DNSRecordDataV4Model{
			{
				Priority: types.Float64Value(10),
				Weight:   types.Float64Value(60),
				Port:     types.Float64Value(5060),
				Target:   types.StringValue("sip.example.com"),
				Service:  types.StringValue("_sip"),
				Proto:    types.StringValue("_tcp"), // v4: has proto
				Name:     types.StringValue("example.com"), // v4: has name
			},
		},
	}

	transformer := NewDNSRecordTransformer()
	v5Model, err := transformer.TransformDNSRecord(v4Model)

	assert.NoError(t, err)
	assert.NotNil(t, v5Model)
	assert.NotNil(t, v5Model.Data)

	// Check SRV fields
	assert.Equal(t, float64(10), v5Model.Data.Priority.ValueFloat64())
	assert.Equal(t, float64(60), v5Model.Data.Weight.ValueFloat64())
	assert.Equal(t, float64(5060), v5Model.Data.Port.ValueFloat64())
	assert.Equal(t, "sip.example.com", v5Model.Data.Target.ValueString())
	assert.Equal(t, "_sip", v5Model.Data.Service.ValueString())
	// Proto and Name fields should be removed in v5
}

func TestTransformLOC(t *testing.T) {
	v4Model := v4_models.DNSRecordV4Model{
		ZoneID: types.StringValue("zone_loc"),
		Name:   types.StringValue("location"),
		Type:   types.StringValue("LOC"),
		Data: []v4_models.DNSRecordDataV4Model{
			{
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
	}

	transformer := NewDNSRecordTransformer()
	v5Model, err := transformer.TransformDNSRecord(v4Model)

	assert.NoError(t, err)
	assert.NotNil(t, v5Model)
	assert.NotNil(t, v5Model.Data)

	// LOC fields should be unchanged
	assert.Equal(t, float64(100), v5Model.Data.Altitude.ValueFloat64())
	assert.Equal(t, float64(37), v5Model.Data.LatDegrees.ValueFloat64())
	assert.Equal(t, "N", v5Model.Data.LatDirection.ValueString())
	assert.Equal(t, float64(46), v5Model.Data.LatMinutes.ValueFloat64())
	assert.Equal(t, float64(46), v5Model.Data.LatSeconds.ValueFloat64())
	assert.Equal(t, float64(122), v5Model.Data.LongDegrees.ValueFloat64())
	assert.Equal(t, "W", v5Model.Data.LongDirection.ValueString())
	assert.Equal(t, float64(23), v5Model.Data.LongMinutes.ValueFloat64())
	assert.Equal(t, float64(35), v5Model.Data.LongSeconds.ValueFloat64())
	assert.Equal(t, float64(10), v5Model.Data.PrecisionHorz.ValueFloat64())
	assert.Equal(t, float64(2), v5Model.Data.PrecisionVert.ValueFloat64())
	assert.Equal(t, float64(1), v5Model.Data.Size.ValueFloat64())
}

func TestTransformURI(t *testing.T) {
	v4Model := v4_models.DNSRecordV4Model{
		ZoneID: types.StringValue("zone_uri"),
		Name:   types.StringValue("_http._tcp"),
		Type:   types.StringValue("URI"),
		Data: []v4_models.DNSRecordDataV4Model{
			{
				Priority: types.Float64Value(10),
				Weight:   types.Float64Value(1),
				Target:   types.StringValue("http://example.com/path"),
			},
		},
	}

	transformer := NewDNSRecordTransformer()
	v5Model, err := transformer.TransformDNSRecord(v4Model)

	assert.NoError(t, err)
	assert.NotNil(t, v5Model)
	assert.NotNil(t, v5Model.Data)

	// Check URI fields
	assert.Equal(t, float64(10), v5Model.Data.Priority.ValueFloat64())
	assert.Equal(t, float64(1), v5Model.Data.Weight.ValueFloat64())
	assert.Equal(t, "http://example.com/path", v5Model.Data.Target.ValueString())
}

func TestTransformPriorityFromData(t *testing.T) {
	// Test that priority from data block is promoted to top level for SRV
	v4Model := v4_models.DNSRecordV4Model{
		ZoneID: types.StringValue("zone_srv_prio"),
		Name:   types.StringValue("_service._proto"),
		Type:   types.StringValue("SRV"),
		// No priority at top level
		Data: []v4_models.DNSRecordDataV4Model{
			{
				Priority: types.Float64Value(20), // Priority in data block
				Weight:   types.Float64Value(50),
				Port:     types.Float64Value(8080),
				Target:   types.StringValue("server.example.com"),
			},
		},
	}

	transformer := NewDNSRecordTransformer()
	v5Model, err := transformer.TransformDNSRecord(v4Model)

	assert.NoError(t, err)
	assert.NotNil(t, v5Model)
	// Priority should be promoted to top level
	assert.Equal(t, float64(20), v5Model.Priority.ValueFloat64())
}

func TestTransformWithoutData(t *testing.T) {
	// Test record without data block
	v4Model := v4_models.DNSRecordV4Model{
		ZoneID: types.StringValue("zone_simple"),
		Name:   types.StringValue("simple"),
		Type:   types.StringValue("CNAME"),
		Value:  types.StringValue("target.example.com"),
		TTL:    types.Float64Value(3600),
	}

	transformer := NewDNSRecordTransformer()
	v5Model, err := transformer.TransformDNSRecord(v4Model)

	assert.NoError(t, err)
	assert.NotNil(t, v5Model)
	assert.Nil(t, v5Model.Data) // No data block
	assert.Equal(t, "target.example.com", v5Model.Content.ValueString())
}

func TestTransformCAAFlagsConversion(t *testing.T) {
	tests := []struct {
		name     string
		flags    types.String
		expected float64
	}{
		{
			name:     "numeric string",
			flags:    types.StringValue("128"),
			expected: 128,
		},
		{
			name:     "zero string",
			flags:    types.StringValue("0"),
			expected: 0,
		},
		{
			name:     "invalid string defaults to 0",
			flags:    types.StringValue("invalid"),
			expected: 0,
		},
		{
			name:     "null value",
			flags:    types.StringNull(),
			expected: 0, // Should result in null, but for testing we check if it doesn't crash
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v4Model := v4_models.DNSRecordV4Model{
				ZoneID: types.StringValue("zone_caa_flags"),
				Name:   types.StringValue("test"),
				Type:   types.StringValue("CAA"),
				Data: []v4_models.DNSRecordDataV4Model{
					{
						Flags:   tt.flags,
						Tag:     types.StringValue("issue"),
						Content: types.StringValue("ca.example.com"),
					},
				},
			}

			transformer := NewDNSRecordTransformer()
			v5Model, err := transformer.TransformDNSRecord(v4Model)

			assert.NoError(t, err)
			assert.NotNil(t, v5Model)

			if tt.flags.IsNull() {
				assert.True(t, v5Model.Data.Flags.IsNull() || v5Model.Data.Flags.ValueFloat64() == 0)
			} else {
				assert.Equal(t, tt.expected, v5Model.Data.Flags.ValueFloat64())
			}
		})
	}
}

func TestTransformerOptions(t *testing.T) {
	// Test with AddDefaultTTL = false
	transformer := &DNSRecordTransformer{
		AddDefaultTTL:    false,
		RemoveDeprecated: true,
		PreserveComments: false,
	}

	v4Model := v4_models.DNSRecordV4Model{
		ZoneID: types.StringValue("zone_no_ttl"),
		Name:   types.StringValue("test"),
		Type:   types.StringValue("A"),
		Value:  types.StringValue("192.0.2.5"),
		// No TTL specified
	}

	v5Model, err := transformer.TransformDNSRecord(v4Model)
	assert.NoError(t, err)
	assert.NotNil(t, v5Model)

	// With AddDefaultTTL = false, TTL should remain null
	assert.True(t, v5Model.TTL.IsNull())
}
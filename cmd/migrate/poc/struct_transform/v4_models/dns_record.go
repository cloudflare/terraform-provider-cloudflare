package v4_models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DNSRecordV4Model represents the v4 schema for cloudflare_record
// This is reconstructed from v4 provider documentation
type DNSRecordV4Model struct {
	ID             types.String             `tfsdk:"id"`
	ZoneID         types.String             `tfsdk:"zone_id"`
	Name           types.String             `tfsdk:"name"`
	Type           types.String             `tfsdk:"type"`
	Value          types.String             `tfsdk:"value"`          // v4: renamed to content in v5
	Content        types.String             `tfsdk:"content"`        // v4: some records used content
	TTL            types.Float64            `tfsdk:"ttl"`            // v4: optional, v5: required
	Proxied        types.Bool               `tfsdk:"proxied"`
	Priority       types.Float64            `tfsdk:"priority"`       // For MX, SRV records
	AllowOverwrite types.Bool               `tfsdk:"allow_overwrite"` // v4: deprecated in v5
	Hostname       types.String             `tfsdk:"hostname"`        // v4: deprecated in v5
	Data           []DNSRecordDataV4Model   `tfsdk:"data"`           // v4: array of data blocks
}

// DNSRecordDataV4Model represents the data block in v4
type DNSRecordDataV4Model struct {
	// CAA record fields
	Flags   types.String  `tfsdk:"flags"`   // v4: string, v5: dynamic (string or number)
	Tag     types.String  `tfsdk:"tag"`
	Content types.String  `tfsdk:"content"` // v4: content, v5: renamed to value

	// SRV record fields
	Priority types.Float64 `tfsdk:"priority"`
	Weight   types.Float64 `tfsdk:"weight"`
	Port     types.Float64 `tfsdk:"port"`
	Target   types.String  `tfsdk:"target"`
	Service  types.String  `tfsdk:"service"`
	Proto    types.String  `tfsdk:"proto"`
	Name     types.String  `tfsdk:"name"`

	// LOC record fields
	Altitude      types.Float64 `tfsdk:"altitude"`
	LatDegrees    types.Float64 `tfsdk:"lat_degrees"`
	LatDirection  types.String  `tfsdk:"lat_direction"`
	LatMinutes    types.Float64 `tfsdk:"lat_minutes"`
	LatSeconds    types.Float64 `tfsdk:"lat_seconds"`
	LongDegrees   types.Float64 `tfsdk:"long_degrees"`
	LongDirection types.String  `tfsdk:"long_direction"`
	LongMinutes   types.Float64 `tfsdk:"long_minutes"`
	LongSeconds   types.Float64 `tfsdk:"long_seconds"`
	PrecisionHorz types.Float64 `tfsdk:"precision_horz"`
	PrecisionVert types.Float64 `tfsdk:"precision_vert"`
	Size          types.Float64 `tfsdk:"size"`
}

// Helper methods

func (m *DNSRecordV4Model) HasData() bool {
	return len(m.Data) > 0
}

func (m *DNSRecordV4Model) GetPrimaryValue() types.String {
	// In v4, simple records use 'value', some use 'content'
	if !m.Value.IsNull() && !m.Value.IsUnknown() {
		return m.Value
	}
	return m.Content
}

func (m *DNSRecordV4Model) IsComplexRecord() bool {
	recordType := m.Type.ValueString()
	complexTypes := map[string]bool{
		"CAA": true,
		"SRV": true,
		"LOC": true,
		"URI": true,
	}
	return complexTypes[recordType]
}
package v5_models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// DNSRecordV5Model represents the v5 schema for cloudflare_dns_record
// This is a simplified version - in production, we'd import from the provider
type DNSRecordV5Model struct {
	ID       types.String           `tfsdk:"id"`
	ZoneID   types.String           `tfsdk:"zone_id"`
	Name     types.String           `tfsdk:"name"`
	Type     types.String           `tfsdk:"type"`
	Content  types.String           `tfsdk:"content"`  // v5: renamed from value
	TTL      types.Float64          `tfsdk:"ttl"`      // v5: required with default 1
	Proxied  types.Bool             `tfsdk:"proxied"`
	Priority types.Float64          `tfsdk:"priority"` // For MX, SRV records
	Data     *DNSRecordDataV5Model  `tfsdk:"data"`     // v5: single object, not array
	Tags     []types.String         `tfsdk:"tags"`
	Comment  types.String           `tfsdk:"comment"`
}

// DNSRecordDataV5Model represents the data object in v5
type DNSRecordDataV5Model struct {
	// CAA record fields
	Flags types.Float64 `tfsdk:"flags"` // v5: dynamic, but we'll use float64 for simplicity
	Tag   types.String  `tfsdk:"tag"`
	Value types.String  `tfsdk:"value"` // v5: renamed from content

	// SRV record fields
	Priority types.Float64 `tfsdk:"priority"`
	Weight   types.Float64 `tfsdk:"weight"`
	Port     types.Float64 `tfsdk:"port"`
	Target   types.String  `tfsdk:"target"`
	Service  types.String  `tfsdk:"service"`

	// LOC record fields (same as v4)
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

func (m *DNSRecordV5Model) SetDefaults() {
	// Set default TTL if not set
	if m.TTL.IsNull() {
		m.TTL = types.Float64Value(1) // Automatic TTL
	}

	// Set default proxied if not set
	if m.Proxied.IsNull() {
		m.Proxied = types.BoolValue(false)
	}
}

func (m *DNSRecordV5Model) Validate() error {
	// Add validation logic here
	return nil
}
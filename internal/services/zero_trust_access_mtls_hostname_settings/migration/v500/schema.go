package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceMTLSHostnameSettingsSchema returns the minimal schema for v4 cloudflare_zero_trust_access_mtls_hostname_settings.
// Schema version: 0 (SDKv2 default)
//
// This minimal schema is used only for reading v4 state during migration.
func SourceMTLSHostnameSettingsSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"settings": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"china_network": schema.BoolAttribute{
							Optional: true, // Was Optional in v4, Required in v5
						},
						"client_certificate_forwarding": schema.BoolAttribute{
							Optional: true, // Was Optional in v4, Required in v5
						},
						"hostname": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			// Root-level computed fields
			"china_network": schema.BoolAttribute{
				Computed: true,
			},
			"client_certificate_forwarding": schema.BoolAttribute{
				Computed: true,
			},
			"hostname": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

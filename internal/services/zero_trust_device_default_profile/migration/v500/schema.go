// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceDeviceProfileSchema returns the minimal v4 schema needed to read legacy state
// This matches the v4 cloudflare_zero_trust_device_profiles / cloudflare_device_settings_policy resources
// Schema version: 0 (SDKv2 default)
func SourceDeviceProfileSchema() schema.Schema {
	return schema.Schema{
		Version: 0,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"default": schema.BoolAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Required: true,
			},
			"precedence": schema.Int64Attribute{
				Optional: true,
			},
			"match": schema.StringAttribute{
				Optional: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"disable_auto_fallback": schema.BoolAttribute{
				Optional: true,
			},
			"captive_portal": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"allow_mode_switch": schema.BoolAttribute{
				Optional: true,
			},
			"switch_locked": schema.BoolAttribute{
				Optional: true,
			},
			"allow_updates": schema.BoolAttribute{
				Optional: true,
			},
			"auto_connect": schema.Int64Attribute{
				Optional: true,
			},
			"allowed_to_leave": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"support_url": schema.StringAttribute{
				Optional: true,
			},
			"service_mode_v2_mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"service_mode_v2_port": schema.Int64Attribute{
				Optional: true,
			},
			"exclude_office_ips": schema.BoolAttribute{
				Optional: true,
			},
			"tunnel_protocol": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"lan_allow_minutes": schema.Float64Attribute{
				Optional: true,
			},
			"lan_allow_subnet_size": schema.Float64Attribute{
				Optional: true,
			},
		},
	}
}

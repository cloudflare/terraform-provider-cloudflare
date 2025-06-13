// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceSettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"disable_for_time": schema.Float64Attribute{
				Description: "Sets the time limit, in seconds, that a user can use an override code to bypass WARP.",
				Computed:    true,
			},
			"gateway_proxy_enabled": schema.BoolAttribute{
				Description: "Enable gateway proxy filtering on TCP.",
				Computed:    true,
			},
			"gateway_udp_proxy_enabled": schema.BoolAttribute{
				Description: "Enable gateway proxy filtering on UDP.",
				Computed:    true,
			},
			"root_certificate_installation_enabled": schema.BoolAttribute{
				Description: "Enable installation of cloudflare managed root certificate.",
				Computed:    true,
			},
			"use_zt_virtual_ip": schema.BoolAttribute{
				Description: "Enable using CGNAT virtual IPv4.",
				Computed:    true,
			},
		},
	}
}

func (d *ZeroTrustDeviceSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceSettingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

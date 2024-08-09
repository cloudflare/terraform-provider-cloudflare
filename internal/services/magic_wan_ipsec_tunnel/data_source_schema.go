// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_ipsec_tunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &MagicWANIPSECTunnelDataSource{}

func (d *MagicWANIPSECTunnelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"ipsec_tunnel_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"ipsec_tunnel": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (d *MagicWANIPSECTunnelDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

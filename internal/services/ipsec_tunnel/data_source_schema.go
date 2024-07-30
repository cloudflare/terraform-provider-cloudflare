// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ipsec_tunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &IPSECTunnelDataSource{}
var _ datasource.DataSourceWithValidateConfig = &IPSECTunnelDataSource{}

func (r IPSECTunnelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

func (r *IPSECTunnelDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *IPSECTunnelDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

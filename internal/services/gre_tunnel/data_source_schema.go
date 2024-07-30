// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gre_tunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &GRETunnelDataSource{}
var _ datasource.DataSourceWithValidateConfig = &GRETunnelDataSource{}

func (r GRETunnelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"gre_tunnel_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"gre_tunnel": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *GRETunnelDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *GRETunnelDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

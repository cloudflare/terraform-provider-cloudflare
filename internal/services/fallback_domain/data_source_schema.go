// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fallback_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &FallbackDomainDataSource{}
var _ datasource.DataSourceWithValidateConfig = &FallbackDomainDataSource{}

func (r FallbackDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"policy_id": schema.StringAttribute{
				Description: "Device ID.",
				Optional:    true,
			},
			"suffix": schema.StringAttribute{
				Description: "The domain suffix to match when resolving locally.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the fallback domain, displayed in the client UI.",
				Optional:    true,
			},
			"dns_server": schema.StringAttribute{
				Description: "A list of IP addresses to handle domain resolution.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (r *FallbackDomainDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *FallbackDomainDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

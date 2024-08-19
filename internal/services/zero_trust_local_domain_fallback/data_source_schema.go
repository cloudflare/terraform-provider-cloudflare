// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &ZeroTrustLocalDomainFallbackDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"policy_id": schema.StringAttribute{
				Description: "Device ID.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the fallback domain, displayed in the client UI.",
				Optional:    true,
			},
			"suffix": schema.StringAttribute{
				Description: "The domain suffix to match when resolving locally.",
				Optional:    true,
			},
			"dns_server": schema.ListAttribute{
				Description: "A list of IP addresses to handle domain resolution.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
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

func (d *ZeroTrustLocalDomainFallbackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustLocalDomainFallbackDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

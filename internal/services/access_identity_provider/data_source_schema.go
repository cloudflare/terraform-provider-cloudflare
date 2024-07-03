// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_identity_provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &AccessIdentityProviderDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessIdentityProviderDataSource{}

func (r AccessIdentityProviderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"identity_provider_id": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
						Optional:    true,
					},
					"zone_id": schema.StringAttribute{
						Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *AccessIdentityProviderDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessIdentityProviderDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

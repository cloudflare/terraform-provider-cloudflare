// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessApplicationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"app_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Validators: []validator.Object{
					objectvalidator.ExactlyOneOf(path.MatchRelative().AtName("account_id"), path.MatchRelative().AtName("zone_id")),
				},
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
						Optional:    true,
					},
					"zone_id": schema.StringAttribute{
						Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
						Optional:    true,
					},
					"aud": schema.StringAttribute{
						Description: "The aud of the app.",
						Optional:    true,
					},
					"domain": schema.StringAttribute{
						Description: "The domain of the app.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the app.",
						Optional:    true,
					},
					"search": schema.StringAttribute{
						Description: "Search for apps by other listed query parameters.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessApplicationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessApplicationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("app_id")),
		datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
		datasourcevalidator.Conflicting(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.Conflicting(path.MatchRoot("filter"), path.MatchRoot("zone_id")),
	}
}

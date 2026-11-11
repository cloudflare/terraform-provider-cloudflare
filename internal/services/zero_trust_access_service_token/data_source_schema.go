// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_service_token

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessServiceTokenDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Access: Service Tokens Read",
				"Access: Service Tokens Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "UUID.",
				Computed:    true,
			},
			"service_token_id": schema.StringAttribute{
				Description: "UUID.",
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
			"client_id": schema.StringAttribute{
				Description: "The Client ID for the service token. Access will check for this value in the `CF-Access-Client-ID` request header.",
				Computed:    true,
			},
			"duration": schema.StringAttribute{
				Description: "The duration for how long the service token will be valid. Must be in the format `300ms` or `2h45m`, or the special value `forever` for non-expiring tokens. Valid time units are: ns, us (or µs), ms, s, m, h. The default is 1 year in hours (8760h).",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the service token is enabled. A disabled service token cannot be used to authenticate; both its current and previous `client_secret` stop being accepted, but the token itself is preserved and can be re-enabled. A service token reported as leaked is disabled automatically and cannot be re-enabled until its secret is rotated; attempting to do so returns a 409. Defaults to enabled when omitted on create.",
				Computed:    true,
			},
			"expires_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"leak_report_url": schema.StringAttribute{
				Description: "Where the leaked `client_secret` was found, if provided by the reporter. Omitted if the token has not been reported as leaked.",
				Computed:    true,
			},
			"leak_reported_at": schema.StringAttribute{
				Description: "When a secret scanning partner reported this service token's `client_secret` as leaked. A leaked service token is disabled and must be rotated before it can be re-enabled. Cleared once the secret is rotated. Omitted if the token has not been reported as leaked.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"leak_reporter": schema.StringAttribute{
				Description: "The secret scanning partner that reported the leak. Omitted if the token has not been reported as leaked.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the service token.",
				Computed:    true,
			},
			"tags": schema.MapAttribute{
				Description: "Tags associated with the service token, as key-value pairs. On update, replaces all existing tags.",
				Computed:    true,
				CustomType:  customfield.NewMapType[types.String](ctx),
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "The name of the service token.",
						Optional:    true,
					},
					"search": schema.StringAttribute{
						Description: "Search for service tokens by other listed query parameters.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessServiceTokenDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessServiceTokenDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("service_token_id"), path.MatchRoot("filter")),
		datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}

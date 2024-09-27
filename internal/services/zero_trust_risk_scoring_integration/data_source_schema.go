// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_scoring_integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustRiskScoringIntegrationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"integration_id": schema.StringAttribute{
				Optional: true,
			},
			"account_tag": schema.StringAttribute{
				Description: "The Cloudflare account tag.",
				Computed:    true,
			},
			"active": schema.BoolAttribute{
				Description: "Whether this integration is enabled and should export changes in risk score.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the integration was created in RFC3339 format.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "The id of the integration, a UUIDv4.",
				Computed:    true,
			},
			"integration_type": schema.StringAttribute{
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Okta"),
				},
			},
			"reference_id": schema.StringAttribute{
				Description: "A reference ID defined by the client.\nShould be set to the Access-Okta IDP integration ID.\nUseful when the risk-score integration needs to be associated with a secondary asset and recalled using that ID.",
				Computed:    true,
			},
			"tenant_url": schema.StringAttribute{
				Description: "The base URL for the tenant. E.g. \"https://tenant.okta.com\"",
				Computed:    true,
			},
			"well_known_url": schema.StringAttribute{
				Description: "The URL for the Shared Signals Framework configuration, e.g. \"/.well-known/sse-configuration/{integration_uuid}/\". https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1",
				Computed:    true,
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

func (d *ZeroTrustRiskScoringIntegrationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustRiskScoringIntegrationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("integration_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("integration_id")),
	}
}

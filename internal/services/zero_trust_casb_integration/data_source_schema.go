// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_integration

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCasbIntegrationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Integration ID to look up. Exactly one of `id` or `filter` must be configured.",
				Computed:    true,
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"created": schema.StringAttribute{
				Description: "When the integration was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"credentials_expiry": schema.StringAttribute{
				Description: "Credentials expiry time.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"is_paused": schema.BoolAttribute{
				Description: "Whether the user paused the integration.",
				Computed:    true,
			},
			"last_hydrated": schema.StringAttribute{
				Description: "Last time the integration was hydrated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Name of the integration.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Integration status.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "When the integration was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"application": schema.MapAttribute{
				Computed:    true,
				CustomType:  customfield.NewMapType[types.String](ctx),
				ElementType: types.StringType,
			},
			"auth_method": schema.MapAttribute{
				Description: "The integration's authentication method.",
				Computed:    true,
				CustomType:  customfield.NewMapType[types.String](ctx),
				ElementType: types.StringType,
			},
			"dlp_profiles": schema.ListAttribute{
				Description: "DLP Profiles enabled for the integration.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"health_details": schema.ListAttribute{
				Description: "Health details with remediation hints.",
				Computed:    true,
				CustomType:  customfield.NewListType[customfield.Map[jsontypes.Normalized]](ctx),
				ElementType: types.MapType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
			"use_cases": schema.ListAttribute{
				Description: "Use cases enabled for the integration.",
				Computed:    true,
				CustomType:  customfield.NewListType[customfield.Map[jsontypes.Normalized]](ctx),
				ElementType: types.MapType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
			"authorization_link": schema.SingleNestedAttribute{
				Description: "Authorization link for the integration.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustCasbIntegrationAuthorizationLinkDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"components": schema.MapAttribute{
						Computed:    true,
						CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
						ElementType: jsontypes.NormalizedType{},
					},
					"link": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"application": schema.StringAttribute{
						Description: "Filter by application/vendor (e.g., GOOGLE_WORKSPACE, MICROSOFT_INTERNAL).",
						Optional:    true,
					},
					"direction": schema.StringAttribute{
						Description: "Direction to order results.\nAvailable values: \"asc\", \"desc\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"dlp_enabled": schema.BoolAttribute{
						Description: "Filter by DLP enabled status (true/false).",
						Optional:    true,
					},
					"order": schema.StringAttribute{
						Description: "Field to order results by.\nAvailable values: \"application\", \"created\", \"name\", \"status\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"application",
								"created",
								"name",
								"status",
							),
						},
					},
					"page": schema.Int64Attribute{
						Description: "Page number within the paginated result set.",
						Optional:    true,
					},
					"page_size": schema.Int64Attribute{
						Description: "Number of results per page.",
						Optional:    true,
					},
					"search": schema.StringAttribute{
						Description: "Search integrations by name or application.",
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Description: "Filter by integration status.\nAvailable values: \"Healthy\", \"Initializing\", \"Offline\", \"Unhealthy\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"Healthy",
								"Initializing",
								"Offline",
								"Unhealthy",
							),
						},
					},
					"use_cases": schema.StringAttribute{
						Description: "Filter by one enabled use case (for example, casb or ces).",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustCasbIntegrationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustCasbIntegrationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("filter")),
	}
}

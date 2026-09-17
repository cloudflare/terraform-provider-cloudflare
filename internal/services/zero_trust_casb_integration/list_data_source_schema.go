// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_integration

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCasbIntegrationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
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
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustCasbIntegrationsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Integration ID.",
							Computed:    true,
						},
						"application": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"created": schema.StringAttribute{
							Description: "When the integration was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"is_paused": schema.BoolAttribute{
							Description: "Whether the user paused the integration.",
							Computed:    true,
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
					},
				},
			},
		},
	}
}

func (d *ZeroTrustCasbIntegrationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustCasbIntegrationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

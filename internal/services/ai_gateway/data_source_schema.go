// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*AIGatewayDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "gateway id",
				Computed:    true,
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"authentication": schema.BoolAttribute{
				Computed: true,
			},
			"cache_invalidate_on_update": schema.BoolAttribute{
				Computed: true,
			},
			"cache_ttl": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"collect_logs": schema.BoolAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"is_default": schema.BoolAttribute{
				Computed: true,
			},
			"log_management": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.Between(10000, 10000000),
				},
			},
			"log_management_strategy": schema.StringAttribute{
				Description: `Available values: "STOP_INSERTING", "DELETE_OLDEST".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("STOP_INSERTING", "DELETE_OLDEST"),
				},
			},
			"logpush": schema.BoolAttribute{
				Computed: true,
			},
			"logpush_public_key": schema.StringAttribute{
				Computed: true,
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"rate_limiting_interval": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"rate_limiting_limit": schema.Int64Attribute{
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"rate_limiting_technique": schema.StringAttribute{
				Description: `Available values: "fixed", "sliding".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("fixed", "sliding"),
				},
			},
			"retry_backoff": schema.StringAttribute{
				Description: "Backoff strategy for retry delays\nAvailable values: \"constant\", \"linear\", \"exponential\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"constant",
						"linear",
						"exponential",
					),
				},
			},
			"retry_delay": schema.Int64Attribute{
				Description: "Delay between retry attempts in milliseconds (0-5000)",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 5000),
				},
			},
			"retry_max_attempts": schema.Int64Attribute{
				Description: "Maximum number of retry attempts for failed requests (1-5)",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 5),
				},
			},
			"store_id": schema.StringAttribute{
				Computed: true,
			},
			"workers_ai_billing_mode": schema.StringAttribute{
				Description: "Controls how Workers AI inference calls routed through this gateway are billed. Only 'postpaid' is currently supported.\nAvailable values: \"postpaid\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("postpaid"),
				},
			},
			"zdr": schema.BoolAttribute{
				Computed: true,
			},
			"dlp": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AIGatewayDLPDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description: `Available values: "BLOCK", "FLAG".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("BLOCK", "FLAG"),
						},
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
					},
					"profiles": schema.ListAttribute{
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"policies": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[AIGatewayDLPPoliciesDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"action": schema.StringAttribute{
									Description: `Available values: "FLAG", "BLOCK".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
									},
								},
								"check": schema.ListAttribute{
									Computed: true,
									Validators: []validator.List{
										listvalidator.ValueStringsAre(
											stringvalidator.OneOfCaseInsensitive("REQUEST", "RESPONSE"),
										),
									},
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"enabled": schema.BoolAttribute{
									Computed: true,
								},
								"profiles": schema.ListAttribute{
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},
			"otel": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[AIGatewayOtelDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"authorization": schema.StringAttribute{
							Computed: true,
						},
						"headers": schema.MapAttribute{
							Computed:    true,
							CustomType:  customfield.NewMapType[types.String](ctx),
							ElementType: types.StringType,
						},
						"url": schema.StringAttribute{
							Computed: true,
						},
						"content_type": schema.StringAttribute{
							Description: `Available values: "json", "protobuf".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("json", "protobuf"),
							},
						},
					},
				},
			},
			"stripe": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AIGatewayStripeDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"authorization": schema.StringAttribute{
						Computed: true,
					},
					"usage_events": schema.ListNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectListType[AIGatewayStripeUsageEventsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"payload": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"search": schema.StringAttribute{
						Description: "Search by id",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *AIGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AIGatewayDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("filter")),
	}
}

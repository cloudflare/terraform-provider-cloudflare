// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AIGatewayResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "gateway id",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cache_invalidate_on_update": schema.BoolAttribute{
				Required: true,
			},
			"cache_ttl": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"collect_logs": schema.BoolAttribute{
				Required: true,
			},
			"rate_limiting_interval": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"rate_limiting_limit": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"authentication": schema.BoolAttribute{
				Optional: true,
			},
			"log_management": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.Between(10000, 10000000),
				},
			},
			"log_management_strategy": schema.StringAttribute{
				Description: `Available values: "STOP_INSERTING", "DELETE_OLDEST".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("STOP_INSERTING", "DELETE_OLDEST"),
				},
			},
			"logpush": schema.BoolAttribute{
				Optional: true,
			},
			"logpush_public_key": schema.StringAttribute{
				Optional: true,
			},
			"rate_limiting_technique": schema.StringAttribute{
				Description: `Available values: "fixed", "sliding".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("fixed", "sliding"),
				},
			},
			"retry_backoff": schema.StringAttribute{
				Description: "Backoff strategy for retry delays\nAvailable values: \"constant\", \"linear\", \"exponential\".",
				Optional:    true,
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
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 5000),
				},
			},
			"retry_max_attempts": schema.Int64Attribute{
				Description: "Maximum number of retry attempts for failed requests (1-5)",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 5),
				},
			},
			"store_id": schema.StringAttribute{
				Optional: true,
			},
			"zdr": schema.BoolAttribute{
				Optional: true,
			},
			"dlp": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description: `Available values: "BLOCK", "FLAG".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("BLOCK", "FLAG"),
						},
					},
					"enabled": schema.BoolAttribute{
						Required: true,
					},
					"profiles": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"policies": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Required: true,
								},
								"action": schema.StringAttribute{
									Description: `Available values: "FLAG", "BLOCK".`,
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
									},
								},
								"check": schema.ListAttribute{
									Required: true,
									Validators: []validator.List{
										listvalidator.ValueStringsAre(
											stringvalidator.OneOfCaseInsensitive("REQUEST", "RESPONSE"),
										),
									},
									ElementType: types.StringType,
								},
								"enabled": schema.BoolAttribute{
									Required: true,
								},
								"profiles": schema.ListAttribute{
									Required:    true,
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},
			"stripe": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"authorization": schema.StringAttribute{
						Required: true,
					},
					"usage_events": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"payload": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
				},
			},
			"workers_ai_billing_mode": schema.StringAttribute{
				Description: "Controls how Workers AI inference calls routed through this gateway are billed. Only 'postpaid' is currently supported.\nAvailable values: \"postpaid\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("postpaid"),
				},
				Default: stringdefault.StaticString("postpaid"),
			},
			"otel": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[AIGatewayOtelModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"authorization": schema.StringAttribute{
							Required: true,
						},
						"headers": schema.MapAttribute{
							Required:    true,
							ElementType: types.StringType,
						},
						"url": schema.StringAttribute{
							Required: true,
						},
						"content_type": schema.StringAttribute{
							Description: `Available values: "json", "protobuf".`,
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("json", "protobuf"),
							},
							Default: stringdefault.StaticString("json"),
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"is_default": schema.BoolAttribute{
				Computed: true,
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *AIGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AIGatewayResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

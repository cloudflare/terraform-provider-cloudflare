// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"AI Gateway Read",
				"AI Gateway Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "gateway id",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
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
			"guardrails": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"prompt": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"p1": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s1": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s10": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s11": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s12": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s13": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s2": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s3": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s4": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s5": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s6": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s7": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s8": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s9": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
						},
					},
					"response": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"p1": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s1": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s10": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s11": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s12": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s13": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s2": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s3": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s4": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s5": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s6": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s7": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s8": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
								},
							},
							"s9": schema.StringAttribute{
								Description: `Available values: "FLAG", "BLOCK".`,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("FLAG", "BLOCK"),
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
						"headers": schema.MapAttribute{
							Required:    true,
							ElementType: types.StringType,
						},
						"url": schema.StringAttribute{
							Required: true,
						},
						"authorization": schema.StringAttribute{
							Optional: true,
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
			"spend_limits": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[AIGatewaySpendLimitsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
					"rules": schema.ListNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectListType[AIGatewaySpendLimitsRulesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"limit": schema.Float64Attribute{
									Required: true,
									Validators: []validator.Float64{
										float64validator.AtLeast(0),
									},
								},
								"limit_type": schema.StringAttribute{
									Description: `Available values: "cost".`,
									Required:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("cost"),
									},
								},
								"window": schema.Int64Attribute{
									Required: true,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
								"id": schema.StringAttribute{
									Computed: true,
									Optional: true,
									Default:  stringdefault.StaticString("8d8fda51"),
								},
								"enabled": schema.BoolAttribute{
									Computed: true,
									Optional: true,
									Default:  booldefault.StaticBool(true),
								},
								"metadata": schema.MapNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Description: `Available values: "partition", "filter".`,
												Required:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("partition", "filter"),
												},
											},
											"values": schema.ListAttribute{
												Optional:    true,
												ElementType: types.StringType,
											},
										},
									},
								},
								"model": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description: `Available values: "filter".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("filter"),
											},
										},
										"values": schema.ListAttribute{
											Required:    true,
											ElementType: types.StringType,
										},
									},
								},
								"ai_gateway_provider": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description: `Available values: "filter".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("filter"),
											},
										},
										"values": schema.ListAttribute{
											Required:    true,
											ElementType: types.StringType,
										},
									},
								},
								"technique": schema.StringAttribute{
									Description: `Available values: "fixed", "sliding".`,
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("fixed", "sliding"),
									},
									Default: stringdefault.StaticString("sliding"),
								},
							},
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

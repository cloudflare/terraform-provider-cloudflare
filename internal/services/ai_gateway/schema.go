// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "AI Gateway identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for the AI Gateway.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cache_invalidate_on_update": schema.BoolAttribute{
				Description: "Invalidate the cache on update.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"cache_ttl": schema.Int64Attribute{
				Description: "Cache TTL in seconds.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"collect_logs": schema.BoolAttribute{
				Description: "Collect logs from the gateway.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"created_at": schema.StringAttribute{
				Description:   "The timestamp when the AI Gateway was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"modified_at": schema.StringAttribute{
				Description:   "The timestamp when the AI Gateway was last modified.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"rate_limiting_interval": schema.Int64Attribute{
				Description: "Rate limiting interval in seconds.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"rate_limiting_limit": schema.Int64Attribute{
				Description: "Rate limiting limit.",
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(0),
			},
			"authentication": schema.BoolAttribute{
				Description: "Enable authentication on the gateway.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"dlp": schema.SingleNestedAttribute{
				Description: "Data Loss Prevention configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description: `Available values: "BLOCK", "WARN".`,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString("WARN"),
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("BLOCK", "WARN"),
						},
					},
					"enabled": schema.BoolAttribute{
						Description: "Enable DLP.",
						Optional:    true,
						Computed:    true,
						Default:     booldefault.StaticBool(false),
					},
					"profiles": schema.ListAttribute{
						Description: "List of DLP profile IDs.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"is_default": schema.BoolAttribute{
				Description: "Whether this is the default gateway.",
				Computed:    true,
				Optional:    true,
			},
			"log_management": schema.Int64Attribute{
				Description: "Log management setting.",
				Optional:    true,
				Computed:    true,
			},
			"log_management_strategy": schema.StringAttribute{
				Description: `Available values: "STOP_INSERTING", "DROP_WHEN_FULL", "NO_SQL".`,
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("STOP_INSERTING", "DROP_WHEN_FULL", "NO_SQL"),
				},
			},
			"logpush": schema.BoolAttribute{
				Description: "Enable logpush.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"logpush_public_key": schema.StringAttribute{
				Description: "Logpush public key.",
				Computed:    true,
				Optional:    true,
			},
			"otel": schema.ListNestedAttribute{
				Description: "OpenTelemetry configuration.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"authorization": schema.StringAttribute{
							Description: "Authorization header value.",
							Optional:    true,
						},
						"headers": schema.MapAttribute{
							Description: "Additional headers to include in OTEL requests.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"url": schema.StringAttribute{
							Description: "OTEL collector URL.",
							Optional:    true,
						},
						"content_type": schema.StringAttribute{
							Description: `Available values: "json", "proto".`,
							Optional:    true,
							Computed:    true,
							Default:     stringdefault.StaticString("json"),
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("json", "proto"),
							},
						},
					},
				},
			},
			"rate_limiting_technique": schema.StringAttribute{
				Description: `Available values: "fixed", "sliding".`,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("fixed"),
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("fixed", "sliding"),
				},
			},
			"store_id": schema.StringAttribute{
				Description: "Store ID for logs.",
				Optional:    true,
				Computed:    true,
			},
			"stripe": schema.SingleNestedAttribute{
				Description: "Stripe configuration for billing.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"authorization": schema.StringAttribute{
						Description: "Stripe authorization token.",
						Optional:    true,
						Sensitive:   true,
					},
					"usage_events": schema.ListNestedAttribute{
						Description: "Usage events configuration.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"payload": schema.StringAttribute{
									Description: "Event payload.",
									Optional:    true,
								},
							},
						},
					},
				},
			},
			"workers_ai_billing_mode": schema.StringAttribute{
				Description: `Available values: "postpaid", "prepaid".`,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("postpaid"),
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("postpaid", "prepaid"),
				},
			},
			"zdr": schema.BoolAttribute{
				Description: "Enable Zero Downtime Requirements.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func (r *AIGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

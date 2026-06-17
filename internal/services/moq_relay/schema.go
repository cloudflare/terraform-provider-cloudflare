// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq_relay

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*MoqRelayResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Server-generated unique identifier (32 hex chars).",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"uid": schema.StringAttribute{
				Description:   "Server-generated unique identifier (32 hex chars).",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Cloudflare account identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Human-readable name for the relay.",
				Required:    true,
			},
			"config": schema.SingleNestedAttribute{
				Description: "origin_fallback and lingering_subscribe are mutually exclusive.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[MoqRelayConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"lingering_subscribe": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[MoqRelayConfigLingeringSubscribeModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
								Optional: true,
								Default:  booldefault.StaticBool(true),
							},
							"max_timeout_ms": schema.Int64Attribute{
								Description: "Relay-level ceiling on lingering subscribe timeout (ms). Default 30000.",
								Computed:    true,
								Optional:    true,
								Validators: []validator.Int64{
									int64validator.Between(0, 300000),
								},
								Default: int64default.StaticInt64(30000),
							},
						},
					},
					"origin_fallback": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[MoqRelayConfigOriginFallbackModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
								Optional: true,
								Default:  booldefault.StaticBool(false),
							},
							"origins": schema.ListNestedAttribute{
								Description: "Ordered list of upstream origin relays. Each entry is an object\n(not a bare string) so per-origin configuration can be added in\nthe future without another breaking change.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectListType[MoqRelayConfigOriginFallbackOriginsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"url": schema.StringAttribute{
											Description: "Upstream origin relay URL.",
											Optional:    true,
										},
									},
								},
							},
						},
					},
				},
			},
			"created": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "\"connected\" when active, omitted otherwise.\nAvailable values: \"connected\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("connected"),
				},
			},
			"token_publish_subscribe": schema.StringAttribute{
				Description: "Full access token (publish + subscribe). Treat as sensitive.",
				Computed:    true,
				Sensitive:   true,
			},
			"token_subscribe": schema.StringAttribute{
				Description: "Subscribe-only token. Treat as sensitive.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func (r *MoqRelayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MoqRelayResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

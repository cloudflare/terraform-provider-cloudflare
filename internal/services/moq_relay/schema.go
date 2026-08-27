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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*MoQRelayResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
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
				Description: "upstreams and lingering_subscribe are mutually exclusive.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[MoQRelayConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"lingering_subscribe": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[MoQRelayConfigLingeringSubscribeModel](ctx),
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
					"upstreams": schema.SingleNestedAttribute{
						Description: "Upstreams are external MOQT server publishers that a relay falls back\nto when it has no local publisher for a requested namespace/track.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[MoQRelayConfigUpstreamsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
								Optional: true,
								Default:  booldefault.StaticBool(false),
							},
							"upstreams": schema.ListNestedAttribute{
								Description: "Ordered list of upstream MOQT server publishers. Each entry is an\nobject (not a bare string) so per-upstream configuration can be\nadded in the future without another breaking change.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectListType[MoQRelayConfigUpstreamsUpstreamsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"url": schema.StringAttribute{
											Description: "Upstream MOQT server publisher URL. Must be an absolute URL with a\nhost and a scheme the relay can dial: moqt:// (raw QUIC) or https://\n(WebTransport). Validated on update (PUT); rejected with 21013.",
											Optional:    true,
										},
									},
								},
								PlanModifiers: []planmodifier.List{listplanmodifier.UseNonNullStateForUnknown()},
							},
						},
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
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
			"issuers": schema.ListNestedAttribute{
				Description: "Token collection (discriminated union on `type`). On create this\nholds the auto-created default pair, each including its one-time\nsecret.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MoQRelayIssuersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cloudflare_tokens": schema.ListNestedAttribute{
							Description: "Always present ([] when empty).",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[MoQRelayIssuersCloudflareTokensModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"created": schema.StringAttribute{
										Computed:      true,
										CustomType:    timetypes.RFC3339Type{},
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
									},
									"expires": schema.StringAttribute{
										Description: "Mandatory; no more than 1 year after `created`.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"jti": schema.StringAttribute{
										Description: "Token identity and registry key (32 hex chars).",
										Computed:    true,
									},
									"operations": schema.ListAttribute{
										Description: "Signed allowlist of what the token may do. V1 coarse roles; the array\nform extends to fine-grained MoQT message names later without a\nbreaking change.",
										Computed:    true,
										Validators: []validator.List{
											listvalidator.ValueStringsAre(
												stringvalidator.OneOfCaseInsensitive("publish", "subscribe"),
											),
										},
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"label": schema.StringAttribute{
										Description: "Optional, customer-set.",
										Computed:    true,
									},
									"secret": schema.StringAttribute{
										Description: "The signed JWT. Present ONLY in create / auto-create responses (shown\nonce); never returned by list, never stored.",
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"issuer": schema.StringAttribute{
							Description: `Available values: "cloudflare".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("cloudflare"),
							},
						},
						"type": schema.StringAttribute{
							Description: `Available values: "cloudflare_jwt".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("cloudflare_jwt"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *MoQRelayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MoQRelayResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

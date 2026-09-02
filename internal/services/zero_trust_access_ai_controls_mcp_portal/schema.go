// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_portal

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessAIControlsMcpPortalResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 501,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"MCP Portals Read",
				"MCP Portals Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the MCP portal.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Description: "Hostname where the MCP portal is available.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Display name for the MCP portal.",
				Required:    true,
			},
			"allow_code_mode": schema.BoolAttribute{
				Description:        "Deprecated: use `code_mode` for new integrations. `true` maps to any non-off Code Mode policy; `false` maps to `code_mode: off`. If both fields are sent, they must be consistent or the request returns a 400.",
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
			},
			"code_mode": schema.StringAttribute{
				Description: "Code Mode policy for this portal. `off`: Code Mode is unavailable; query parameters are ignored. `opt_in`: Code Mode is off by default; clients turn it on with `?codemode=search_and_execute`. `default_on`: Code Mode is on by default; clients can opt out with `?codemode=off`. `enforced`: Code Mode is always on; query parameters are ignored. Defaults to `opt_in` when omitted on create. If both `code_mode` and `allow_code_mode` are sent, they must be consistent or the request returns a 400.\nAvailable values: \"off\", \"opt_in\", \"default_on\", \"enforced\".",
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"off",
						"opt_in",
						"default_on",
						"enforced",
					),
				},
			},
			"description": schema.StringAttribute{
				Description: "Optional description of the MCP portal.",
				Optional:    true,
			},
			"secure_web_gateway": schema.BoolAttribute{
				Description: "Route outbound MCP traffic through Zero Trust Secure Web Gateway.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"servers": schema.SetNestedAttribute{
				Description: "MCP servers attached to the portal and their portal-specific settings.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectSetType[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"server_id": schema.StringAttribute{
							Description: "Unique identifier for the MCP server.",
							Required:    true,
						},
						"default_disabled": schema.BoolAttribute{
							Description: "Disable this server by default for clients connecting through the portal.",
							Computed:    true,
							Optional:    true,
							Default:     booldefault.StaticBool(false),
						},
						"on_behalf": schema.BoolAttribute{
							Description: "Use end-user OAuth credentials when connecting this server to the portal.",
							Computed:    true,
							Optional:    true,
							Default:     booldefault.StaticBool(true),
						},
						"updated_prompts": schema.ListNestedAttribute{
							Description: "Portal-specific prompt overrides.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description: "Name of the tool or prompt capability to override.",
										Required:    true,
									},
									"alias": schema.StringAttribute{
										Description: "Custom name exposed for the capability.",
										Optional:    true,
									},
									"description": schema.StringAttribute{
										Description: "Custom description exposed for the capability.",
										Optional:    true,
									},
									"enabled": schema.BoolAttribute{
										Description: "Whether the capability is available through the MCP server.",
										Optional:    true,
									},
								},
							},
						},
						"updated_tools": schema.ListNestedAttribute{
							Description: "Portal-specific tool overrides.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description: "Name of the tool or prompt capability to override.",
										Required:    true,
									},
									"alias": schema.StringAttribute{
										Description: "Custom name exposed for the capability.",
										Optional:    true,
									},
									"description": schema.StringAttribute{
										Description: "Custom description exposed for the capability.",
										Optional:    true,
									},
									"enabled": schema.BoolAttribute{
										Description: "Whether the capability is available through the MCP server.",
										Optional:    true,
									},
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
			"created_by": schema.StringAttribute{
				Computed: true,
			},
			"modified_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_by": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ZeroTrustAccessAIControlsMcpPortalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessAIControlsMcpPortalResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_server

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessAIControlsMcpServerResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"MCP Portals Read",
				"MCP Portals Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the MCP server.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"auth_type": schema.StringAttribute{
				Description: "Authentication method used to connect to the upstream MCP server.\nAvailable values: \"oauth\", \"bearer\", \"unauthenticated\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"oauth",
						"bearer",
						"unauthenticated",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Description:   "URL of the upstream MCP endpoint.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Display name for the MCP server.",
				Required:    true,
			},
			"auth_credentials": schema.StringAttribute{
				Description: "Static credential for the upstream MCP server. For auth_type \"bearer\", either a raw token string (e.g. \"sk-abc123\"), which is wrapped server-side as `Authorization: Bearer <token>`, or a JSON-encoded object of the form `{\"headers\":{\"Header-Name\":\"value\",...}}` for custom or multiple static headers (e.g. Cloudflare Access service tokens: `{\"headers\":{\"cf-access-client-id\":\"...\",\"cf-access-client-secret\":\"...\"}}`).",
				Optional:    true,
				Sensitive:   true,
			},
			"client_secret": schema.StringAttribute{
				Description: "Pre-registered OAuth client_secret. Write-only - accepted on create/update when auth_credentials.auth_mode is 'manual'. Stored AES-GCM-encrypted in server_oauth_secrets; never returned by read endpoints.",
				Optional:    true,
				Sensitive:   true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description of the MCP server.",
				Optional:    true,
			},
			"updated_prompts": schema.ListNestedAttribute{
				Description: "Server-wide prompt capability overrides.",
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
				Description: "Server-wide tool capability overrides.",
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
			"is_shared_oauth_callback_enabled": schema.BoolAttribute{
				Description: "When true, the gateway worker uses the shared Cloudflare-owned OAuth callback endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the customer portal hostname. Defaults to false (off); opt in per server by setting true.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"secure_web_gateway": schema.BoolAttribute{
				Description: "Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"authentication_status": schema.StringAttribute{
				Description: "Whether administrative authentication is required before capabilities can be synced. Manual OAuth is user-managed and has no administrative authentication flow.\nAvailable values: \"not_required\", \"required\", \"connected\", \"stale\", \"manual\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"not_required",
						"required",
						"connected",
						"stale",
						"manual",
					),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created_by": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"error": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"last_successful_sync": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"last_synced": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"modified_by": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"prompts": schema.ListAttribute{
				Computed:   true,
				CustomType: customfield.NewListType[customfield.Map[jsontypes.Normalized]](ctx),
				ElementType: types.MapType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
			"tools": schema.ListAttribute{
				Computed:   true,
				CustomType: customfield.NewListType[customfield.Map[jsontypes.Normalized]](ctx),
				ElementType: types.MapType{
					ElemType: jsontypes.NormalizedType{},
				},
			},
			"auth_config_summary": schema.SingleNestedAttribute{
				Description: "Safe subset of auth_credentials surfaced to the dashboard. Includes auth_mode (dcr|manual), has_client_secret, client_secret_version, and the OAuth endpoints + client_id for manual servers. Never includes the secret value.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryModel](ctx),
				Attributes: map[string]schema.Attribute{
					"auth_mode": schema.StringAttribute{
						Description: `Available values: "dcr", "manual".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("dcr", "manual"),
						},
					},
					"client_secret_version": schema.Float64Attribute{
						Computed: true,
					},
					"config": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryConfigModel](ctx),
						Attributes: map[string]schema.Attribute{
							"authorization_endpoint": schema.StringAttribute{
								Computed: true,
							},
							"issuer": schema.StringAttribute{
								Computed: true,
							},
							"resource": schema.StringAttribute{
								Computed: true,
							},
							"revocation_endpoint": schema.StringAttribute{
								Computed: true,
							},
							"token_endpoint": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"has_client_secret": schema.BoolAttribute{
						Computed: true,
					},
					"registration_info": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryRegistrationInfoModel](ctx),
						Attributes: map[string]schema.Attribute{
							"client_id": schema.StringAttribute{
								Computed: true,
							},
							"redirect_uris": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"scope": schema.StringAttribute{
								Computed: true,
							},
							"token_endpoint_auth_method": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
			"error_details": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustAccessAIControlsMcpServerErrorDetailsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cause": schema.StringAttribute{
						Description: "Underlying error message",
						Computed:    true,
					},
					"is_upstream": schema.BoolAttribute{
						Description: "True = MCP server returned an error. False = couldn't reach the server",
						Computed:    true,
					},
					"mcp_code": schema.Float64Attribute{
						Description: "MCP protocol error code",
						Computed:    true,
					},
					"retryable": schema.BoolAttribute{
						Description: "Whether the error is transient and worth retrying",
						Computed:    true,
					},
					"status_code": schema.Float64Attribute{
						Description: "HTTP status code from the server",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *ZeroTrustAccessAIControlsMcpServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessAIControlsMcpServerResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

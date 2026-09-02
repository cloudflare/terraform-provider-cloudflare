// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_server

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
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

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessAIControlsMcpServerDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"MCP Portals Read",
				"MCP Portals Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the MCP server.",
				Computed:    true,
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Optional:    true,
			},
			"auth_type": schema.StringAttribute{
				Description: "Authentication method used to connect to the upstream MCP server.\nAvailable values: \"oauth\", \"bearer\", \"unauthenticated\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"oauth",
						"bearer",
						"unauthenticated",
					),
				},
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
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created_by": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description of the MCP server.",
				Computed:    true,
			},
			"error": schema.StringAttribute{
				Computed: true,
			},
			"hostname": schema.StringAttribute{
				Description: "URL of the upstream MCP endpoint.",
				Computed:    true,
			},
			"is_shared_oauth_callback_enabled": schema.BoolAttribute{
				Description: "When true, the gateway worker uses the shared Cloudflare-owned OAuth callback endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the customer portal hostname. New public server creates default to true; existing servers default to false from migration until explicitly updated. Effective behavior is gated by the gateway worker's per-env rollout mode KV key.",
				Computed:    true,
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
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_by": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Description: "Display name for the MCP server.",
				Computed:    true,
			},
			"secure_web_gateway": schema.BoolAttribute{
				Description: "Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Current sync state of the server\nAvailable values: \"waiting\", \"ready\", \"stale\", \"error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"waiting",
						"ready",
						"stale",
						"error",
					),
				},
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
				CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryConfigDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[ZeroTrustAccessAIControlsMcpServerAuthConfigSummaryRegistrationInfoDataSourceModel](ctx),
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
				CustomType: customfield.NewNestedObjectType[ZeroTrustAccessAIControlsMcpServerErrorDetailsDataSourceModel](ctx),
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
			"updated_prompts": schema.ListNestedAttribute{
				Description: "Server-wide prompt capability overrides.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessAIControlsMcpServerUpdatedPromptsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name of the tool or prompt capability to override.",
							Computed:    true,
						},
						"alias": schema.StringAttribute{
							Description: "Custom name exposed for the capability.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Custom description exposed for the capability.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the capability is available through the MCP server.",
							Computed:    true,
						},
					},
				},
			},
			"updated_tools": schema.ListNestedAttribute{
				Description: "Server-wide tool capability overrides.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessAIControlsMcpServerUpdatedToolsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name of the tool or prompt capability to override.",
							Computed:    true,
						},
						"alias": schema.StringAttribute{
							Description: "Custom name exposed for the capability.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Custom description exposed for the capability.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the capability is available through the MCP server.",
							Computed:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"search": schema.StringAttribute{
						Description: "Search by id, name",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessAIControlsMcpServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessAIControlsMcpServerDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("filter")),
	}
}

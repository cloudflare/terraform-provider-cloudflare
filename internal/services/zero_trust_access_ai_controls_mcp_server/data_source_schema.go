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
				Description: "server id",
				Computed:    true,
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"auth_type": schema.StringAttribute{
				Description: `Available values: "oauth", "bearer", "unauthenticated".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"oauth",
						"bearer",
						"unauthenticated",
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
				Computed: true,
			},
			"error": schema.StringAttribute{
				Computed: true,
			},
			"hostname": schema.StringAttribute{
				Computed: true,
			},
			"is_shared_oauth_callback_enabled": schema.BoolAttribute{
				Description: "When true, the gateway worker uses the shared Cloudflare-owned OAuth callback endpoint as the redirect_uri for upstream on-behalf OAuth, instead of the customer portal hostname. Defaults to false (off); opt in per server by setting true. Effective behavior is gated by the gateway worker's per-env rollout mode KV key.",
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
				Computed: true,
			},
			"secure_web_gateway": schema.BoolAttribute{
				Description: "Route outbound traffic to this MCP server through Zero Trust Secure Web Gateway",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Computed: true,
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
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessAIControlsMcpServerUpdatedPromptsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed: true,
						},
						"alias": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
			"updated_tools": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessAIControlsMcpServerUpdatedToolsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed: true,
						},
						"alias": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
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

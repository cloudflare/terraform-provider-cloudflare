// This file is intentionally NOT code-generated.
// It implements the v500 → v501 state upgrade for cloudflare_zero_trust_access_ai_controls_mcp_portal,
// which migrates the `servers` attribute from a List to a Set.

package zero_trust_access_ai_controls_mcp_portal

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessAIControlsMcpPortalResource)(nil)

// zeroTrustAccessAIControlsMcpPortalModelV500 is the state shape for schema version 500.
// It is identical to ZeroTrustAccessAIControlsMcpPortalModel except that Servers is a
// NestedObjectList (the old representation) rather than a NestedObjectSet.
type zeroTrustAccessAIControlsMcpPortalModelV500 struct {
	ID               types.String                                                                 `tfsdk:"id"`
	AccountID        types.String                                                                 `tfsdk:"account_id"`
	Hostname         types.String                                                                 `tfsdk:"hostname"`
	Name             types.String                                                                 `tfsdk:"name"`
	Description      types.String                                                                 `tfsdk:"description"`
	AllowCodeMode    types.Bool                                                                   `tfsdk:"allow_code_mode"`
	SecureWebGateway types.Bool                                                                   `tfsdk:"secure_web_gateway"`
	Servers          customfield.NestedObjectList[ZeroTrustAccessAIControlsMcpPortalServersModel] `tfsdk:"servers"`
	CreatedAt        timetypes.RFC3339                                                            `tfsdk:"created_at"`
	CreatedBy        types.String                                                                 `tfsdk:"created_by"`
	ModifiedAt       timetypes.RFC3339                                                            `tfsdk:"modified_at"`
	ModifiedBy       types.String                                                                 `tfsdk:"modified_by"`
}

// resourceSchemaV500 returns the schema as it existed at version 500, where
// `servers` was a ListNestedAttribute. Used by the state upgrader to decode
// prior state before converting to the v501 SetNestedAttribute shape.
func resourceSchemaV500(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "portal id",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"allow_code_mode": schema.BoolAttribute{
				Description: "Allow remote code execution in Dynamic Workers (beta)",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"secure_web_gateway": schema.BoolAttribute{
				Description: "Route outbound MCP traffic through Zero Trust Secure Web Gateway",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"servers": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"server_id": schema.StringAttribute{
							Description: "server id",
							Required:    true,
						},
						"default_disabled": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(false),
						},
						"on_behalf": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							Default:  booldefault.StaticBool(true),
						},
						"updated_prompts": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Required: true,
									},
									"alias": schema.StringAttribute{
										Optional: true,
									},
									"description": schema.StringAttribute{
										Optional: true,
									},
									"enabled": schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
						"updated_tools": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Required: true,
									},
									"alias": schema.StringAttribute{
										Optional: true,
									},
									"description": schema.StringAttribute{
										Optional: true,
									},
									"enabled": schema.BoolAttribute{
										Optional: true,
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

// upgradeMcpPortalV500ToV501 converts a v500 state (servers as List) to v501
// (servers as Set). All scalar fields are copied through unchanged.
func upgradeMcpPortalV500ToV501(ctx context.Context, prior zeroTrustAccessAIControlsMcpPortalModelV500) (ZeroTrustAccessAIControlsMcpPortalModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	out := ZeroTrustAccessAIControlsMcpPortalModel{
		ID:               prior.ID,
		AccountID:        prior.AccountID,
		Hostname:         prior.Hostname,
		Name:             prior.Name,
		Description:      prior.Description,
		AllowCodeMode:    prior.AllowCodeMode,
		SecureWebGateway: prior.SecureWebGateway,
		CreatedAt:        prior.CreatedAt,
		CreatedBy:        prior.CreatedBy,
		ModifiedAt:       prior.ModifiedAt,
		ModifiedBy:       prior.ModifiedBy,
	}

	switch {
	case prior.Servers.IsNull():
		out.Servers = customfield.NullObjectSet[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx)
	case prior.Servers.IsUnknown():
		out.Servers = customfield.UnknownObjectSet[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx)
	default:
		elems := prior.Servers.Elements()
		set, d := customfield.NewObjectSetFromAttributes[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx, elems)
		diags.Append(d...)
		out.Servers = set
	}

	return out, diags
}

func (r *ZeroTrustAccessAIControlsMcpPortalResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	priorSchema := resourceSchemaV500(ctx)
	return map[int64]resource.StateUpgrader{
		500: {
			PriorSchema: &priorSchema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior zeroTrustAccessAIControlsMcpPortalModelV500
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeMcpPortalV500ToV501(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

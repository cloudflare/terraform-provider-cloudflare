// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_ai_controls_mcp_portal

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessAIControlsMcpPortalResource)(nil)

// zeroTrustAccessAIControlsMcpPortalModelV500 mirrors the pre-501 state shape,
// where servers was modeled as a List instead of a Set. Only the Servers field
// type differs from the current model; every other attribute is unchanged.
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

func (r *ZeroTrustAccessAIControlsMcpPortalResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	priorSchema := resourceSchemaV500(ctx)
	upgrader := resource.StateUpgrader{
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

			resp.Diagnostics.Append(resp.State.Set(ctx, upgraded)...)
		},
	}

	// The resource only ever shipped at schema version 500 (servers as a List),
	// so 500 is the real prior version. The 0 entry is kept for safety: any such
	// state is list-shaped too and converts identically.
	return map[int64]resource.StateUpgrader{
		0:   upgrader,
		500: upgrader,
	}
}

// upgradeMcpPortalV500ToV501 converts a pre-501 portal state (servers as a List)
// to the current model (servers as a Set). It is a pure function so it can be
// unit-tested without constructing a tfsdk.State.
func upgradeMcpPortalV500ToV501(
	ctx context.Context,
	prior zeroTrustAccessAIControlsMcpPortalModelV500,
) (ZeroTrustAccessAIControlsMcpPortalModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	upgraded := ZeroTrustAccessAIControlsMcpPortalModel{
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
		upgraded.Servers = customfield.NullObjectSet[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx)
	case prior.Servers.IsUnknown():
		upgraded.Servers = customfield.UnknownObjectSet[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx)
	default:
		servers, d := prior.Servers.AsStructSliceT(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return upgraded, diags
		}
		set, d := customfield.NewObjectSet(ctx, servers)
		diags.Append(d...)
		if diags.HasError() {
			return upgraded, diags
		}
		upgraded.Servers = set
	}

	return upgraded, diags
}

// resourceSchemaV500 returns the pre-501 resource schema, where servers was a
// ListNestedAttribute. Used as the PriorSchema when upgrading state to 501. It
// reuses the current nested-object definition so the two stay in sync.
func resourceSchemaV500(ctx context.Context) schema.Schema {
	s := ResourceSchema(ctx)
	s.Version = 500

	if set, ok := s.Attributes["servers"].(schema.SetNestedAttribute); ok {
		s.Attributes["servers"] = schema.ListNestedAttribute{
			Computed:     true,
			Optional:     true,
			CustomType:   customfield.NewNestedObjectListType[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx),
			NestedObject: set.NestedObject,
		}
	}

	return s
}

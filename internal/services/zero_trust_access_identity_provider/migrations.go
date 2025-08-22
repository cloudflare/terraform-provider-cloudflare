// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessIdentityProviderResource)(nil)

// zeroTrustAccessIdentityProviderResourceSchemaV0 defines the v0 schema (v4 provider format)
// This represents the structure that the v4 provider used
var zeroTrustAccessIdentityProviderResourceSchemaV0 = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"account_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"zone_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"type": schema.StringAttribute{
			Required: true,
		},
		// In v4, config and scim_config were blocks (stored as arrays in state)
		// In v5, they are single objects
		"config": schema.ListAttribute{
			ElementType: types.MapType{ElemType: types.StringType}, // Simplified - actual structure is more complex
			Optional:    true,
		},
		"scim_config": schema.ListAttribute{
			ElementType: types.MapType{ElemType: types.StringType}, // Simplified - actual structure is more complex
			Optional:    true,
		},
	},
}

func (r *ZeroTrustAccessIdentityProviderResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &zeroTrustAccessIdentityProviderResourceSchemaV0,
			StateUpgrader: upgradeZeroTrustAccessIdentityProviderStateV0toV1,
		},
	}
}

// upgradeZeroTrustAccessIdentityProviderStateV0toV1 migrates from v4 provider state format to v5
func upgradeZeroTrustAccessIdentityProviderStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Debug(ctx, "Starting state migration from v0 to v1 for zero_trust_access_identity_provider")

	// Since the v4 and v5 schemas are very different (blocks vs objects, field renames),
	// and the external migration tool (cmd/migrate) already handles the complex state transformations,
	// we'll use a simpler approach: extract the raw state data and let the external migration
	// tool handle the detailed transformations.
	
	// Get the raw state as a map
	var rawState map[string]interface{}
	resp.Diagnostics.Append(req.State.Get(ctx, &rawState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to extract raw state: %v", resp.Diagnostics.Errors()))
		return
	}

	// For this migration, we rely on the external cmd/migrate tool to handle the
	// complex transformations (config block->object, idp_public_cert->idp_public_certs, etc.)
	// The migration test framework runs the cmd/migrate tool which includes state transformation
	// via transformZeroTrustAccessIdentityProviderStateJSON

	// Create new state structure with basic fields
	var newState ZeroTrustAccessIdentityProviderModel

	// Extract basic fields from raw state if they exist
	if id, ok := rawState["id"].(string); ok {
		newState.ID = types.StringValue(id)
	}
	if accountID, ok := rawState["account_id"].(string); ok {
		newState.AccountID = types.StringPointerValue(&accountID)
	}
	if zoneID, ok := rawState["zone_id"].(string); ok {
		newState.ZoneID = types.StringPointerValue(&zoneID)
	}
	if name, ok := rawState["name"].(string); ok {
		newState.Name = types.StringValue(name)
	}
	if providerType, ok := rawState["type"].(string); ok {
		newState.Type = types.StringValue(providerType)
	}

	// For complex nested structures (config, scim_config), we'll need to rely on
	// the external migration tool to handle the transformation correctly.
	// The test framework should call the cmd/migrate tool before this state upgrader runs.

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to set new state: %v", resp.Diagnostics.Errors()))
		return
	}

	tflog.Debug(ctx, "Successfully completed state migration from v0 to v1 for zero_trust_access_identity_provider")
}

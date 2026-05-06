package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v0 to v500.
// This is triggered for resources that were created with schema version 0
// (either from the old cloudflare_access_group or early cloudflare_zero_trust_access_group).
//
// NOTE: When called from the parent package dispatcher (PriorSchema=nil), req.State is empty
// and we must unmarshal from req.RawState using the v4 source schema directly.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil for schema version 0 migration")
		return
	}

	tflog.Info(ctx, "Upgrading zero_trust_access_group state from v4 SDKv2 provider (schema_version=0)")

	// Parse source state (v4 SDKv2 format) from RawState using v4 schema
	var sourceState SourceV4ZeroTrustAccessGroupModel
	resp.Diagnostics.Append(unmarshalV4State(ctx, req.RawState, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target state (v5 Plugin Framework format)
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// unmarshalV4State decodes RawState bytes into the v4 source model using the v4 source schema.
// This is necessary when PriorSchema is nil (dispatcher pattern) so req.State isn't pre-populated.
func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceV4ZeroTrustAccessGroupModel) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourceV4ZeroTrustAccessGroupSchema()
	sourceType := sourceSchema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(sourceType)
	if err != nil {
		diags.AddError("Failed to unmarshal v4 state", "Could not parse raw state as v4 format: "+err.Error())
		return diags
	}

	state := tfsdk.State{Raw: rawValue, Schema: sourceSchema}
	diags.Append(state.Get(ctx, target)...)
	return diags
}

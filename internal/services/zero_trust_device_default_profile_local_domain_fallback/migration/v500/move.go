package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveFallbackDomainToDefaultProfile handles resource renames from legacy fallback domain resources
// (without policy_id) to the default profile resource.
//
// This function is called by Terraform when it encounters a `moved` block that references
// the old resource type (cloudflare_zero_trust_local_fallback_domain or cloudflare_fallback_domain)
// and determines the resource should migrate to cloudflare_zero_trust_device_default_profile_local_domain_fallback
// (i.e., resources WITHOUT policy_id or with policy_id=null).
//
// Source resources:
// - cloudflare_zero_trust_local_fallback_domain (current v4 name)
// - cloudflare_fallback_domain (deprecated v4 name)
//
// Target resource:
// - cloudflare_zero_trust_device_default_profile_local_domain_fallback (v5)
func MoveFallbackDomainToDefaultProfile(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Validate source type name
	if req.SourceTypeName != "cloudflare_zero_trust_local_fallback_domain" && req.SourceTypeName != "cloudflare_fallback_domain" {
		tflog.Warn(ctx, "MoveFallbackDomainToDefaultProfile called with unexpected source type",
			map[string]any{
				"source_type":          req.SourceTypeName,
				"expected_types":       []string{"cloudflare_zero_trust_local_fallback_domain", "cloudflare_fallback_domain"},
				"target_resource_type": "cloudflare_zero_trust_device_default_profile_local_domain_fallback",
			})
		// Don't handle this move - let Terraform try other handlers or fail
		return
	}

	tflog.Info(ctx, "Moving state from legacy fallback domain to default profile",
		map[string]any{
			"source_type": req.SourceTypeName,
		})

	// Parse the source state (legacy v4 format)
	var sourceState SourceCloudflareFallbackDomainModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse source state", map[string]any{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	// Validate that policy_id is null or absent (default profile requirement)
	// If policy_id is present and not null, this should have migrated to custom profile instead
	if !sourceState.PolicyID.IsNull() && !sourceState.PolicyID.IsUnknown() {
		resp.Diagnostics.AddError(
			"Invalid migration path",
			"This resource has a policy_id and should migrate to cloudflare_zero_trust_device_custom_profile_local_domain_fallback instead. "+
				"The moved block is pointing to the wrong target resource type. "+
				"Please update the moved block to use cloudflare_zero_trust_device_custom_profile_local_domain_fallback as the target.",
		)
		tflog.Error(ctx, "Resource has policy_id but is trying to migrate to default profile", map[string]any{
			"policy_id": sourceState.PolicyID.ValueString(),
		})
		return
	}

	tflog.Debug(ctx, "Source state validation passed - no policy_id present",
		map[string]any{
			"account_id": sourceState.AccountID.ValueString(),
		})

	// Transform to target (current v5 default profile format)
	targetState, diags := TransformToDefaultProfile(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to transform state", map[string]any{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	// Set the moved state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set target state", map[string]any{
			"diagnostics": resp.Diagnostics,
		})
	} else {
		tflog.Info(ctx, "State move to default profile completed successfully")
	}
}

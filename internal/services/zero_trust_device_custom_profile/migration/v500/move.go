// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveDeviceProfilesToCustomProfile handles resource renames from legacy device profile resources
// to the custom profile resource.
//
// This function is called by Terraform when it encounters a `moved` block that references
// the old resource type (cloudflare_zero_trust_device_profiles or cloudflare_device_settings_policy)
// and determines the resource should migrate to cloudflare_zero_trust_device_custom_profile
// (i.e., resources WITH match AND precedence).
//
// Source resources:
// - cloudflare_zero_trust_device_profiles (current v4 name)
// - cloudflare_device_settings_policy (deprecated v4 name)
//
// Target resource:
// - cloudflare_zero_trust_device_custom_profile (v5)
func MoveDeviceProfilesToCustomProfile(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Validate source type name
	if req.SourceTypeName != "cloudflare_zero_trust_device_profiles" && req.SourceTypeName != "cloudflare_device_settings_policy" {
		tflog.Warn(ctx, "MoveDeviceProfilesToCustomProfile called with unexpected source type",
			map[string]any{
				"source_type":          req.SourceTypeName,
				"expected_types":       []string{"cloudflare_zero_trust_device_profiles", "cloudflare_device_settings_policy"},
				"target_resource_type": "cloudflare_zero_trust_device_custom_profile",
			})
		// Don't handle this move - let Terraform try other handlers or fail
		return
	}

	tflog.Info(ctx, "Moving state from legacy device profiles to custom profile",
		map[string]any{
			"source_type": req.SourceTypeName,
		})

	// Read the source state using the v4 schema
	var source SourceDeviceProfileModel
	diags := req.SourceState.Get(ctx, &source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to read source state", map[string]any{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	// Transform the source state to the target custom profile structure
	target, transformDiags := TransformToCustomProfile(ctx, source)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to transform state", map[string]any{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	// Set the transformed state into the target state
	diags = resp.TargetState.Set(ctx, target)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set target state", map[string]any{
			"diagnostics": resp.Diagnostics,
		})
	} else {
		tflog.Info(ctx, "State move to custom profile completed successfully")
	}
}

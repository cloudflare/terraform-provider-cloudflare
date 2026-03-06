// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// MoveDeviceProfilesToDefaultProfile handles resource renames from legacy device profile resources
// to the default profile resource.
//
// This function is called by Terraform when it encounters a `moved` block that references
// the old resource type (cloudflare_zero_trust_device_profiles or cloudflare_device_settings_policy)
// and determines the resource should migrate to cloudflare_zero_trust_device_default_profile
// (i.e., resources WITHOUT match+precedence or WITH default=true).
//
// Source resources:
// - cloudflare_zero_trust_device_profiles (current v4 name)
// - cloudflare_device_settings_policy (deprecated v4 name)
//
// Target resource:
// - cloudflare_zero_trust_device_default_profile (v5)
func MoveDeviceProfilesToDefaultProfile(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Read the source state using the v4 schema
	var source SourceDeviceProfileModel
	diags := req.SourceState.Get(ctx, &source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform the source state to the target default profile structure
	target, transformDiags := TransformToDefaultProfile(ctx, source)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the transformed state into the target state
	diags = resp.TargetState.Set(ctx, target)
	resp.Diagnostics.Append(diags...)
}

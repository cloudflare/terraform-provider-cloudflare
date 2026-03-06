package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveFallbackDomainToCustomProfile handles moving state from legacy fallback domain resources
// to cloudflare_zero_trust_device_custom_profile_local_domain_fallback.
//
// This is triggered by Terraform 1.8+ when it encounters a `moved` block:
//
//   moved {
//     from = cloudflare_zero_trust_local_fallback_domain.example
//     to   = cloudflare_zero_trust_device_custom_profile_local_domain_fallback.example
//   }
//
// or for the deprecated alias:
//
//   moved {
//     from = cloudflare_fallback_domain.example
//     to   = cloudflare_zero_trust_device_custom_profile_local_domain_fallback.example
//   }
//
// This handler is called when:
// - The legacy resource has a policy_id (custom profile path)
// - tf-migrate generates moved blocks pointing to this resource for states with policy_id
//
// For states WITHOUT policy_id, the moved block points to the default profile resource instead.
func MoveFallbackDomainToCustomProfile(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	// Validate source type name
	if req.SourceTypeName != "cloudflare_zero_trust_local_fallback_domain" && req.SourceTypeName != "cloudflare_fallback_domain" {
		tflog.Warn(ctx, "MoveFallbackDomainToCustomProfile called with unexpected source type",
			map[string]any{
				"source_type":           req.SourceTypeName,
				"expected_types":        []string{"cloudflare_zero_trust_local_fallback_domain", "cloudflare_fallback_domain"},
				"target_resource_type":  "cloudflare_zero_trust_device_custom_profile_local_domain_fallback",
			})
		// Don't handle this move - let Terraform try other handlers or fail
		return
	}

	tflog.Info(ctx, "Moving state from legacy fallback domain to custom profile",
		map[string]any{
			"source_type": req.SourceTypeName,
		})

	// Parse the source state (legacy v4 format)
	var sourceState SourceCloudflareFallbackDomainModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target (current v5 custom profile format)
	targetState, diags := TransformToCustomProfile(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the moved state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, targetState)...)

	tflog.Info(ctx, "State move to custom profile completed successfully")
}

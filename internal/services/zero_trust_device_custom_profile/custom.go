package zero_trust_device_custom_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func (r *ZeroTrustDeviceCustomProfileResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Nothing to modify on destroy, nothing to preserve from state on create.
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

	var plan ZeroTrustDeviceCustomProfileModel
	var state ZeroTrustDeviceCustomProfileModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve include from state when plan is unknown. The framework's
	// UseStateForUnknown skips null state, but null is a valid value for
	// include (it means "not set, exclude is used instead").
	if plan.Include.IsUnknown() {
		plan.Include = state.Include
	}

	// Same for target_tests: purely computed, null is valid (no DEX tests).
	if plan.TargetTests.IsUnknown() {
		plan.TargetTests = state.TargetTests
	}

	// Same for fallback_domains: purely computed.
	if plan.FallbackDomains.IsUnknown() {
		plan.FallbackDomains = state.FallbackDomains
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

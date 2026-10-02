package zero_trust_device_custom_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func (r *ZeroTrustDeviceCustomProfileResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Nothing to modify on destroy.
	if req.Plan.Raw.IsNull() {
		return
	}

	// Nothing to preserve from state on create.
	if req.State.Raw.IsNull() {
		// On create, normalize unknown nested attributes in exclude/include to null
		// when the user didn't set them (config is null). This prevents cosmetic
		// "(known after apply)" noise on the initial plan.
		var plan ZeroTrustDeviceCustomProfileModel
		var diags diag.Diagnostics
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}

		plan.Exclude, diags = normalizeSplitTunnelList(ctx, plan.Exclude, customfield.NullObjectList[ZeroTrustDeviceCustomProfileExcludeModel](ctx))
		resp.Diagnostics.Append(diags...)
		plan.Include, diags = normalizeSplitTunnelList(ctx, plan.Include, customfield.NullObjectList[ZeroTrustDeviceCustomProfileIncludeModel](ctx))
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
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

	// Normalize unknown nested attributes in exclude/include entries.
	// When a user specifies only `address` on an exclude entry, the `host`
	// and `description` attributes become unknown at plan time because they
	// are Optional+Computed. We resolve them to their state values (for
	// existing entries) or null (for new entries) since the API does not
	// compute these fields.
	var diags diag.Diagnostics
	plan.Exclude, diags = normalizeSplitTunnelList(ctx, plan.Exclude, state.Exclude)
	resp.Diagnostics.Append(diags...)
	plan.Include, diags = normalizeSplitTunnelList(ctx, plan.Include, state.Include)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

// splitTunnelEntry is a constraint interface for exclude/include entry models.
type splitTunnelEntry interface {
	ZeroTrustDeviceCustomProfileExcludeModel | ZeroTrustDeviceCustomProfileIncludeModel
}

// splitTunnelMutuallyExclusive maps each field to the field it cannot coexist
// with in a single split-tunnel entry. The API returns error 2049 when both
// are non-null in the same object.
var splitTunnelMutuallyExclusive = map[string]string{
	"address": "host",
	"host":    "address",
}

// normalizeSplitTunnelList resolves unknown nested attributes in exclude/include
// list elements. Matching is index-based (plan[i] ↔ state[i]). For unknown
// fields: if the mutually exclusive counterpart (address↔host) is already set
// in the plan entry, resolve to null to avoid API error 2049; otherwise fall
// back to the state value (or null for new entries beyond state length).
func normalizeSplitTunnelList[T splitTunnelEntry](
	ctx context.Context,
	planList customfield.NestedObjectList[T],
	stateList customfield.NestedObjectList[T],
) (customfield.NestedObjectList[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	if planList.IsNull() || planList.IsUnknown() {
		return planList, diags
	}

	planElements := planList.Elements()
	if len(planElements) == 0 {
		return planList, diags
	}

	stateElements := []attr.Value{}
	if !stateList.IsNull() && !stateList.IsUnknown() {
		stateElements = stateList.Elements()
	}

	changed := false
	normalized := make([]attr.Value, len(planElements))

	for i, elem := range planElements {
		obj, ok := elem.(types.Object)
		if !ok || obj.IsNull() || obj.IsUnknown() {
			normalized[i] = elem
			continue
		}

		attrs := obj.Attributes()
		newAttrs := make(map[string]attr.Value, len(attrs))
		elementChanged := false

		// Get corresponding state element attributes if available (index-based).
		var stateAttrs map[string]attr.Value
		if i < len(stateElements) {
			if stateObj, ok := stateElements[i].(types.Object); ok && !stateObj.IsNull() && !stateObj.IsUnknown() {
				stateAttrs = stateObj.Attributes()
			}
		}

		for key, val := range attrs {
			if val.IsUnknown() {
				elementChanged = true
				// Mutually exclusive counterpart already set — resolve to null.
				if counterpart, isMutex := splitTunnelMutuallyExclusive[key]; isMutex {
					if cv, exists := attrs[counterpart]; exists && !cv.IsNull() && !cv.IsUnknown() {
						newAttrs[key] = basetypes.NewStringNull()
						continue
					}
				}
				// Try to use state value; fall back to null.
				// NOTE: all nested attributes in exclude/include models are strings;
				// update this fallback if non-string Optional+Computed fields are added.
				if stateAttrs != nil {
					if sv, exists := stateAttrs[key]; exists {
						newAttrs[key] = sv
					} else {
						newAttrs[key] = basetypes.NewStringNull()
					}
				} else {
					newAttrs[key] = basetypes.NewStringNull()
				}
			} else {
				newAttrs[key] = val
			}
		}

		if elementChanged {
			changed = true
			newObj, d := types.ObjectValue(obj.AttributeTypes(ctx), newAttrs)
			diags.Append(d...)
			if d.HasError() {
				// If we can't construct the normalized object, keep original.
				normalized[i] = elem
			} else {
				normalized[i] = newObj
			}
		} else {
			normalized[i] = elem
		}
	}

	if !changed {
		return planList, diags
	}

	result, d := customfield.NewObjectListFromAttributes[T](ctx, normalized)
	diags.Append(d...)
	if d.HasError() {
		return planList, diags
	}
	return result, diags
}

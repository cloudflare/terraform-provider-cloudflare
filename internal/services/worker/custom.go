package worker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NormalizeFloat64() planmodifier.Float64 {
	return normalizeFloat64Modifier{}
}

type normalizeFloat64Modifier struct{}

func (m normalizeFloat64Modifier) Description(_ context.Context) string {
	return "Normalize float64 values to prevent spurious changes"
}

func (m normalizeFloat64Modifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// For some reason sometimes req.StateValue.Equal(req.PlanValue) is false, but
// req.StateValue.ValueFloat64() == req.PlanValue.ValueFloat64() is true. This
// appears to be caused by a difference in how float64 values are parsed in
// different places. This plan modifier normalizes the float64 value in the plan
// such that req.StateValue.Equal(req.PlanValue) is true in this case.
func (m normalizeFloat64Modifier) PlanModifyFloat64(ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response) {
	if req.PlanValue.IsUnknown() || req.PlanValue.IsNull() {
		return
	}
	resp.PlanValue = types.Float64Value(req.PlanValue.ValueFloat64())
}

func DefaultSubdomainPreviewsEnabledToEnabled() planmodifier.Bool {
	return defaultSubdomainPreviewsEnabledToEnabledModifier{}
}

type defaultSubdomainPreviewsEnabledToEnabledModifier struct{}

func (m defaultSubdomainPreviewsEnabledToEnabledModifier) Description(_ context.Context) string {
	return "Defaults subdomain.previews_enabled to the value of subdomain.enabled if subdomain.previews_enabled is not explicitly set"
}

func (m defaultSubdomainPreviewsEnabledToEnabledModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// PlanModifyBool sets subdomain.previews_enabled to the value of
// subdomain.enabled when subdomain.previews_enabled is null in the config.
func (m defaultSubdomainPreviewsEnabledToEnabledModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	// If the config value is not null, the user has explicitly set it, so don't modify
	if !req.ConfigValue.IsNull() {
		return
	}

	// If we're destroying, don't modify
	if req.Plan.Raw.IsNull() {
		return
	}

	// Get the sibling "enabled" attribute value
	enabledPath := req.Path.ParentPath().AtName("enabled")
	var enabled types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, enabledPath, &enabled)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If subdomain.enabled is null or unknown, we can't use it as a default
	if enabled.IsNull() || enabled.IsUnknown() {
		// Fall back to false as the default
		resp.PlanValue = types.BoolValue(false)
		return
	}

	// Set subdomain.previews_enabled to the value of subdomain.enabled
	resp.PlanValue = enabled
}

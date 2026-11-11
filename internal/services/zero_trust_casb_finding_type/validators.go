package zero_trust_casb_finding_type

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// modifyIntegrationScopePlan mirrors posture-api's integration scope rule at plan time:
// integration_ids must be non-empty when applies_to_all_integrations is false, and empty
// when it is true.
//
// integration_ids is optional and computed, and its plan modifier keeps the IDs from state
// when it is omitted from configuration. That is what lets an existing scoped finding type
// omit the attribute, but it would also carry stale IDs into a plan that switches to all
// integrations, which the API rejects. Such plans mark the IDs unknown instead, so they are
// left out of the request and read back from the response.
func modifyIntegrationScopePlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	// Destroy plans have no configuration to validate.
	if req.Plan.Raw.IsNull() {
		return
	}

	var appliesToAll types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("applies_to_all_integrations"), &appliesToAll)...)

	var configIDs customfield.List[types.String]
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("integration_ids"), &configIDs)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if appliesToAllWithIDs(appliesToAll, configIDs) {
		resp.Diagnostics.AddAttributeError(
			path.Root("integration_ids"),
			"Unexpected integration IDs",
			"integration_ids must be omitted or empty when applies_to_all_integrations is true.",
		)
		return
	}

	var plannedIDs customfield.List[types.String]
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("integration_ids"), &plannedIDs)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if preservedIDsForAllIntegrations(appliesToAll, configIDs, plannedIDs) {
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("integration_ids"),
			customfield.UnknownList[types.String](ctx))...)
		return
	}

	// On create there is no prior state, so configuration is authoritative.
	// On update the resolved plan already carries any preserved IDs; when state has none
	// (for example a finding type that applied to all integrations), the plan leaves them
	// unknown and configuration is authoritative again.
	integrationIDs := plannedIDs
	if req.State.Raw.IsNull() || plannedIDs.IsUnknown() {
		integrationIDs = configIDs
	}

	if !integrationScopeInvalid(appliesToAll, integrationIDs) {
		return
	}

	resp.Diagnostics.AddAttributeError(
		path.Root("integration_ids"),
		"Missing integration IDs",
		"integration_ids must contain at least one value when applies_to_all_integrations is false.",
	)
}

func integrationScopeInvalid(appliesToAll types.Bool, integrationIDs customfield.List[types.String]) bool {
	if appliesToAll.IsNull() || appliesToAll.IsUnknown() || integrationIDs.IsUnknown() {
		return false
	}

	return !appliesToAll.ValueBool() && (integrationIDs.IsNull() || len(integrationIDs.Elements()) == 0)
}

// preservedIDsForAllIntegrations reports whether a plan for all integrations still carries
// integration IDs that came from state rather than configuration.
func preservedIDsForAllIntegrations(
	appliesToAll types.Bool,
	configIDs customfield.List[types.String],
	plannedIDs customfield.List[types.String],
) bool {
	if appliesToAll.IsNull() || appliesToAll.IsUnknown() || !appliesToAll.ValueBool() || !configIDs.IsNull() {
		return false
	}

	return !plannedIDs.IsNull() && !plannedIDs.IsUnknown() && len(plannedIDs.Elements()) > 0
}

func appliesToAllWithIDs(appliesToAll types.Bool, configIDs customfield.List[types.String]) bool {
	if appliesToAll.IsNull() || appliesToAll.IsUnknown() || configIDs.IsNull() || configIDs.IsUnknown() {
		return false
	}

	return appliesToAll.ValueBool() && len(configIDs.Elements()) > 0
}

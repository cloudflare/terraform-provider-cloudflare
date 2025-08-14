package zero_trust_access_identity_provider

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, res *resource.ModifyPlanResponse) {
	var planApp, stateApp *ZeroTrustAccessIdentityProviderModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &planApp)...)
	res.Diagnostics.Append(req.State.Get(ctx, &stateApp)...)
	if res.Diagnostics.HasError() || planApp == nil {
		return
	}

	// Handle SCIM secret behavior
	if !planApp.SCIMConfig.IsNull() {
		planModel, diag := planApp.SCIMConfig.Value(ctx)
		if diag.HasError() || planModel == nil {
			return
		}
		
		// Check if SCIM config exists in the current state (not the plan)
		if stateApp != nil && !stateApp.SCIMConfig.IsNull() {
			// SCIM config already exists in state - preserve existing secret behavior
			stateModel, diag := stateApp.SCIMConfig.Value(ctx)
			if !diag.HasError() && stateModel != nil {
				if !stateModel.Secret.IsNull() && !stateModel.Secret.IsUnknown() {
					// Secret exists in state - preserve it
					planModel.Secret = stateModel.Secret
				} else {
					// Secret is null in state (e.g., after import) - keep it null to avoid unwanted diffs
					planModel.Secret = types.StringNull()
				}
			}
		} else {
			// No SCIM config in state OR state is null - this is first time enabling SCIM, secret will be generated
			planModel.Secret = types.StringUnknown()
		}
		
		planApp.SCIMConfig, diag = customfield.NewObject(ctx, planModel)
		if diag.HasError() {
			res.Diagnostics.Append(diag...)
			return
		}
	}

	res.Diagnostics.Append(res.Plan.Set(ctx, planApp)...)
}

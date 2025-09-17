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

	// Handle scim_config.secret based on different scenarios
	if stateApp != nil && !stateApp.SCIMConfig.IsNull() && !planApp.SCIMConfig.IsNull() {
		stateModel, _ := stateApp.SCIMConfig.Value(ctx)
		planModel, _ := planApp.SCIMConfig.Value(ctx)

		// Check if enabled is changing from false/null to true - regenerate secret
		enabledChangingToTrue := (stateModel.Enabled.IsNull() || !stateModel.Enabled.ValueBool()) && 
			planModel.Enabled.ValueBool()

		if enabledChangingToTrue {
			// Set secret to unknown when enabling SCIM, so it gets computed by the API
			planModel.Secret = types.StringUnknown()
		} else {
			// Preserve existing secret value to prevent unnecessary diffs
			planModel.Secret = stateModel.Secret
		}

		planApp.SCIMConfig, _ = customfield.NewObject(ctx, planModel)
	}

	res.Plan.Set(ctx, &planApp)
}

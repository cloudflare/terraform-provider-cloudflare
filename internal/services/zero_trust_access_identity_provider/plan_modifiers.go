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
		
		if stateApp != nil && !stateApp.SCIMConfig.IsNull() {
			// SCIM was already enabled - preserve existing secret from state
			stateModel, diag := stateApp.SCIMConfig.Value(ctx)
			if !diag.HasError() && stateModel != nil && !stateModel.Secret.IsNull() && !stateModel.Secret.IsUnknown() {
				planModel.Secret = stateModel.Secret
			}
		} else {
			// SCIM is being enabled for the first time - secret will be generated
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

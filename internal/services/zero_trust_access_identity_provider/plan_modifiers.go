package zero_trust_access_identity_provider

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, res *resource.ModifyPlanResponse) {
	var planApp, stateApp *ZeroTrustAccessIdentityProviderModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &planApp)...)
	res.Diagnostics.Append(req.State.Get(ctx, &stateApp)...)
	if res.Diagnostics.HasError() || planApp == nil {
		return
	}

	// Secret value is computed from create but does not change. Set to state value to not show changes in plan
	if stateApp != nil && !stateApp.SCIMConfig.IsNull() && !planApp.SCIMConfig.IsNull() {
		stateModel, _ := stateApp.SCIMConfig.Value(ctx)
		planModel, _ := planApp.SCIMConfig.Value(ctx)

		planModel.Secret = stateModel.Secret
		planApp.SCIMConfig, _ = customfield.NewObject(ctx, planModel)
	}

	res.Plan.Set(ctx, &planApp)
}

package zero_trust_access_service_token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, res *resource.ModifyPlanResponse) {
	var planApp, stateApp *ZeroTrustAccessServiceTokenModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &planApp)...)
	res.Diagnostics.Append(req.State.Get(ctx, &stateApp)...)
	if res.Diagnostics.HasError() || planApp == nil {
		return
	}

	// Client secret can't ever be updated - it's only set when the token is first created.
	// Tell TF to set the value to Null if it's unknown.
	if stateApp != nil && stateApp.ClientSecret.IsNull() && planApp.ClientSecret.IsUnknown() {
		planApp.ClientSecret = stateApp.ClientSecret
	}

	res.Plan.Set(ctx, &planApp)
}

package zero_trust_access_service_token

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, res *resource.ModifyPlanResponse) {
	var planApp, stateApp *ZeroTrustAccessServiceTokenModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &planApp)...)
	res.Diagnostics.Append(req.State.Get(ctx, &stateApp)...)
	if res.Diagnostics.HasError() || planApp == nil {
		return
	}

	// If ClientSecretVersion changes, then ClientSecret will change.
	if stateApp != nil && !stateApp.ClientSecretVersion.Equal(planApp.ClientSecretVersion) {
		planApp.ClientSecret = types.StringUnknown()
	} else if stateApp != nil && stateApp.ClientSecret.IsNull() && planApp.ClientSecret.IsUnknown() {
		// Otherwise maintain the ClientSecret value set in state since API won't return it.
		planApp.ClientSecret = stateApp.ClientSecret
	}

	res.Plan.Set(ctx, &planApp)
}

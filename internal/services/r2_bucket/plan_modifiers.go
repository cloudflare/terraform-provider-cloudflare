package r2_bucket

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, res *resource.ModifyPlanResponse) {
	var planApp, stateApp *R2BucketModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &planApp)...)
	res.Diagnostics.Append(req.State.Get(ctx, &stateApp)...)
	if res.Diagnostics.HasError() || planApp == nil {
		return
	}
	if stateApp != nil && !planApp.Location.IsNull() && !stateApp.Location.IsNull() &&
		!planApp.Location.IsUnknown() && !stateApp.Location.IsUnknown() {
		planAppLocation := planApp.Location.ValueString()
		stateAppLocation := stateApp.Location.ValueString()
		if strings.EqualFold(planAppLocation, stateAppLocation) {
			planApp.Location = stateApp.Location
		}
	}

	res.Diagnostics.Append(res.Plan.Set(ctx, &planApp)...)
}

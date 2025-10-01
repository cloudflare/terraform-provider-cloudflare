package r2_bucket

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, res *resource.ModifyPlanResponse) {
	var planApp, stateApp *R2BucketModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &planApp)...)
	res.Diagnostics.Append(req.State.Get(ctx, &stateApp)...)
	if res.Diagnostics.HasError() || planApp == nil {
		return
	}

	// Handle case-insensitive location comparison
	if stateApp != nil && !planApp.Location.IsNull() && !stateApp.Location.IsNull() &&
		!planApp.Location.IsUnknown() && !stateApp.Location.IsUnknown() {
		planAppLocation := planApp.Location.ValueString()
		stateAppLocation := stateApp.Location.ValueString()
		// If locations match case-insensitively, preserve the state's case
		if strings.EqualFold(planAppLocation, stateAppLocation) {
			res.Diagnostics.Append(res.Plan.SetAttribute(ctx, path.Root("location"), stateApp.Location)...)
		}
	}

	if stateApp != nil {
		// Preserve creation_date from state to avoid drift detection
		res.Diagnostics.Append(res.Plan.SetAttribute(ctx, path.Root("creation_date"), stateApp.CreationDate)...)
	}
}

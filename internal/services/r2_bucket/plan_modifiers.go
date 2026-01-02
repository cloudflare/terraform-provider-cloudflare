package r2_bucket

import (
	"context"

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

	// Location of a bucket cannot be changed after it is already created
	if stateApp != nil && !stateApp.Location.IsNull() && !stateApp.Location.IsUnknown() {
		res.Diagnostics.Append(res.Plan.SetAttribute(ctx, path.Root("location"), stateApp.Location)...)
	}

	if stateApp != nil {
		// Preserve creation_date from state to avoid drift detection
		res.Diagnostics.Append(res.Plan.SetAttribute(ctx, path.Root("creation_date"), stateApp.CreationDate)...)
	}
}

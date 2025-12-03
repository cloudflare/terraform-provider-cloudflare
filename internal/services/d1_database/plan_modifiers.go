package d1_database

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, res *resource.ModifyPlanResponse) {
	var plan, state *D1DatabaseModel
	res.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	res.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if res.Diagnostics.HasError() || plan == nil {
		return
	}

	// read_replication can only be set during updates, not during creation
	// Only error if the user explicitly set it (not null and not unknown)
	if state == nil && !plan.ReadReplication.IsNull() && !plan.ReadReplication.IsUnknown() {
		res.Diagnostics.AddAttributeError(
			path.Root("read_replication"),
			"Invalid Attribute Configuration",
			"read_replication can only be configured during updates, not during initial resource creation. Please remove this attribute from the initial configuration and add it in a subsequent update.",
		)
		return
	}

	// If this is an update and read_replication was removed from config (plan is null)
	// but exists in state (state is not null), set it to disabled, we can revisit this to err if we want 
	// the users to do this on another step instead
	if state != nil && plan.ReadReplication.IsNull() && !state.ReadReplication.IsNull() {
		plan.ReadReplication = customfield.NewObjectMust(
			ctx,
			&D1DatabaseReadReplicationModel{
				Mode: types.StringValue("disabled"),
			},
		)
		res.Diagnostics.Append(res.Plan.Set(ctx, &plan)...)
	}
}

package v501

import (
	"context"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV500 handles state upgrades from schema version 500 to version 501.
//
// v500 state uses Set-typed policies and permission_groups. v501 switches to
// List-typed with FastSetType for O(n) performance. This handler reads the v500
// Set state, sorts policies and permission_groups canonically (for stable List
// ordering), and writes back as List state.
func UpgradeFromV500(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account_token state from v500 to v501")

	var state AccountTokenModelV500
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Sort policies and permission_groups canonically for stable List ordering
	sortPolicies(state.Policies)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	tflog.Info(ctx, "State upgrade from v500 to v501 completed successfully")
}

// sortPolicies applies canonical sort: permission_groups by ID within each
// policy, then policies by effect then resources.
func sortPolicies(policies []PolicyV500) {
	if len(policies) == 0 {
		return
	}

	// Sort permission groups within each policy
	for i := range policies {
		pgs := policies[i].PermissionGroups
		sort.SliceStable(pgs, func(a, b int) bool {
			return pgs[a].ID.ValueString() < pgs[b].ID.ValueString()
		})
	}

	// Sort policies by effect, then resources
	sort.SliceStable(policies, func(i, j int) bool {
		ei := policies[i].Effect.ValueString()
		ej := policies[j].Effect.ValueString()
		if ei != ej {
			return ei < ej
		}
		return policies[i].Resources.ValueString() < policies[j].Resources.ValueString()
	})
}

// Package v500 implements state migration from legacy provider (v4) to current provider (v5)
// for the cloudflare_load_balancer_monitor resource.
package v500

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 disambiguates v4 and early-v5 state that both report version 0.
func UpgradeFromV0(targetSchema schema.Schema) func(context.Context, resource.UpgradeStateRequest, *resource.UpgradeStateResponse) {
	return func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
		if req.RawState == nil || len(req.RawState.JSON) == 0 {
			resp.Diagnostics.AddError("Missing raw state", "Cannot determine whether schema version 0 state uses the v4 header set or the early-v5 header map")
			return
		}

		// Prefer v4 when both schemas accept state without a header. The v4 path
		// also applies the defaults and zero-value normalization required by v5.
		v4Resp := resource.UpgradeStateResponse{State: resp.State}
		v4Err := upgradeRawV4(ctx, req.RawState, &v4Resp)
		if v4Err == nil {
			resp.State = v4Resp.State
			resp.Diagnostics.Append(v4Resp.Diagnostics...)
			return
		}

		targetRaw, targetErr := req.RawState.Unmarshal(targetSchema.Type().TerraformType(ctx))
		if targetErr == nil {
			resp.State.Raw = targetRaw
			return
		}

		resp.Diagnostics.AddError("Unrecognized load_balancer_monitor state",
			fmt.Sprintf("State could not be decoded as v4 collection-shaped state (%s) or early-v5 object-shaped state (%s)", v4Err, targetErr))
	}
}

func upgradeRawV4(ctx context.Context, rawState *tfprotov6.RawState, resp *resource.UpgradeStateResponse) error {
	sourceSchema := SourceLoadBalancerMonitorSchema()
	rawValue, err := rawState.Unmarshal(sourceSchema.Type().TerraformType(ctx))
	if err != nil {
		return err
	}

	var v4State SourceLoadBalancerMonitorModel
	state := tfsdk.State{Raw: rawValue, Schema: sourceSchema}
	diags := state.Get(ctx, &v4State)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return fmt.Errorf("v4 state model decode failed")
	}

	v5State, transformDiags := Transform(ctx, &v4State)
	resp.Diagnostics.Append(transformDiags...)
	if transformDiags.HasError() {
		return fmt.Errorf("v4-to-v5 transformation failed")
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	if resp.Diagnostics.HasError() {
		return fmt.Errorf("setting transformed state failed")
	}
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
	return nil
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
//
// CRITICAL: For no-op upgrades, we copy raw state directly to preserve all data without transformation.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer_monitor state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

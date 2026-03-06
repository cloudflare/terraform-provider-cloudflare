package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// V5SchemaV0 is set by parent package init() to avoid circular imports.
// It should return the custom_hostname v5 schema with version 0.
var V5SchemaV0 func(context.Context) *schema.Schema

func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil for schema version 0 migration")
		return
	}

	isV4, detectErr := detectV4State(req)
	if detectErr != nil {
		resp.Diagnostics.AddError("Failed to detect state format", fmt.Sprintf("Could not determine V4 vs V5 state: %s", detectErr))
		return
	}

	if isV4 {
		var v4State SourceCustomHostnameModel
		v4Diags := unmarshalV4State(ctx, req.RawState, &v4State)
		resp.Diagnostics.Append(v4Diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		tflog.Info(ctx, "Detected V4 custom_hostname state, performing v4->v5 transform")
		targetState, diags := Transform(ctx, v4State)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)
		return
	}

	tflog.Info(ctx, "Detected early V5 custom_hostname state at version 0, performing no-op version bump")
	v5Schema := V5SchemaV0(ctx)
	v5Type := v5Schema.Type().TerraformType(ctx)
	v5RawValue, err := req.RawState.Unmarshal(v5Type)
	if err != nil {
		resp.Diagnostics.AddError("Failed to unmarshal v5 state", "Could not parse raw state as v5 format: "+err.Error())
		return
	}

	resp.State.Raw = v5RawValue
}

func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading custom_hostname state from version=1 to version=500 (no-op)")
	resp.State.Raw = req.State.Raw
}

func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceCustomHostnameModel) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourceCustomHostnameSchema()
	sourceType := sourceSchema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(sourceType)
	if err != nil {
		diags.AddError("Failed to unmarshal v4 state", "Could not parse raw state as v4 format: "+err.Error())
		return diags
	}

	state := tfsdk.State{Raw: rawValue, Schema: sourceSchema}
	diags.Append(state.Get(ctx, target)...)
	return diags
}

func detectV4State(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState != nil && len(req.RawState.JSON) > 0 {
		var rawJSON map[string]interface{}
		if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err == nil {
			if _, ok := rawJSON["wait_for_ssl_pending_validation"]; ok {
				return true, nil
			}

			if sslRaw, ok := rawJSON["ssl"]; ok {
				switch sslRaw.(type) {
				case []interface{}:
					return true, nil
				}
			}

			return false, nil
		}
	}

	var rawState map[string]tftypes.Value
	if err := req.State.Raw.As(&rawState); err != nil {
		return false, fmt.Errorf("failed to read raw state as object: %w", err)
	}

	waitVal, ok := rawState["wait_for_ssl_pending_validation"]
	if ok && waitVal.IsKnown() && !waitVal.IsNull() {
		return true, nil
	}

	return false, nil
}

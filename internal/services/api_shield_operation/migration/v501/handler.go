// Package v501 strips the incorrectly promoted GET query-parameter attributes
// feature and with_schemas from any v5.26.0 state (schema version 500).
package v501

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func UpgradeFromV500(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		return
	}

	var rawAttrs map[string]json.RawMessage
	if err := json.Unmarshal(req.RawState.JSON, &rawAttrs); err != nil {
		resp.Diagnostics.AddError("failed to parse state during v500→v501 migration", err.Error())
		return
	}

	delete(rawAttrs, "feature")
	delete(rawAttrs, "with_schemas")

	cleaned, err := json.Marshal(rawAttrs)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialise state during v500→v501 migration", err.Error())
		return
	}

	targetType := resp.State.Schema.Type().TerraformType(ctx)
	cleanedRaw := &tfprotov6.RawState{JSON: cleaned}
	tfVal, err := cleanedRaw.Unmarshal(targetType)
	if err != nil {
		resp.Diagnostics.AddError("failed to unmarshal cleaned state during v500→v501 migration", err.Error())
		return
	}

	resp.State.Raw = tfVal
}

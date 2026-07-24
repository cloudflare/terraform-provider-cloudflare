package v501

import (
	"context"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV500 moves priority and the MX target to their canonical data
// fields. The top-level priority remains in state as a computed compatibility
// value while the API continues returning it.
func UpgradeFromV500(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading DNS record state from v500 to v501")
	if copyUnrelatedState(ctx, req, resp) {
		return
	}

	var state v500.TargetDNSRecordModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	CanonicalizePriorityData(&state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// UpgradeFromV0 handles earlier v5 state that used schema version 0 with the
// same attribute types as v500.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	UpgradeFromV500(ctx, req, resp)
}

// copyUnrelatedState avoids typed round-tripping of custom fields for records
// unaffected by this migration. The old and new schemas have identical
// Terraform value types; only the priority configuration mode changed.
func copyUnrelatedState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) bool {
	var recordType types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("type"), &recordType)...)
	if resp.Diagnostics.HasError() {
		return true
	}
	typeValue := strings.ToUpper(recordType.ValueString())
	if typeValue == "MX" || typeValue == "SRV" || typeValue == "URI" {
		return false
	}
	resp.State.Raw = req.State.Raw
	return true
}

// CanonicalizePriorityData applies the v501 state representation. Existing
// nested values win over compatibility fields returned at the root.
func CanonicalizePriorityData(state *v500.TargetDNSRecordModel) {
	recordType := strings.ToUpper(state.Type.ValueString())
	if recordType != "MX" && recordType != "SRV" && recordType != "URI" {
		return
	}

	if state.Data == nil {
		state.Data = &v500.TargetDNSRecordDataModel{}
	}
	if state.Data.Priority.IsNull() {
		state.Data.Priority = state.Priority
	}
	if recordType == "MX" && state.Data.Target.IsNull() {
		state.Data.Target = state.Content
	}

	// Avoid unknown zero values in freshly materialized data objects. These are
	// nullable Terraform values; assigning explicit nulls keeps state encoding
	// stable when only priority/target are populated.
	initializeNullDataFields(state.Data)
}

func initializeNullDataFields(data *v500.TargetDNSRecordDataModel) {
	// A zero-value framework value is already null. Assigning the two fields
	// involved in this migration explicitly documents that zero is preserved.
	if data.Priority.IsNull() {
		data.Priority = types.Float64Null()
	}
	if data.Target.IsNull() {
		data.Target = types.StringNull()
	}
}

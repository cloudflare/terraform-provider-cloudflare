package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles the state upgrade from v4 (schema version 0) to v5 (schema version 500)
// This is for per-zone certificates only. If the state has type="per-hostname", it will skip the upgrade.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Debug(ctx, "Starting state upgrade from v0 to v500 for per-zone authenticated_origin_pulls_certificate")

	// Read v4 state into v4 model
	var v4Model V4Model
	resp.Diagnostics.Append(req.State.Get(ctx, &v4Model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// CRITICAL: Filter based on type field
	// This StateUpgrader is ONLY for per-zone certificates
	// If this is a per-hostname certificate, skip it (should never happen in practice)
	if !v4Model.Type.IsNull() && v4Model.Type.ValueString() != "per-zone" {
		tflog.Warn(ctx, "State has type != 'per-zone', skipping upgrade for per-zone resource", map[string]interface{}{
			"type": v4Model.Type.ValueString(),
		})
		resp.Diagnostics.AddError(
			"Invalid State for Per-Zone Resource",
			"This state appears to be for a per-hostname certificate but is being migrated to a per-zone resource. "+
				"This should not happen. The state may need to be manually corrected.",
		)
		return
	}

	// Transform v4 model to v5 model
	v5Model := V5Model{
		ID:           v4Model.ID, // Preserve ID from v4
		ZoneID:       v4Model.ZoneID,
		Certificate:  v4Model.Certificate,
		PrivateKey:   v4Model.PrivateKey,
		Issuer:       v4Model.Issuer,
		SerialNumber: v4Model.SerialNumber, // Preserve serial_number from v4
		Signature:    v4Model.Signature,
		ExpiresOn:    v4Model.ExpiresOn,
		Status:       v4Model.Status,
		UploadedOn:   v4Model.UploadedOn,
		// Note: certificate_id will be set from ID by the resource Read
		// Note: enabled is computed and will be populated by Read
		// Note: type is removed (not in v5 schema)
	}

	tflog.Debug(ctx, "Successfully upgraded state from v0 to v500 for per-zone certificate")

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5Model)...)
}

// UpgradeFromV1 handles the no-op upgrade from v1 to v500
// This is called when the resource is already in v5 format but needs to move to schema version 500
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading authenticated_origin_pulls_certificate state from version=1 to version=500 (no-op)")
	resp.State.Raw = req.State.Raw
	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

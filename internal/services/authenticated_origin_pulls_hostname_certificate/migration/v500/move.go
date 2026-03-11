package v500

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MoveState implements the state mover for moving from cloudflare_authenticated_origin_pulls_certificate
// (with type="per-hostname") to cloudflare_authenticated_origin_pulls_hostname_certificate.
// This is called when a `moved` block is detected in the Terraform configuration.
func MoveState(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	tflog.Info(ctx, "Starting state move from cloudflare_authenticated_origin_pulls_certificate to cloudflare_authenticated_origin_pulls_hostname_certificate")

	// Read v4 state into v4 model
	var v4Model V4Model
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &v4Model)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to get source state during move")
		return
	}

	tflog.Debug(ctx, "Source state retrieved", map[string]interface{}{
		"zone_id": v4Model.ZoneID.ValueString(),
		"type":    v4Model.Type.ValueString(),
		"id":      v4Model.ID.ValueString(),
	})

	// Validate that this is a per-hostname certificate (the only type we should move)
	if !v4Model.Type.IsNull() && v4Model.Type.ValueString() != "per-hostname" {
		resp.Diagnostics.AddError(
			"Invalid State Move",
			fmt.Sprintf("Cannot move cloudflare_authenticated_origin_pulls_certificate with type='%s' to cloudflare_authenticated_origin_pulls_hostname_certificate. Only type='per-hostname' should be moved to this resource type.", v4Model.Type.ValueString()),
		)
		return
	}

	// Transform v4 model to v5 model
	// Map all fields directly - the schemas are compatible except for the type field which is removed
	v5Model := V5Model{
		ID:           v4Model.ID,
		ZoneID:       v4Model.ZoneID,
		Certificate:  v4Model.Certificate,
		PrivateKey:   v4Model.PrivateKey,
		Issuer:       v4Model.Issuer,
		Signature:    v4Model.Signature,
		SerialNumber: v4Model.SerialNumber,
		ExpiresOn:    v4Model.ExpiresOn,
		Status:       v4Model.Status,
		UploadedOn:   v4Model.UploadedOn,
	}

	tflog.Info(ctx, "State move transformation completed", map[string]interface{}{
		"source_type": v4Model.Type.ValueString(),
		"zone_id":     v5Model.ZoneID.ValueString(),
	})

	// Set the target state
	resp.Diagnostics.Append(resp.TargetState.Set(ctx, &v5Model)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set target state during move")
		return
	}

	tflog.Info(ctx, "State move from cloudflare_authenticated_origin_pulls_certificate to cloudflare_authenticated_origin_pulls_hostname_certificate completed successfully")
}

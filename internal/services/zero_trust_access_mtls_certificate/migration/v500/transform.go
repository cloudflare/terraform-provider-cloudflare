package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function is shared by both UpgradeFromV0 and MoveState handlers.
func Transform(ctx context.Context, source SourceAccessMutualTLSCertificateModel) (*TargetAccessMTLSCertificateModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	tflog.Debug(ctx, "Transforming access_mtls_certificate state from v4 to v5",
		map[string]interface{}{
			"source_name": source.Name.ValueString(),
		})

	target := &TargetAccessMTLSCertificateModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		ZoneID:      source.ZoneID,
		Certificate: source.Certificate,
		Name:        source.Name,
		Fingerprint: source.Fingerprint,
	}

	// Convert associated_hostnames: types.Set → *[]types.String
	// If null/missing in v4 → empty slice (matches v5 default of empty set)
	if !source.AssociatedHostnames.IsNull() && !source.AssociatedHostnames.IsUnknown() {
		var elements []types.String
		diags.Append(source.AssociatedHostnames.ElementsAs(ctx, &elements, false)...)
		if diags.HasError() {
			return nil, diags
		}
		target.AssociatedHostnames = &elements
	} else {
		empty := []types.String{}
		target.AssociatedHostnames = &empty
	}

	// Convert expires_on: types.String → timetypes.RFC3339
	if !source.ExpiresOn.IsNull() && !source.ExpiresOn.IsUnknown() {
		expiresOn, expiresOnDiags := timetypes.NewRFC3339Value(source.ExpiresOn.ValueString())
		diags.Append(expiresOnDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.ExpiresOn = expiresOn
	} else {
		target.ExpiresOn = timetypes.NewRFC3339Null()
	}

	return target, diags
}

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4 SDKv2) state to target (current v5 Plugin Framework) state.
// This handles all field transformations and type conversions for zero_trust_gateway_certificate.
func Transform(_ context.Context, source SourceCloudflareZeroTrustGatewayCertificateModel) (*TargetZeroTrustGatewayCertificateModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for zero_trust_gateway_certificate migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct pass-through fields
	target := &TargetZeroTrustGatewayCertificateModel{
		ID:            source.ID,
		AccountID:     source.AccountID,
		Activate:      source.Activate,
		InUse:         source.InUse,
		BindingStatus: source.BindingStatus,
	}

	// Step 3: Handle validity_period_days
	// v4 had Default: 1826; SDKv2 always writes the default to state even if not in config.
	// Since the state upgrader cannot inspect config, we null out the value if it equals 1826
	// (assume it was the default). If it differs, the user explicitly set it, so preserve it.
	// Edge case: a user who explicitly set 1826 will lose it, but this is an acceptable
	// trade-off since their cert validity won't change and they can re-specify in v5 config.
	vpd := source.ValidityPeriodDays
	if !vpd.IsNull() && !vpd.IsUnknown() && vpd.ValueInt64() != 1826 {
		target.ValidityPeriodDays = vpd
	} else {
		target.ValidityPeriodDays = types.Int64Null()
	}

	// Step 4: Handle date field conversions (types.String → timetypes.RFC3339)
	// v4 stored dates as plain strings already in RFC3339 format (via time.RFC3339Nano).
	target.CreatedAt = convertStringToRFC3339(source.CreatedAt, &diags)
	target.ExpiresOn = convertStringToRFC3339(source.ExpiresOn, &diags)
	target.UploadedOn = convertStringToRFC3339(source.UploadedOn, &diags)
	if diags.HasError() {
		return nil, diags
	}

	// Step 5: Set new v5 computed fields to null (will refresh from API after migration)
	// These fields did not exist in v4 and must not be populated from v4 state.
	target.Certificate = types.StringNull()
	target.Fingerprint = types.StringNull()
	target.IssuerOrg = types.StringNull()
	target.IssuerRaw = types.StringNull()
	target.Type = types.StringNull()
	target.UpdatedAt = timetypes.NewRFC3339Null()

	// Note: custom, gateway_managed, qs_pack_id are intentionally not copied (removed in v5)

	return target, diags
}

// convertStringToRFC3339 converts a types.String to timetypes.RFC3339.
// If the string is null, unknown, or empty, returns timetypes.NewRFC3339Null().
func convertStringToRFC3339(val types.String, diags *diag.Diagnostics) timetypes.RFC3339 {
	if val.IsNull() || val.IsUnknown() || val.ValueString() == "" {
		return timetypes.NewRFC3339Null()
	}
	result, parseDiags := timetypes.NewRFC3339Value(val.ValueString())
	diags.Append(parseDiags...)
	return result
}

package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (v4 SDKv2) state to target (v5 Plugin Framework) state.
//
// mtls_certificate is one of the simplest migrations:
// - No field renames
// - No type conversions for string/bool fields
// - uploaded_on and expires_on: types.String → timetypes.RFC3339 (same RFC3339 value)
// - updated_at: new computed field in v5, set to null (API will populate on first Read)
//
// This function is used by UpgradeFromV4.
func Transform(ctx context.Context, source SourceCloudflareMTLSCertificateModel) (*TargetMTLSCertificateModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for mtls_certificate migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	if source.CA.IsNull() || source.CA.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"ca is required for mtls_certificate migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	if source.Certificates.IsNull() || source.Certificates.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"certificates is required for mtls_certificate migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Initialize target model
	target := &TargetMTLSCertificateModel{}

	// Step 1: Direct field copies (unchanged between v4 and v5)
	target.ID = source.ID
	target.AccountID = source.AccountID
	target.CA = source.CA
	target.Certificates = source.Certificates
	// v4 SDKv2 stores unset optional strings as "" rather than null.
	// Normalize to null so v5 optional fields behave correctly (avoids spurious ForceNew diffs).
	target.Name = migrations.FalseyStringToNull(source.Name)
	target.PrivateKey = migrations.FalseyStringToNull(source.PrivateKey)
	target.Issuer = source.Issuer
	target.Signature = source.Signature
	target.SerialNumber = source.SerialNumber

	// Step 2: Convert types.String timestamps → timetypes.RFC3339
	// The v4 provider stored these as plain TypeString in RFC3339 format.
	// The v5 provider uses timetypes.RFC3339 (a custom type wrapping the same string).
	if !source.UploadedOn.IsNull() && !source.UploadedOn.IsUnknown() {
		uploadedOn, uploadedOnDiags := timetypes.NewRFC3339Value(source.UploadedOn.ValueString())
		diags.Append(uploadedOnDiags...)
		if !diags.HasError() {
			target.UploadedOn = uploadedOn
		}
	} else {
		target.UploadedOn = timetypes.NewRFC3339Null()
	}

	if !source.ExpiresOn.IsNull() && !source.ExpiresOn.IsUnknown() {
		expiresOn, expiresOnDiags := timetypes.NewRFC3339Value(source.ExpiresOn.ValueString())
		diags.Append(expiresOnDiags...)
		if !diags.HasError() {
			target.ExpiresOn = expiresOn
		}
	} else {
		target.ExpiresOn = timetypes.NewRFC3339Null()
	}

	// Step 3: New v5 computed field - not present in v4, set to null.
	// The API will populate updated_at on the first Read after migration.
	target.UpdatedAt = timetypes.NewRFC3339Null()

	return target, diags
}

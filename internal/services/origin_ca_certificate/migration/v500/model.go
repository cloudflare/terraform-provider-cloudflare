package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x SDKv2)
// ============================================================================

// SourceCloudflareOriginCACertificateModel represents the source cloudflare_origin_ca_certificate state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by UpgradeFromV4 to parse legacy v4 state.
//
// Key differences from target v5 model:
// - Hostnames: types.Set (v4) → *[]types.String (v5) - Set to List conversion needed
// - RequestedValidity: types.Int64 (v4) → types.Float64 (v5) - Type conversion needed
// - MinDaysForRenewal: Exists in v4, REMOVED in v5 - Will be dropped during migration
type SourceCloudflareOriginCACertificateModel struct {
	// Resource identifier (implicit in SDKv2 but present in state)
	ID types.String `tfsdk:"id"`

	// Required user-provided fields
	Csr         types.String `tfsdk:"csr"`          // Certificate Signing Request (no change in v5)
	RequestType types.String `tfsdk:"request_type"` // Signature type: origin-rsa, origin-ecc, keyless-certificate (no change in v5)
	Hostnames   types.Set    `tfsdk:"hostnames"`    // TypeSet in v4 → *[]types.String in v5 (SET TO LIST CONVERSION)

	// Optional user-provided field
	RequestedValidity  types.Int64 `tfsdk:"requested_validity"`   // TypeInt in v4 → Float64 in v5 (TYPE CONVERSION + DEFAULT 5475)
	MinDaysForRenewal  types.Int64 `tfsdk:"min_days_for_renewal"` // REMOVED in v5 (DROP DURING MIGRATION)

	// Computed fields (API-provided, do not modify during migration)
	Certificate types.String `tfsdk:"certificate"` // API-generated certificate (no change)
	ExpiresOn   types.String `tfsdk:"expires_on"`  // API-generated expiration timestamp (no change)
}

// ============================================================================
// Target Models (Current Provider - v5.x+ Framework)
// ============================================================================

// TargetOriginCACertificateModel represents the target cloudflare_origin_ca_certificate state structure (v500).
// This corresponds to schema_version=500 in the current cloudflare provider.
//
// Note: This matches the model in the parent package's model.go file (OriginCACertificateModel).
// We duplicate it here to keep the migration package self-contained.
type TargetOriginCACertificateModel struct {
	// Resource identifier
	ID types.String `tfsdk:"id"`

	// Required user-provided fields
	Csr         types.String    `tfsdk:"csr"`          // Certificate Signing Request
	RequestType types.String    `tfsdk:"request_type"` // Signature type: origin-rsa, origin-ecc, keyless-certificate
	Hostnames   *[]types.String `tfsdk:"hostnames"`    // List of hostnames (converted from Set)

	// Optional user-provided field with default
	RequestedValidity types.Float64 `tfsdk:"requested_validity"` // Number of days for validity (default: 5475)

	// Computed fields (API-provided)
	Certificate types.String `tfsdk:"certificate"` // The Origin CA certificate (newline-encoded)
	ExpiresOn   types.String `tfsdk:"expires_on"`  // When the certificate will expire
}

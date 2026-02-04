package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourcePagesDomainModel represents the legacy pages_domain resource state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_pages_domain
//
// v4 schema location: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_pages_domain.go
type SourcePagesDomainModel struct {
	// Required fields
	AccountID   types.String `tfsdk:"account_id"`
	Domain      types.String `tfsdk:"domain"`       // v4 field name (renamed to "name" in v5)
	ProjectName types.String `tfsdk:"project_name"`

	// Computed fields
	Status types.String `tfsdk:"status"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetPagesDomainModel represents the current pages_domain resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_pages_domain
//
// This matches the model in the parent package's model.go file.
// We duplicate it here to keep the migration package self-contained.
type TargetPagesDomainModel struct {
	ID                   types.String                                                  `tfsdk:"id"`
	Name                 types.String                                                  `tfsdk:"name"` // Renamed from "domain" in v4
	AccountID            types.String                                                  `tfsdk:"account_id"`
	ProjectName          types.String                                                  `tfsdk:"project_name"`
	CertificateAuthority types.String                                                  `tfsdk:"certificate_authority"`
	CreatedOn            types.String                                                  `tfsdk:"created_on"`
	DomainID             types.String                                                  `tfsdk:"domain_id"`
	Status               types.String                                                  `tfsdk:"status"`
	ZoneTag              types.String                                                  `tfsdk:"zone_tag"`
	ValidationData       customfield.NestedObject[TargetPagesDomainValidationDataModel]   `tfsdk:"validation_data"`
	VerificationData     customfield.NestedObject[TargetPagesDomainVerificationDataModel] `tfsdk:"verification_data"`
}

// TargetPagesDomainValidationDataModel represents the validation_data nested object in v5.
type TargetPagesDomainValidationDataModel struct {
	Method       types.String `tfsdk:"method"`
	Status       types.String `tfsdk:"status"`
	ErrorMessage types.String `tfsdk:"error_message"`
	TXTName      types.String `tfsdk:"txt_name"`
	TXTValue     types.String `tfsdk:"txt_value"`
}

// TargetPagesDomainVerificationDataModel represents the verification_data nested object in v5.
type TargetPagesDomainVerificationDataModel struct {
	Status       types.String `tfsdk:"status"`
	ErrorMessage types.String `tfsdk:"error_message"`
}

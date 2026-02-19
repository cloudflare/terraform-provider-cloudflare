package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareCertificatePackModel represents the legacy cloudflare_certificate_pack resource state from v4.x provider.
// Schema version: 0 (SDKv2 provider)
type SourceCloudflareCertificatePackModel struct {
	ID                   types.String                   `tfsdk:"id"`
	ZoneID               types.String                   `tfsdk:"zone_id"`
	CertificateAuthority types.String                   `tfsdk:"certificate_authority"`
	Type                 types.String                   `tfsdk:"type"`
	ValidationMethod     types.String                   `tfsdk:"validation_method"`
	ValidityDays         types.Int64                    `tfsdk:"validity_days"` // Int in v4 schema, but stored as number
	CloudflareBranding   types.Bool                     `tfsdk:"cloudflare_branding"`
	Hosts                types.Set                      `tfsdk:"hosts"` // Regular Set in v4, not customfield.Set
	ValidationRecords    []SourceValidationRecordsModel `tfsdk:"validation_records"`
	ValidationErrors     []SourceValidationErrorsModel  `tfsdk:"validation_errors"`
	WaitForActiveStatus  types.Bool                     `tfsdk:"wait_for_active_status"` // REMOVED in v5
}

// SourceValidationRecordsModel represents validation_records item structure from v4.x provider.
type SourceValidationRecordsModel struct {
	// Fields that exist in both v4 and v5
	Emails   types.List   `tfsdk:"emails"`
	HTTPBody types.String `tfsdk:"http_body"`
	HTTPURL  types.String `tfsdk:"http_url"`
	TXTName  types.String `tfsdk:"txt_name"`
	TXTValue types.String `tfsdk:"txt_value"`

	// Fields REMOVED in v5 (only in v4)
	CNAMETarget types.String `tfsdk:"cname_target"`
	CNAMEName   types.String `tfsdk:"cname_name"`
}

// SourceValidationErrorsModel represents validation_errors item structure from v4.x provider.
type SourceValidationErrorsModel struct {
	Message types.String `tfsdk:"message"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetCertificatePackModel represents the current cloudflare_certificate_pack resource state from v5.x+ provider.
// Schema version: 500
type TargetCertificatePackModel struct {
	ID                   types.String                                               `tfsdk:"id"`
	ZoneID               types.String                                               `tfsdk:"zone_id"`
	CertificateAuthority types.String                                               `tfsdk:"certificate_authority"`
	Type                 types.String                                               `tfsdk:"type"`
	ValidationMethod     types.String                                               `tfsdk:"validation_method"`
	ValidityDays         types.Int64                                                `tfsdk:"validity_days"`
	CloudflareBranding   types.Bool                                                 `tfsdk:"cloudflare_branding"`
	Hosts                customfield.Set[types.String]                              `tfsdk:"hosts"`
	PrimaryCertificate   types.String                                               `tfsdk:"primary_certificate"`
	Status               types.String                                               `tfsdk:"status"`
	Certificates         customfield.NestedObjectList[TargetCertificatesModel]      `tfsdk:"certificates"`
	ValidationErrors     customfield.NestedObjectList[TargetValidationErrorsModel]  `tfsdk:"validation_errors"`
	ValidationRecords    customfield.NestedObjectList[TargetValidationRecordsModel] `tfsdk:"validation_records"`
}

// TargetValidationRecordsModel represents validation_records item structure from v5.x+ provider.
type TargetValidationRecordsModel struct {
	Emails   customfield.List[types.String] `tfsdk:"emails"`
	HTTPBody types.String                   `tfsdk:"http_body"`
	HTTPURL  types.String                   `tfsdk:"http_url"`
	TXTName  types.String                   `tfsdk:"txt_name"`
	TXTValue types.String                   `tfsdk:"txt_value"`
}

// TargetValidationErrorsModel represents validation_errors item structure from v5.x+ provider.
type TargetValidationErrorsModel struct {
	Message types.String `tfsdk:"message"`
}

// TargetCertificatesModel represents certificates item structure from v5.x+ provider.
type TargetCertificatesModel struct {
	ID              types.String                                                     `tfsdk:"id"`
	Hosts           customfield.List[types.String]                                   `tfsdk:"hosts"`
	Status          types.String                                                     `tfsdk:"status"`
	BundleMethod    types.String                                                     `tfsdk:"bundle_method"`
	ExpiresOn       timetypes.RFC3339                                                `tfsdk:"expires_on"`
	GeoRestrictions customfield.NestedObject[TargetCertificatesGeoRestrictionsModel] `tfsdk:"geo_restrictions"`
	Issuer          types.String                                                     `tfsdk:"issuer"`
	ModifiedOn      timetypes.RFC3339                                                `tfsdk:"modified_on"`
	Priority        types.Float64                                                    `tfsdk:"priority"`
	Signature       types.String                                                     `tfsdk:"signature"`
	UploadedOn      timetypes.RFC3339                                                `tfsdk:"uploaded_on"`
	ZoneID          types.String                                                     `tfsdk:"zone_id"`
}

// TargetCertificatesGeoRestrictionsModel represents geo_restrictions nested structure.
type TargetCertificatesGeoRestrictionsModel struct {
	Label types.String `tfsdk:"label"`
}

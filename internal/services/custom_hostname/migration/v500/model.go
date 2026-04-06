package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCustomHostnameModel represents legacy v4 state from cloudflare_custom_hostname.
type SourceCustomHostnameModel struct {
	ID                          types.String              `tfsdk:"id"`
	ZoneID                      types.String              `tfsdk:"zone_id"`
	Hostname                    types.String              `tfsdk:"hostname"`
	SSL                         []SourceCustomHostnameSSL `tfsdk:"ssl"`
	CustomOriginServer          types.String              `tfsdk:"custom_origin_server"`
	CustomOriginSNI             types.String              `tfsdk:"custom_origin_sni"`
	CustomMetadata              types.Map                 `tfsdk:"custom_metadata"`
	Status                      types.String              `tfsdk:"status"`
	OwnershipVerification       types.Map                 `tfsdk:"ownership_verification"`
	OwnershipVerificationHTTP   types.Map                 `tfsdk:"ownership_verification_http"`
	WaitForSSLPendingValidation types.Bool                `tfsdk:"wait_for_ssl_pending_validation"`
}

type SourceCustomHostnameSSL struct {
	Status               types.String                            `tfsdk:"status"`
	BundleMethod         types.String                            `tfsdk:"bundle_method"`
	Method               types.String                            `tfsdk:"method"`
	Type                 types.String                            `tfsdk:"type"`
	CertificateAuthority types.String                            `tfsdk:"certificate_authority"`
	ValidationRecords    []SourceCustomHostnameSSLValidationInfo `tfsdk:"validation_records"`
	ValidationErrors     []SourceCustomHostnameSSLValidationErr  `tfsdk:"validation_errors"`
	Wildcard             types.Bool                              `tfsdk:"wildcard"`
	CustomCertificate    types.String                            `tfsdk:"custom_certificate"`
	CustomKey            types.String                            `tfsdk:"custom_key"`
	Settings             []SourceCustomHostnameSSLSettings       `tfsdk:"settings"`
}

type SourceCustomHostnameSSLValidationErr struct {
	Message types.String `tfsdk:"message"`
}

type SourceCustomHostnameSSLValidationInfo struct {
	CNAMETarget types.String `tfsdk:"cname_target"`
	CNAMEName   types.String `tfsdk:"cname_name"`
	TXTName     types.String `tfsdk:"txt_name"`
	TXTValue    types.String `tfsdk:"txt_value"`
	HTTPURL     types.String `tfsdk:"http_url"`
	HTTPBody    types.String `tfsdk:"http_body"`
	Emails      types.List   `tfsdk:"emails"`
}

type SourceCustomHostnameSSLSettings struct {
	HTTP2         types.String `tfsdk:"http2"`
	TLS13         types.String `tfsdk:"tls13"`
	MinTLSVersion types.String `tfsdk:"min_tls_version"`
	Ciphers       types.Set    `tfsdk:"ciphers"`
	EarlyHints    types.String `tfsdk:"early_hints"`
}

// TargetCustomHostnameModel mirrors the v5 model for migration output.
type TargetCustomHostnameModel struct {
	ID                        types.String                                                            `tfsdk:"id"`
	ZoneID                    types.String                                                            `tfsdk:"zone_id"`
	Hostname                  types.String                                                            `tfsdk:"hostname"`
	SSL                       *TargetCustomHostnameSSLModel                                           `tfsdk:"ssl"`
	CustomOriginServer        types.String                                                            `tfsdk:"custom_origin_server"`
	CustomOriginSNI           types.String                                                            `tfsdk:"custom_origin_sni"`
	CustomMetadata            *map[string]types.String                                                `tfsdk:"custom_metadata"`
	CreatedAt                 timetypes.RFC3339                                                       `tfsdk:"created_at"`
	Status                    types.String                                                            `tfsdk:"status"`
	VerificationErrors        customfield.List[types.String]                                          `tfsdk:"verification_errors"`
	OwnershipVerification     customfield.NestedObject[TargetCustomHostnameOwnershipVerification]     `tfsdk:"ownership_verification"`
	OwnershipVerificationHTTP customfield.NestedObject[TargetCustomHostnameOwnershipVerificationHTTP] `tfsdk:"ownership_verification_http"`
}

type TargetCustomHostnameSSLModel struct {
	BundleMethod         types.String                               `tfsdk:"bundle_method"`
	CertificateAuthority types.String                               `tfsdk:"certificate_authority"`
	CloudflareBranding   types.Bool                                 `tfsdk:"cloudflare_branding"`
	CustomCERTBundle     *[]*TargetCustomHostnameSSLCustomCERTModel `tfsdk:"custom_cert_bundle"`
	CustomCertificate    types.String                               `tfsdk:"custom_certificate"`
	CustomCSRID          types.String                               `tfsdk:"custom_csr_id"`
	CustomKey            types.String                               `tfsdk:"custom_key"`
	Method               types.String                               `tfsdk:"method"`
	Settings             *TargetCustomHostnameSSLSettingsModel      `tfsdk:"settings"`
	Type                 types.String                               `tfsdk:"type"`
	Wildcard             types.Bool                                 `tfsdk:"wildcard"`
}

type TargetCustomHostnameSSLCustomCERTModel struct {
	CustomCertificate types.String `tfsdk:"custom_certificate"`
	CustomKey         types.String `tfsdk:"custom_key"`
}

type TargetCustomHostnameSSLSettingsModel struct {
	Ciphers       *[]types.String `tfsdk:"ciphers"`
	EarlyHints    types.String    `tfsdk:"early_hints"`
	HTTP2         types.String    `tfsdk:"http2"`
	MinTLSVersion types.String    `tfsdk:"min_tls_version"`
	TLS1_3        types.String    `tfsdk:"tls_1_3"`
}

type TargetCustomHostnameOwnershipVerification struct {
	Name  types.String `tfsdk:"name"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type TargetCustomHostnameOwnershipVerificationHTTP struct {
	HTTPBody types.String `tfsdk:"http_body"`
	HTTPURL  types.String `tfsdk:"http_url"`
}

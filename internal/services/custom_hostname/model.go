// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameResultEnvelope struct {
	Result CustomHostnameModel `json:"result,computed"`
}

type CustomHostnameModel struct {
	ID                        types.String                                                           `tfsdk:"id" json:"id,computed"`
	ZoneID                    types.String                                                           `tfsdk:"zone_id" path:"zone_id"`
	SSL                       *CustomHostnameSSLModel                                                `tfsdk:"ssl" json:"ssl"`
	CustomMetadata            *CustomHostnameCustomMetadataModel                                     `tfsdk:"custom_metadata" json:"custom_metadata"`
	Hostname                  types.String                                                           `tfsdk:"hostname" json:"hostname"`
	CustomOriginServer        types.String                                                           `tfsdk:"custom_origin_server" json:"custom_origin_server"`
	CustomOriginSNI           types.String                                                           `tfsdk:"custom_origin_sni" json:"custom_origin_sni"`
	CreatedAt                 timetypes.RFC3339                                                      `tfsdk:"created_at" json:"created_at,computed"`
	OwnershipVerification     customfield.NestedObject[CustomHostnameOwnershipVerificationModel]     `tfsdk:"ownership_verification" json:"ownership_verification,computed"`
	OwnershipVerificationHTTP customfield.NestedObject[CustomHostnameOwnershipVerificationHTTPModel] `tfsdk:"ownership_verification_http" json:"ownership_verification_http,computed"`
	Status                    types.String                                                           `tfsdk:"status" json:"status,computed"`
	VerificationErrors        *[]jsontypes.Normalized                                                `tfsdk:"verification_errors" json:"verification_errors,computed"`
}

type CustomHostnameSSLModel struct {
	BundleMethod         types.String                    `tfsdk:"bundle_method" json:"bundle_method"`
	CertificateAuthority types.String                    `tfsdk:"certificate_authority" json:"certificate_authority"`
	CustomCertificate    types.String                    `tfsdk:"custom_certificate" json:"custom_certificate"`
	CustomKey            types.String                    `tfsdk:"custom_key" json:"custom_key"`
	Method               types.String                    `tfsdk:"method" json:"method"`
	Settings             *CustomHostnameSSLSettingsModel `tfsdk:"settings" json:"settings"`
	Type                 types.String                    `tfsdk:"type" json:"type"`
	Wildcard             types.Bool                      `tfsdk:"wildcard" json:"wildcard"`
}

type CustomHostnameSSLSettingsModel struct {
	Ciphers       *[]types.String `tfsdk:"ciphers" json:"ciphers"`
	EarlyHints    types.String    `tfsdk:"early_hints" json:"early_hints"`
	HTTP2         types.String    `tfsdk:"http2" json:"http2"`
	MinTLSVersion types.String    `tfsdk:"min_tls_version" json:"min_tls_version"`
	TLS1_3        types.String    `tfsdk:"tls_1_3" json:"tls_1_3"`
}

type CustomHostnameCustomMetadataModel struct {
	Key types.String `tfsdk:"key" json:"key"`
}

type CustomHostnameOwnershipVerificationModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type CustomHostnameOwnershipVerificationHTTPModel struct {
	HTTPBody types.String `tfsdk:"http_body" json:"http_body"`
	HTTPURL  types.String `tfsdk:"http_url" json:"http_url"`
}

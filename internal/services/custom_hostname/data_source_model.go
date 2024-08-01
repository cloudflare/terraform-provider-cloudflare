// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameResultDataSourceEnvelope struct {
	Result CustomHostnameDataSourceModel `json:"result,computed"`
}

type CustomHostnameResultListDataSourceEnvelope struct {
	Result *[]*CustomHostnameDataSourceModel `json:"result,computed"`
}

type CustomHostnameDataSourceModel struct {
	CustomHostnameID          types.String                                               `tfsdk:"custom_hostname_id" path:"custom_hostname_id"`
	ZoneID                    types.String                                               `tfsdk:"zone_id" path:"zone_id"`
	CreatedAt                 timetypes.RFC3339                                          `tfsdk:"created_at" json:"created_at,computed"`
	Hostname                  types.String                                               `tfsdk:"hostname" json:"hostname,computed"`
	ID                        types.String                                               `tfsdk:"id" json:"id,computed"`
	SSL                       customfield.NestedObject[CustomHostnameSSLDataSourceModel] `tfsdk:"ssl" json:"ssl,computed"`
	CustomOriginServer        types.String                                               `tfsdk:"custom_origin_server" json:"custom_origin_server"`
	CustomOriginSNI           types.String                                               `tfsdk:"custom_origin_sni" json:"custom_origin_sni"`
	Status                    types.String                                               `tfsdk:"status" json:"status"`
	VerificationErrors        *[]jsontypes.Normalized                                    `tfsdk:"verification_errors" json:"verification_errors"`
	CustomMetadata            *CustomHostnameCustomMetadataDataSourceModel               `tfsdk:"custom_metadata" json:"custom_metadata"`
	OwnershipVerification     *CustomHostnameOwnershipVerificationDataSourceModel        `tfsdk:"ownership_verification" json:"ownership_verification"`
	OwnershipVerificationHTTP *CustomHostnameOwnershipVerificationHTTPDataSourceModel    `tfsdk:"ownership_verification_http" json:"ownership_verification_http"`
	Filter                    *CustomHostnameFindOneByDataSourceModel                    `tfsdk:"filter"`
}

type CustomHostnameSSLDataSourceModel struct {
	ID                   types.String                                          `tfsdk:"id" json:"id"`
	BundleMethod         types.String                                          `tfsdk:"bundle_method" json:"bundle_method,computed"`
	CertificateAuthority types.String                                          `tfsdk:"certificate_authority" json:"certificate_authority"`
	CustomCertificate    types.String                                          `tfsdk:"custom_certificate" json:"custom_certificate"`
	CustomCsrID          types.String                                          `tfsdk:"custom_csr_id" json:"custom_csr_id"`
	CustomKey            types.String                                          `tfsdk:"custom_key" json:"custom_key"`
	ExpiresOn            timetypes.RFC3339                                     `tfsdk:"expires_on" json:"expires_on"`
	Hosts                *[]jsontypes.Normalized                               `tfsdk:"hosts" json:"hosts"`
	Issuer               types.String                                          `tfsdk:"issuer" json:"issuer"`
	Method               types.String                                          `tfsdk:"method" json:"method"`
	SerialNumber         types.String                                          `tfsdk:"serial_number" json:"serial_number"`
	Settings             *CustomHostnameSSLSettingsDataSourceModel             `tfsdk:"settings" json:"settings"`
	Signature            types.String                                          `tfsdk:"signature" json:"signature"`
	Status               types.String                                          `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                          `tfsdk:"type" json:"type"`
	UploadedOn           timetypes.RFC3339                                     `tfsdk:"uploaded_on" json:"uploaded_on"`
	ValidationErrors     *[]*CustomHostnameSSLValidationErrorsDataSourceModel  `tfsdk:"validation_errors" json:"validation_errors"`
	ValidationRecords    *[]*CustomHostnameSSLValidationRecordsDataSourceModel `tfsdk:"validation_records" json:"validation_records"`
	Wildcard             types.Bool                                            `tfsdk:"wildcard" json:"wildcard"`
}

type CustomHostnameSSLSettingsDataSourceModel struct {
	Ciphers       *[]types.String `tfsdk:"ciphers" json:"ciphers"`
	EarlyHints    types.String    `tfsdk:"early_hints" json:"early_hints"`
	HTTP2         types.String    `tfsdk:"http2" json:"http2"`
	MinTLSVersion types.String    `tfsdk:"min_tls_version" json:"min_tls_version"`
	TLS1_3        types.String    `tfsdk:"tls_1_3" json:"tls_1_3"`
}

type CustomHostnameSSLValidationErrorsDataSourceModel struct {
	Message types.String `tfsdk:"message" json:"message"`
}

type CustomHostnameSSLValidationRecordsDataSourceModel struct {
	Emails   *[]jsontypes.Normalized `tfsdk:"emails" json:"emails"`
	HTTPBody types.String            `tfsdk:"http_body" json:"http_body"`
	HTTPURL  types.String            `tfsdk:"http_url" json:"http_url"`
	TXTName  types.String            `tfsdk:"txt_name" json:"txt_name"`
	TXTValue types.String            `tfsdk:"txt_value" json:"txt_value"`
}

type CustomHostnameCustomMetadataDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key"`
}

type CustomHostnameOwnershipVerificationDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type CustomHostnameOwnershipVerificationHTTPDataSourceModel struct {
	HTTPBody types.String `tfsdk:"http_body" json:"http_body"`
	HTTPURL  types.String `tfsdk:"http_url" json:"http_url"`
}

type CustomHostnameFindOneByDataSourceModel struct {
	ZoneID    types.String  `tfsdk:"zone_id" path:"zone_id"`
	ID        types.String  `tfsdk:"id" query:"id"`
	Direction types.String  `tfsdk:"direction" query:"direction"`
	Hostname  types.String  `tfsdk:"hostname" query:"hostname"`
	Order     types.String  `tfsdk:"order" query:"order"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
	SSL       types.Float64 `tfsdk:"ssl" query:"ssl"`
}

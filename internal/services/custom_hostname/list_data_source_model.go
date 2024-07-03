// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnamesResultListDataSourceEnvelope struct {
	Result *[]*CustomHostnamesItemsDataSourceModel `json:"result,computed"`
}

type CustomHostnamesDataSourceModel struct {
	ZoneID    types.String                            `tfsdk:"zone_id" path:"zone_id"`
	ID        types.String                            `tfsdk:"id" query:"id"`
	Direction types.String                            `tfsdk:"direction" query:"direction"`
	Hostname  types.String                            `tfsdk:"hostname" query:"hostname"`
	Order     types.String                            `tfsdk:"order" query:"order"`
	Page      types.Float64                           `tfsdk:"page" query:"page"`
	PerPage   types.Float64                           `tfsdk:"per_page" query:"per_page"`
	SSL       types.Float64                           `tfsdk:"ssl" query:"ssl"`
	MaxItems  types.Int64                             `tfsdk:"max_items"`
	Items     *[]*CustomHostnamesItemsDataSourceModel `tfsdk:"items"`
}

type CustomHostnamesItemsDataSourceModel struct {
	ID                 types.String    `tfsdk:"id" json:"id,computed"`
	Hostname           types.String    `tfsdk:"hostname" json:"hostname,computed"`
	CreatedAt          types.String    `tfsdk:"created_at" json:"created_at,computed"`
	CustomOriginServer types.String    `tfsdk:"custom_origin_server" json:"custom_origin_server,computed"`
	CustomOriginSNI    types.String    `tfsdk:"custom_origin_sni" json:"custom_origin_sni,computed"`
	Status             types.String    `tfsdk:"status" json:"status,computed"`
	VerificationErrors *[]types.String `tfsdk:"verification_errors" json:"verification_errors,computed"`
}

type CustomHostnamesItemsSSLDataSourceModel struct {
	ID                   types.String                                                `tfsdk:"id" json:"id,computed"`
	BundleMethod         types.String                                                `tfsdk:"bundle_method" json:"bundle_method,computed"`
	CertificateAuthority types.String                                                `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CustomCertificate    types.String                                                `tfsdk:"custom_certificate" json:"custom_certificate,computed"`
	CustomCsrID          types.String                                                `tfsdk:"custom_csr_id" json:"custom_csr_id,computed"`
	CustomKey            types.String                                                `tfsdk:"custom_key" json:"custom_key,computed"`
	ExpiresOn            types.String                                                `tfsdk:"expires_on" json:"expires_on,computed"`
	Hosts                *[]types.String                                             `tfsdk:"hosts" json:"hosts,computed"`
	Issuer               types.String                                                `tfsdk:"issuer" json:"issuer,computed"`
	Method               types.String                                                `tfsdk:"method" json:"method,computed"`
	SerialNumber         types.String                                                `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature            types.String                                                `tfsdk:"signature" json:"signature,computed"`
	Status               types.String                                                `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                                `tfsdk:"type" json:"type,computed"`
	UploadedOn           types.String                                                `tfsdk:"uploaded_on" json:"uploaded_on,computed"`
	ValidationErrors     *[]*CustomHostnamesItemsSSLValidationErrorsDataSourceModel  `tfsdk:"validation_errors" json:"validation_errors,computed"`
	ValidationRecords    *[]*CustomHostnamesItemsSSLValidationRecordsDataSourceModel `tfsdk:"validation_records" json:"validation_records,computed"`
	Wildcard             types.Bool                                                  `tfsdk:"wildcard" json:"wildcard,computed"`
}

type CustomHostnamesItemsSSLSettingsDataSourceModel struct {
	Ciphers       *[]types.String `tfsdk:"ciphers" json:"ciphers,computed"`
	EarlyHints    types.String    `tfsdk:"early_hints" json:"early_hints,computed"`
	HTTP2         types.String    `tfsdk:"http2" json:"http2,computed"`
	MinTLSVersion types.String    `tfsdk:"min_tls_version" json:"min_tls_version,computed"`
	TLS1_3        types.String    `tfsdk:"tls_1_3" json:"tls_1_3,computed"`
}

type CustomHostnamesItemsSSLValidationErrorsDataSourceModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type CustomHostnamesItemsSSLValidationRecordsDataSourceModel struct {
	Emails   *[]types.String `tfsdk:"emails" json:"emails,computed"`
	HTTPBody types.String    `tfsdk:"http_body" json:"http_body,computed"`
	HTTPURL  types.String    `tfsdk:"http_url" json:"http_url,computed"`
	TXTName  types.String    `tfsdk:"txt_name" json:"txt_name,computed"`
	TXTValue types.String    `tfsdk:"txt_value" json:"txt_value,computed"`
}

type CustomHostnamesItemsCustomMetadataDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type CustomHostnamesItemsOwnershipVerificationDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type CustomHostnamesItemsOwnershipVerificationHTTPDataSourceModel struct {
	HTTPBody types.String `tfsdk:"http_body" json:"http_body,computed"`
	HTTPURL  types.String `tfsdk:"http_url" json:"http_url,computed"`
}

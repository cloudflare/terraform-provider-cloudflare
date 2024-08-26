// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/custom_hostnames"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnamesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CustomHostnamesResultDataSourceModel] `json:"result,computed"`
}

type CustomHostnamesDataSourceModel struct {
	ZoneID    types.String                                                       `tfsdk:"zone_id" path:"zone_id"`
	Direction types.String                                                       `tfsdk:"direction" query:"direction"`
	Hostname  types.String                                                       `tfsdk:"hostname" query:"hostname"`
	ID        types.String                                                       `tfsdk:"id" query:"id"`
	SSL       types.Float64                                                      `tfsdk:"ssl" query:"ssl"`
	Order     types.String                                                       `tfsdk:"order" query:"order,computed_optional"`
	MaxItems  types.Int64                                                        `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[CustomHostnamesResultDataSourceModel] `tfsdk:"result"`
}

func (m *CustomHostnamesDataSourceModel) toListParams() (params custom_hostnames.CustomHostnameListParams, diags diag.Diagnostics) {
	params = custom_hostnames.CustomHostnameListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(custom_hostnames.CustomHostnameListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Hostname.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(custom_hostnames.CustomHostnameListParamsOrder(m.Order.ValueString()))
	}
	if !m.SSL.IsNull() {
		params.SSL = cloudflare.F(custom_hostnames.CustomHostnameListParamsSSL(m.SSL.ValueFloat64()))
	}

	return
}

type CustomHostnamesResultDataSourceModel struct {
	ID                        types.String                                                `tfsdk:"id" json:"id,computed"`
	Hostname                  types.String                                                `tfsdk:"hostname" json:"hostname,computed"`
	SSL                       customfield.NestedObject[CustomHostnamesSSLDataSourceModel] `tfsdk:"ssl" json:"ssl,computed"`
	CreatedAt                 timetypes.RFC3339                                           `tfsdk:"created_at" json:"created_at,computed"`
	CustomMetadata            *CustomHostnamesCustomMetadataDataSourceModel               `tfsdk:"custom_metadata" json:"custom_metadata,computed_optional"`
	CustomOriginServer        types.String                                                `tfsdk:"custom_origin_server" json:"custom_origin_server,computed_optional"`
	CustomOriginSNI           types.String                                                `tfsdk:"custom_origin_sni" json:"custom_origin_sni,computed_optional"`
	OwnershipVerification     *CustomHostnamesOwnershipVerificationDataSourceModel        `tfsdk:"ownership_verification" json:"ownership_verification,computed_optional"`
	OwnershipVerificationHTTP *CustomHostnamesOwnershipVerificationHTTPDataSourceModel    `tfsdk:"ownership_verification_http" json:"ownership_verification_http,computed_optional"`
	Status                    types.String                                                `tfsdk:"status" json:"status,computed_optional"`
	VerificationErrors        *[]types.String                                             `tfsdk:"verification_errors" json:"verification_errors,computed_optional"`
}

type CustomHostnamesSSLDataSourceModel struct {
	ID                   types.String                                           `tfsdk:"id" json:"id,computed_optional"`
	BundleMethod         types.String                                           `tfsdk:"bundle_method" json:"bundle_method,computed"`
	CertificateAuthority types.String                                           `tfsdk:"certificate_authority" json:"certificate_authority,computed_optional"`
	CustomCertificate    types.String                                           `tfsdk:"custom_certificate" json:"custom_certificate,computed_optional"`
	CustomCsrID          types.String                                           `tfsdk:"custom_csr_id" json:"custom_csr_id,computed_optional"`
	CustomKey            types.String                                           `tfsdk:"custom_key" json:"custom_key,computed_optional"`
	ExpiresOn            timetypes.RFC3339                                      `tfsdk:"expires_on" json:"expires_on,computed_optional"`
	Hosts                *[]types.String                                        `tfsdk:"hosts" json:"hosts,computed_optional"`
	Issuer               types.String                                           `tfsdk:"issuer" json:"issuer,computed_optional"`
	Method               types.String                                           `tfsdk:"method" json:"method,computed_optional"`
	SerialNumber         types.String                                           `tfsdk:"serial_number" json:"serial_number,computed_optional"`
	Settings             *CustomHostnamesSSLSettingsDataSourceModel             `tfsdk:"settings" json:"settings,computed_optional"`
	Signature            types.String                                           `tfsdk:"signature" json:"signature,computed_optional"`
	Status               types.String                                           `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                           `tfsdk:"type" json:"type,computed_optional"`
	UploadedOn           timetypes.RFC3339                                      `tfsdk:"uploaded_on" json:"uploaded_on,computed_optional"`
	ValidationErrors     *[]*CustomHostnamesSSLValidationErrorsDataSourceModel  `tfsdk:"validation_errors" json:"validation_errors,computed_optional"`
	ValidationRecords    *[]*CustomHostnamesSSLValidationRecordsDataSourceModel `tfsdk:"validation_records" json:"validation_records,computed_optional"`
	Wildcard             types.Bool                                             `tfsdk:"wildcard" json:"wildcard,computed_optional"`
}

type CustomHostnamesSSLSettingsDataSourceModel struct {
	Ciphers       *[]types.String `tfsdk:"ciphers" json:"ciphers,computed_optional"`
	EarlyHints    types.String    `tfsdk:"early_hints" json:"early_hints,computed_optional"`
	HTTP2         types.String    `tfsdk:"http2" json:"http2,computed_optional"`
	MinTLSVersion types.String    `tfsdk:"min_tls_version" json:"min_tls_version,computed_optional"`
	TLS1_3        types.String    `tfsdk:"tls_1_3" json:"tls_1_3,computed_optional"`
}

type CustomHostnamesSSLValidationErrorsDataSourceModel struct {
	Message types.String `tfsdk:"message" json:"message,computed_optional"`
}

type CustomHostnamesSSLValidationRecordsDataSourceModel struct {
	Emails   *[]types.String `tfsdk:"emails" json:"emails,computed_optional"`
	HTTPBody types.String    `tfsdk:"http_body" json:"http_body,computed_optional"`
	HTTPURL  types.String    `tfsdk:"http_url" json:"http_url,computed_optional"`
	TXTName  types.String    `tfsdk:"txt_name" json:"txt_name,computed_optional"`
	TXTValue types.String    `tfsdk:"txt_value" json:"txt_value,computed_optional"`
}

type CustomHostnamesCustomMetadataDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed_optional"`
}

type CustomHostnamesOwnershipVerificationDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed_optional"`
	Type  types.String `tfsdk:"type" json:"type,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type CustomHostnamesOwnershipVerificationHTTPDataSourceModel struct {
	HTTPBody types.String `tfsdk:"http_body" json:"http_body,computed_optional"`
	HTTPURL  types.String `tfsdk:"http_url" json:"http_url,computed_optional"`
}

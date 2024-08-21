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
	Result *[]*CustomHostnamesResultDataSourceModel `json:"result,computed"`
}

type CustomHostnamesDataSourceModel struct {
	ZoneID    types.String                             `tfsdk:"zone_id" path:"zone_id"`
	Direction types.String                             `tfsdk:"direction" query:"direction"`
	Hostname  types.String                             `tfsdk:"hostname" query:"hostname"`
	ID        types.String                             `tfsdk:"id" query:"id"`
	SSL       types.Float64                            `tfsdk:"ssl" query:"ssl"`
	Order     types.String                             `tfsdk:"order" query:"order"`
	MaxItems  types.Int64                              `tfsdk:"max_items"`
	Result    *[]*CustomHostnamesResultDataSourceModel `tfsdk:"result"`
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
	CustomMetadata            *CustomHostnamesCustomMetadataDataSourceModel               `tfsdk:"custom_metadata" json:"custom_metadata"`
	CustomOriginServer        types.String                                                `tfsdk:"custom_origin_server" json:"custom_origin_server"`
	CustomOriginSNI           types.String                                                `tfsdk:"custom_origin_sni" json:"custom_origin_sni"`
	OwnershipVerification     *CustomHostnamesOwnershipVerificationDataSourceModel        `tfsdk:"ownership_verification" json:"ownership_verification"`
	OwnershipVerificationHTTP *CustomHostnamesOwnershipVerificationHTTPDataSourceModel    `tfsdk:"ownership_verification_http" json:"ownership_verification_http"`
	Status                    types.String                                                `tfsdk:"status" json:"status"`
	VerificationErrors        *[]types.String                                             `tfsdk:"verification_errors" json:"verification_errors"`
}

type CustomHostnamesSSLDataSourceModel struct {
	ID                   types.String                                           `tfsdk:"id" json:"id"`
	BundleMethod         types.String                                           `tfsdk:"bundle_method" json:"bundle_method,computed"`
	CertificateAuthority types.String                                           `tfsdk:"certificate_authority" json:"certificate_authority"`
	CustomCertificate    types.String                                           `tfsdk:"custom_certificate" json:"custom_certificate"`
	CustomCsrID          types.String                                           `tfsdk:"custom_csr_id" json:"custom_csr_id"`
	CustomKey            types.String                                           `tfsdk:"custom_key" json:"custom_key"`
	ExpiresOn            timetypes.RFC3339                                      `tfsdk:"expires_on" json:"expires_on"`
	Hosts                *[]types.String                                        `tfsdk:"hosts" json:"hosts"`
	Issuer               types.String                                           `tfsdk:"issuer" json:"issuer"`
	Method               types.String                                           `tfsdk:"method" json:"method"`
	SerialNumber         types.String                                           `tfsdk:"serial_number" json:"serial_number"`
	Settings             *CustomHostnamesSSLSettingsDataSourceModel             `tfsdk:"settings" json:"settings"`
	Signature            types.String                                           `tfsdk:"signature" json:"signature"`
	Status               types.String                                           `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                           `tfsdk:"type" json:"type"`
	UploadedOn           timetypes.RFC3339                                      `tfsdk:"uploaded_on" json:"uploaded_on"`
	ValidationErrors     *[]*CustomHostnamesSSLValidationErrorsDataSourceModel  `tfsdk:"validation_errors" json:"validation_errors"`
	ValidationRecords    *[]*CustomHostnamesSSLValidationRecordsDataSourceModel `tfsdk:"validation_records" json:"validation_records"`
	Wildcard             types.Bool                                             `tfsdk:"wildcard" json:"wildcard"`
}

type CustomHostnamesSSLSettingsDataSourceModel struct {
	Ciphers       *[]types.String `tfsdk:"ciphers" json:"ciphers"`
	EarlyHints    types.String    `tfsdk:"early_hints" json:"early_hints"`
	HTTP2         types.String    `tfsdk:"http2" json:"http2"`
	MinTLSVersion types.String    `tfsdk:"min_tls_version" json:"min_tls_version"`
	TLS1_3        types.String    `tfsdk:"tls_1_3" json:"tls_1_3"`
}

type CustomHostnamesSSLValidationErrorsDataSourceModel struct {
	Message types.String `tfsdk:"message" json:"message"`
}

type CustomHostnamesSSLValidationRecordsDataSourceModel struct {
	Emails   *[]types.String `tfsdk:"emails" json:"emails"`
	HTTPBody types.String    `tfsdk:"http_body" json:"http_body"`
	HTTPURL  types.String    `tfsdk:"http_url" json:"http_url"`
	TXTName  types.String    `tfsdk:"txt_name" json:"txt_name"`
	TXTValue types.String    `tfsdk:"txt_value" json:"txt_value"`
}

type CustomHostnamesCustomMetadataDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key"`
}

type CustomHostnamesOwnershipVerificationDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name"`
	Type  types.String `tfsdk:"type" json:"type"`
	Value types.String `tfsdk:"value" json:"value"`
}

type CustomHostnamesOwnershipVerificationHTTPDataSourceModel struct {
	HTTPBody types.String `tfsdk:"http_body" json:"http_body"`
	HTTPURL  types.String `tfsdk:"http_url" json:"http_url"`
}

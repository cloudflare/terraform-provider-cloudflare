// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/custom_hostnames"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnamesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CustomHostnamesResultDataSourceModel] `json:"result,computed"`
}

type CustomHostnamesDataSourceModel struct {
	ZoneID    types.String                                                       `tfsdk:"zone_id" path:"zone_id,required"`
	Direction types.String                                                       `tfsdk:"direction" query:"direction,optional"`
	Hostname  types.String                                                       `tfsdk:"hostname" query:"hostname,optional"`
	ID        types.String                                                       `tfsdk:"id" query:"id,optional"`
	SSL       types.Float64                                                      `tfsdk:"ssl" query:"ssl,optional"`
	Order     types.String                                                       `tfsdk:"order" query:"order,computed_optional"`
	MaxItems  types.Int64                                                        `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[CustomHostnamesResultDataSourceModel] `tfsdk:"result"`
}

func (m *CustomHostnamesDataSourceModel) toListParams(_ context.Context) (params custom_hostnames.CustomHostnameListParams, diags diag.Diagnostics) {
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
	ID                        types.String                                                                      `tfsdk:"id" json:"id,computed"`
	Hostname                  types.String                                                                      `tfsdk:"hostname" json:"hostname,computed"`
	SSL                       customfield.NestedObject[CustomHostnamesSSLDataSourceModel]                       `tfsdk:"ssl" json:"ssl,computed"`
	CreatedAt                 timetypes.RFC3339                                                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomMetadata            customfield.Map[types.String]                                                     `tfsdk:"custom_metadata" json:"custom_metadata,computed"`
	CustomOriginServer        types.String                                                                      `tfsdk:"custom_origin_server" json:"custom_origin_server,computed"`
	CustomOriginSNI           types.String                                                                      `tfsdk:"custom_origin_sni" json:"custom_origin_sni,computed"`
	OwnershipVerification     customfield.NestedObject[CustomHostnamesOwnershipVerificationDataSourceModel]     `tfsdk:"ownership_verification" json:"ownership_verification,computed"`
	OwnershipVerificationHTTP customfield.NestedObject[CustomHostnamesOwnershipVerificationHTTPDataSourceModel] `tfsdk:"ownership_verification_http" json:"ownership_verification_http,computed"`
	Status                    types.String                                                                      `tfsdk:"status" json:"status,computed"`
	VerificationErrors        customfield.List[types.String]                                                    `tfsdk:"verification_errors" json:"verification_errors,computed"`
}

type CustomHostnamesSSLDataSourceModel struct {
	ID                   types.String                                                                     `tfsdk:"id" json:"id,computed"`
	BundleMethod         types.String                                                                     `tfsdk:"bundle_method" json:"bundle_method,computed"`
	CertificateAuthority types.String                                                                     `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CustomCertificate    types.String                                                                     `tfsdk:"custom_certificate" json:"custom_certificate,computed"`
	CustomCsrID          types.String                                                                     `tfsdk:"custom_csr_id" json:"custom_csr_id,computed"`
	CustomKey            types.String                                                                     `tfsdk:"custom_key" json:"custom_key,computed"`
	ExpiresOn            timetypes.RFC3339                                                                `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Hosts                customfield.List[types.String]                                                   `tfsdk:"hosts" json:"hosts,computed"`
	Issuer               types.String                                                                     `tfsdk:"issuer" json:"issuer,computed"`
	Method               types.String                                                                     `tfsdk:"method" json:"method,computed"`
	SerialNumber         types.String                                                                     `tfsdk:"serial_number" json:"serial_number,computed"`
	Settings             customfield.NestedObject[CustomHostnamesSSLSettingsDataSourceModel]              `tfsdk:"settings" json:"settings,computed"`
	Signature            types.String                                                                     `tfsdk:"signature" json:"signature,computed"`
	Status               types.String                                                                     `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                                                     `tfsdk:"type" json:"type,computed"`
	UploadedOn           timetypes.RFC3339                                                                `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	ValidationErrors     customfield.NestedObjectList[CustomHostnamesSSLValidationErrorsDataSourceModel]  `tfsdk:"validation_errors" json:"validation_errors,computed"`
	ValidationRecords    customfield.NestedObjectList[CustomHostnamesSSLValidationRecordsDataSourceModel] `tfsdk:"validation_records" json:"validation_records,computed"`
	Wildcard             types.Bool                                                                       `tfsdk:"wildcard" json:"wildcard,computed"`
}

type CustomHostnamesSSLSettingsDataSourceModel struct {
	Ciphers       customfield.List[types.String] `tfsdk:"ciphers" json:"ciphers,computed"`
	EarlyHints    types.String                   `tfsdk:"early_hints" json:"early_hints,computed"`
	HTTP2         types.String                   `tfsdk:"http2" json:"http2,computed"`
	MinTLSVersion types.String                   `tfsdk:"min_tls_version" json:"min_tls_version,computed"`
	TLS1_3        types.String                   `tfsdk:"tls_1_3" json:"tls_1_3,computed"`
}

type CustomHostnamesSSLValidationErrorsDataSourceModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type CustomHostnamesSSLValidationRecordsDataSourceModel struct {
	Emails   customfield.List[types.String] `tfsdk:"emails" json:"emails,computed"`
	HTTPBody types.String                   `tfsdk:"http_body" json:"http_body,computed"`
	HTTPURL  types.String                   `tfsdk:"http_url" json:"http_url,computed"`
	TXTName  types.String                   `tfsdk:"txt_name" json:"txt_name,computed"`
	TXTValue types.String                   `tfsdk:"txt_value" json:"txt_value,computed"`
}

type CustomHostnamesOwnershipVerificationDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type CustomHostnamesOwnershipVerificationHTTPDataSourceModel struct {
	HTTPBody types.String `tfsdk:"http_body" json:"http_body,computed"`
	HTTPURL  types.String `tfsdk:"http_url" json:"http_url,computed"`
}

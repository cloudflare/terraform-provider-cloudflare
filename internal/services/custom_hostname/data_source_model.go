// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/custom_hostnames"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameResultDataSourceEnvelope struct {
Result CustomHostnameDataSourceModel `json:"result,computed"`
}

type CustomHostnameDataSourceModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
CustomHostnameID types.String `tfsdk:"custom_hostname_id" path:"custom_hostname_id,optional"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
CustomOriginServer types.String `tfsdk:"custom_origin_server" json:"custom_origin_server,computed"`
CustomOriginSNI types.String `tfsdk:"custom_origin_sni" json:"custom_origin_sni,computed"`
Hostname types.String `tfsdk:"hostname" json:"hostname,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
CustomMetadata customfield.Map[types.String] `tfsdk:"custom_metadata" json:"custom_metadata,computed"`
VerificationErrors customfield.List[types.String] `tfsdk:"verification_errors" json:"verification_errors,computed"`
OwnershipVerification customfield.NestedObject[CustomHostnameOwnershipVerificationDataSourceModel] `tfsdk:"ownership_verification" json:"ownership_verification,computed"`
OwnershipVerificationHTTP customfield.NestedObject[CustomHostnameOwnershipVerificationHTTPDataSourceModel] `tfsdk:"ownership_verification_http" json:"ownership_verification_http,computed"`
SSL customfield.NestedObject[CustomHostnameSSLDataSourceModel] `tfsdk:"ssl" json:"ssl,computed"`
Filter *CustomHostnameFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CustomHostnameDataSourceModel) toReadParams(_ context.Context) (params custom_hostnames.CustomHostnameGetParams, diags diag.Diagnostics) {
  params = custom_hostnames.CustomHostnameGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

func (m *CustomHostnameDataSourceModel) toListParams(_ context.Context) (params custom_hostnames.CustomHostnameListParams, diags diag.Diagnostics) {
  params = custom_hostnames.CustomHostnameListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  if !m.Filter.ID.IsNull() {
    params.ID = cloudflare.F(m.Filter.ID.ValueString())
  }
  if !m.Filter.Direction.IsNull() {
    params.Direction = cloudflare.F(custom_hostnames.CustomHostnameListParamsDirection(m.Filter.Direction.ValueString()))
  }
  if !m.Filter.Hostname.IsNull() {
    params.Hostname = cloudflare.F(m.Filter.Hostname.ValueString())
  }
  if !m.Filter.Order.IsNull() {
    params.Order = cloudflare.F(custom_hostnames.CustomHostnameListParamsOrder(m.Filter.Order.ValueString()))
  }
  if !m.Filter.SSL.IsNull() {
    params.SSL = cloudflare.F(custom_hostnames.CustomHostnameListParamsSSL(m.Filter.SSL.ValueFloat64()))
  }

  return
}

type CustomHostnameOwnershipVerificationDataSourceModel struct {
Name types.String `tfsdk:"name" json:"name,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

type CustomHostnameOwnershipVerificationHTTPDataSourceModel struct {
HTTPBody types.String `tfsdk:"http_body" json:"http_body,computed"`
HTTPURL types.String `tfsdk:"http_url" json:"http_url,computed"`
}

type CustomHostnameSSLDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
BundleMethod types.String `tfsdk:"bundle_method" json:"bundle_method,computed"`
CertificateAuthority types.String `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
CustomCertificate types.String `tfsdk:"custom_certificate" json:"custom_certificate,computed"`
CustomCsrID types.String `tfsdk:"custom_csr_id" json:"custom_csr_id,computed"`
CustomKey types.String `tfsdk:"custom_key" json:"custom_key,computed"`
ExpiresOn timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
Hosts customfield.List[types.String] `tfsdk:"hosts" json:"hosts,computed"`
Issuer types.String `tfsdk:"issuer" json:"issuer,computed"`
Method types.String `tfsdk:"method" json:"method,computed"`
SerialNumber types.String `tfsdk:"serial_number" json:"serial_number,computed"`
Settings customfield.NestedObject[CustomHostnameSSLSettingsDataSourceModel] `tfsdk:"settings" json:"settings,computed"`
Signature types.String `tfsdk:"signature" json:"signature,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
UploadedOn timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
ValidationErrors customfield.NestedObjectList[CustomHostnameSSLValidationErrorsDataSourceModel] `tfsdk:"validation_errors" json:"validation_errors,computed"`
ValidationRecords customfield.NestedObjectList[CustomHostnameSSLValidationRecordsDataSourceModel] `tfsdk:"validation_records" json:"validation_records,computed"`
Wildcard types.Bool `tfsdk:"wildcard" json:"wildcard,computed"`
}

type CustomHostnameSSLSettingsDataSourceModel struct {
Ciphers customfield.List[types.String] `tfsdk:"ciphers" json:"ciphers,computed"`
EarlyHints types.String `tfsdk:"early_hints" json:"early_hints,computed"`
HTTP2 types.String `tfsdk:"http2" json:"http2,computed"`
MinTLSVersion types.String `tfsdk:"min_tls_version" json:"min_tls_version,computed"`
TLS1_3 types.String `tfsdk:"tls_1_3" json:"tls_1_3,computed"`
}

type CustomHostnameSSLValidationErrorsDataSourceModel struct {
Message types.String `tfsdk:"message" json:"message,computed"`
}

type CustomHostnameSSLValidationRecordsDataSourceModel struct {
Emails customfield.List[types.String] `tfsdk:"emails" json:"emails,computed"`
HTTPBody types.String `tfsdk:"http_body" json:"http_body,computed"`
HTTPURL types.String `tfsdk:"http_url" json:"http_url,computed"`
TXTName types.String `tfsdk:"txt_name" json:"txt_name,computed"`
TXTValue types.String `tfsdk:"txt_value" json:"txt_value,computed"`
}

type CustomHostnameFindOneByDataSourceModel struct {
ID types.String `tfsdk:"id" query:"id,optional"`
Direction types.String `tfsdk:"direction" query:"direction,optional"`
Hostname types.String `tfsdk:"hostname" query:"hostname,optional"`
Order types.String `tfsdk:"order" query:"order,computed_optional"`
SSL types.Float64 `tfsdk:"ssl" query:"ssl,optional"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnameResultEnvelope struct {
Result CustomHostnameModel `json:"result"`
}

type CustomHostnameModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Hostname types.String `tfsdk:"hostname" json:"hostname,required"`
SSL *CustomHostnameSSLModel `tfsdk:"ssl" json:"ssl,required"`
CustomOriginServer types.String `tfsdk:"custom_origin_server" json:"custom_origin_server,optional"`
CustomOriginSNI types.String `tfsdk:"custom_origin_sni" json:"custom_origin_sni,optional"`
CustomMetadata *map[string]types.String `tfsdk:"custom_metadata" json:"custom_metadata,optional"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
Status types.String `tfsdk:"status" json:"status,computed"`
VerificationErrors customfield.List[types.String] `tfsdk:"verification_errors" json:"verification_errors,computed"`
OwnershipVerification customfield.NestedObject[CustomHostnameOwnershipVerificationModel] `tfsdk:"ownership_verification" json:"ownership_verification,computed"`
OwnershipVerificationHTTP customfield.NestedObject[CustomHostnameOwnershipVerificationHTTPModel] `tfsdk:"ownership_verification_http" json:"ownership_verification_http,computed"`
}

func (m CustomHostnameModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m CustomHostnameModel) MarshalJSONForUpdate(state CustomHostnameModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}

type CustomHostnameSSLModel struct {
BundleMethod types.String `tfsdk:"bundle_method" json:"bundle_method,computed_optional"`
CertificateAuthority types.String `tfsdk:"certificate_authority" json:"certificate_authority,optional"`
CloudflareBranding types.Bool `tfsdk:"cloudflare_branding" json:"cloudflare_branding,optional"`
CustomCERTBundle *[]*CustomHostnameSSLCustomCERTBundleModel `tfsdk:"custom_cert_bundle" json:"custom_cert_bundle,optional"`
CustomCertificate types.String `tfsdk:"custom_certificate" json:"custom_certificate,optional"`
CustomKey types.String `tfsdk:"custom_key" json:"custom_key,optional"`
Method types.String `tfsdk:"method" json:"method,optional"`
Settings *CustomHostnameSSLSettingsModel `tfsdk:"settings" json:"settings,optional"`
Type types.String `tfsdk:"type" json:"type,optional"`
Wildcard types.Bool `tfsdk:"wildcard" json:"wildcard,optional"`
}

type CustomHostnameSSLCustomCERTBundleModel struct {
CustomCertificate types.String `tfsdk:"custom_certificate" json:"custom_certificate,required"`
CustomKey types.String `tfsdk:"custom_key" json:"custom_key,required"`
}

type CustomHostnameSSLSettingsModel struct {
Ciphers *[]types.String `tfsdk:"ciphers" json:"ciphers,optional"`
EarlyHints types.String `tfsdk:"early_hints" json:"early_hints,optional"`
HTTP2 types.String `tfsdk:"http2" json:"http2,optional"`
MinTLSVersion types.String `tfsdk:"min_tls_version" json:"min_tls_version,optional"`
TLS1_3 types.String `tfsdk:"tls_1_3" json:"tls_1_3,optional"`
}

type CustomHostnameOwnershipVerificationModel struct {
Name types.String `tfsdk:"name" json:"name,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

type CustomHostnameOwnershipVerificationHTTPModel struct {
HTTPBody types.String `tfsdk:"http_body" json:"http_body,computed"`
HTTPURL types.String `tfsdk:"http_url" json:"http_url,computed"`
}

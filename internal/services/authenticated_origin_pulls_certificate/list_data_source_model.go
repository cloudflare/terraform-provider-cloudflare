// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/origin_tls_client_auth"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsCertificatesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[AuthenticatedOriginPullsCertificatesResultDataSourceModel] `json:"result,computed"`
}

type AuthenticatedOriginPullsCertificatesDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[AuthenticatedOriginPullsCertificatesResultDataSourceModel] `tfsdk:"result"`
}

func (m *AuthenticatedOriginPullsCertificatesDataSourceModel) toListParams(_ context.Context) (params origin_tls_client_auth.OriginTLSClientAuthListParams, diags diag.Diagnostics) {
  params = origin_tls_client_auth.OriginTLSClientAuthListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

type AuthenticatedOriginPullsCertificatesResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Certificate types.String `tfsdk:"certificate" json:"certificate,computed"`
ExpiresOn timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
Issuer types.String `tfsdk:"issuer" json:"issuer,computed"`
Signature types.String `tfsdk:"signature" json:"signature,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
UploadedOn timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
PrivateKey types.String `tfsdk:"private_key" json:"private_key,computed"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/keyless_certificates"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificatesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[KeylessCertificatesResultDataSourceModel] `json:"result,computed"`
}

type KeylessCertificatesDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[KeylessCertificatesResultDataSourceModel] `tfsdk:"result"`
}

func (m *KeylessCertificatesDataSourceModel) toListParams(_ context.Context) (params keyless_certificates.KeylessCertificateListParams, diags diag.Diagnostics) {
  params = keyless_certificates.KeylessCertificateListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

type KeylessCertificatesResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
Host types.String `tfsdk:"host" json:"host,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
Permissions customfield.List[types.String] `tfsdk:"permissions" json:"permissions,computed"`
Port types.Float64 `tfsdk:"port" json:"port,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
Tunnel customfield.NestedObject[KeylessCertificatesTunnelDataSourceModel] `tfsdk:"tunnel" json:"tunnel,computed"`
}

type KeylessCertificatesTunnelDataSourceModel struct {
PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
VnetID types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/ssl"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultDataSourceEnvelope struct {
Result CertificatePackDataSourceModel `json:"result,computed"`
}

type CertificatePackDataSourceModel struct {
CertificatePackID types.String `tfsdk:"certificate_pack_id" path:"certificate_pack_id,required"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *CertificatePackDataSourceModel) toReadParams(_ context.Context) (params ssl.CertificatePackGetParams, diags diag.Diagnostics) {
  params = ssl.CertificatePackGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

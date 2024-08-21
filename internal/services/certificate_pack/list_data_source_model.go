// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/ssl"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePacksResultListDataSourceEnvelope struct {
	Result *[]*CertificatePacksResultDataSourceModel `json:"result,computed"`
}

type CertificatePacksDataSourceModel struct {
	ZoneID   types.String                              `tfsdk:"zone_id" path:"zone_id"`
	Status   types.String                              `tfsdk:"status" query:"status"`
	MaxItems types.Int64                               `tfsdk:"max_items"`
	Result   *[]*CertificatePacksResultDataSourceModel `tfsdk:"result"`
}

func (m *CertificatePacksDataSourceModel) toListParams() (params ssl.CertificatePackListParams, diags diag.Diagnostics) {
	params = ssl.CertificatePackListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Status.IsNull() {
		params.Status = cloudflare.F(ssl.CertificatePackListParamsStatus(m.Status.ValueString()))
	}

	return
}

type CertificatePacksResultDataSourceModel struct {
}

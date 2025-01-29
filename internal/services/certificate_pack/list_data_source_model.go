// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/ssl"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePacksResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CertificatePacksResultDataSourceModel] `json:"result,computed"`
}

type CertificatePacksDataSourceModel struct {
	ZoneID   types.String                                                        `tfsdk:"zone_id" path:"zone_id,required"`
	Status   types.String                                                        `tfsdk:"status" query:"status,optional"`
	MaxItems types.Int64                                                         `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[CertificatePacksResultDataSourceModel] `tfsdk:"result"`
}

func (m *CertificatePacksDataSourceModel) toListParams(_ context.Context) (params ssl.CertificatePackListParams, diags diag.Diagnostics) {
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

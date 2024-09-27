// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/ssl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultDataSourceEnvelope struct {
	Result CertificatePackDataSourceModel `json:"result,computed"`
}

type CertificatePackResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CertificatePackDataSourceModel] `json:"result,computed"`
}

type CertificatePackDataSourceModel struct {
	CertificatePackID types.String                             `tfsdk:"certificate_pack_id" path:"certificate_pack_id,optional"`
	ZoneID            types.String                             `tfsdk:"zone_id" path:"zone_id,optional"`
	Filter            *CertificatePackFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CertificatePackDataSourceModel) toReadParams(_ context.Context) (params ssl.CertificatePackGetParams, diags diag.Diagnostics) {
	params = ssl.CertificatePackGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *CertificatePackDataSourceModel) toListParams(_ context.Context) (params ssl.CertificatePackListParams, diags diag.Diagnostics) {
	params = ssl.CertificatePackListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(ssl.CertificatePackListParamsStatus(m.Filter.Status.ValueString()))
	}

	return
}

type CertificatePackFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Status types.String `tfsdk:"status" query:"status,optional"`
}

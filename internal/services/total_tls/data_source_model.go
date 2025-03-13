// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package total_tls

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/acm"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TotalTLSResultDataSourceEnvelope struct {
	Result TotalTLSDataSourceModel `json:"result,computed"`
}

type TotalTLSDataSourceModel struct {
	ZoneID               types.String `tfsdk:"zone_id" path:"zone_id,required"`
	CertificateAuthority types.String `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	Enabled              types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	ValidityPeriod       types.Int64  `tfsdk:"validity_period" json:"validity_period,computed"`
}

func (m *TotalTLSDataSourceModel) toReadParams(_ context.Context) (params acm.TotalTLSGetParams, diags diag.Diagnostics) {
	params = acm.TotalTLSGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

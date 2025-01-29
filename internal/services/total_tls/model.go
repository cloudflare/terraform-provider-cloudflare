// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package total_tls

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TotalTLSResultEnvelope struct {
	Result TotalTLSModel `json:"result"`
}

type TotalTLSModel struct {
	ID                   types.String `tfsdk:"id" json:"-,computed"`
	ZoneID               types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled              types.Bool   `tfsdk:"enabled" json:"enabled,required"`
	CertificateAuthority types.String `tfsdk:"certificate_authority" json:"certificate_authority,optional"`
	ValidityPeriod       types.Int64  `tfsdk:"validity_period" json:"validity_period,computed"`
}

func (m TotalTLSModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TotalTLSModel) MarshalJSONForUpdate(state TotalTLSModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

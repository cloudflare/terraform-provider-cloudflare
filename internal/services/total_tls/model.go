// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package total_tls

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TotalTLSResultEnvelope struct {
	Result TotalTLSModel `json:"result"`
}

type TotalTLSModel struct {
	ID                   types.String `tfsdk:"id" json:"-,computed"`
	ZoneID               types.String `tfsdk:"zone_id" path:"zone_id"`
	CertificateAuthority types.String `tfsdk:"certificate_authority" json:"certificate_authority,computed_optional"`
	Enabled              types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	ValidityPeriod       types.Int64  `tfsdk:"validity_period" json:"validity_period,computed"`
}

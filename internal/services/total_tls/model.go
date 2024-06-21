// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package total_tls

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TotalTLSResultEnvelope struct {
	Result TotalTLSModel `json:"result,computed"`
}

type TotalTLSModel struct {
	ZoneID               types.String `tfsdk:"zone_id" path:"zone_id"`
	Enabled              types.Bool   `tfsdk:"enabled" json:"enabled"`
	CertificateAuthority types.String `tfsdk:"certificate_authority" json:"certificate_authority"`
}

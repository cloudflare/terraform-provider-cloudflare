// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package total_tls

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TotalTLSResultDataSourceEnvelope struct {
	Result TotalTLSDataSourceModel `json:"result,computed"`
}

type TotalTLSDataSourceModel struct {
	ZoneID               types.String `tfsdk:"zone_id" path:"zone_id"`
	CertificateAuthority types.String `tfsdk:"certificate_authority" json:"certificate_authority"`
	Enabled              types.Bool   `tfsdk:"enabled" json:"enabled"`
	ValidityDays         types.Int64  `tfsdk:"validity_days" json:"validity_days"`
}

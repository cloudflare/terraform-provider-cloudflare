// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultEnvelope struct {
	Result CertificatePackModel `json:"result,computed"`
}

type CertificatePackModel struct {
	ID                types.String `tfsdk:"id" json:"id,computed"`
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id"`
	CertificatePackID types.String `tfsdk:"certificate_pack_id" path:"certificate_pack_id"`
}

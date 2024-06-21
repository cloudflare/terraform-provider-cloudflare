// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsCertificateResultEnvelope struct {
	Result AuthenticatedOriginPullsCertificateModel `json:"result,computed"`
}

type AuthenticatedOriginPullsCertificateModel struct {
	ZoneID        types.String `tfsdk:"zone_id" path:"zone_id"`
	CertificateID types.String `tfsdk:"certificate_id" path:"certificate_id"`
	Certificate   types.String `tfsdk:"certificate" json:"certificate"`
	PrivateKey    types.String `tfsdk:"private_key" json:"private_key"`
}

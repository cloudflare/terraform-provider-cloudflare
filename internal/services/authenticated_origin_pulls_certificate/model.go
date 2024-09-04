// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsCertificateResultEnvelope struct {
	Result AuthenticatedOriginPullsCertificateModel `json:"result"`
}

type AuthenticatedOriginPullsCertificateModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String      `tfsdk:"zone_id" path:"zone_id"`
	PrivateKey  types.String      `tfsdk:"private_key" json:"private_key"`
	Certificate types.String      `tfsdk:"certificate" json:"certificate,computed_optional"`
	ExpiresOn   timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer      types.String      `tfsdk:"issuer" json:"issuer,computed"`
	Signature   types.String      `tfsdk:"signature" json:"signature,computed"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
	UploadedOn  timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
}

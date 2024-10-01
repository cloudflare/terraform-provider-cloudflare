// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsCertificateResultEnvelope struct {
	Result AuthenticatedOriginPullsCertificateModel `json:"result"`
}

type AuthenticatedOriginPullsCertificateModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Certificate types.String      `tfsdk:"certificate" json:"certificate,required"`
	PrivateKey  types.String      `tfsdk:"private_key" json:"private_key,required"`
	ExpiresOn   timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer      types.String      `tfsdk:"issuer" json:"issuer,computed"`
	Signature   types.String      `tfsdk:"signature" json:"signature,computed"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
	UploadedOn  timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
}

func (m AuthenticatedOriginPullsCertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AuthenticatedOriginPullsCertificateModel) MarshalJSONForUpdate(state AuthenticatedOriginPullsCertificateModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

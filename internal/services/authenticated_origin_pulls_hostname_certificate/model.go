// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_hostname_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsHostnameCertificateResultEnvelope struct {
	Result AuthenticatedOriginPullsHostnameCertificateModel `json:"result"`
}

type AuthenticatedOriginPullsHostnameCertificateModel struct {
	ZoneID        types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	CertificateID types.String      `tfsdk:"certificate_id" path:"certificate_id,computed"`
	Certificate   types.String      `tfsdk:"certificate" json:"certificate,required"`
	PrivateKey    types.String      `tfsdk:"private_key" json:"private_key,required"`
	ExpiresOn     timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	Issuer        types.String      `tfsdk:"issuer" json:"issuer,computed"`
	SerialNumber  types.String      `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature     types.String      `tfsdk:"signature" json:"signature,computed"`
	Status        types.String      `tfsdk:"status" json:"status,computed"`
	UploadedOn    timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
}

func (m AuthenticatedOriginPullsHostnameCertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AuthenticatedOriginPullsHostnameCertificateModel) MarshalJSONForUpdate(state AuthenticatedOriginPullsHostnameCertificateModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

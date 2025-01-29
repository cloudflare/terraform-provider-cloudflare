// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MTLSCertificateResultEnvelope struct {
	Result MTLSCertificateModel `json:"result"`
}

type MTLSCertificateModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	AccountID    types.String      `tfsdk:"account_id" path:"account_id,required"`
	CA           types.Bool        `tfsdk:"ca" json:"ca,required"`
	Certificates types.String      `tfsdk:"certificates" json:"certificates,required"`
	Name         types.String      `tfsdk:"name" json:"name,optional"`
	PrivateKey   types.String      `tfsdk:"private_key" json:"private_key,optional"`
	ExpiresOn    timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer       types.String      `tfsdk:"issuer" json:"issuer,computed"`
	SerialNumber types.String      `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature    types.String      `tfsdk:"signature" json:"signature,computed"`
	UpdatedAt    timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UploadedOn   timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
}

func (m MTLSCertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MTLSCertificateModel) MarshalJSONForUpdate(state MTLSCertificateModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

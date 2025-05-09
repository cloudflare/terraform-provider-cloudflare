// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessMTLSCertificateResultEnvelope struct {
	Result ZeroTrustAccessMTLSCertificateModel `json:"result"`
}

type ZeroTrustAccessMTLSCertificateModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	AccountID           types.String      `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID              types.String      `tfsdk:"zone_id" path:"zone_id,optional"`
	Certificate         types.String      `tfsdk:"certificate" json:"certificate,required,no_refresh"`
	Name                types.String      `tfsdk:"name" json:"name,required"`
	AssociatedHostnames *[]types.String   `tfsdk:"associated_hostnames" json:"associated_hostnames,optional"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresOn           timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Fingerprint         types.String      `tfsdk:"fingerprint" json:"fingerprint,computed"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustAccessMTLSCertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessMTLSCertificateModel) MarshalJSONForUpdate(state ZeroTrustAccessMTLSCertificateModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

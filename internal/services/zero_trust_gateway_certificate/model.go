// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayCertificateResultEnvelope struct {
	Result ZeroTrustGatewayCertificateModel `json:"result"`
}

type ZeroTrustGatewayCertificateModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	AccountID          types.String      `tfsdk:"account_id" path:"account_id,required"`
	ValidityPeriodDays types.Int64       `tfsdk:"validity_period_days" json:"validity_period_days,optional"`
	BindingStatus      types.String      `tfsdk:"binding_status" json:"binding_status,computed"`
	Certificate        types.String      `tfsdk:"certificate" json:"certificate,computed"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresOn          timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Fingerprint        types.String      `tfsdk:"fingerprint" json:"fingerprint,computed"`
	InUse              types.Bool        `tfsdk:"in_use" json:"in_use,computed"`
	IssuerOrg          types.String      `tfsdk:"issuer_org" json:"issuer_org,computed"`
	IssuerRaw          types.String      `tfsdk:"issuer_raw" json:"issuer_raw,computed"`
	Type               types.String      `tfsdk:"type" json:"type,computed"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UploadedOn         timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
}

func (m ZeroTrustGatewayCertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustGatewayCertificateModel) MarshalJSONForUpdate(state ZeroTrustGatewayCertificateModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

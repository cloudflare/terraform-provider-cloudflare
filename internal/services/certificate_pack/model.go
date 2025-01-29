// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultEnvelope struct {
	Result CertificatePackModel `json:"result"`
}

type CertificatePackModel struct {
	ID                   types.String    `tfsdk:"id" json:"id,computed"`
	ZoneID               types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	CertificateAuthority types.String    `tfsdk:"certificate_authority" json:"certificate_authority,required"`
	Type                 types.String    `tfsdk:"type" json:"type,required"`
	ValidationMethod     types.String    `tfsdk:"validation_method" json:"validation_method,required"`
	ValidityDays         types.Int64     `tfsdk:"validity_days" json:"validity_days,required"`
	Hosts                *[]types.String `tfsdk:"hosts" json:"hosts,required"`
	CloudflareBranding   types.Bool      `tfsdk:"cloudflare_branding" json:"cloudflare_branding,optional"`
	Status               types.String    `tfsdk:"status" json:"status,computed"`
}

func (m CertificatePackModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CertificatePackModel) MarshalJSONForUpdate(state CertificatePackModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

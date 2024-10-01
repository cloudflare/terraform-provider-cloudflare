// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultEnvelope struct {
	Result CertificatePackModel `json:"result"`
}

type CertificatePackModel struct {
	ID                   types.String                   `tfsdk:"id" json:"-,computed"`
	CertificatePackID    types.String                   `tfsdk:"certificate_pack_id" path:"certificate_pack_id,required"`
	ZoneID               types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	CertificateAuthority types.String                   `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CloudflareBranding   types.Bool                     `tfsdk:"cloudflare_branding" json:"cloudflare_branding,computed"`
	Status               types.String                   `tfsdk:"status" json:"status,computed"`
	Type                 types.String                   `tfsdk:"type" json:"type,computed"`
	ValidationMethod     types.String                   `tfsdk:"validation_method" json:"validation_method,computed"`
	ValidityDays         types.Int64                    `tfsdk:"validity_days" json:"validity_days,computed"`
	Hosts                customfield.List[types.String] `tfsdk:"hosts" json:"hosts,computed"`
}

func (m CertificatePackModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CertificatePackModel) MarshalJSONForUpdate(state CertificatePackModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

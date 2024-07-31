// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultEnvelope struct {
	Result CertificatePackModel `json:"result,computed"`
}

type CertificatePackModel struct {
	ID                   types.String    `tfsdk:"id" json:"-,computed"`
	ZoneID               types.String    `tfsdk:"zone_id" path:"zone_id"`
	CertificatePackID    types.String    `tfsdk:"certificate_pack_id" path:"certificate_pack_id"`
	CertificateAuthority types.String    `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CloudflareBranding   types.Bool      `tfsdk:"cloudflare_branding" json:"cloudflare_branding,computed"`
	Hosts                *[]types.String `tfsdk:"hosts" json:"hosts,computed"`
	Status               types.String    `tfsdk:"status" json:"status,computed"`
	Type                 types.String    `tfsdk:"type" json:"type,computed"`
	ValidationMethod     types.String    `tfsdk:"validation_method" json:"validation_method,computed"`
	ValidityDays         types.Int64     `tfsdk:"validity_days" json:"validity_days,computed"`
}

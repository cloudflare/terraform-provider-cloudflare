// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ClientCertificateResultEnvelope struct {
	Result ClientCertificateModel `json:"result"`
}

type ClientCertificateModel struct {
	ID                   types.String                                                         `tfsdk:"id" json:"id,computed"`
	ZoneID               types.String                                                         `tfsdk:"zone_id" path:"zone_id,required"`
	Csr                  types.String                                                         `tfsdk:"csr" json:"csr,required"`
	ValidityDays         types.Int64                                                          `tfsdk:"validity_days" json:"validity_days,required"`
	Reactivate           types.Bool                                                           `tfsdk:"reactivate" json:"reactivate,optional,no_refresh"`
	Certificate          types.String                                                         `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           types.String                                                         `tfsdk:"common_name" json:"common_name,computed"`
	Country              types.String                                                         `tfsdk:"country" json:"country,computed"`
	ExpiresOn            types.String                                                         `tfsdk:"expires_on" json:"expires_on,computed"`
	FingerprintSha256    types.String                                                         `tfsdk:"fingerprint_sha256" json:"fingerprint_sha256,computed"`
	IssuedOn             types.String                                                         `tfsdk:"issued_on" json:"issued_on,computed"`
	Location             types.String                                                         `tfsdk:"location" json:"location,computed"`
	Organization         types.String                                                         `tfsdk:"organization" json:"organization,computed"`
	OrganizationalUnit   types.String                                                         `tfsdk:"organizational_unit" json:"organizational_unit,computed"`
	SerialNumber         types.String                                                         `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature            types.String                                                         `tfsdk:"signature" json:"signature,computed"`
	Ski                  types.String                                                         `tfsdk:"ski" json:"ski,computed"`
	State                types.String                                                         `tfsdk:"state" json:"state,computed"`
	Status               types.String                                                         `tfsdk:"status" json:"status,computed"`
	CertificateAuthority customfield.NestedObject[ClientCertificateCertificateAuthorityModel] `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
}

func (m ClientCertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ClientCertificateModel) MarshalJSONForUpdate(state ClientCertificateModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type ClientCertificateCertificateAuthorityModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

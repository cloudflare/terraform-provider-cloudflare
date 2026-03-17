// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_authorities_hostname_associations

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificateAuthoritiesHostnameAssociationsResultEnvelope struct {
	Result CertificateAuthoritiesHostnameAssociationsModel `json:"result"`
}

type CertificateAuthoritiesHostnameAssociationsModel struct {
	ID                types.String    `tfsdk:"id" json:"-,computed"`
	ZoneID            types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	MTLSCertificateID types.String    `tfsdk:"mtls_certificate_id" json:"mtls_certificate_id,optional,no_refresh"`
	Hostnames         *[]types.String `tfsdk:"hostnames" json:"hostnames,optional"`
}

func (m CertificateAuthoritiesHostnameAssociationsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CertificateAuthoritiesHostnameAssociationsModel) MarshalJSONForUpdate(state CertificateAuthoritiesHostnameAssociationsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

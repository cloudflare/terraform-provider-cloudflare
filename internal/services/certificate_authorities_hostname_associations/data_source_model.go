// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_authorities_hostname_associations

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/certificate_authorities"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificateAuthoritiesHostnameAssociationsResultDataSourceEnvelope struct {
	Result CertificateAuthoritiesHostnameAssociationsDataSourceModel `json:"result,computed"`
}

type CertificateAuthoritiesHostnameAssociationsDataSourceModel struct {
	ID                types.String                   `tfsdk:"id" path:"zone_id,computed"`
	ZoneID            types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	MTLSCertificateID types.String                   `tfsdk:"mtls_certificate_id" query:"mtls_certificate_id,optional"`
	Hostnames         customfield.List[types.String] `tfsdk:"hostnames" json:"hostnames,computed"`
}

func (m *CertificateAuthoritiesHostnameAssociationsDataSourceModel) toReadParams(_ context.Context) (params certificate_authorities.HostnameAssociationGetParams, diags diag.Diagnostics) {
	params = certificate_authorities.HostnameAssociationGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.MTLSCertificateID.IsNull() {
		params.MTLSCertificateID = cloudflare.F(m.MTLSCertificateID.ValueString())
	}

	return
}

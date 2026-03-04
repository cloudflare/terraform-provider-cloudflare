// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate_associations

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/mtls_certificates"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MTLSCertificateAssociationsResultDataSourceEnvelope struct {
	Result MTLSCertificateAssociationsDataSourceModel `json:"result,computed"`
}

type MTLSCertificateAssociationsDataSourceModel struct {
	AccountID         types.String `tfsdk:"account_id" path:"account_id,required"`
	MTLSCertificateID types.String `tfsdk:"mtls_certificate_id" path:"mtls_certificate_id,required"`
	Service           types.String `tfsdk:"service" json:"service,computed"`
	Status            types.String `tfsdk:"status" json:"status,computed"`
}

func (m *MTLSCertificateAssociationsDataSourceModel) toReadParams(_ context.Context) (params mtls_certificates.AssociationGetParams, diags diag.Diagnostics) {
	params = mtls_certificates.AssociationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

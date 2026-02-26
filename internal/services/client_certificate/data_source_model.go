// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_certificate

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/client_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ClientCertificateResultDataSourceEnvelope struct {
	Result ClientCertificateDataSourceModel `json:"result,computed"`
}

type ClientCertificateDataSourceModel struct {
	ID                   types.String                                                                   `tfsdk:"id" path:"client_certificate_id,computed"`
	ClientCertificateID  types.String                                                                   `tfsdk:"client_certificate_id" path:"client_certificate_id,optional"`
	ZoneID               types.String                                                                   `tfsdk:"zone_id" path:"zone_id,required"`
	Certificate          types.String                                                                   `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           types.String                                                                   `tfsdk:"common_name" json:"common_name,computed"`
	Country              types.String                                                                   `tfsdk:"country" json:"country,computed"`
	Csr                  types.String                                                                   `tfsdk:"csr" json:"csr,computed"`
	ExpiresOn            types.String                                                                   `tfsdk:"expires_on" json:"expires_on,computed"`
	FingerprintSha256    types.String                                                                   `tfsdk:"fingerprint_sha256" json:"fingerprint_sha256,computed"`
	IssuedOn             types.String                                                                   `tfsdk:"issued_on" json:"issued_on,computed"`
	Location             types.String                                                                   `tfsdk:"location" json:"location,computed"`
	Organization         types.String                                                                   `tfsdk:"organization" json:"organization,computed"`
	OrganizationalUnit   types.String                                                                   `tfsdk:"organizational_unit" json:"organizational_unit,computed"`
	SerialNumber         types.String                                                                   `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature            types.String                                                                   `tfsdk:"signature" json:"signature,computed"`
	Ski                  types.String                                                                   `tfsdk:"ski" json:"ski,computed"`
	State                types.String                                                                   `tfsdk:"state" json:"state,computed"`
	Status               types.String                                                                   `tfsdk:"status" json:"status,computed"`
	ValidityDays         types.Int64                                                                    `tfsdk:"validity_days" json:"validity_days,computed"`
	CertificateAuthority customfield.NestedObject[ClientCertificateCertificateAuthorityDataSourceModel] `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	Filter               *ClientCertificateFindOneByDataSourceModel                                     `tfsdk:"filter"`
}

func (m *ClientCertificateDataSourceModel) toReadParams(_ context.Context) (params client_certificates.ClientCertificateGetParams, diags diag.Diagnostics) {
	params = client_certificates.ClientCertificateGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *ClientCertificateDataSourceModel) toListParams(_ context.Context) (params client_certificates.ClientCertificateListParams, diags diag.Diagnostics) {
	params = client_certificates.ClientCertificateListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Filter.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Filter.Limit.ValueInt64())
	}
	if !m.Filter.Offset.IsNull() {
		params.Offset = cloudflare.F(m.Filter.Offset.ValueInt64())
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(client_certificates.ClientCertificateListParamsStatus(m.Filter.Status.ValueString()))
	}

	return
}

type ClientCertificateCertificateAuthorityDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type ClientCertificateFindOneByDataSourceModel struct {
	Limit  types.Int64  `tfsdk:"limit" query:"limit,optional"`
	Offset types.Int64  `tfsdk:"offset" query:"offset,optional"`
	Status types.String `tfsdk:"status" query:"status,optional"`
}

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

type ClientCertificatesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ClientCertificatesResultDataSourceModel] `json:"result,computed"`
}

type ClientCertificatesDataSourceModel struct {
	ZoneID   types.String                                                          `tfsdk:"zone_id" path:"zone_id,required"`
	Limit    types.Int64                                                           `tfsdk:"limit" query:"limit,optional"`
	Offset   types.Int64                                                           `tfsdk:"offset" query:"offset,optional"`
	Status   types.String                                                          `tfsdk:"status" query:"status,optional"`
	MaxItems types.Int64                                                           `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[ClientCertificatesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ClientCertificatesDataSourceModel) toListParams(_ context.Context) (params client_certificates.ClientCertificateListParams, diags diag.Diagnostics) {
	params = client_certificates.ClientCertificateListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Limit.ValueInt64())
	}
	if !m.Offset.IsNull() {
		params.Offset = cloudflare.F(m.Offset.ValueInt64())
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(client_certificates.ClientCertificateListParamsStatus(m.Status.ValueString()))
	}

	return
}

type ClientCertificatesResultDataSourceModel struct {
	ID                   types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Certificate          types.String                                                                    `tfsdk:"certificate" json:"certificate,computed"`
	CertificateAuthority customfield.NestedObject[ClientCertificatesCertificateAuthorityDataSourceModel] `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CommonName           types.String                                                                    `tfsdk:"common_name" json:"common_name,computed"`
	Country              types.String                                                                    `tfsdk:"country" json:"country,computed"`
	Csr                  types.String                                                                    `tfsdk:"csr" json:"csr,computed"`
	ExpiresOn            types.String                                                                    `tfsdk:"expires_on" json:"expires_on,computed"`
	FingerprintSha256    types.String                                                                    `tfsdk:"fingerprint_sha256" json:"fingerprint_sha256,computed"`
	IssuedOn             types.String                                                                    `tfsdk:"issued_on" json:"issued_on,computed"`
	Location             types.String                                                                    `tfsdk:"location" json:"location,computed"`
	Organization         types.String                                                                    `tfsdk:"organization" json:"organization,computed"`
	OrganizationalUnit   types.String                                                                    `tfsdk:"organizational_unit" json:"organizational_unit,computed"`
	SerialNumber         types.String                                                                    `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature            types.String                                                                    `tfsdk:"signature" json:"signature,computed"`
	Ski                  types.String                                                                    `tfsdk:"ski" json:"ski,computed"`
	State                types.String                                                                    `tfsdk:"state" json:"state,computed"`
	Status               types.String                                                                    `tfsdk:"status" json:"status,computed"`
	ValidityDays         types.Int64                                                                     `tfsdk:"validity_days" json:"validity_days,computed"`
}

type ClientCertificatesCertificateAuthorityDataSourceModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

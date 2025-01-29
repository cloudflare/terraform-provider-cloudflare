// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/origin_ca_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificateResultDataSourceEnvelope struct {
	Result OriginCACertificateDataSourceModel `json:"result,computed"`
}

type OriginCACertificateResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[OriginCACertificateDataSourceModel] `json:"result,computed"`
}

type OriginCACertificateDataSourceModel struct {
	ID                types.String                                 `tfsdk:"id" json:"-,computed"`
	CertificateID     types.String                                 `tfsdk:"certificate_id" path:"certificate_id,optional"`
	Certificate       types.String                                 `tfsdk:"certificate" json:"certificate,computed"`
	Csr               types.String                                 `tfsdk:"csr" json:"csr,computed"`
	ExpiresOn         types.String                                 `tfsdk:"expires_on" json:"expires_on,computed"`
	RequestType       types.String                                 `tfsdk:"request_type" json:"request_type,computed"`
	RequestedValidity types.Float64                                `tfsdk:"requested_validity" json:"requested_validity,computed"`
	Hostnames         customfield.List[types.String]               `tfsdk:"hostnames" json:"hostnames,computed"`
	Filter            *OriginCACertificateFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *OriginCACertificateDataSourceModel) toListParams(_ context.Context) (params origin_ca_certificates.OriginCACertificateListParams, diags diag.Diagnostics) {
	params = origin_ca_certificates.OriginCACertificateListParams{}

	if !m.Filter.ZoneID.IsNull() {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type OriginCACertificateFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" query:"zone_id,optional"`
}

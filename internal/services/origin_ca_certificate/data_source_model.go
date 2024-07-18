// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OriginCACertificateResultDataSourceEnvelope struct {
	Result OriginCACertificateDataSourceModel `json:"result,computed"`
}

type OriginCACertificateResultListDataSourceEnvelope struct {
	Result *[]*OriginCACertificateDataSourceModel `json:"result,computed"`
}

type OriginCACertificateDataSourceModel struct {
	CertificateID     types.String                                 `tfsdk:"certificate_id" path:"certificate_id"`
	Csr               types.String                                 `tfsdk:"csr" json:"csr"`
	Hostnames         *[]types.String                              `tfsdk:"hostnames" json:"hostnames"`
	RequestType       types.String                                 `tfsdk:"request_type" json:"request_type"`
	RequestedValidity types.Float64                                `tfsdk:"requested_validity" json:"requested_validity"`
	ID                types.String                                 `tfsdk:"id" json:"id"`
	Certificate       types.String                                 `tfsdk:"certificate" json:"certificate"`
	ExpiresOn         types.String                                 `tfsdk:"expires_on" json:"expires_on"`
	FindOneBy         *OriginCACertificateFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type OriginCACertificateFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" query:"zone_id"`
}

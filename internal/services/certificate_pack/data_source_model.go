// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultDataSourceEnvelope struct {
	Result CertificatePackDataSourceModel `json:"result,computed"`
}

type CertificatePackResultListDataSourceEnvelope struct {
	Result *[]*CertificatePackDataSourceModel `json:"result,computed"`
}

type CertificatePackDataSourceModel struct {
	ZoneID            types.String                             `tfsdk:"zone_id" path:"zone_id"`
	CertificatePackID types.String                             `tfsdk:"certificate_pack_id" path:"certificate_pack_id"`
	FindOneBy         *CertificatePackFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type CertificatePackFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Status types.String `tfsdk:"status" query:"status"`
}

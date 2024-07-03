// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificateResultDataSourceEnvelope struct {
	Result KeylessCertificateDataSourceModel `json:"result,computed"`
}

type KeylessCertificateResultListDataSourceEnvelope struct {
	Result *[]*KeylessCertificateDataSourceModel `json:"result,computed"`
}

type KeylessCertificateDataSourceModel struct {
	ZoneID               types.String                                `tfsdk:"zone_id" path:"zone_id"`
	KeylessCertificateID types.String                                `tfsdk:"keyless_certificate_id" path:"keyless_certificate_id"`
	ID                   types.String                                `tfsdk:"id" json:"id"`
	CreatedOn            types.String                                `tfsdk:"created_on" json:"created_on"`
	Enabled              types.Bool                                  `tfsdk:"enabled" json:"enabled"`
	Host                 types.String                                `tfsdk:"host" json:"host"`
	ModifiedOn           types.String                                `tfsdk:"modified_on" json:"modified_on"`
	Name                 types.String                                `tfsdk:"name" json:"name"`
	Permissions          types.String                                `tfsdk:"permissions" json:"permissions"`
	Port                 types.Float64                               `tfsdk:"port" json:"port"`
	Status               types.String                                `tfsdk:"status" json:"status"`
	Tunnel               *KeylessCertificateTunnelDataSourceModel    `tfsdk:"tunnel" json:"tunnel"`
	FindOneBy            *KeylessCertificateFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type KeylessCertificateTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id"`
}

type KeylessCertificateFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}

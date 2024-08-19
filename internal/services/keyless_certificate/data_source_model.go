// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificateResultDataSourceEnvelope struct {
	Result KeylessCertificateDataSourceModel `json:"result,computed"`
}

type KeylessCertificateResultListDataSourceEnvelope struct {
	Result *[]*KeylessCertificateDataSourceModel `json:"result,computed"`
}

type KeylessCertificateDataSourceModel struct {
	KeylessCertificateID types.String                                `tfsdk:"keyless_certificate_id" path:"keyless_certificate_id"`
	ZoneID               types.String                                `tfsdk:"zone_id" path:"zone_id"`
	CreatedOn            timetypes.RFC3339                           `tfsdk:"created_on" json:"created_on,computed"`
	Enabled              types.Bool                                  `tfsdk:"enabled" json:"enabled,computed"`
	Host                 types.String                                `tfsdk:"host" json:"host,computed"`
	ID                   types.String                                `tfsdk:"id" json:"id,computed"`
	ModifiedOn           timetypes.RFC3339                           `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                 types.String                                `tfsdk:"name" json:"name,computed"`
	Port                 types.Float64                               `tfsdk:"port" json:"port,computed"`
	Status               types.String                                `tfsdk:"status" json:"status,computed"`
	Permissions          *[]types.String                             `tfsdk:"permissions" json:"permissions,computed"`
	Tunnel               *KeylessCertificateTunnelDataSourceModel    `tfsdk:"tunnel" json:"tunnel"`
	Filter               *KeylessCertificateFindOneByDataSourceModel `tfsdk:"filter"`
}

type KeylessCertificateTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

type KeylessCertificateFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}

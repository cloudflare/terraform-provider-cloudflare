// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificateResultEnvelope struct {
	Result KeylessCertificateModel `json:"result,computed"`
}

type KeylessCertificateResultDataSourceEnvelope struct {
	Result KeylessCertificateDataSourceModel `json:"result,computed"`
}

type KeylessCertificatesResultDataSourceEnvelope struct {
	Result KeylessCertificatesDataSourceModel `json:"result,computed"`
}

type KeylessCertificateModel struct {
	ID           types.String                   `tfsdk:"id" json:"id,computed"`
	ZoneID       types.String                   `tfsdk:"zone_id" path:"zone_id"`
	Host         types.String                   `tfsdk:"host" json:"host"`
	Port         types.Float64                  `tfsdk:"port" json:"port"`
	Name         types.String                   `tfsdk:"name" json:"name"`
	Tunnel       *KeylessCertificateTunnelModel `tfsdk:"tunnel" json:"tunnel"`
	Certificate  types.String                   `tfsdk:"certificate" json:"certificate"`
	BundleMethod types.String                   `tfsdk:"bundle_method" json:"bundle_method"`
	Enabled      types.Bool                     `tfsdk:"enabled" json:"enabled"`
	CreatedOn    types.String                   `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn   types.String                   `tfsdk:"modified_on" json:"modified_on,computed"`
	Permissions  *[]types.String                `tfsdk:"permissions" json:"permissions,computed"`
	Status       types.String                   `tfsdk:"status" json:"status,computed"`
}

type KeylessCertificateTunnelModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id"`
}

type KeylessCertificateDataSourceModel struct {
}

type KeylessCertificatesDataSourceModel struct {
}

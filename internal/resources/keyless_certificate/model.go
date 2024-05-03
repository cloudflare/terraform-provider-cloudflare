// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificateResultEnvelope struct {
	Result KeylessCertificateModel `json:"result,computed"`
}

type KeylessCertificateModel struct {
	ID           types.String                   `tfsdk:"id" json:"id,computed"`
	ZoneID       types.String                   `tfsdk:"zone_id" path:"zone_id"`
	Certificate  types.String                   `tfsdk:"certificate" json:"certificate"`
	Host         types.String                   `tfsdk:"host" json:"host"`
	Port         types.Float64                  `tfsdk:"port" json:"port"`
	BundleMethod types.String                   `tfsdk:"bundle_method" json:"bundle_method"`
	Name         types.String                   `tfsdk:"name" json:"name"`
	Tunnel       *KeylessCertificateTunnelModel `tfsdk:"tunnel" json:"tunnel"`
}

type KeylessCertificateTunnelModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id"`
}

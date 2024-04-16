// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificatesResultEnvelope struct {
	Result KeylessCertificatesModel `json:"result,computed"`
}

type KeylessCertificatesModel struct {
	ID           types.String                    `tfsdk:"id" json:"id,computed"`
	ZoneID       types.String                    `tfsdk:"zone_id" path:"zone_id"`
	Certificate  types.String                    `tfsdk:"certificate" json:"certificate"`
	Host         types.String                    `tfsdk:"host" json:"host"`
	Port         types.Float64                   `tfsdk:"port" json:"port"`
	BundleMethod types.String                    `tfsdk:"bundle_method" json:"bundle_method"`
	Name         types.String                    `tfsdk:"name" json:"name"`
	Tunnel       *KeylessCertificatesTunnelModel `tfsdk:"tunnel" json:"tunnel"`
}

type KeylessCertificatesTunnelModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id"`
}

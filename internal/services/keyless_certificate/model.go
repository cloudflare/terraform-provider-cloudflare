// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificateResultEnvelope struct {
	Result KeylessCertificateModel `json:"result"`
}

type KeylessCertificateModel struct {
	ID           types.String                                            `tfsdk:"id" json:"id,computed"`
	ZoneID       types.String                                            `tfsdk:"zone_id" path:"zone_id,required"`
	Certificate  types.String                                            `tfsdk:"certificate" json:"certificate,required"`
	BundleMethod types.String                                            `tfsdk:"bundle_method" json:"bundle_method,computed_optional"`
	Enabled      types.Bool                                              `tfsdk:"enabled" json:"enabled,computed_optional"`
	Host         types.String                                            `tfsdk:"host" json:"host,computed_optional"`
	Name         types.String                                            `tfsdk:"name" json:"name,computed_optional"`
	Port         types.Float64                                           `tfsdk:"port" json:"port,computed_optional"`
	Tunnel       customfield.NestedObject[KeylessCertificateTunnelModel] `tfsdk:"tunnel" json:"tunnel,computed_optional"`
	CreatedOn    timetypes.RFC3339                                       `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn   timetypes.RFC3339                                       `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Status       types.String                                            `tfsdk:"status" json:"status,computed"`
	Permissions  customfield.List[types.String]                          `tfsdk:"permissions" json:"permissions,computed"`
}

type KeylessCertificateTunnelModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed_optional"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed_optional"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnameResultEnvelope struct {
	Result Web3HostnameModel `json:"result"`
}

type Web3HostnameModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Target      types.String      `tfsdk:"target" json:"target,required"`
	Description types.String      `tfsdk:"description" json:"description,optional"`
	Dnslink     types.String      `tfsdk:"dnslink" json:"dnslink,optional"`
	CreatedOn   timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
}

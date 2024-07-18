// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnamesResultListDataSourceEnvelope struct {
	Result *[]*Web3HostnamesItemsDataSourceModel `json:"result,computed"`
}

type Web3HostnamesDataSourceModel struct {
	ZoneIdentifier types.String                          `tfsdk:"zone_identifier" path:"zone_identifier"`
	MaxItems       types.Int64                           `tfsdk:"max_items"`
	Items          *[]*Web3HostnamesItemsDataSourceModel `tfsdk:"items"`
}

type Web3HostnamesItemsDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	CreatedOn   types.String `tfsdk:"created_on" json:"created_on,computed"`
	Description types.String `tfsdk:"description" json:"description"`
	Dnslink     types.String `tfsdk:"dnslink" json:"dnslink"`
	ModifiedOn  types.String `tfsdk:"modified_on" json:"modified_on,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Status      types.String `tfsdk:"status" json:"status,computed"`
	Target      types.String `tfsdk:"target" json:"target"`
}

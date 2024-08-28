// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnamesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[Web3HostnamesResultDataSourceModel] `json:"result,computed"`
}

type Web3HostnamesDataSourceModel struct {
	ZoneIdentifier types.String                                                     `tfsdk:"zone_identifier" path:"zone_identifier"`
	MaxItems       types.Int64                                                      `tfsdk:"max_items"`
	Result         customfield.NestedObjectList[Web3HostnamesResultDataSourceModel] `tfsdk:"result"`
}

type Web3HostnamesResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed_optional"`
	Dnslink     types.String      `tfsdk:"dnslink" json:"dnslink,computed_optional"`
	ModifiedOn  timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
	Target      types.String      `tfsdk:"target" json:"target,computed_optional"`
}
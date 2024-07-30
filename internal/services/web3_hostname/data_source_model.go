// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnameResultDataSourceEnvelope struct {
	Result Web3HostnameDataSourceModel `json:"result,computed"`
}

type Web3HostnameResultListDataSourceEnvelope struct {
	Result *[]*Web3HostnameDataSourceModel `json:"result,computed"`
}

type Web3HostnameDataSourceModel struct {
	ZoneIdentifier types.String                          `tfsdk:"zone_identifier" path:"zone_identifier"`
	Identifier     types.String                          `tfsdk:"identifier" path:"identifier"`
	ID             types.String                          `tfsdk:"id" json:"id,computed"`
	CreatedOn      timetypes.RFC3339                     `tfsdk:"created_on" json:"created_on,computed"`
	Description    types.String                          `tfsdk:"description" json:"description"`
	Dnslink        types.String                          `tfsdk:"dnslink" json:"dnslink"`
	ModifiedOn     timetypes.RFC3339                     `tfsdk:"modified_on" json:"modified_on,computed"`
	Name           types.String                          `tfsdk:"name" json:"name,computed"`
	Status         types.String                          `tfsdk:"status" json:"status,computed"`
	Target         types.String                          `tfsdk:"target" json:"target"`
	Filter         *Web3HostnameFindOneByDataSourceModel `tfsdk:"filter"`
}

type Web3HostnameFindOneByDataSourceModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
}

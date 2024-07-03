// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
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
	ID             types.String                          `tfsdk:"id" json:"id"`
	CreatedOn      types.String                          `tfsdk:"created_on" json:"created_on"`
	Description    types.String                          `tfsdk:"description" json:"description"`
	Dnslink        types.String                          `tfsdk:"dnslink" json:"dnslink"`
	ModifiedOn     types.String                          `tfsdk:"modified_on" json:"modified_on"`
	Name           types.String                          `tfsdk:"name" json:"name"`
	Status         types.String                          `tfsdk:"status" json:"status"`
	Target         types.String                          `tfsdk:"target" json:"target"`
	FindOneBy      *Web3HostnameFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type Web3HostnameFindOneByDataSourceModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
}

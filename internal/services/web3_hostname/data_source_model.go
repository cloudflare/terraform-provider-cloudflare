// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Web3HostnameResultDataSourceEnvelope struct {
	Result Web3HostnameDataSourceModel `json:"result,computed"`
}

type Web3HostnameResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[Web3HostnameDataSourceModel] `json:"result,computed"`
}

type Web3HostnameDataSourceModel struct {
	Identifier     types.String                          `tfsdk:"identifier" path:"identifier"`
	ZoneIdentifier types.String                          `tfsdk:"zone_identifier" path:"zone_identifier"`
	CreatedOn      timetypes.RFC3339                     `tfsdk:"created_on" json:"created_on,computed"`
	ID             types.String                          `tfsdk:"id" json:"id,computed"`
	ModifiedOn     timetypes.RFC3339                     `tfsdk:"modified_on" json:"modified_on,computed"`
	Name           types.String                          `tfsdk:"name" json:"name,computed"`
	Status         types.String                          `tfsdk:"status" json:"status,computed"`
	Description    types.String                          `tfsdk:"description" json:"description,computed_optional"`
	Dnslink        types.String                          `tfsdk:"dnslink" json:"dnslink,computed_optional"`
	Target         types.String                          `tfsdk:"target" json:"target,computed_optional"`
	Filter         *Web3HostnameFindOneByDataSourceModel `tfsdk:"filter"`
}

type Web3HostnameFindOneByDataSourceModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
}

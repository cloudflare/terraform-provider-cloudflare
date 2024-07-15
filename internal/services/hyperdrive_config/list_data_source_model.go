// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigsResultListDataSourceEnvelope struct {
	Result *[]*HyperdriveConfigsItemsDataSourceModel `json:"result,computed"`
}

type HyperdriveConfigsDataSourceModel struct {
	AccountID types.String                              `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                               `tfsdk:"max_items"`
	Items     *[]*HyperdriveConfigsItemsDataSourceModel `tfsdk:"items"`
}

type HyperdriveConfigsItemsDataSourceModel struct {
	Caching *HyperdriveConfigsItemsCachingDataSourceModel `tfsdk:"caching" json:"caching"`
	Name    types.String                                  `tfsdk:"name" json:"name"`
	Origin  *HyperdriveConfigsItemsOriginDataSourceModel  `tfsdk:"origin" json:"origin"`
}

type HyperdriveConfigsItemsCachingDataSourceModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate"`
}

type HyperdriveConfigsItemsOriginDataSourceModel struct {
	Database       types.String `tfsdk:"database" json:"database,computed"`
	Host           types.String `tfsdk:"host" json:"host,computed"`
	Scheme         types.String `tfsdk:"scheme" json:"scheme,computed"`
	User           types.String `tfsdk:"user" json:"user,computed"`
	AccessClientID types.String `tfsdk:"access_client_id" json:"access_client_id"`
	Port           types.Int64  `tfsdk:"port" json:"port"`
}

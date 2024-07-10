// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigResultDataSourceEnvelope struct {
	Result HyperdriveConfigDataSourceModel `json:"result,computed"`
}

type HyperdriveConfigResultListDataSourceEnvelope struct {
	Result *[]*HyperdriveConfigDataSourceModel `json:"result,computed"`
}

type HyperdriveConfigDataSourceModel struct {
	AccountID    types.String                              `tfsdk:"account_id" path:"account_id"`
	HyperdriveID types.String                              `tfsdk:"hyperdrive_id" path:"hyperdrive_id"`
	Caching      *HyperdriveConfigCachingDataSourceModel   `tfsdk:"caching" json:"caching"`
	Name         types.String                              `tfsdk:"name" json:"name"`
	Origin       *HyperdriveConfigOriginDataSourceModel    `tfsdk:"origin" json:"origin"`
	FindOneBy    *HyperdriveConfigFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type HyperdriveConfigCachingDataSourceModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate"`
}

type HyperdriveConfigOriginDataSourceModel struct {
	Database types.String `tfsdk:"database" json:"database,computed"`
	Host     types.String `tfsdk:"host" json:"host,computed"`
	Port     types.Int64  `tfsdk:"port" json:"port,computed"`
	Scheme   types.String `tfsdk:"scheme" json:"scheme,computed"`
	User     types.String `tfsdk:"user" json:"user,computed"`
}

type HyperdriveConfigFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}

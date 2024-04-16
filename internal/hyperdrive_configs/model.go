// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_configs

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigsResultEnvelope struct {
	Result HyperdriveConfigsModel `json:"result,computed"`
}

type HyperdriveConfigsModel struct {
	AccountID    types.String                   `tfsdk:"account_id" path:"account_id"`
	HyperdriveID types.String                   `tfsdk:"hyperdrive_id" path:"hyperdrive_id"`
	Name         types.String                   `tfsdk:"name" json:"name"`
	Origin       *HyperdriveConfigsOriginModel  `tfsdk:"origin" json:"origin"`
	Caching      *HyperdriveConfigsCachingModel `tfsdk:"caching" json:"caching"`
}

type HyperdriveConfigsOriginModel struct {
	Database types.String `tfsdk:"database" json:"database"`
	Host     types.String `tfsdk:"host" json:"host"`
	Port     types.Int64  `tfsdk:"port" json:"port"`
	Scheme   types.String `tfsdk:"scheme" json:"scheme"`
	User     types.String `tfsdk:"user" json:"user"`
}

type HyperdriveConfigsCachingModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate"`
}

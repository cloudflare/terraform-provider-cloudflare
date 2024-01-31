package hyperdrive_config

import "github.com/hashicorp/terraform-plugin-framework/types"

type HyperdriveConfigModel struct {
	AccountID types.String                  `tfsdk:"account_id"`
	ID        types.String                  `tfsdk:"id"`
	Name      types.String                  `tfsdk:"name"`
	Password  types.String                  `tfsdk:"password"`
	Origin    *HyperdriveConfigOriginModel  `tfsdk:"origin"`
	Caching   *HyperdriveConfigCachingModel `tfsdk:"caching"`
}

type HyperdriveConfigOriginModel struct {
	Database types.String `tfsdk:"database"`
	Host     types.String `tfsdk:"host"`
	Port     types.Int64  `tfsdk:"port"`
	Scheme   types.String `tfsdk:"scheme"`
	User     types.String `tfsdk:"user"`
}

type HyperdriveConfigCachingModel struct {
	Disabled             types.Bool  `tfsdk:"disabled"`
	MaxAge               types.Int64 `tfsdk:"max_age"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate"`
}

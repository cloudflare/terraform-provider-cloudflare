package hyperdrive_config

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigModel struct {
	AccountID types.String                 `tfsdk:"account_id"`
	ID        types.String                 `tfsdk:"id"`
	Name      types.String                 `tfsdk:"name"`
	Origin    *HyperdriveConfigOriginModel `tfsdk:"origin"`
	Caching   types.Object                 `tfsdk:"caching"`
}

type HyperdriveConfigOriginModel struct {
	Database           types.String `tfsdk:"database"`
	Password           types.String `tfsdk:"password"`
	Host               types.String `tfsdk:"host"`
	Port               types.Int64  `tfsdk:"port"`
	Scheme             types.String `tfsdk:"scheme"`
	User               types.String `tfsdk:"user"`
	AccessClientID     types.String `tfsdk:"access_client_id"`
	AccessClientSecret types.String `tfsdk:"access_client_secret"`
}

type HyperdriveConfigCachingModel struct {
	Disabled             types.Bool  `tfsdk:"disabled"`
	MaxAge               types.Int64 `tfsdk:"max_age"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate"`
}

func (m HyperdriveConfigCachingModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"disabled":               types.BoolType,
		"max_age":                types.Int64Type,
		"stale_while_revalidate": types.Int64Type,
	}
}

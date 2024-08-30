// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigResultEnvelope struct {
	Result HyperdriveConfigModel `json:"result"`
}

type HyperdriveConfigModel struct {
	ID        types.String                  `tfsdk:"id" json:"-,computed"`
	Name      types.String                  `tfsdk:"name" json:"name"`
	AccountID types.String                  `tfsdk:"account_id" path:"account_id"`
	Origin    *HyperdriveConfigOriginModel  `tfsdk:"origin" json:"origin"`
	Caching   *HyperdriveConfigCachingModel `tfsdk:"caching" json:"caching"`
}

type HyperdriveConfigOriginModel struct {
	Database       types.String `tfsdk:"database" json:"database"`
	Host           types.String `tfsdk:"host" json:"host"`
	Scheme         types.String `tfsdk:"scheme" json:"scheme,computed_optional"`
	User           types.String `tfsdk:"user" json:"user"`
	AccessClientID types.String `tfsdk:"access_client_id" json:"access_client_id,computed_optional"`
	Port           types.Int64  `tfsdk:"port" json:"port,computed_optional"`
}

type HyperdriveConfigCachingModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled,computed_optional"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age,computed_optional"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate,computed_optional"`
}

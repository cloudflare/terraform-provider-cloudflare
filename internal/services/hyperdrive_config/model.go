// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigResultEnvelope struct {
	Result HyperdriveConfigModel `json:"result"`
}

type HyperdriveConfigModel struct {
	ID        types.String                                           `tfsdk:"id" json:"-,computed"`
	Name      types.String                                           `tfsdk:"name" json:"name,required"`
	AccountID types.String                                           `tfsdk:"account_id" path:"account_id,required"`
	Origin    *HyperdriveConfigOriginModel                           `tfsdk:"origin" json:"origin,required"`
	Caching   customfield.NestedObject[HyperdriveConfigCachingModel] `tfsdk:"caching" json:"caching,computed_optional"`
}

type HyperdriveConfigOriginModel struct {
	Database       types.String `tfsdk:"database" json:"database,required"`
	Host           types.String `tfsdk:"host" json:"host,required"`
	Scheme         types.String `tfsdk:"scheme" json:"scheme,computed_optional"`
	User           types.String `tfsdk:"user" json:"user,required"`
	AccessClientID types.String `tfsdk:"access_client_id" json:"access_client_id,optional"`
	Port           types.Int64  `tfsdk:"port" json:"port,optional"`
}

type HyperdriveConfigCachingModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled,optional"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age,optional"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate,optional"`
}

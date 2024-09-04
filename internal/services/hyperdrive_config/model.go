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
	Name      types.String                                           `tfsdk:"name" json:"name,computed_optional"`
	AccountID types.String                                           `tfsdk:"account_id" path:"account_id"`
	Caching   customfield.NestedObject[HyperdriveConfigCachingModel] `tfsdk:"caching" json:"caching,computed_optional"`
	Origin    customfield.NestedObject[HyperdriveConfigOriginModel]  `tfsdk:"origin" json:"origin,computed_optional"`
}

type HyperdriveConfigCachingModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled,computed_optional"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age,computed_optional"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate,computed_optional"`
}

type HyperdriveConfigOriginModel struct {
	Database       types.String `tfsdk:"database" json:"database,computed_optional"`
	Host           types.String `tfsdk:"host" json:"host,computed_optional"`
	Scheme         types.String `tfsdk:"scheme" json:"scheme,computed_optional"`
	User           types.String `tfsdk:"user" json:"user,computed_optional"`
	AccessClientID types.String `tfsdk:"access_client_id" json:"access_client_id,computed_optional"`
	Port           types.Int64  `tfsdk:"port" json:"port,computed_optional"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigResultEnvelope struct {
	Result HyperdriveConfigModel `json:"result"`
}

type HyperdriveConfigModel struct {
	ID         types.String                                           `tfsdk:"id" json:"id,computed"`
	AccountID  types.String                                           `tfsdk:"account_id" path:"account_id,required"`
	Name       types.String                                           `tfsdk:"name" json:"name,required"`
	Origin     *HyperdriveConfigOriginModel                           `tfsdk:"origin" json:"origin,required"`
	Caching    customfield.NestedObject[HyperdriveConfigCachingModel] `tfsdk:"caching" json:"caching,computed_optional"`
	CreatedOn  timetypes.RFC3339                                      `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                                      `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m HyperdriveConfigModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m HyperdriveConfigModel) MarshalJSONForUpdate(state HyperdriveConfigModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type HyperdriveConfigOriginModel struct {
	Database           types.String `tfsdk:"database" json:"database,required"`
	Host               types.String `tfsdk:"host" json:"host,required"`
	Password           types.String `tfsdk:"password" json:"password,required"`
	Port               types.Int64  `tfsdk:"port" json:"port,optional"`
	Scheme             types.String `tfsdk:"scheme" json:"scheme,required"`
	User               types.String `tfsdk:"user" json:"user,required"`
	AccessClientID     types.String `tfsdk:"access_client_id" json:"access_client_id,optional"`
	AccessClientSecret types.String `tfsdk:"access_client_secret" json:"access_client_secret,optional"`
}

type HyperdriveConfigCachingModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled,optional"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age,optional"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate,optional"`
}

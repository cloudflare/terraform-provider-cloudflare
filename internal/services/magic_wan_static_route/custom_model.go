package magic_wan_static_route

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomMagicWANStaticRouteResultEnvelope struct {
	Result CustomMagicWANStaticRouteModel `json:"result"`
}

type CustomMagicWANStaticRouteModel struct {
	ID          types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                   `tfsdk:"account_id" path:"account_id,required"`
	Nexthop     types.String                   `tfsdk:"nexthop" json:"nexthop,required"`
	Prefix      types.String                   `tfsdk:"prefix" json:"prefix,required"`
	Priority    types.Int64                    `tfsdk:"priority" json:"priority,required"`
	Description types.String                   `tfsdk:"description" json:"description,optional"`
	Weight      types.Int64                    `tfsdk:"weight" json:"weight,optional"`
	Scope       *MagicWANStaticRouteScopeModel `tfsdk:"scope" json:"scope,optional"`
	CreatedOn   timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m CustomMagicWANStaticRouteModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomMagicWANStaticRouteModel) MarshalJSONForUpdate(state CustomMagicWANStaticRouteModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

func unmarshalStaticRouteModel(bytes []byte, env *CustomMagicWANStaticRouteResultEnvelope, wrapperField string, unmarshalComputedOnly bool) (err error) {
	return utils.UnmarshalMagicModel(bytes, env, wrapperField, unmarshalComputedOnly)
}

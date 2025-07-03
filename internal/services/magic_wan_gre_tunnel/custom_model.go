package magic_wan_gre_tunnel

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomMagicWANGRETunnelResultEnvelope struct {
	Result CustomMagicWANGRETunnelModel `json:"result"`
}

type CustomMagicWANGRETunnelModel struct {
	ID                    types.String                                                `tfsdk:"id" json:"id,computed"`
	AccountID             types.String                                                `tfsdk:"account_id" path:"account_id,required"`
	CloudflareGREEndpoint types.String                                                `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,required"`
	CustomerGREEndpoint   types.String                                                `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,required"`
	InterfaceAddress      types.String                                                `tfsdk:"interface_address" json:"interface_address,required"`
	Name                  types.String                                                `tfsdk:"name" json:"name,required"`
	Description           types.String                                                `tfsdk:"description" json:"description,optional"`
	Mtu                   types.Int64                                                 `tfsdk:"mtu" json:"mtu,computed_optional"`
	TTL                   types.Int64                                                 `tfsdk:"ttl" json:"ttl,computed_optional"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed_optional"`
	CreatedOn             timetypes.RFC3339                                           `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn            timetypes.RFC3339                                           `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m CustomMagicWANGRETunnelModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomMagicWANGRETunnelModel) MarshalJSONForUpdate(state CustomMagicWANGRETunnelModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

func unmarshalGRETunnelModel(bytes []byte, env *CustomMagicWANGRETunnelResultEnvelope, wrapperField string, unmarshalComputedOnly bool) (err error) {
	return utils.UnmarshalMagicModel(bytes, env, wrapperField, unmarshalComputedOnly)
}

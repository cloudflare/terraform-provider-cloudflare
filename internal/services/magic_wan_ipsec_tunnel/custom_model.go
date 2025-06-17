package magic_wan_ipsec_tunnel

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomMagicWANIPSECTunnelResultEnvelope struct {
	Result CustomMagicWANIPSECTunnelModel `json:"result"`
}

type CustomMagicWANIPSECTunnelModel struct {
	ID                 types.String                                                  `tfsdk:"id" json:"id,computed"`
	AccountID          types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	CloudflareEndpoint types.String                                                  `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint,required"`
	InterfaceAddress   types.String                                                  `tfsdk:"interface_address" json:"interface_address,required"`
	Name               types.String                                                  `tfsdk:"name" json:"name,required"`
	CustomerEndpoint   types.String                                                  `tfsdk:"customer_endpoint" json:"customer_endpoint,optional"`
	Description        types.String                                                  `tfsdk:"description" json:"description,optional"`
	PSK                types.String                                                  `tfsdk:"psk" json:"psk,optional,no_refresh"`
	ReplayProtection   types.Bool                                                    `tfsdk:"replay_protection" json:"replay_protection,computed_optional"`
	HealthCheck        customfield.NestedObject[MagicWANIPSECTunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed_optional"`
	AllowNullCipher    types.Bool                                                    `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed_optional,no_refresh"`
	CreatedOn          timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn         timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PSKMetadata        customfield.NestedObject[MagicWANIPSECTunnelPSKMetadataModel] `tfsdk:"psk_metadata" json:"psk_metadata,computed_optional"`
}

func (m CustomMagicWANIPSECTunnelModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomMagicWANIPSECTunnelModel) MarshalJSONForUpdate(state CustomMagicWANIPSECTunnelModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

func unmarshalIPSECTunnelModel(bytes []byte, env *CustomMagicWANIPSECTunnelResultEnvelope, wrapperField string, unmarshalComputedOnly bool) (err error) {
	return utils.UnmarshalMagicModel(bytes, env, wrapperField, unmarshalComputedOnly)
}

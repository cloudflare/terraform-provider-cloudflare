// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq_relay

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MoqRelayResultEnvelope struct {
	Result MoqRelayModel `json:"result"`
}

type MoqRelayModel struct {
	ID                    types.String                                  `tfsdk:"id" json:"-,computed"`
	UID                   types.String                                  `tfsdk:"uid" json:"uid,computed"`
	AccountID             types.String                                  `tfsdk:"account_id" path:"account_id,required"`
	Name                  types.String                                  `tfsdk:"name" json:"name,required"`
	Config                customfield.NestedObject[MoqRelayConfigModel] `tfsdk:"config" json:"config,computed_optional"`
	Created               timetypes.RFC3339                             `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified              timetypes.RFC3339                             `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Status                types.String                                  `tfsdk:"status" json:"status,computed"`
	TokenPublishSubscribe types.String                                  `tfsdk:"token_publish_subscribe" json:"token_publish_subscribe,computed,no_refresh"`
	TokenSubscribe        types.String                                  `tfsdk:"token_subscribe" json:"token_subscribe,computed,no_refresh"`
}

func (m MoqRelayModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MoqRelayModel) MarshalJSONForUpdate(state MoqRelayModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MoqRelayConfigModel struct {
	LingeringSubscribe customfield.NestedObject[MoqRelayConfigLingeringSubscribeModel] `tfsdk:"lingering_subscribe" json:"lingering_subscribe,computed_optional"`
	OriginFallback     customfield.NestedObject[MoqRelayConfigOriginFallbackModel]     `tfsdk:"origin_fallback" json:"origin_fallback,computed_optional"`
}

type MoqRelayConfigLingeringSubscribeModel struct {
	Enabled      types.Bool  `tfsdk:"enabled" json:"enabled,computed_optional"`
	MaxTimeoutMs types.Int64 `tfsdk:"max_timeout_ms" json:"max_timeout_ms,computed_optional"`
}

type MoqRelayConfigOriginFallbackModel struct {
	Enabled types.Bool                                                             `tfsdk:"enabled" json:"enabled,computed_optional"`
	Origins customfield.NestedObjectList[MoqRelayConfigOriginFallbackOriginsModel] `tfsdk:"origins" json:"origins,computed_optional"`
}

type MoqRelayConfigOriginFallbackOriginsModel struct {
	URL types.String `tfsdk:"url" json:"url,optional"`
}

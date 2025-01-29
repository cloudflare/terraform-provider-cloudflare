// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicNetworkMonitoringRuleResultEnvelope struct {
	Result MagicNetworkMonitoringRuleModel `json:"result"`
}

type MagicNetworkMonitoringRuleModel struct {
	ID                     types.String    `tfsdk:"id" json:"id,required"`
	AccountID              types.String    `tfsdk:"account_id" path:"account_id,required"`
	Name                   types.String    `tfsdk:"name" json:"name,required"`
	AutomaticAdvertisement types.Bool      `tfsdk:"automatic_advertisement" json:"automatic_advertisement,optional"`
	Bandwidth              types.Float64   `tfsdk:"bandwidth" json:"bandwidth,optional"`
	PacketThreshold        types.Float64   `tfsdk:"packet_threshold" json:"packet_threshold,optional"`
	Prefixes               *[]types.String `tfsdk:"prefixes" json:"prefixes,optional"`
	Duration               types.String    `tfsdk:"duration" json:"duration,computed_optional"`
	BandwidthThreshold     types.Float64   `tfsdk:"bandwidth_threshold" json:"bandwidth_threshold,computed"`
}

func (m MagicNetworkMonitoringRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MagicNetworkMonitoringRuleModel) MarshalJSONForUpdate(state MagicNetworkMonitoringRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

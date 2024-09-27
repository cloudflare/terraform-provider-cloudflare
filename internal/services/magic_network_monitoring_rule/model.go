// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicNetworkMonitoringRuleResultEnvelope struct {
	Result MagicNetworkMonitoringRuleModel `json:"result"`
}

type MagicNetworkMonitoringRuleModel struct {
	ID                     types.String                   `tfsdk:"id" json:"id,required"`
	AccountID              types.String                   `tfsdk:"account_id" path:"account_id,required"`
	Name                   types.String                   `tfsdk:"name" json:"name,required"`
	Bandwidth              types.Float64                  `tfsdk:"bandwidth" json:"bandwidth,optional"`
	AutomaticAdvertisement types.Bool                     `tfsdk:"automatic_advertisement" json:"automatic_advertisement,computed_optional"`
	Duration               types.String                   `tfsdk:"duration" json:"duration,computed_optional"`
	PacketThreshold        types.Float64                  `tfsdk:"packet_threshold" json:"packet_threshold,computed_optional"`
	Prefixes               customfield.List[types.String] `tfsdk:"prefixes" json:"prefixes,computed_optional"`
	BandwidthThreshold     types.Float64                  `tfsdk:"bandwidth_threshold" json:"bandwidth_threshold,computed"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/magic_network_monitoring"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicNetworkMonitoringRulesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicNetworkMonitoringRulesResultDataSourceModel] `json:"result,computed"`
}

type MagicNetworkMonitoringRulesDataSourceModel struct {
	AccountID types.String                                                                   `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                    `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[MagicNetworkMonitoringRulesResultDataSourceModel] `tfsdk:"result"`
}

func (m *MagicNetworkMonitoringRulesDataSourceModel) toListParams(_ context.Context) (params magic_network_monitoring.RuleListParams, diags diag.Diagnostics) {
	params = magic_network_monitoring.RuleListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicNetworkMonitoringRulesResultDataSourceModel struct {
	AutomaticAdvertisement types.Bool                     `tfsdk:"automatic_advertisement" json:"automatic_advertisement,computed"`
	Duration               types.String                   `tfsdk:"duration" json:"duration,computed"`
	Name                   types.String                   `tfsdk:"name" json:"name,computed"`
	Prefixes               customfield.List[types.String] `tfsdk:"prefixes" json:"prefixes,computed"`
	ID                     types.String                   `tfsdk:"id" json:"id,computed"`
	BandwidthThreshold     types.Float64                  `tfsdk:"bandwidth_threshold" json:"bandwidth_threshold,computed"`
	PacketThreshold        types.Float64                  `tfsdk:"packet_threshold" json:"packet_threshold,computed"`
}

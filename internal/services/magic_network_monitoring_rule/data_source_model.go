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

type MagicNetworkMonitoringRuleResultDataSourceEnvelope struct {
	Result MagicNetworkMonitoringRuleDataSourceModel `json:"result,computed"`
}

type MagicNetworkMonitoringRuleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicNetworkMonitoringRuleDataSourceModel] `json:"result,computed"`
}

type MagicNetworkMonitoringRuleDataSourceModel struct {
	AccountID              types.String                                        `tfsdk:"account_id" path:"account_id,optional"`
	RuleID                 types.String                                        `tfsdk:"rule_id" path:"rule_id,optional"`
	AutomaticAdvertisement types.Bool                                          `tfsdk:"automatic_advertisement" json:"automatic_advertisement,computed"`
	BandwidthThreshold     types.Float64                                       `tfsdk:"bandwidth_threshold" json:"bandwidth_threshold,computed"`
	Duration               types.String                                        `tfsdk:"duration" json:"duration,computed"`
	ID                     types.String                                        `tfsdk:"id" json:"id,computed"`
	Name                   types.String                                        `tfsdk:"name" json:"name,computed"`
	PacketThreshold        types.Float64                                       `tfsdk:"packet_threshold" json:"packet_threshold,computed"`
	Prefixes               customfield.List[types.String]                      `tfsdk:"prefixes" json:"prefixes,computed"`
	Filter                 *MagicNetworkMonitoringRuleFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *MagicNetworkMonitoringRuleDataSourceModel) toReadParams(_ context.Context) (params magic_network_monitoring.RuleGetParams, diags diag.Diagnostics) {
	params = magic_network_monitoring.RuleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *MagicNetworkMonitoringRuleDataSourceModel) toListParams(_ context.Context) (params magic_network_monitoring.RuleListParams, diags diag.Diagnostics) {
	params = magic_network_monitoring.RuleListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type MagicNetworkMonitoringRuleFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

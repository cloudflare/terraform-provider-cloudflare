// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns_firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSFirewallsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[DNSFirewallsResultDataSourceModel] `json:"result,computed"`
}

type DNSFirewallsDataSourceModel struct {
	AccountID types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                     `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[DNSFirewallsResultDataSourceModel] `tfsdk:"result"`
}

func (m *DNSFirewallsDataSourceModel) toListParams(_ context.Context) (params dns_firewall.DNSFirewallListParams, diags diag.Diagnostics) {
	params = dns_firewall.DNSFirewallListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type DNSFirewallsResultDataSourceModel struct {
	ID                   types.String                                                          `tfsdk:"id" json:"id,computed"`
	DeprecateAnyRequests types.Bool                                                            `tfsdk:"deprecate_any_requests" json:"deprecate_any_requests,computed"`
	DNSFirewallIPs       customfield.List[types.String]                                        `tfsdk:"dns_firewall_ips" json:"dns_firewall_ips,computed"`
	ECSFallback          types.Bool                                                            `tfsdk:"ecs_fallback" json:"ecs_fallback,computed"`
	MaximumCacheTTL      types.Float64                                                         `tfsdk:"maximum_cache_ttl" json:"maximum_cache_ttl,computed"`
	MinimumCacheTTL      types.Float64                                                         `tfsdk:"minimum_cache_ttl" json:"minimum_cache_ttl,computed"`
	ModifiedOn           timetypes.RFC3339                                                     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name                 types.String                                                          `tfsdk:"name" json:"name,computed"`
	NegativeCacheTTL     types.Float64                                                         `tfsdk:"negative_cache_ttl" json:"negative_cache_ttl,computed"`
	Ratelimit            types.Float64                                                         `tfsdk:"ratelimit" json:"ratelimit,computed"`
	Retries              types.Float64                                                         `tfsdk:"retries" json:"retries,computed"`
	UpstreamIPs          customfield.List[types.String]                                        `tfsdk:"upstream_ips" json:"upstream_ips,computed"`
	AttackMitigation     customfield.NestedObject[DNSFirewallsAttackMitigationDataSourceModel] `tfsdk:"attack_mitigation" json:"attack_mitigation,computed"`
}

type DNSFirewallsAttackMitigationDataSourceModel struct {
	Enabled                   types.Bool `tfsdk:"enabled" json:"enabled,computed"`
	OnlyWhenUpstreamUnhealthy types.Bool `tfsdk:"only_when_upstream_unhealthy" json:"only_when_upstream_unhealthy,computed"`
}

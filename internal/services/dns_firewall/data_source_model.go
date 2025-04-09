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

type DNSFirewallResultDataSourceEnvelope struct {
Result DNSFirewallDataSourceModel `json:"result,computed"`
}

type DNSFirewallDataSourceModel struct {
ID types.String `tfsdk:"id" path:"dns_firewall_id,computed"`
DNSFirewallID types.String `tfsdk:"dns_firewall_id" path:"dns_firewall_id,optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
DeprecateAnyRequests types.Bool `tfsdk:"deprecate_any_requests" json:"deprecate_any_requests,computed"`
ECSFallback types.Bool `tfsdk:"ecs_fallback" json:"ecs_fallback,computed"`
MaximumCacheTTL types.Float64 `tfsdk:"maximum_cache_ttl" json:"maximum_cache_ttl,computed"`
MinimumCacheTTL types.Float64 `tfsdk:"minimum_cache_ttl" json:"minimum_cache_ttl,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
NegativeCacheTTL types.Float64 `tfsdk:"negative_cache_ttl" json:"negative_cache_ttl,computed"`
Ratelimit types.Float64 `tfsdk:"ratelimit" json:"ratelimit,computed"`
Retries types.Float64 `tfsdk:"retries" json:"retries,computed"`
DNSFirewallIPs customfield.List[types.String] `tfsdk:"dns_firewall_ips" json:"dns_firewall_ips,computed"`
UpstreamIPs customfield.List[types.String] `tfsdk:"upstream_ips" json:"upstream_ips,computed"`
AttackMitigation customfield.NestedObject[DNSFirewallAttackMitigationDataSourceModel] `tfsdk:"attack_mitigation" json:"attack_mitigation,computed"`
}

func (m *DNSFirewallDataSourceModel) toReadParams(_ context.Context) (params dns_firewall.DNSFirewallGetParams, diags diag.Diagnostics) {
  params = dns_firewall.DNSFirewallGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type DNSFirewallAttackMitigationDataSourceModel struct {
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
OnlyWhenUpstreamUnhealthy types.Bool `tfsdk:"only_when_upstream_unhealthy" json:"only_when_upstream_unhealthy,computed"`
}

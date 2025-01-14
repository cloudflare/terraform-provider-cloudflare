// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSFirewallResultEnvelope struct {
	Result DNSFirewallModel `json:"result"`
}

type DNSFirewallModel struct {
	ID                   types.String                                               `tfsdk:"id" json:"id,computed"`
	AccountID            types.String                                               `tfsdk:"account_id" path:"account_id,required"`
	DeprecateAnyRequests types.Bool                                                 `tfsdk:"deprecate_any_requests" json:"deprecate_any_requests,optional"`
	ECSFallback          types.Bool                                                 `tfsdk:"ecs_fallback" json:"ecs_fallback,optional"`
	Name                 types.String                                               `tfsdk:"name" json:"name,optional"`
	NegativeCacheTTL     types.Float64                                              `tfsdk:"negative_cache_ttl" json:"negative_cache_ttl,optional"`
	Ratelimit            types.Float64                                              `tfsdk:"ratelimit" json:"ratelimit,optional"`
	UpstreamIPs          *[]types.String                                            `tfsdk:"upstream_ips" json:"upstream_ips,optional"`
	MaximumCacheTTL      types.Float64                                              `tfsdk:"maximum_cache_ttl" json:"maximum_cache_ttl,computed_optional"`
	MinimumCacheTTL      types.Float64                                              `tfsdk:"minimum_cache_ttl" json:"minimum_cache_ttl,computed_optional"`
	Retries              types.Float64                                              `tfsdk:"retries" json:"retries,computed_optional"`
	AttackMitigation     customfield.NestedObject[DNSFirewallAttackMitigationModel] `tfsdk:"attack_mitigation" json:"attack_mitigation,computed_optional"`
	ModifiedOn           timetypes.RFC3339                                          `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	DNSFirewallIPs       customfield.List[types.String]                             `tfsdk:"dns_firewall_ips" json:"dns_firewall_ips,computed"`
}

func (m DNSFirewallModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSFirewallModel) MarshalJSONForUpdate(state DNSFirewallModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type DNSFirewallAttackMitigationModel struct {
	Enabled                   types.Bool `tfsdk:"enabled" json:"enabled,optional"`
	OnlyWhenUpstreamUnhealthy types.Bool `tfsdk:"only_when_upstream_unhealthy" json:"only_when_upstream_unhealthy,computed_optional"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/firewall"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRuleResultDataSourceEnvelope struct {
	Result UserAgentBlockingRuleDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRuleDataSourceModel struct {
	UARuleID types.String `tfsdk:"ua_rule_id" path:"ua_rule_id,required"`
	ZoneID   types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *UserAgentBlockingRuleDataSourceModel) toReadParams(_ context.Context) (params firewall.UARuleGetParams, diags diag.Diagnostics) {
	params = firewall.UARuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

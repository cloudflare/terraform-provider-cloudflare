// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_routing"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRuleResultDataSourceEnvelope struct {
	Result EmailRoutingRuleDataSourceModel `json:"result,computed"`
}

type EmailRoutingRuleDataSourceModel struct {
	RuleIdentifier types.String `tfsdk:"rule_identifier" path:"rule_identifier,required"`
	ZoneID         types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *EmailRoutingRuleDataSourceModel) toReadParams(_ context.Context) (params email_routing.RuleGetParams, diags diag.Diagnostics) {
	params = email_routing.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

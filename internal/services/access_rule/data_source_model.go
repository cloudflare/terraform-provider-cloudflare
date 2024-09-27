// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRuleResultDataSourceEnvelope struct {
	Result AccessRuleDataSourceModel `json:"result,computed"`
}

type AccessRuleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccessRuleDataSourceModel] `json:"result,computed"`
}

type AccessRuleDataSourceModel struct {
	AccountID  types.String                        `tfsdk:"account_id" path:"account_id,optional"`
	Identifier types.String                        `tfsdk:"identifier" path:"identifier,optional"`
	ZoneID     types.String                        `tfsdk:"zone_id" path:"zone_id,optional"`
	Filter     *AccessRuleFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *AccessRuleDataSourceModel) toReadParams(_ context.Context) (params firewall.AccessRuleGetParams, diags diag.Diagnostics) {
	params = firewall.AccessRuleGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *AccessRuleDataSourceModel) toListParams(_ context.Context) (params firewall.AccessRuleListParams, diags diag.Diagnostics) {
	params = firewall.AccessRuleListParams{}

	if m.Filter.Configuration != nil {
		paramsConfiguration := firewall.AccessRuleListParamsConfiguration{}
		if !m.Filter.Configuration.Target.IsNull() {
			paramsConfiguration.Target = cloudflare.F(firewall.AccessRuleListParamsConfigurationTarget(m.Filter.Configuration.Target.ValueString()))
		}
		if !m.Filter.Configuration.Value.IsNull() {
			paramsConfiguration.Value = cloudflare.F(m.Filter.Configuration.Value.ValueString())
		}
		params.Configuration = cloudflare.F(paramsConfiguration)
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(firewall.AccessRuleListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Match.IsNull() {
		params.Match = cloudflare.F(firewall.AccessRuleListParamsMatch(m.Filter.Match.ValueString()))
	}
	if !m.Filter.Mode.IsNull() {
		params.Mode = cloudflare.F(firewall.AccessRuleListParamsMode(m.Filter.Mode.ValueString()))
	}
	if !m.Filter.Notes.IsNull() {
		params.Notes = cloudflare.F(m.Filter.Notes.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(firewall.AccessRuleListParamsOrder(m.Filter.Order.ValueString()))
	}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type AccessRuleFindOneByDataSourceModel struct {
	AccountID     types.String                            `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID        types.String                            `tfsdk:"zone_id" path:"zone_id,optional"`
	Configuration *AccessRuleConfigurationDataSourceModel `tfsdk:"configuration" query:"configuration,optional"`
	Direction     types.String                            `tfsdk:"direction" query:"direction,optional"`
	Match         types.String                            `tfsdk:"match" query:"match,computed_optional"`
	Mode          types.String                            `tfsdk:"mode" query:"mode,optional"`
	Notes         types.String                            `tfsdk:"notes" query:"notes,optional"`
	Order         types.String                            `tfsdk:"order" query:"order,optional"`
}

type AccessRuleConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target,computed_optional"`
	Value  types.String `tfsdk:"value" json:"value,computed_optional"`
}

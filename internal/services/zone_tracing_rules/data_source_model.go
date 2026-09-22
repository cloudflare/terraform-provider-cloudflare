// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tracing_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zones"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneTracingRulesResultDataSourceEnvelope struct {
	Result ZoneTracingRulesDataSourceModel `json:"result,computed"`
}

type ZoneTracingRulesDataSourceModel struct {
	ID     types.String                                                       `tfsdk:"id" path:"zone_id,computed"`
	ZoneID types.String                                                       `tfsdk:"zone_id" path:"zone_id,required"`
	Rules  customfield.NestedObjectList[ZoneTracingRulesRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

func (m *ZoneTracingRulesDataSourceModel) toReadParams(_ context.Context) (params zones.ObservabilityTracingRuleGetParams, diags diag.Diagnostics) {
	params = zones.ObservabilityTracingRuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type ZoneTracingRulesRulesDataSourceModel struct {
	Action           types.String                                                                   `tfsdk:"action" json:"action,computed"`
	ActionParameters customfield.NestedObject[ZoneTracingRulesRulesActionParametersDataSourceModel] `tfsdk:"action_parameters" json:"action_parameters,computed"`
	Description      types.String                                                                   `tfsdk:"description" json:"description,computed"`
	Enabled          types.Bool                                                                     `tfsdk:"enabled" json:"enabled,computed"`
	Expression       types.String                                                                   `tfsdk:"expression" json:"expression,computed"`
}

type ZoneTracingRulesRulesActionParametersDataSourceModel struct {
	SamplingRatio types.Float64 `tfsdk:"sampling_ratio" json:"sampling_ratio,computed"`
}

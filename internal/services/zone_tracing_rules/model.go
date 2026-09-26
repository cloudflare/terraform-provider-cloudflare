// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tracing_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneTracingRulesResultEnvelope struct {
	Result ZoneTracingRulesModel `json:"result"`
}

type ZoneTracingRulesModel struct {
	ID     types.String                   `tfsdk:"id" json:"-,computed"`
	ZoneID types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	Rules  *[]*ZoneTracingRulesRulesModel `tfsdk:"rules" json:"rules,required"`
}

func (m ZoneTracingRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneTracingRulesModel) MarshalJSONForUpdate(state ZoneTracingRulesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZoneTracingRulesRulesModel struct {
	Action           types.String                                `tfsdk:"action" json:"action,required"`
	ActionParameters *ZoneTracingRulesRulesActionParametersModel `tfsdk:"action_parameters" json:"action_parameters,required"`
	Description      types.String                                `tfsdk:"description" json:"description,required"`
	Enabled          types.Bool                                  `tfsdk:"enabled" json:"enabled,required"`
	Expression       types.String                                `tfsdk:"expression" json:"expression,required"`
}

type ZoneTracingRulesRulesActionParametersModel struct {
	SamplingRatio types.Float64 `tfsdk:"sampling_ratio" json:"sampling_ratio,required"`
}

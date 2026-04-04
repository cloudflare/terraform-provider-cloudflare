// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// type RulesResultDataSourceEnvelope struct {
// 	Result RulesDataSourceModel `json:"result"`
// }

type RulesDataSourceModel struct {
	ID    types.String                                            `tfsdk:"id" json:"id,computed"`
	Phase types.String                                            `tfsdk:"phase" json:"phase,computed"`
	Rules customfield.NestedObjectList[RulesRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed,decode_null_to_zero"`
}

type RulesRulesDataSourceModel struct {
	Action                 types.String                                                                `tfsdk:"action" json:"action,computed"`
	ActionParameters       customfield.NestedObject[RulesetRulesActionParametersDataSourceModel]       `tfsdk:"action_parameters" json:"action_parameters,computed,decode_null_to_zero"`
	Description            types.String                                                                `tfsdk:"description" json:"description,computed,decode_null_to_zero"`
	Enabled                types.Bool                                                                  `tfsdk:"enabled" json:"enabled,computed"`
	ExposedCredentialCheck customfield.NestedObject[RulesetRulesExposedCredentialCheckDataSourceModel] `tfsdk:"exposed_credential_check" json:"exposed_credential_check,computed"`
	Expression             types.String                                                                `tfsdk:"expression" json:"expression,computed"`
	Logging                customfield.NestedObject[RulesetRulesLoggingDataSourceModel]                `tfsdk:"logging" json:"logging,computed"`
	Ratelimit              customfield.NestedObject[RulesetRulesRatelimitDataSourceModel]              `tfsdk:"ratelimit" json:"ratelimit,computed"`
}

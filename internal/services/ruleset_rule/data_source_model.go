// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetRuleResultDataSourceEnvelope struct {
}

// RulesetRuleDataSourceModel represents the data source model for a single ruleset rule.
type RulesetRuleDataSourceModel struct {
	// Rule fields from ruleset (includes ID)
	ID                     types.String                                                               `tfsdk:"id" json:"id,computed"`
	Action                 types.String                                                               `tfsdk:"action" json:"action,computed"`
	ActionParameters       customfield.NestedObject[RulesetRuleActionParametersDataSourceModel]       `tfsdk:"action_parameters" json:"action_parameters,computed,decode_null_to_zero"`
	Description            types.String                                                               `tfsdk:"description" json:"description,computed,decode_null_to_zero"`
	Enabled                types.Bool                                                                 `tfsdk:"enabled" json:"enabled,computed"`
	ExposedCredentialCheck customfield.NestedObject[RulesetRuleExposedCredentialCheckDataSourceModel] `tfsdk:"exposed_credential_check" json:"exposed_credential_check,computed"`
	Expression             types.String                                                               `tfsdk:"expression" json:"expression,computed"`
	Logging                customfield.NestedObject[RulesetRuleLoggingDataSourceModel]                `tfsdk:"logging" json:"logging,computed"`
	Ratelimit              customfield.NestedObject[RulesetRuleRatelimitDataSourceModel]              `tfsdk:"ratelimit" json:"ratelimit,computed"`
	Ref                    types.String                                                               `tfsdk:"ref" json:"ref,computed"`
	Categories             customfield.List[types.String]                                             `tfsdk:"categories" json:"categories,computed"`

	// Additional fields for lookup
	RuleID    types.String `tfsdk:"rule_id"`
	RulesetID types.String `tfsdk:"ruleset_id"`
	AccountID types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`
}

func (m *RulesetRuleDataSourceModel) toReadParams(_ context.Context) (params rulesets.RulesetGetParams, diags diag.Diagnostics) {
	params = rulesets.RulesetGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

// Type aliases for action parameters and nested models from ruleset package
type RulesetRuleActionParametersDataSourceModel = ruleset.RulesetRulesActionParametersDataSourceModel
type RulesetRuleActionParametersResponseDataSourceModel = ruleset.RulesetRulesActionParametersResponseDataSourceModel
type RulesetRuleActionParametersAlgorithmsDataSourceModel = ruleset.RulesetRulesActionParametersAlgorithmsDataSourceModel
type RulesetRuleActionParametersMatchedDataDataSourceModel = ruleset.RulesetRulesActionParametersMatchedDataDataSourceModel
type RulesetRuleActionParametersOverridesDataSourceModel = ruleset.RulesetRulesActionParametersOverridesDataSourceModel
type RulesetRuleActionParametersOverridesCategoriesDataSourceModel = ruleset.RulesetRulesActionParametersOverridesCategoriesDataSourceModel
type RulesetRuleActionParametersOverridesRulesDataSourceModel = ruleset.RulesetRulesActionParametersOverridesRulesDataSourceModel
type RulesetRuleActionParametersFromListDataSourceModel = ruleset.RulesetRulesActionParametersFromListDataSourceModel
type RulesetRuleActionParametersFromValueDataSourceModel = ruleset.RulesetRulesActionParametersFromValueDataSourceModel
type RulesetRuleActionParametersFromValueTargetURLDataSourceModel = ruleset.RulesetRulesActionParametersFromValueTargetURLDataSourceModel
type RulesetRuleActionParametersHeadersDataSourceModel = ruleset.RulesetRulesActionParametersHeadersDataSourceModel
type RulesetRuleActionParametersURIDataSourceModel = ruleset.RulesetRulesActionParametersURIDataSourceModel
type RulesetRuleActionParametersURIPathDataSourceModel = ruleset.RulesetRulesActionParametersURIPathDataSourceModel
type RulesetRuleActionParametersURIQueryDataSourceModel = ruleset.RulesetRulesActionParametersURIQueryDataSourceModel
type RulesetRuleActionParametersOriginDataSourceModel = ruleset.RulesetRulesActionParametersOriginDataSourceModel
type RulesetRuleActionParametersSNIDataSourceModel = ruleset.RulesetRulesActionParametersSNIDataSourceModel
type RulesetRuleActionParametersAutominifyDataSourceModel = ruleset.RulesetRulesActionParametersAutominifyDataSourceModel
type RulesetRuleActionParametersBrowserTTLDataSourceModel = ruleset.RulesetRulesActionParametersBrowserTTLDataSourceModel
type RulesetRuleActionParametersCacheKeyDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyDataSourceModel
type RulesetRuleActionParametersCacheKeyCustomKeyDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyCustomKeyDataSourceModel
type RulesetRuleActionParametersCacheKeyCustomKeyCookieDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyCustomKeyCookieDataSourceModel
type RulesetRuleActionParametersCacheKeyCustomKeyHeaderDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyCustomKeyHeaderDataSourceModel
type RulesetRuleActionParametersCacheKeyCustomKeyHostDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyCustomKeyHostDataSourceModel
type RulesetRuleActionParametersCacheKeyCustomKeyQueryStringDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyCustomKeyQueryStringDataSourceModel
type RulesetRuleActionParametersCacheKeyCustomKeyQueryStringIncludeDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeDataSourceModel
type RulesetRuleActionParametersCacheKeyCustomKeyQueryStringExcludeDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeDataSourceModel
type RulesetRuleActionParametersCacheKeyCustomKeyUserDataSourceModel = ruleset.RulesetRulesActionParametersCacheKeyCustomKeyUserDataSourceModel
type RulesetRuleActionParametersCacheReserveDataSourceModel = ruleset.RulesetRulesActionParametersCacheReserveDataSourceModel
type RulesetRuleActionParametersEdgeTTLDataSourceModel = ruleset.RulesetRulesActionParametersEdgeTTLDataSourceModel
type RulesetRuleActionParametersEdgeTTLStatusCodeTTLDataSourceModel = ruleset.RulesetRulesActionParametersEdgeTTLStatusCodeTTLDataSourceModel
type RulesetRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeDataSourceModel = ruleset.RulesetRulesActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeDataSourceModel
type RulesetRuleActionParametersServeStaleDataSourceModel = ruleset.RulesetRulesActionParametersServeStaleDataSourceModel
type RulesetRuleActionParametersCookieFieldsDataSourceModel = ruleset.RulesetRulesActionParametersCookieFieldsDataSourceModel
type RulesetRuleActionParametersRawResponseFieldsDataSourceModel = ruleset.RulesetRulesActionParametersRawResponseFieldsDataSourceModel
type RulesetRuleActionParametersRequestFieldsDataSourceModel = ruleset.RulesetRulesActionParametersRequestFieldsDataSourceModel
type RulesetRuleActionParametersResponseFieldsDataSourceModel = ruleset.RulesetRulesActionParametersResponseFieldsDataSourceModel
type RulesetRuleActionParametersTransformedRequestFieldsDataSourceModel = ruleset.RulesetRulesActionParametersTransformedRequestFieldsDataSourceModel
type RulesetRuleExposedCredentialCheckDataSourceModel = ruleset.RulesetRulesExposedCredentialCheckDataSourceModel
type RulesetRuleLoggingDataSourceModel = ruleset.RulesetRulesLoggingDataSourceModel
type RulesetRuleRatelimitDataSourceModel = ruleset.RulesetRulesRatelimitDataSourceModel

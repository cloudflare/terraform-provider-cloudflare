// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetResultEnvelope struct {
	Result RulesetModel `json:"result"`
}

type RulesetModel struct {
	ID          types.String          `tfsdk:"id" json:"id,computed"`
	AccountID   types.String          `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID      types.String          `tfsdk:"zone_id" path:"zone_id,optional"`
	Kind        types.String          `tfsdk:"kind" json:"kind,required"`
	Name        types.String          `tfsdk:"name" json:"name,required"`
	Phase       types.String          `tfsdk:"phase" json:"phase,required"`
	Rules       *[]*RulesetRulesModel `tfsdk:"rules" json:"rules,required"`
	Description types.String          `tfsdk:"description" json:"description,computed_optional"`
	LastUpdated timetypes.RFC3339     `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Version     types.String          `tfsdk:"version" json:"version,computed"`
}

type RulesetRulesModel struct {
	LastUpdated      timetypes.RFC3339                                           `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Version          types.String                                                `tfsdk:"version" json:"version,computed"`
	ID               types.String                                                `tfsdk:"id" json:"id,computed_optional"`
	Action           types.String                                                `tfsdk:"action" json:"action,computed_optional"`
	ActionParameters customfield.NestedObject[RulesetRulesActionParametersModel] `tfsdk:"action_parameters" json:"action_parameters,computed_optional"`
	Categories       customfield.List[types.String]                              `tfsdk:"categories" json:"categories,computed"`
	Description      types.String                                                `tfsdk:"description" json:"description,computed_optional"`
	Enabled          types.Bool                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Expression       types.String                                                `tfsdk:"expression" json:"expression,computed_optional"`
	Logging          customfield.NestedObject[RulesetRulesLoggingModel]          `tfsdk:"logging" json:"logging,computed_optional"`
	Ref              types.String                                                `tfsdk:"ref" json:"ref,computed_optional"`
}

type RulesetRulesActionParametersModel struct {
	Response customfield.NestedObject[RulesetRulesActionParametersResponseModel] `tfsdk:"response" json:"response,computed_optional"`
}

type RulesetRulesActionParametersResponseModel struct {
	Content     types.String `tfsdk:"content" json:"content,required"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,required"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,required"`
}

type RulesetRulesLoggingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,required"`
}

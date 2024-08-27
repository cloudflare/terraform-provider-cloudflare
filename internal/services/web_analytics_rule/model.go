// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_rule

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebAnalyticsRuleResultEnvelope struct {
	Result WebAnalyticsRuleModel `json:"result"`
}

type WebAnalyticsRuleModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id"`
	RulesetID types.String      `tfsdk:"ruleset_id" path:"ruleset_id"`
	Host      types.String      `tfsdk:"host" json:"host"`
	Inclusive types.Bool        `tfsdk:"inclusive" json:"inclusive"`
	IsPaused  types.Bool        `tfsdk:"is_paused" json:"is_paused"`
	Paths     *[]types.String   `tfsdk:"paths" json:"paths"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Priority  types.Float64     `tfsdk:"priority" json:"priority,computed"`
}

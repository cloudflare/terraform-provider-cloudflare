// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageRuleResultDataSourceEnvelope struct {
	Result PageRuleDataSourceModel `json:"result,computed"`
}

type PageRuleDataSourceModel struct {
	PageruleID types.String `tfsdk:"pagerule_id" path:"pagerule_id"`
	ZoneID     types.String `tfsdk:"zone_id" path:"zone_id"`
}

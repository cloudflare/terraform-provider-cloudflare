// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/page_rules"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type PageRuleResultDataSourceEnvelope struct {
Result PageRuleDataSourceModel `json:"result,computed"`
}

type PageRuleDataSourceModel struct {
PageruleID types.String `tfsdk:"pagerule_id" path:"pagerule_id,required"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
ID types.String `tfsdk:"id" json:"id,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Priority types.Int64 `tfsdk:"priority" json:"priority,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
}

func (m *PageRuleDataSourceModel) toReadParams(_ context.Context) (params page_rules.PageRuleGetParams, diags diag.Diagnostics) {
  params = page_rules.PageRuleGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

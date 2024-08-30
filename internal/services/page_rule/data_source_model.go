// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/pagerules"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageRuleResultDataSourceEnvelope struct {
	Result PageRuleDataSourceModel `json:"result,computed"`
}

type PageRuleDataSourceModel struct {
	PageruleID types.String `tfsdk:"pagerule_id" path:"pagerule_id"`
	ZoneID     types.String `tfsdk:"zone_id" path:"zone_id"`
}

func (m *PageRuleDataSourceModel) toReadParams(_ context.Context) (params pagerules.PageruleGetParams, diags diag.Diagnostics) {
	params = pagerules.PageruleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

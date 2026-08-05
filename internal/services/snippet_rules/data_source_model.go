// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/snippets"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetRulesResultDataSourceEnvelope struct {
	Result SnippetRulesDataSourceModel `json:"result,computed"`
}

type SnippetRulesDataSourceModel struct {
	ID     types.String `tfsdk:"id" path:"zone_id,computed"`
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m *SnippetRulesDataSourceModel) toReadParams(_ context.Context) (params snippets.RuleGetParams, diags diag.Diagnostics) {
	params = snippets.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

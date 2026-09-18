// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/snippets"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetRulesResultDataSourceEnvelope struct {
	Result SnippetRulesDataSourceModel `json:"result,computed"`
}

type SnippetRulesDataSourceModel struct {
	ID          types.String      `tfsdk:"id" path:"zone_id,computed"`
	ZoneID      types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Expression  types.String      `tfsdk:"expression" json:"expression,computed"`
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	SnippetName types.String      `tfsdk:"snippet_name" json:"snippet_name,computed"`
}

func (m *SnippetRulesDataSourceModel) toReadParams(_ context.Context) (params snippets.RuleGetParams, diags diag.Diagnostics) {
	params = snippets.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

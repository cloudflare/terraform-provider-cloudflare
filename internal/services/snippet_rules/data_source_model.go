// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetRulesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SnippetRuleDataSourceModel] `json:"result,computed"`
}

type SnippetRulesDataSourceModel struct {
	ZoneID   types.String                                             `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                              `tfsdk:"max_items"`
	Rules    customfield.NestedObjectList[SnippetRuleDataSourceModel] `tfsdk:"rules"`
}

func (m *SnippetRulesDataSourceModel) toListParams(_ context.Context) (params snippets.RuleListParams, diags diag.Diagnostics) {
	params = snippets.RuleListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type SnippetRuleDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	Expression  types.String      `tfsdk:"expression" json:"expression,computed"`
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	SnippetName types.String      `tfsdk:"snippet_name" json:"snippet_name,computed"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
}

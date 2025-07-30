// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetRulesResultEnvelope struct {
	Result *[]*SnippetRulesBodyModel `json:"result"`
}

type SnippetRulesModel struct {
	ID     types.String              `tfsdk:"id" json:"id,computed"`
	ZoneID types.String              `tfsdk:"zone_id" path:"zone_id,required"`
	Body   *[]*SnippetRulesBodyModel `tfsdk:"body" json:"body,required"`
}

func (m SnippetRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Body)
}

func (m SnippetRulesModel) MarshalJSONForUpdate(state SnippetRulesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Body, state.Body)
}

type SnippetRulesBodyModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	Expression  types.String      `tfsdk:"expression" json:"expression,required"`
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	SnippetName types.String      `tfsdk:"snippet_name" json:"snippet_name,required"`
	Description types.String      `tfsdk:"description" json:"description,computed_optional"`
	Enabled     types.Bool        `tfsdk:"enabled" json:"enabled,computed_optional"`
}

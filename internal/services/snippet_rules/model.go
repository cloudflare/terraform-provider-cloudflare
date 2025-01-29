// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetRulesResultEnvelope struct {
	Result SnippetRulesModel `json:"result"`
}

type SnippetRulesModel struct {
	ZoneID types.String                                         `tfsdk:"zone_id" path:"zone_id,required"`
	Rules  customfield.NestedObjectList[SnippetRulesRulesModel] `tfsdk:"rules" json:"rules,computed_optional"`
}

func (m SnippetRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SnippetRulesModel) MarshalJSONForUpdate(state SnippetRulesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type SnippetRulesRulesModel struct {
	Description types.String `tfsdk:"description" json:"description,optional"`
	Enabled     types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	Expression  types.String `tfsdk:"expression" json:"expression,optional"`
	SnippetName types.String `tfsdk:"snippet_name" json:"snippet_name,optional"`
}

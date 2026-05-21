// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_namespace

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_search"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchNamespacesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AISearchNamespacesResultDataSourceModel] `json:"result,computed"`
}

type AISearchNamespacesDataSourceModel struct {
	AccountID types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	Search    types.String                                                          `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                           `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AISearchNamespacesResultDataSourceModel] `tfsdk:"result"`
}

func (m *AISearchNamespacesDataSourceModel) toListParams(_ context.Context) (params ai_search.NamespaceListParams, diags diag.Diagnostics) {
	params = ai_search.NamespaceListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type AISearchNamespacesResultDataSourceModel struct {
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
}

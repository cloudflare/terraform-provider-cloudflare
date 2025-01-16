// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetsResultDataSourceEnvelope struct {
	Result SnippetsDataSourceModel `json:"result,computed"`
}

type SnippetsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SnippetsDataSourceModel] `json:"result,computed"`
}

type SnippetsDataSourceModel struct {
	ZoneID      types.String                      `tfsdk:"zone_id" path:"zone_id,optional"`
	SnippetName types.String                      `tfsdk:"snippet_name" path:"snippet_name,computed_optional"`
	CreatedOn   types.String                      `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn  types.String                      `tfsdk:"modified_on" json:"modified_on,computed"`
	Filter      *SnippetsFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *SnippetsDataSourceModel) toReadParams(_ context.Context) (params snippets.SnippetGetParams, diags diag.Diagnostics) {
	params = snippets.SnippetGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *SnippetsDataSourceModel) toListParams(_ context.Context) (params snippets.SnippetListParams, diags diag.Diagnostics) {
	params = snippets.SnippetListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	return
}

type SnippetsFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

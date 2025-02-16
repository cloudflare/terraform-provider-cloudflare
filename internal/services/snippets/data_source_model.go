// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/snippets"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetsResultDataSourceEnvelope struct {
	Result SnippetsDataSourceModel `json:"result,computed"`
}

type SnippetsDataSourceModel struct {
	SnippetName types.String `tfsdk:"snippet_name" path:"snippet_name,required"`
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	CreatedOn   types.String `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn  types.String `tfsdk:"modified_on" json:"modified_on,computed"`
}

func (m *SnippetsDataSourceModel) toReadParams(_ context.Context) (params snippets.SnippetGetParams, diags diag.Diagnostics) {
	params = snippets.SnippetGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

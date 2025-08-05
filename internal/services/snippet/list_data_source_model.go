// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnippetsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SnippetsResultDataSourceModel] `json:"result,computed"`
}

type SnippetsDataSourceModel struct {
	ZoneID   types.String                                                `tfsdk:"zone_id" path:"zone_id,required"`
	MaxItems types.Int64                                                 `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[SnippetsResultDataSourceModel] `tfsdk:"result"`
}

func (m *SnippetsDataSourceModel) toListParams(_ context.Context) (params snippets.SnippetListParams, diags diag.Diagnostics) {
	params = snippets.SnippetListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type SnippetsResultDataSourceModel struct {
	CreatedOn   timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	SnippetName types.String      `tfsdk:"snippet_name" json:"snippet_name,computed"`
	ModifiedOn  timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

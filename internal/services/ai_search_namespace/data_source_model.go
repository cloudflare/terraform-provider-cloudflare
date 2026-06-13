// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_namespace

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_search"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchNamespaceResultDataSourceEnvelope struct {
	Result AISearchNamespaceDataSourceModel `json:"result,computed"`
}

type AISearchNamespaceDataSourceModel struct {
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String      `tfsdk:"name" path:"name,required"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
}

func (m *AISearchNamespaceDataSourceModel) toReadParams(_ context.Context) (params ai_search.NamespaceReadParams, diags diag.Diagnostics) {
	params = ai_search.NamespaceReadParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

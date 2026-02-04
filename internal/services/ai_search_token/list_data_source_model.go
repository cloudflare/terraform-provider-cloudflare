// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search_token

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ai_search"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISearchTokensResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AISearchTokensResultDataSourceModel] `json:"result,computed"`
}

type AISearchTokensDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AISearchTokensResultDataSourceModel] `tfsdk:"result"`
}

func (m *AISearchTokensDataSourceModel) toListParams(_ context.Context) (params ai_search.TokenListParams, diags diag.Diagnostics) {
	params = ai_search.TokenListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AISearchTokensResultDataSourceModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	AccountID  types.String      `tfsdk:"account_id" json:"account_id,computed"`
	AccountTag types.String      `tfsdk:"account_tag" json:"account_tag,computed"`
	CfAPIID    types.String      `tfsdk:"cf_api_id" json:"cf_api_id,computed"`
	CfAPIKey   types.String      `tfsdk:"cf_api_key" json:"cf_api_key,computed"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ModifiedAt timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Name       types.String      `tfsdk:"name" json:"name,computed"`
	CreatedBy  types.String      `tfsdk:"created_by" json:"created_by,computed"`
	Enabled    types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Legacy     types.Bool        `tfsdk:"legacy" json:"legacy,computed"`
	ModifiedBy types.String      `tfsdk:"modified_by" json:"modified_by,computed"`
	SyncedAt   timetypes.RFC3339 `tfsdk:"synced_at" json:"synced_at,computed" format:"date-time"`
}

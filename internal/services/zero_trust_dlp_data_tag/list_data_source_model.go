// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_data_tag

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPDataTagsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPDataTagsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPDataTagsDataSourceModel struct {
	AccountID  types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	CategoryID types.String                                                            `tfsdk:"category_id" path:"category_id,required"`
	MaxItems   types.Int64                                                             `tfsdk:"max_items"`
	Result     customfield.NestedObjectList[ZeroTrustDLPDataTagsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDLPDataTagsDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPDataTagCategoryDataTagListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPDataTagCategoryDataTagListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPDataTagsResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
}

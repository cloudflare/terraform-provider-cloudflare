// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_data_tag

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPDataTagResultDataSourceEnvelope struct {
	Result ZeroTrustDLPDataTagDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPDataTagDataSourceModel struct {
	ID          types.String      `tfsdk:"id" path:"tag_id,computed"`
	TagID       types.String      `tfsdk:"tag_id" path:"tag_id,required"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	CategoryID  types.String      `tfsdk:"category_id" path:"category_id,required"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m *ZeroTrustDLPDataTagDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DLPDataTagCategoryDataTagGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPDataTagCategoryDataTagGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

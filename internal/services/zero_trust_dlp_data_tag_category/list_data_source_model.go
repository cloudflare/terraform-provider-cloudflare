// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_data_tag_category

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPDataTagCategoriesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustDLPDataTagCategoriesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustDLPDataTagCategoriesDataSourceModel struct {
	AccountID types.String                                                                     `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustDLPDataTagCategoriesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustDLPDataTagCategoriesDataSourceModel) toListParams(_ context.Context) (params zero_trust.DLPDataTagCategoryListParams, diags diag.Diagnostics) {
	params = zero_trust.DLPDataTagCategoryListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustDLPDataTagCategoriesResultDataSourceModel struct {
	ID          types.String                                                                   `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339                                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String                                                                   `tfsdk:"name" json:"name,computed"`
	Tags        customfield.NestedObjectList[ZeroTrustDLPDataTagCategoriesTagsDataSourceModel] `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt   timetypes.RFC3339                                                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String                                                                   `tfsdk:"description" json:"description,computed"`
	TemplateID  types.String                                                                   `tfsdk:"template_id" json:"template_id,computed"`
}

type ZeroTrustDLPDataTagCategoriesTagsDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
}

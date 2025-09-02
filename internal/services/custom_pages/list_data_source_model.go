// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/custom_pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomPagesListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CustomPagesListResultDataSourceModel] `json:"result,computed"`
}

type CustomPagesListDataSourceModel struct {
	AccountID types.String                                                       `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                       `tfsdk:"zone_id" path:"zone_id,optional"`
	MaxItems  types.Int64                                                        `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[CustomPagesListResultDataSourceModel] `tfsdk:"result"`
}

func (m *CustomPagesListDataSourceModel) toListParams(_ context.Context) (params custom_pages.CustomPageListParams, diags diag.Diagnostics) {
	params = custom_pages.CustomPageListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type CustomPagesListResultDataSourceModel struct {
	ID             types.String                   `tfsdk:"id" json:"id,computed"`
	CreatedOn      timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description    types.String                   `tfsdk:"description" json:"description,computed"`
	ModifiedOn     timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PreviewTarget  types.String                   `tfsdk:"preview_target" json:"preview_target,computed"`
	RequiredTokens customfield.List[types.String] `tfsdk:"required_tokens" json:"required_tokens,computed"`
	State          types.String                   `tfsdk:"state" json:"state,computed"`
	URL            types.String                   `tfsdk:"url" json:"url,computed"`
}

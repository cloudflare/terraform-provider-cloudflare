// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/custom_pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomPagesResultDataSourceEnvelope struct {
	Result CustomPagesDataSourceModel `json:"result,computed"`
}

type CustomPagesDataSourceModel struct {
	Identifier     types.String                   `tfsdk:"identifier" path:"identifier,required"`
	AccountID      types.String                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID         types.String                   `tfsdk:"zone_id" path:"zone_id,optional"`
	CreatedOn      timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description    types.String                   `tfsdk:"description" json:"description,computed"`
	ID             types.String                   `tfsdk:"id" json:"id,computed"`
	ModifiedOn     timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PreviewTarget  types.String                   `tfsdk:"preview_target" json:"preview_target,computed"`
	State          types.String                   `tfsdk:"state" json:"state,computed"`
	URL            types.String                   `tfsdk:"url" json:"url,computed"`
	RequiredTokens customfield.List[types.String] `tfsdk:"required_tokens" json:"required_tokens,computed"`
}

func (m *CustomPagesDataSourceModel) toReadParams(_ context.Context) (params custom_pages.CustomPageGetParams, diags diag.Diagnostics) {
	params = custom_pages.CustomPageGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

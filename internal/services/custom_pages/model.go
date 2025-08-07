// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomPagesResultEnvelope struct {
	Result CustomPagesModel `json:"result"`
}

type CustomPagesModel struct {
	ID             types.String                   `tfsdk:"id" json:"-,computed"`
	Identifier     types.String                   `tfsdk:"identifier" path:"identifier,required"`
	AccountID      types.String                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID         types.String                   `tfsdk:"zone_id" path:"zone_id,optional"`
	State          types.String                   `tfsdk:"state" json:"state,required"`
	URL            types.String                   `tfsdk:"url" json:"url,computed_optional"`
	CreatedOn      timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description    types.String                   `tfsdk:"description" json:"description,computed"`
	ModifiedOn     timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PreviewTarget  types.String                   `tfsdk:"preview_target" json:"preview_target,computed"`
	RequiredTokens customfield.List[types.String] `tfsdk:"required_tokens" json:"required_tokens,computed"`
}

func (m CustomPagesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomPagesModel) MarshalJSONForUpdate(state CustomPagesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

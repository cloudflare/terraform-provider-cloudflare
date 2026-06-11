// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_data_tag

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPDataTagResultEnvelope struct {
	Result ZeroTrustDLPDataTagModel `json:"result"`
}

type ZeroTrustDLPDataTagModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	CategoryID  types.String      `tfsdk:"category_id" path:"category_id,required"`
	Name        types.String      `tfsdk:"name" json:"name,required"`
	Description types.String      `tfsdk:"description" json:"description,optional"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustDLPDataTagModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustDLPDataTagModel) MarshalJSONForUpdate(state ZeroTrustDLPDataTagModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_bgp_filter_profile

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANBGPFilterProfileResultEnvelope struct {
	Result MagicWANBGPFilterProfileModel `json:"result"`
}

type MagicWANBGPFilterProfileModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	MatchAction types.String      `tfsdk:"match_action" json:"match_action,required"`
	Name        types.String      `tfsdk:"name" json:"name,required"`
	Targets     *[]types.String   `tfsdk:"targets" json:"targets,required"`
	Description types.String      `tfsdk:"description" json:"description,computed_optional"`
	CreatedOn   timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m MagicWANBGPFilterProfileModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MagicWANBGPFilterProfileModel) MarshalJSONForUpdate(state MagicWANBGPFilterProfileModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

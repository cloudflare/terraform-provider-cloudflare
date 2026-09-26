// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_bgp_filter_profile

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANBGPFilterProfilesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicWANBGPFilterProfilesResultDataSourceModel] `json:"result,computed"`
}

type MagicWANBGPFilterProfilesDataSourceModel struct {
	AccountID types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                  `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[MagicWANBGPFilterProfilesResultDataSourceModel] `tfsdk:"result"`
}

func (m *MagicWANBGPFilterProfilesDataSourceModel) toListParams(_ context.Context) (params magic_transit.BGPFilterProfileListParams, diags diag.Diagnostics) {
	params = magic_transit.BGPFilterProfileListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicWANBGPFilterProfilesResultDataSourceModel struct {
	ID          types.String                   `tfsdk:"id" json:"id,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	MatchAction types.String                   `tfsdk:"match_action" json:"match_action,computed"`
	Name        types.String                   `tfsdk:"name" json:"name,computed"`
	Targets     customfield.List[types.String] `tfsdk:"targets" json:"targets,computed"`
	CreatedOn   timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

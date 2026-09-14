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

type MagicWANBGPFilterProfileResultDataSourceEnvelope struct {
	Result MagicWANBGPFilterProfileDataSourceModel `json:"result,computed"`
}

type MagicWANBGPFilterProfileDataSourceModel struct {
	ID          types.String                   `tfsdk:"id" path:"profile_id,computed"`
	ProfileID   types.String                   `tfsdk:"profile_id" path:"profile_id,required"`
	AccountID   types.String                   `tfsdk:"account_id" path:"account_id,required"`
	CreatedOn   timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	MatchAction types.String                   `tfsdk:"match_action" json:"match_action,computed"`
	ModifiedOn  timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String                   `tfsdk:"name" json:"name,computed"`
	Targets     customfield.List[types.String] `tfsdk:"targets" json:"targets,computed"`
}

func (m *MagicWANBGPFilterProfileDataSourceModel) toReadParams(_ context.Context) (params magic_transit.BGPFilterProfileGetParams, diags diag.Diagnostics) {
	params = magic_transit.BGPFilterProfileGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

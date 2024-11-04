// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package call_app_turn_key

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/calls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallAppTURNKeysResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CallAppTURNKeysResultDataSourceModel] `json:"result,computed"`
}

type CallAppTURNKeysDataSourceModel struct {
	AccountID types.String                                                       `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                        `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[CallAppTURNKeysResultDataSourceModel] `tfsdk:"result"`
}

func (m *CallAppTURNKeysDataSourceModel) toListParams(_ context.Context) (params calls.TURNKeyListParams, diags diag.Diagnostics) {
	params = calls.TURNKeyListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type CallAppTURNKeysResultDataSourceModel struct {
	Created  timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name     types.String      `tfsdk:"name" json:"name,computed"`
	UID      types.String      `tfsdk:"uid" json:"uid,computed"`
}

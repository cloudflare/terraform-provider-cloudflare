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

type CallAppTURNKeyResultDataSourceEnvelope struct {
	Result CallAppTURNKeyDataSourceModel `json:"result,computed"`
}

type CallAppTURNKeyResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CallAppTURNKeyDataSourceModel] `json:"result,computed"`
}

type CallAppTURNKeyDataSourceModel struct {
	AccountID types.String                            `tfsdk:"account_id" path:"account_id,optional"`
	KeyID     types.String                            `tfsdk:"key_id" path:"key_id,optional"`
	Created   timetypes.RFC3339                       `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339                       `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String                            `tfsdk:"name" json:"name,computed"`
	UID       types.String                            `tfsdk:"uid" json:"uid,computed"`
	Filter    *CallAppTURNKeyFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CallAppTURNKeyDataSourceModel) toReadParams(_ context.Context) (params calls.TURNKeyGetParams, diags diag.Diagnostics) {
	params = calls.TURNKeyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *CallAppTURNKeyDataSourceModel) toListParams(_ context.Context) (params calls.TURNKeyListParams, diags diag.Diagnostics) {
	params = calls.TURNKeyListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type CallAppTURNKeyFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

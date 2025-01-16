// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_turn_app

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/calls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallsTURNAppResultDataSourceEnvelope struct {
	Result CallsTURNAppDataSourceModel `json:"result,computed"`
}

type CallsTURNAppResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CallsTURNAppDataSourceModel] `json:"result,computed"`
}

type CallsTURNAppDataSourceModel struct {
	AccountID types.String                          `tfsdk:"account_id" path:"account_id,optional"`
	KeyID     types.String                          `tfsdk:"key_id" path:"key_id,optional"`
	Created   timetypes.RFC3339                     `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339                     `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String                          `tfsdk:"name" json:"name,computed"`
	UID       types.String                          `tfsdk:"uid" json:"uid,computed"`
	Filter    *CallsTURNAppFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CallsTURNAppDataSourceModel) toReadParams(_ context.Context) (params calls.TURNGetParams, diags diag.Diagnostics) {
	params = calls.TURNGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *CallsTURNAppDataSourceModel) toListParams(_ context.Context) (params calls.TURNListParams, diags diag.Diagnostics) {
	params = calls.TURNListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type CallsTURNAppFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

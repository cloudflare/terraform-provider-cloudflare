// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package call_app

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/calls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallAppResultDataSourceEnvelope struct {
	Result CallAppDataSourceModel `json:"result,computed"`
}

type CallAppResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CallAppDataSourceModel] `json:"result,computed"`
}

type CallAppDataSourceModel struct {
	AccountID types.String                     `tfsdk:"account_id" path:"account_id,optional"`
	AppID     types.String                     `tfsdk:"app_id" path:"app_id,optional"`
	Created   timetypes.RFC3339                `tfsdk:"created" json:"created,optional" format:"date-time"`
	Modified  timetypes.RFC3339                `tfsdk:"modified" json:"modified,optional" format:"date-time"`
	UID       types.String                     `tfsdk:"uid" json:"uid,optional"`
	Name      types.String                     `tfsdk:"name" json:"name,computed_optional"`
	Filter    *CallAppFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CallAppDataSourceModel) toReadParams(_ context.Context) (params calls.CallGetParams, diags diag.Diagnostics) {
	params = calls.CallGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *CallAppDataSourceModel) toListParams(_ context.Context) (params calls.CallListParams, diags diag.Diagnostics) {
	params = calls.CallListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type CallAppFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

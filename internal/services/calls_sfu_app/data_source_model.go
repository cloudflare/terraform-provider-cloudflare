// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_sfu_app

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/calls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallsSfuAppResultDataSourceEnvelope struct {
	Result CallsSfuAppDataSourceModel `json:"result,computed"`
}

type CallsSfuAppResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CallsSfuAppDataSourceModel] `json:"result,computed"`
}

type CallsSfuAppDataSourceModel struct {
	AccountID types.String                         `tfsdk:"account_id" path:"account_id,optional"`
	AppID     types.String                         `tfsdk:"app_id" path:"app_id,optional"`
	Created   timetypes.RFC3339                    `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339                    `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String                         `tfsdk:"name" json:"name,computed"`
	UID       types.String                         `tfsdk:"uid" json:"uid,computed"`
	Filter    *CallsSfuAppFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CallsSfuAppDataSourceModel) toReadParams(_ context.Context) (params calls.SfuGetParams, diags diag.Diagnostics) {
	params = calls.SfuGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *CallsSfuAppDataSourceModel) toListParams(_ context.Context) (params calls.SfuListParams, diags diag.Diagnostics) {
	params = calls.SfuListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type CallsSfuAppFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

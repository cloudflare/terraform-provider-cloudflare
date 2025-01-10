// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_sfu_app

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/calls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallsSFUAppResultDataSourceEnvelope struct {
	Result CallsSFUAppDataSourceModel `json:"result,computed"`
}

type CallsSFUAppResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CallsSFUAppDataSourceModel] `json:"result,computed"`
}

type CallsSFUAppDataSourceModel struct {
	AccountID types.String                         `tfsdk:"account_id" path:"account_id,optional"`
	AppID     types.String                         `tfsdk:"app_id" path:"app_id,optional"`
	Created   timetypes.RFC3339                    `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339                    `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name      types.String                         `tfsdk:"name" json:"name,computed"`
	UID       types.String                         `tfsdk:"uid" json:"uid,computed"`
	Filter    *CallsSFUAppFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CallsSFUAppDataSourceModel) toReadParams(_ context.Context) (params calls.SFUGetParams, diags diag.Diagnostics) {
	params = calls.SFUGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *CallsSFUAppDataSourceModel) toListParams(_ context.Context) (params calls.SFUListParams, diags diag.Diagnostics) {
	params = calls.SFUListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type CallsSFUAppFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

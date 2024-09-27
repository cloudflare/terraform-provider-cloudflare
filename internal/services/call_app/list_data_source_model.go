// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package call_app

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/calls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallAppsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CallAppsResultDataSourceModel] `json:"result,computed"`
}

type CallAppsDataSourceModel struct {
	AccountID types.String                                                `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                 `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[CallAppsResultDataSourceModel] `tfsdk:"result"`
}

func (m *CallAppsDataSourceModel) toListParams(_ context.Context) (params calls.CallListParams, diags diag.Diagnostics) {
	params = calls.CallListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type CallAppsResultDataSourceModel struct {
}

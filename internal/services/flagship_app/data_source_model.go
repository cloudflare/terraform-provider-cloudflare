// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_app

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/flagship"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FlagshipAppResultDataSourceEnvelope struct {
	Result FlagshipAppDataSourceModel `json:"result,computed"`
}

type FlagshipAppDataSourceModel struct {
	ID        types.String `tfsdk:"id" path:"app_id,computed"`
	AppID     types.String `tfsdk:"app_id" path:"app_id,required"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
	UpdatedBy types.String `tfsdk:"updated_by" json:"updated_by,computed"`
}

func (m *FlagshipAppDataSourceModel) toReadParams(_ context.Context) (params flagship.AppGetParams, diags diag.Diagnostics) {
	params = flagship.AppGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

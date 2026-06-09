// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_app

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/flagship"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FlagshipAppsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[FlagshipAppsResultDataSourceModel] `json:"result,computed"`
}

type FlagshipAppsDataSourceModel struct {
	AccountID types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                     `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[FlagshipAppsResultDataSourceModel] `tfsdk:"result"`
}

func (m *FlagshipAppsDataSourceModel) toListParams(_ context.Context) (params flagship.AppListParams, diags diag.Diagnostics) {
	params = flagship.AppListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type FlagshipAppsResultDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
	UpdatedBy types.String `tfsdk:"updated_by" json:"updated_by,computed"`
}

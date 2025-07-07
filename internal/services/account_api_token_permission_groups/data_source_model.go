// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_api_token_permission_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountAPITokenPermissionGroupsResultDataSourceEnvelope struct {
	Result AccountAPITokenPermissionGroupsDataSourceModel `json:"result,computed"`
}

type AccountAPITokenPermissionGroupsDataSourceModel struct {
	AccountID types.String                   `tfsdk:"account_id" path:"account_id,required"`
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	Name      types.String                   `tfsdk:"name" json:"name,computed"`
	Scopes    customfield.List[types.String] `tfsdk:"scopes" json:"scopes,computed"`
}

func (m *AccountAPITokenPermissionGroupsDataSourceModel) toReadParams(_ context.Context) (params accounts.TokenPermissionGroupGetParams, diags diag.Diagnostics) {
	params = accounts.TokenPermissionGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

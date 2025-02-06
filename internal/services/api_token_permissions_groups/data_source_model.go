// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token_permissions_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenPermissionsGroupsResultDataSourceEnvelope struct {
	Result APITokenPermissionsGroupsDataSourceModel `json:"result,computed"`
}

type APITokenPermissionsGroupsDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

func (m *APITokenPermissionsGroupsDataSourceModel) toReadParams(_ context.Context) (params accounts.TokenPermissionGroupGetParams, diags diag.Diagnostics) {
	params = accounts.TokenPermissionGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

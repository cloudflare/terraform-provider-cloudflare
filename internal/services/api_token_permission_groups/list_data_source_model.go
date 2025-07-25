// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token_permission_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/user"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenPermissionGroupsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APITokenPermissionGroupsListResultDataSourceModel] `json:"result,computed"`
}

type APITokenPermissionGroupsListDataSourceModel struct {
	Name     types.String                                                                    `tfsdk:"name" query:"name,optional"`
	Scope    types.String                                                                    `tfsdk:"scope" query:"scope,optional"`
	MaxItems types.Int64                                                                     `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[APITokenPermissionGroupsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *APITokenPermissionGroupsListDataSourceModel) toListParams(_ context.Context) (params user.TokenPermissionGroupListParams, diags diag.Diagnostics) {
	params = user.TokenPermissionGroupListParams{}

	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Scope.IsNull() {
		params.Scope = cloudflare.F(m.Scope.ValueString())
	}

	return
}

type APITokenPermissionGroupsListResultDataSourceModel struct {
	ID     types.String                   `tfsdk:"id" json:"id,computed"`
	Name   types.String                   `tfsdk:"name" json:"name,computed"`
	Scopes customfield.List[types.String] `tfsdk:"scopes" json:"scopes,computed"`
}

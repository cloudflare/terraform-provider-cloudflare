// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/iam"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PermissionGroupsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PermissionGroupsListResultDataSourceModel] `json:"result,computed"`
}

type PermissionGroupsListDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	ID        types.String                                                            `tfsdk:"id" query:"id,optional"`
	Label     types.String                                                            `tfsdk:"label" query:"label,optional"`
	Name      types.String                                                            `tfsdk:"name" query:"name,optional"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[PermissionGroupsListResultDataSourceModel] `tfsdk:"result"`
}

func (m *PermissionGroupsListDataSourceModel) toListParams(_ context.Context) (params iam.PermissionGroupListParams, diags diag.Diagnostics) {
	params = iam.PermissionGroupListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.Label.IsNull() {
		params.Label = cloudflare.F(m.Label.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}

	return
}

type PermissionGroupsListResultDataSourceModel struct {
}

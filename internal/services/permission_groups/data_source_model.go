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

type PermissionGroupsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[PermissionGroupsDataSourceModel] `json:"result,computed"`
}

type PermissionGroupsDataSourceModel struct {
	AccountID         types.String                              `tfsdk:"account_id" path:"account_id,optional"`
	PermissionGroupID types.String                              `tfsdk:"permission_group_id" path:"permission_group_id,optional"`
	ID                types.String                              `tfsdk:"id" json:"id,optional"`
	Name              types.String                              `tfsdk:"name" json:"name,optional"`
	Meta              *PermissionGroupsMetaDataSourceModel      `tfsdk:"meta" json:"meta,optional"`
	Filter            *PermissionGroupsFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *PermissionGroupsDataSourceModel) toReadParams(_ context.Context) (params iam.PermissionGroupGetParams, diags diag.Diagnostics) {
	params = iam.PermissionGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *PermissionGroupsDataSourceModel) toListParams(_ context.Context) (params iam.PermissionGroupListParams, diags diag.Diagnostics) {
	params = iam.PermissionGroupListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.ID.IsNull() {
		params.ID = cloudflare.F(m.Filter.ID.ValueString())
	}
	if !m.Filter.Label.IsNull() {
		params.Label = cloudflare.F(m.Filter.Label.ValueString())
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}

	return
}

type PermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type PermissionGroupsFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	ID        types.String `tfsdk:"id" query:"id,optional"`
	Label     types.String `tfsdk:"label" query:"label,optional"`
	Name      types.String `tfsdk:"name" query:"name,optional"`
}

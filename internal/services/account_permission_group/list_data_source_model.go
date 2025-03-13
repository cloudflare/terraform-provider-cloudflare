// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_permission_group

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/iam"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountPermissionGroupsResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[AccountPermissionGroupsResultDataSourceModel] `json:"result,computed"`
}

type AccountPermissionGroupsDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
ID types.String `tfsdk:"id" query:"id,optional"`
Label types.String `tfsdk:"label" query:"label,optional"`
Name types.String `tfsdk:"name" query:"name,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[AccountPermissionGroupsResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccountPermissionGroupsDataSourceModel) toListParams(_ context.Context) (params iam.PermissionGroupListParams, diags diag.Diagnostics) {
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

type AccountPermissionGroupsResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Meta customfield.NestedObject[AccountPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}

type AccountPermissionGroupsMetaDataSourceModel struct {
Key types.String `tfsdk:"key" json:"key,computed"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

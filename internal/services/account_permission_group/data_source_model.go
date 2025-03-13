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

type AccountPermissionGroupDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
PermissionGroupID types.String `tfsdk:"permission_group_id" path:"permission_group_id,required"`
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Meta customfield.NestedObject[AccountPermissionGroupMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
}

func (m *AccountPermissionGroupDataSourceModel) toReadParams(_ context.Context) (params iam.PermissionGroupGetParams, diags diag.Diagnostics) {
  params = iam.PermissionGroupGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type AccountPermissionGroupMetaDataSourceModel struct {
Key types.String `tfsdk:"key" json:"key,computed"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

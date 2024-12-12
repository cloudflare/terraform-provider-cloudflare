// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_token

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountTokensResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountTokensResultDataSourceModel] `json:"result,computed"`
}

type AccountTokensDataSourceModel struct {
	AccountID types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	Direction types.String                                                     `tfsdk:"direction" query:"direction,optional"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AccountTokensResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccountTokensDataSourceModel) toListParams(_ context.Context) (params accounts.TokenListParams, diags diag.Diagnostics) {
	params = accounts.TokenListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(accounts.TokenListParamsDirection(m.Direction.ValueString()))
	}

	return
}

type AccountTokensResultDataSourceModel struct {
	ID         types.String                                                       `tfsdk:"id" json:"id,computed"`
	Condition  customfield.NestedObject[AccountTokensConditionDataSourceModel]    `tfsdk:"condition" json:"condition,computed"`
	ExpiresOn  timetypes.RFC3339                                                  `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	IssuedOn   timetypes.RFC3339                                                  `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn timetypes.RFC3339                                                  `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                                                  `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name       types.String                                                       `tfsdk:"name" json:"name,computed"`
	NotBefore  timetypes.RFC3339                                                  `tfsdk:"not_before" json:"not_before,computed" format:"date-time"`
	Policies   customfield.NestedObjectList[AccountTokensPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Status     types.String                                                       `tfsdk:"status" json:"status,computed"`
}

type AccountTokensConditionDataSourceModel struct {
	RequestIP customfield.NestedObject[AccountTokensConditionRequestIPDataSourceModel] `tfsdk:"request_ip" json:"request_ip,computed"`
}

type AccountTokensConditionRequestIPDataSourceModel struct {
	In    customfield.List[types.String] `tfsdk:"in" json:"in,computed"`
	NotIn customfield.List[types.String] `tfsdk:"not_in" json:"not_in,computed"`
}

type AccountTokensPoliciesDataSourceModel struct {
	ID               types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                                                       `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups customfield.NestedObjectList[AccountTokensPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        customfield.Map[types.String]                                                      `tfsdk:"resources" json:"resources,computed"`
}

type AccountTokensPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[AccountTokensPoliciesPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
	Name types.String                                                                       `tfsdk:"name" json:"name,computed"`
}

type AccountTokensPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

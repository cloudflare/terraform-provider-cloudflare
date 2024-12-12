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

type AccountTokenResultDataSourceEnvelope struct {
	Result AccountTokenDataSourceModel `json:"result,computed"`
}

type AccountTokenResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountTokenDataSourceModel] `json:"result,computed"`
}

type AccountTokenDataSourceModel struct {
	AccountID  types.String                                                      `tfsdk:"account_id" path:"account_id,optional"`
	TokenID    types.String                                                      `tfsdk:"token_id" path:"token_id,optional"`
	ExpiresOn  timetypes.RFC3339                                                 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	ID         types.String                                                      `tfsdk:"id" json:"id,computed"`
	IssuedOn   timetypes.RFC3339                                                 `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn timetypes.RFC3339                                                 `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                                                 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name       types.String                                                      `tfsdk:"name" json:"name,computed"`
	NotBefore  timetypes.RFC3339                                                 `tfsdk:"not_before" json:"not_before,computed" format:"date-time"`
	Status     types.String                                                      `tfsdk:"status" json:"status,computed"`
	Condition  customfield.NestedObject[AccountTokenConditionDataSourceModel]    `tfsdk:"condition" json:"condition,computed"`
	Policies   customfield.NestedObjectList[AccountTokenPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Filter     *AccountTokenFindOneByDataSourceModel                             `tfsdk:"filter"`
}

func (m *AccountTokenDataSourceModel) toReadParams(_ context.Context) (params accounts.TokenGetParams, diags diag.Diagnostics) {
	params = accounts.TokenGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AccountTokenDataSourceModel) toListParams(_ context.Context) (params accounts.TokenListParams, diags diag.Diagnostics) {
	params = accounts.TokenListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(accounts.TokenListParamsDirection(m.Filter.Direction.ValueString()))
	}

	return
}

type AccountTokenConditionDataSourceModel struct {
	RequestIP customfield.NestedObject[AccountTokenConditionRequestIPDataSourceModel] `tfsdk:"request_ip" json:"request_ip,computed"`
}

type AccountTokenConditionRequestIPDataSourceModel struct {
	In    customfield.List[types.String] `tfsdk:"in" json:"in,computed"`
	NotIn customfield.List[types.String] `tfsdk:"not_in" json:"not_in,computed"`
}

type AccountTokenPoliciesDataSourceModel struct {
	ID               types.String                                                                      `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                                                      `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups customfield.NestedObjectList[AccountTokenPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        customfield.Map[types.String]                                                     `tfsdk:"resources" json:"resources,computed"`
}

type AccountTokenPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                                                      `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[AccountTokenPoliciesPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
	Name types.String                                                                      `tfsdk:"name" json:"name,computed"`
}

type AccountTokenPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type AccountTokenFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Direction types.String `tfsdk:"direction" query:"direction,optional"`
}

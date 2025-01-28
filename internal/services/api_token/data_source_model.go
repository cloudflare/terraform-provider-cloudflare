// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/user"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultDataSourceEnvelope struct {
	Result APITokenDataSourceModel `json:"result,computed"`
}

type APITokenResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APITokenDataSourceModel] `json:"result,computed"`
}

type APITokenDataSourceModel struct {
	ID         types.String                                                  `tfsdk:"id" json:"-,computed"`
	TokenID    types.String                                                  `tfsdk:"token_id" path:"token_id,optional"`
	ExpiresOn  timetypes.RFC3339                                             `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	IssuedOn   timetypes.RFC3339                                             `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn timetypes.RFC3339                                             `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name       types.String                                                  `tfsdk:"name" json:"name,computed"`
	NotBefore  timetypes.RFC3339                                             `tfsdk:"not_before" json:"not_before,computed" format:"date-time"`
	Status     types.String                                                  `tfsdk:"status" json:"status,computed"`
	Condition  customfield.NestedObject[APITokenConditionDataSourceModel]    `tfsdk:"condition" json:"condition,computed"`
	Policies   customfield.NestedObjectList[APITokenPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Filter     *APITokenFindOneByDataSourceModel                             `tfsdk:"filter"`
}

func (m *APITokenDataSourceModel) toListParams(_ context.Context) (params user.TokenListParams, diags diag.Diagnostics) {
	params = user.TokenListParams{}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(TokenListParamsDirection(m.Filter.Direction.ValueString()))
	}

	return
}

type APITokenConditionDataSourceModel struct {
	RequestIP customfield.NestedObject[APITokenConditionRequestIPDataSourceModel] `tfsdk:"request_ip" json:"request_ip,computed"`
}

type APITokenConditionRequestIPDataSourceModel struct {
	In    customfield.List[types.String] `tfsdk:"in" json:"in,computed"`
	NotIn customfield.List[types.String] `tfsdk:"not_in" json:"not_in,computed"`
}

type APITokenPoliciesDataSourceModel struct {
	ID               types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                                                  `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups customfield.NestedObjectList[APITokenPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        customfield.Map[types.String]                                                 `tfsdk:"resources" json:"resources,computed"`
}

type APITokenPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[APITokenPoliciesPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
	Name types.String                                                                  `tfsdk:"name" json:"name,computed"`
}

type APITokenPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type APITokenFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction,optional"`
}

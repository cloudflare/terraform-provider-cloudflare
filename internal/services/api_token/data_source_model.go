// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/user"
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
	TokenID    types.String                        `tfsdk:"token_id" path:"token_id"`
	ID         types.String                        `tfsdk:"id" json:"id,computed"`
	IssuedOn   timetypes.RFC3339                   `tfsdk:"issued_on" json:"issued_on,computed"`
	LastUsedOn timetypes.RFC3339                   `tfsdk:"last_used_on" json:"last_used_on,computed"`
	ModifiedOn timetypes.RFC3339                   `tfsdk:"modified_on" json:"modified_on,computed"`
	ExpiresOn  timetypes.RFC3339                   `tfsdk:"expires_on" json:"expires_on,computed_optional"`
	Name       types.String                        `tfsdk:"name" json:"name,computed_optional"`
	NotBefore  timetypes.RFC3339                   `tfsdk:"not_before" json:"not_before,computed_optional"`
	Status     types.String                        `tfsdk:"status" json:"status,computed_optional"`
	Condition  *APITokenConditionDataSourceModel   `tfsdk:"condition" json:"condition,computed_optional"`
	Policies   *[]*APITokenPoliciesDataSourceModel `tfsdk:"policies" json:"policies,computed_optional"`
	Filter     *APITokenFindOneByDataSourceModel   `tfsdk:"filter"`
}

func (m *APITokenDataSourceModel) toListParams() (params user.TokenListParams, diags diag.Diagnostics) {
	params = user.TokenListParams{}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(user.TokenListParamsDirection(m.Filter.Direction.ValueString()))
	}

	return
}

type APITokenConditionDataSourceModel struct {
	RequestIP *APITokenConditionRequestIPDataSourceModel `tfsdk:"request_ip" json:"request.ip,computed_optional"`
}

type APITokenConditionRequestIPDataSourceModel struct {
	In    *[]types.String `tfsdk:"in" json:"in,computed_optional"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in,computed_optional"`
}

type APITokenPoliciesDataSourceModel struct {
	ID               types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                                                  `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups customfield.NestedObjectList[APITokenPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        customfield.NestedObject[APITokenPoliciesResourcesDataSourceModel]            `tfsdk:"resources" json:"resources,computed"`
}

type APITokenPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                         `tfsdk:"id" json:"id,computed"`
	Meta *APITokenPoliciesPermissionGroupsMetaDataSourceModel `tfsdk:"meta" json:"meta,computed_optional"`
	Name types.String                                         `tfsdk:"name" json:"name,computed"`
}

type APITokenPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type APITokenPoliciesResourcesDataSourceModel struct {
	Resource types.String `tfsdk:"resource" json:"resource,computed_optional"`
	Scope    types.String `tfsdk:"scope" json:"scope,computed_optional"`
}

type APITokenFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction"`
}

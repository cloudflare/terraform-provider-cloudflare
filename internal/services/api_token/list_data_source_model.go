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

type APITokensResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APITokensResultDataSourceModel] `json:"result,computed"`
}

type APITokensDataSourceModel struct {
	Direction types.String                                                 `tfsdk:"direction" query:"direction"`
	MaxItems  types.Int64                                                  `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[APITokensResultDataSourceModel] `tfsdk:"result"`
}

func (m *APITokensDataSourceModel) toListParams() (params user.TokenListParams, diags diag.Diagnostics) {
	params = user.TokenListParams{}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(user.TokenListParamsDirection(m.Direction.ValueString()))
	}

	return
}

type APITokensResultDataSourceModel struct {
	ID         types.String                         `tfsdk:"id" json:"id,computed"`
	Condition  *APITokensConditionDataSourceModel   `tfsdk:"condition" json:"condition,computed_optional"`
	ExpiresOn  timetypes.RFC3339                    `tfsdk:"expires_on" json:"expires_on,computed_optional" format:"date-time"`
	IssuedOn   timetypes.RFC3339                    `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn timetypes.RFC3339                    `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name       types.String                         `tfsdk:"name" json:"name,computed_optional"`
	NotBefore  timetypes.RFC3339                    `tfsdk:"not_before" json:"not_before,computed_optional" format:"date-time"`
	Policies   *[]*APITokensPoliciesDataSourceModel `tfsdk:"policies" json:"policies,computed_optional"`
	Status     types.String                         `tfsdk:"status" json:"status,computed_optional"`
}

type APITokensConditionDataSourceModel struct {
	RequestIP *APITokensConditionRequestIPDataSourceModel `tfsdk:"request_ip" json:"request.ip,computed_optional"`
}

type APITokensConditionRequestIPDataSourceModel struct {
	In    *[]types.String `tfsdk:"in" json:"in,computed_optional"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in,computed_optional"`
}

type APITokensPoliciesDataSourceModel struct {
	ID               types.String                                                                   `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                                                   `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups customfield.NestedObjectList[APITokensPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        customfield.NestedObject[APITokensPoliciesResourcesDataSourceModel]            `tfsdk:"resources" json:"resources,computed"`
}

type APITokensPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                          `tfsdk:"id" json:"id,computed"`
	Meta *APITokensPoliciesPermissionGroupsMetaDataSourceModel `tfsdk:"meta" json:"meta,computed_optional"`
	Name types.String                                          `tfsdk:"name" json:"name,computed"`
}

type APITokensPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type APITokensPoliciesResourcesDataSourceModel struct {
	Resource types.String `tfsdk:"resource" json:"resource,computed_optional"`
	Scope    types.String `tfsdk:"scope" json:"scope,computed_optional"`
}

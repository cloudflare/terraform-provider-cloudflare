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
	Condition  *APITokensConditionDataSourceModel   `tfsdk:"condition" json:"condition"`
	ExpiresOn  timetypes.RFC3339                    `tfsdk:"expires_on" json:"expires_on"`
	IssuedOn   timetypes.RFC3339                    `tfsdk:"issued_on" json:"issued_on,computed"`
	LastUsedOn timetypes.RFC3339                    `tfsdk:"last_used_on" json:"last_used_on,computed"`
	ModifiedOn timetypes.RFC3339                    `tfsdk:"modified_on" json:"modified_on,computed"`
	Name       types.String                         `tfsdk:"name" json:"name"`
	NotBefore  timetypes.RFC3339                    `tfsdk:"not_before" json:"not_before"`
	Policies   *[]*APITokensPoliciesDataSourceModel `tfsdk:"policies" json:"policies"`
	Status     types.String                         `tfsdk:"status" json:"status"`
}

type APITokensConditionDataSourceModel struct {
	RequestIP *APITokensConditionRequestIPDataSourceModel `tfsdk:"request_ip" json:"request.ip"`
}

type APITokensConditionRequestIPDataSourceModel struct {
	In    *[]types.String `tfsdk:"in" json:"in"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in"`
}

type APITokensPoliciesDataSourceModel struct {
	ID               types.String                                                                   `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                                                   `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups customfield.NestedObjectList[APITokensPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        customfield.NestedObject[APITokensPoliciesResourcesDataSourceModel]            `tfsdk:"resources" json:"resources,computed"`
}

type APITokensPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                          `tfsdk:"id" json:"id,computed"`
	Meta *APITokensPoliciesPermissionGroupsMetaDataSourceModel `tfsdk:"meta" json:"meta"`
	Name types.String                                          `tfsdk:"name" json:"name,computed"`
}

type APITokensPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key"`
	Value types.String `tfsdk:"value" json:"value"`
}

type APITokensPoliciesResourcesDataSourceModel struct {
	Resource types.String `tfsdk:"resource" json:"resource"`
	Scope    types.String `tfsdk:"scope" json:"scope"`
}

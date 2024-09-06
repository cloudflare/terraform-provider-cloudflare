// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"context"

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
	Direction types.String                                                 `tfsdk:"direction" query:"direction,optional"`
	MaxItems  types.Int64                                                  `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[APITokensResultDataSourceModel] `tfsdk:"result"`
}

func (m *APITokensDataSourceModel) toListParams(_ context.Context) (params user.TokenListParams, diags diag.Diagnostics) {
	params = user.TokenListParams{}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(user.TokenListParamsDirection(m.Direction.ValueString()))
	}

	return
}

type APITokensResultDataSourceModel struct {
	ID         types.String                                                   `tfsdk:"id" json:"id,computed"`
	Condition  customfield.NestedObject[APITokensConditionDataSourceModel]    `tfsdk:"condition" json:"condition,computed"`
	ExpiresOn  timetypes.RFC3339                                              `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	IssuedOn   timetypes.RFC3339                                              `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn timetypes.RFC3339                                              `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                                              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name       types.String                                                   `tfsdk:"name" json:"name,computed"`
	NotBefore  timetypes.RFC3339                                              `tfsdk:"not_before" json:"not_before,computed" format:"date-time"`
	Policies   customfield.NestedObjectList[APITokensPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Status     types.String                                                   `tfsdk:"status" json:"status,computed"`
}

type APITokensConditionDataSourceModel struct {
	RequestIP customfield.NestedObject[APITokensConditionRequestIPDataSourceModel] `tfsdk:"request_ip" json:"request.ip,computed"`
}

type APITokensConditionRequestIPDataSourceModel struct {
	In    customfield.List[types.String] `tfsdk:"in" json:"in,computed"`
	NotIn customfield.List[types.String] `tfsdk:"not_in" json:"not_in,computed"`
}

type APITokensPoliciesDataSourceModel struct {
	ID               types.String                                                                   `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                                                   `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups customfield.NestedObjectList[APITokensPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        customfield.NestedObject[APITokensPoliciesResourcesDataSourceModel]            `tfsdk:"resources" json:"resources,computed"`
}

type APITokensPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                                                   `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[APITokensPoliciesPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
	Name types.String                                                                   `tfsdk:"name" json:"name,computed"`
}

type APITokensPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type APITokensPoliciesResourcesDataSourceModel struct {
	Resource types.String `tfsdk:"resource" json:"resource,computed"`
	Scope    types.String `tfsdk:"scope" json:"scope,computed"`
}

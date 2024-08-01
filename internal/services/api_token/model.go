// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultEnvelope struct {
	Result APITokenModel `json:"result,computed"`
}

type APITokenModel struct {
	TokenID   jsontypes.Normalized      `tfsdk:"token_id" path:"token_id"`
	Name      types.String              `tfsdk:"name" json:"name"`
	Policies  *[]*APITokenPoliciesModel `tfsdk:"policies" json:"policies"`
	ExpiresOn timetypes.RFC3339         `tfsdk:"expires_on" json:"expires_on"`
	NotBefore timetypes.RFC3339         `tfsdk:"not_before" json:"not_before"`
	Status    types.String              `tfsdk:"status" json:"status"`
	Condition *APITokenConditionModel   `tfsdk:"condition" json:"condition"`
	ID        types.String              `tfsdk:"id" json:"id,computed"`
	Value     types.String              `tfsdk:"value" json:"value,computed"`
}

type APITokenPoliciesModel struct {
	ID               types.String                              `tfsdk:"id" json:"id,computed"`
	Effect           types.String                              `tfsdk:"effect" json:"effect"`
	PermissionGroups *[]*APITokenPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups"`
	Resources        jsontypes.Normalized                      `tfsdk:"resources" json:"resources"`
}

type APITokenPoliciesPermissionGroupsModel struct {
	ID   types.String         `tfsdk:"id" json:"id,computed"`
	Meta jsontypes.Normalized `tfsdk:"meta" json:"meta"`
	Name types.String         `tfsdk:"name" json:"name,computed"`
}

type APITokenConditionModel struct {
	RequestIP *APITokenConditionRequestIPModel `tfsdk:"request_ip" json:"request_ip"`
}

type APITokenConditionRequestIPModel struct {
	In    *[]types.String `tfsdk:"in" json:"in"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in"`
}

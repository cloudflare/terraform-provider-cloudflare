// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultEnvelope struct {
	Result APITokenModel `json:"result,computed"`
}

type APITokenModel struct {
	ID         types.String              `tfsdk:"id" json:"id,computed"`
	Name       types.String              `tfsdk:"name" json:"name"`
	Policies   *[]*APITokenPoliciesModel `tfsdk:"policies" json:"policies"`
	ExpiresOn  timetypes.RFC3339         `tfsdk:"expires_on" json:"expires_on"`
	NotBefore  timetypes.RFC3339         `tfsdk:"not_before" json:"not_before"`
	Status     types.String              `tfsdk:"status" json:"status"`
	Condition  *APITokenConditionModel   `tfsdk:"condition" json:"condition"`
	IssuedOn   timetypes.RFC3339         `tfsdk:"issued_on" json:"issued_on,computed"`
	LastUsedOn timetypes.RFC3339         `tfsdk:"last_used_on" json:"last_used_on,computed"`
	ModifiedOn timetypes.RFC3339         `tfsdk:"modified_on" json:"modified_on,computed"`
	Value      types.String              `tfsdk:"value" json:"value,computed"`
}

type APITokenPoliciesModel struct {
	ID               types.String                              `tfsdk:"id" json:"id,computed"`
	Effect           types.String                              `tfsdk:"effect" json:"effect"`
	PermissionGroups *[]*APITokenPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups"`
	Resources        *APITokenPoliciesResourcesModel           `tfsdk:"resources" json:"resources"`
}

type APITokenPoliciesPermissionGroupsModel struct {
	ID   types.String                               `tfsdk:"id" json:"id,computed"`
	Meta *APITokenPoliciesPermissionGroupsMetaModel `tfsdk:"meta" json:"meta"`
	Name types.String                               `tfsdk:"name" json:"name,computed"`
}

type APITokenPoliciesPermissionGroupsMetaModel struct {
	Key   types.String `tfsdk:"key" json:"key"`
	Value types.String `tfsdk:"value" json:"value"`
}

type APITokenPoliciesResourcesModel struct {
	Resource types.String `tfsdk:"resource" json:"resource"`
	Scope    types.String `tfsdk:"scope" json:"scope"`
}

type APITokenConditionModel struct {
	RequestIP *APITokenConditionRequestIPModel `tfsdk:"request_ip" json:"request.ip"`
}

type APITokenConditionRequestIPModel struct {
	In    *[]types.String `tfsdk:"in" json:"in"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in"`
}

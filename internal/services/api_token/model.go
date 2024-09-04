// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultEnvelope struct {
	Result APITokenModel `json:"result"`
}

type APITokenModel struct {
	ID         types.String                                        `tfsdk:"id" json:"id,computed"`
	ExpiresOn  timetypes.RFC3339                                   `tfsdk:"expires_on" json:"expires_on,computed_optional" format:"date-time"`
	Name       types.String                                        `tfsdk:"name" json:"name,computed_optional"`
	NotBefore  timetypes.RFC3339                                   `tfsdk:"not_before" json:"not_before,computed_optional" format:"date-time"`
	Status     types.String                                        `tfsdk:"status" json:"status,computed_optional"`
	Condition  customfield.NestedObject[APITokenConditionModel]    `tfsdk:"condition" json:"condition,computed_optional"`
	Policies   customfield.NestedObjectList[APITokenPoliciesModel] `tfsdk:"policies" json:"policies,computed_optional"`
	IssuedOn   timetypes.RFC3339                                   `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn timetypes.RFC3339                                   `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                                   `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      types.String                                        `tfsdk:"value" json:"value,computed"`
}

type APITokenConditionModel struct {
	RequestIP customfield.NestedObject[APITokenConditionRequestIPModel] `tfsdk:"request_ip" json:"request.ip,computed_optional"`
}

type APITokenConditionRequestIPModel struct {
	In    customfield.List[types.String] `tfsdk:"in" json:"in,computed_optional"`
	NotIn customfield.List[types.String] `tfsdk:"not_in" json:"not_in,computed_optional"`
}

type APITokenPoliciesModel struct {
	ID               types.String                                                        `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                                        `tfsdk:"effect" json:"effect,computed_optional"`
	PermissionGroups customfield.NestedObjectList[APITokenPoliciesPermissionGroupsModel] `tfsdk:"permission_groups" json:"permission_groups,computed_optional"`
	Resources        customfield.NestedObject[APITokenPoliciesResourcesModel]            `tfsdk:"resources" json:"resources,computed_optional"`
}

type APITokenPoliciesPermissionGroupsModel struct {
	ID   types.String                                                        `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[APITokenPoliciesPermissionGroupsMetaModel] `tfsdk:"meta" json:"meta,computed_optional"`
	Name types.String                                                        `tfsdk:"name" json:"name,computed"`
}

type APITokenPoliciesPermissionGroupsMetaModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type APITokenPoliciesResourcesModel struct {
	Resource types.String `tfsdk:"resource" json:"resource,computed_optional"`
	Scope    types.String `tfsdk:"scope" json:"scope,computed_optional"`
}

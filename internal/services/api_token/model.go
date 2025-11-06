// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultEnvelope struct {
	Result APITokenModel `json:"result"`
}

type APITokenModel struct {
	ID         types.String              `tfsdk:"id" json:"id,computed"`
	Name       types.String              `tfsdk:"name" json:"name,required"`
	Policies   *[]*APITokenPoliciesModel `tfsdk:"policies" json:"policies,required"`
	ExpiresOn  timetypes.RFC3339         `tfsdk:"expires_on" json:"expires_on,optional" format:"date-time"`
	NotBefore  timetypes.RFC3339         `tfsdk:"not_before" json:"not_before,optional" format:"date-time"`
	Condition  *APITokenConditionModel   `tfsdk:"condition" json:"condition,optional"`
	Status     types.String              `tfsdk:"status" json:"status,computed_optional"`
	IssuedOn   timetypes.RFC3339         `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn timetypes.RFC3339         `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339         `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      types.String              `tfsdk:"value" json:"value,computed,no_refresh"`
}

func (m APITokenModel) MarshalJSON() (data []byte, err error) {
	return MarshalCustom(m)
}

func (m APITokenModel) MarshalJSONForUpdate(state APITokenModel) (data []byte, err error) {
	return MarshalCustom(m)
}

type APITokenPoliciesModel struct {
	Effect           types.String                              `tfsdk:"effect" json:"effect,required"`
	PermissionGroups *[]*APITokenPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups,required"`
	Resources        types.String                              `tfsdk:"resources" json:"resources,required"`
}

type APITokenPoliciesPermissionGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type APITokenPoliciesPermissionGroupsMetaModel struct {
	Key   types.String `tfsdk:"key" json:"key,optional"`
	Value types.String `tfsdk:"value" json:"value,optional"`
}

type APITokenConditionModel struct {
	RequestIP *APITokenConditionRequestIPModel `tfsdk:"request_ip" json:"request_ip,optional"`
}

type APITokenConditionRequestIPModel struct {
	In    *[]types.String `tfsdk:"in" json:"in,optional"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in,optional"`
}

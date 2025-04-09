// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultEnvelope struct {
Result APITokenModel `json:"result"`
}

type APITokenModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,required"`
Policies *[]*APITokenPoliciesModel `tfsdk:"policies" json:"policies,required"`
ExpiresOn timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,optional" format:"date-time"`
NotBefore timetypes.RFC3339 `tfsdk:"not_before" json:"not_before,optional" format:"date-time"`
Status types.String `tfsdk:"status" json:"status,optional"`
Condition customfield.NestedObject[APITokenConditionModel] `tfsdk:"condition" json:"condition,computed_optional"`
IssuedOn timetypes.RFC3339 `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
LastUsedOn timetypes.RFC3339 `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

func (m APITokenModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m APITokenModel) MarshalJSONForUpdate(state APITokenModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type APITokenPoliciesModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
Effect types.String `tfsdk:"effect" json:"effect,required"`
PermissionGroups *[]*APITokenPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups,required"`
Resources *map[string]types.String `tfsdk:"resources" json:"resources,required"`
}

type APITokenPoliciesPermissionGroupsModel struct {
ID types.String `tfsdk:"id" json:"id,required"`
Meta *APITokenPoliciesPermissionGroupsMetaModel `tfsdk:"meta" json:"meta,optional"`
Name types.String `tfsdk:"name" json:"name,computed"`
}

type APITokenPoliciesPermissionGroupsMetaModel struct {
Key types.String `tfsdk:"key" json:"key,optional"`
Value types.String `tfsdk:"value" json:"value,optional"`
}

type APITokenConditionModel struct {
RequestIP customfield.NestedObject[APITokenConditionRequestIPModel] `tfsdk:"request_ip" json:"request_ip,computed_optional"`
}

type APITokenConditionRequestIPModel struct {
In *[]types.String `tfsdk:"in" json:"in,optional"`
NotIn *[]types.String `tfsdk:"not_in" json:"not_in,optional"`
}

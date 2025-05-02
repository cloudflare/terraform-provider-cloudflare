// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_token

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountTokenResultEnvelope struct {
	Result AccountTokenModel `json:"result"`
}

type AccountTokenModel struct {
	ID         types.String                  `tfsdk:"id" json:"id,computed"`
	AccountID  types.String                  `tfsdk:"account_id" path:"account_id,required"`
	Name       types.String                  `tfsdk:"name" json:"name,required"`
	Policies   *[]*AccountTokenPoliciesModel `tfsdk:"policies" json:"policies,required"`
	ExpiresOn  timetypes.RFC3339             `tfsdk:"expires_on" json:"expires_on,optional" format:"date-time"`
	NotBefore  timetypes.RFC3339             `tfsdk:"not_before" json:"not_before,optional" format:"date-time"`
	Status     types.String                  `tfsdk:"status" json:"status,computed"`
	Condition  *AccountTokenConditionModel   `tfsdk:"condition" json:"condition,optional"`
	IssuedOn   timetypes.RFC3339             `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn timetypes.RFC3339             `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Value      types.String                  `tfsdk:"value" json:"value,computed"`
}

func (m AccountTokenModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountTokenModel) MarshalJSONForUpdate(state AccountTokenModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccountTokenPoliciesModel struct {
	ID               types.String                                  `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                  `tfsdk:"effect" json:"effect,required"`
	PermissionGroups *[]*AccountTokenPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups,required"`
	Resources        *map[string]types.String                      `tfsdk:"resources" json:"resources,required"`
}

type AccountTokenPoliciesPermissionGroupsModel struct {
	ID   types.String                                                            `tfsdk:"id" json:"id,required"`
	Meta customfield.NestedObject[AccountTokenPoliciesPermissionGroupsMetaModel] `tfsdk:"meta" json:"meta,computed_optional"`
	Name types.String                                                            `tfsdk:"name" json:"name,computed"`
}

type AccountTokenPoliciesPermissionGroupsMetaModel struct {
	Key   types.String `tfsdk:"key" json:"key,optional"`
	Value types.String `tfsdk:"value" json:"value,optional"`
}

type AccountTokenConditionModel struct {
	RequestIP *AccountTokenConditionRequestIPModel `tfsdk:"request_ip" json:"request_ip,optional"`
}

type AccountTokenConditionRequestIPModel struct {
	In    *[]types.String `tfsdk:"in" json:"in,optional"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in,optional"`
}

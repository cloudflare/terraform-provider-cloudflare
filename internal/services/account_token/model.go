// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_token

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountTokenResultEnvelope struct {
	Result AccountTokenModel `json:"result"`
}

type AccountTokenModel struct {
	ID                     types.String                  `tfsdk:"id" json:"id,computed"`
	AccountID              types.String                  `tfsdk:"account_id" path:"account_id,required"`
	Name                   types.String                  `tfsdk:"name" json:"name,required"`
	Policies               *[]*AccountTokenPoliciesModel `tfsdk:"policies" json:"policies,required"`
	ExpiresOn              timetypes.RFC3339             `tfsdk:"expires_on" json:"expires_on,optional" format:"date-time"`
	NotBefore              timetypes.RFC3339             `tfsdk:"not_before" json:"not_before,optional" format:"date-time"`
	Condition              *AccountTokenConditionModel   `tfsdk:"condition" json:"condition,optional"`
	Status                 types.String                  `tfsdk:"status" json:"status,computed_optional"`
	CreatorEmailAtCreation types.String                  `tfsdk:"creator_email_at_creation" json:"creator_email_at_creation,computed"`
	IssuedOn               timetypes.RFC3339             `tfsdk:"issued_on" json:"issued_on,computed" format:"date-time"`
	LastUsedOn             timetypes.RFC3339             `tfsdk:"last_used_on" json:"last_used_on,computed" format:"date-time"`
	ModifiedOn             timetypes.RFC3339             `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	ProvisionerID          types.String                  `tfsdk:"provisioner_id" json:"provisioner_id,computed"`
	ProvisionerType        types.String                  `tfsdk:"provisioner_type" json:"provisioner_type,computed"`
	Value                  types.String                  `tfsdk:"value" json:"value,computed,no_refresh"`
}

func (m AccountTokenModel) MarshalJSON() (data []byte, err error) {
	return MarshalCustom(m)
}

func (m AccountTokenModel) MarshalJSONForUpdate(state AccountTokenModel) (data []byte, err error) {
	return MarshalCustom(m)
}

type AccountTokenPoliciesModel struct {
	Effect           types.String                                  `tfsdk:"effect" json:"effect,required"`
	PermissionGroups *[]*AccountTokenPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups,required"`
	Resources        types.String                                  `tfsdk:"resources" json:"resources,required"`
}

type AccountTokenPoliciesPermissionGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type AccountTokenPoliciesPermissionGroupsMetaModel struct {
	Category    types.String      `tfsdk:"category" json:"category,optional"`
	Deprecated  types.String      `tfsdk:"deprecated" json:"deprecated,optional"`
	Description types.String      `tfsdk:"description" json:"description,optional"`
	Editable    types.String      `tfsdk:"editable" json:"editable,optional"`
	EolAt       timetypes.RFC3339 `tfsdk:"eol_at" json:"eol_at,optional" format:"date-time"`
	Label       types.String      `tfsdk:"label" json:"label,optional"`
	Scopes      types.String      `tfsdk:"scopes" json:"scopes,optional"`
	Visibility  types.String      `tfsdk:"visibility" json:"visibility,optional"`
}

type AccountTokenConditionModel struct {
	RequestIP *AccountTokenConditionRequestIPModel `tfsdk:"request_ip" json:"request_ip,optional"`
}

type AccountTokenConditionRequestIPModel struct {
	In    *[]types.String `tfsdk:"in" json:"in,optional"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in,optional"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserResultEnvelope struct {
	Result UserModel `json:"result"`
}

type UserModel struct {
	ID                             types.String                                         `tfsdk:"id" json:"id,computed"`
	Country                        types.String                                         `tfsdk:"country" json:"country,optional"`
	FirstName                      types.String                                         `tfsdk:"first_name" json:"first_name,optional"`
	LastName                       types.String                                         `tfsdk:"last_name" json:"last_name,optional"`
	Telephone                      types.String                                         `tfsdk:"telephone" json:"telephone,optional"`
	Zipcode                        types.String                                         `tfsdk:"zipcode" json:"zipcode,optional"`
	HasBusinessZones               types.Bool                                           `tfsdk:"has_business_zones" json:"has_business_zones,computed"`
	HasEnterpriseZones             types.Bool                                           `tfsdk:"has_enterprise_zones" json:"has_enterprise_zones,computed"`
	HasProZones                    types.Bool                                           `tfsdk:"has_pro_zones" json:"has_pro_zones,computed"`
	Suspended                      types.Bool                                           `tfsdk:"suspended" json:"suspended,computed"`
	TwoFactorAuthenticationEnabled types.Bool                                           `tfsdk:"two_factor_authentication_enabled" json:"two_factor_authentication_enabled,computed"`
	TwoFactorAuthenticationLocked  types.Bool                                           `tfsdk:"two_factor_authentication_locked" json:"two_factor_authentication_locked,computed"`
	Betas                          customfield.List[types.String]                       `tfsdk:"betas" json:"betas,computed"`
	Organizations                  customfield.NestedObjectList[UserOrganizationsModel] `tfsdk:"organizations" json:"organizations,computed"`
}

func (m UserModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserModel) MarshalJSONForUpdate(state UserModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type UserOrganizationsModel struct {
	ID          types.String                   `tfsdk:"id" json:"id,computed"`
	Name        types.String                   `tfsdk:"name" json:"name,computed"`
	Permissions customfield.List[types.String] `tfsdk:"permissions" json:"permissions,computed"`
	Roles       customfield.List[types.String] `tfsdk:"roles" json:"roles,computed"`
	Status      types.String                   `tfsdk:"status" json:"status,computed"`
}

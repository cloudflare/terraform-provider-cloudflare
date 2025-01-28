// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMembersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountMembersResultDataSourceModel] `json:"result,computed"`
}

type AccountMembersDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	Direction types.String                                                      `tfsdk:"direction" query:"direction,optional"`
	Order     types.String                                                      `tfsdk:"order" query:"order,optional"`
	Status    types.String                                                      `tfsdk:"status" query:"status,optional"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AccountMembersResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccountMembersDataSourceModel) toListParams(_ context.Context) (params accounts.MemberListParams, diags diag.Diagnostics) {
	params = accounts.MemberListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(MemberListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(MemberListParamsOrder(m.Order.ValueString()))
	}
	if !m.Status.IsNull() {
		params.Status = cloudflare.F(MemberListParamsStatus(m.Status.ValueString()))
	}

	return
}

type AccountMembersResultDataSourceModel struct {
	ID       types.String                                                        `tfsdk:"id" json:"id,computed"`
	Policies customfield.NestedObjectList[AccountMembersPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Roles    customfield.NestedObjectList[AccountMembersRolesDataSourceModel]    `tfsdk:"roles" json:"roles,computed"`
	Status   types.String                                                        `tfsdk:"status" json:"status,computed"`
	User     customfield.NestedObject[AccountMembersUserDataSourceModel]         `tfsdk:"user" json:"user,computed"`
}

type AccountMembersPoliciesDataSourceModel struct {
	ID               types.String                                                                        `tfsdk:"id" json:"id,computed"`
	Access           types.String                                                                        `tfsdk:"access" json:"access,computed"`
	PermissionGroups customfield.NestedObjectList[AccountMembersPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	ResourceGroups   customfield.NestedObjectList[AccountMembersPoliciesResourceGroupsDataSourceModel]   `tfsdk:"resource_groups" json:"resource_groups,computed"`
}

type AccountMembersPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                                                        `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[AccountMembersPoliciesPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
	Name types.String                                                                        `tfsdk:"name" json:"name,computed"`
}

type AccountMembersPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type AccountMembersPoliciesResourceGroupsDataSourceModel struct {
	ID    types.String                                                                           `tfsdk:"id" json:"id,computed"`
	Scope customfield.NestedObjectList[AccountMembersPoliciesResourceGroupsScopeDataSourceModel] `tfsdk:"scope" json:"scope,computed"`
	Meta  customfield.NestedObject[AccountMembersPoliciesResourceGroupsMetaDataSourceModel]      `tfsdk:"meta" json:"meta,computed"`
	Name  types.String                                                                           `tfsdk:"name" json:"name,computed"`
}

type AccountMembersPoliciesResourceGroupsScopeDataSourceModel struct {
	Key     types.String                                                                                  `tfsdk:"key" json:"key,computed"`
	Objects customfield.NestedObjectList[AccountMembersPoliciesResourceGroupsScopeObjectsDataSourceModel] `tfsdk:"objects" json:"objects,computed"`
}

type AccountMembersPoliciesResourceGroupsScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type AccountMembersPoliciesResourceGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type AccountMembersRolesDataSourceModel struct {
	ID          types.String                                                            `tfsdk:"id" json:"id,computed"`
	Description types.String                                                            `tfsdk:"description" json:"description,computed"`
	Name        types.String                                                            `tfsdk:"name" json:"name,computed"`
	Permissions customfield.NestedObject[AccountMembersRolesPermissionsDataSourceModel] `tfsdk:"permissions" json:"permissions,computed"`
}

type AccountMembersRolesPermissionsDataSourceModel struct {
	Analytics    customfield.NestedObject[AccountMembersRolesPermissionsAnalyticsDataSourceModel]    `tfsdk:"analytics" json:"analytics,computed"`
	Billing      customfield.NestedObject[AccountMembersRolesPermissionsBillingDataSourceModel]      `tfsdk:"billing" json:"billing,computed"`
	CachePurge   customfield.NestedObject[AccountMembersRolesPermissionsCachePurgeDataSourceModel]   `tfsdk:"cache_purge" json:"cache_purge,computed"`
	DNS          customfield.NestedObject[AccountMembersRolesPermissionsDNSDataSourceModel]          `tfsdk:"dns" json:"dns,computed"`
	DNSRecords   customfield.NestedObject[AccountMembersRolesPermissionsDNSRecordsDataSourceModel]   `tfsdk:"dns_records" json:"dns_records,computed"`
	LB           customfield.NestedObject[AccountMembersRolesPermissionsLBDataSourceModel]           `tfsdk:"lb" json:"lb,computed"`
	Logs         customfield.NestedObject[AccountMembersRolesPermissionsLogsDataSourceModel]         `tfsdk:"logs" json:"logs,computed"`
	Organization customfield.NestedObject[AccountMembersRolesPermissionsOrganizationDataSourceModel] `tfsdk:"organization" json:"organization,computed"`
	SSL          customfield.NestedObject[AccountMembersRolesPermissionsSSLDataSourceModel]          `tfsdk:"ssl" json:"ssl,computed"`
	WAF          customfield.NestedObject[AccountMembersRolesPermissionsWAFDataSourceModel]          `tfsdk:"waf" json:"waf,computed"`
	ZoneSettings customfield.NestedObject[AccountMembersRolesPermissionsZoneSettingsDataSourceModel] `tfsdk:"zone_settings" json:"zone_settings,computed"`
	Zones        customfield.NestedObject[AccountMembersRolesPermissionsZonesDataSourceModel]        `tfsdk:"zones" json:"zones,computed"`
}

type AccountMembersRolesPermissionsAnalyticsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsBillingDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsCachePurgeDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsDNSDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsDNSRecordsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsLBDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsLogsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsOrganizationDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsSSLDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsWAFDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsZoneSettingsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersRolesPermissionsZonesDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountMembersUserDataSourceModel struct {
	Email                          types.String `tfsdk:"email" json:"email,computed"`
	ID                             types.String `tfsdk:"id" json:"id,computed"`
	FirstName                      types.String `tfsdk:"first_name" json:"first_name,computed"`
	LastName                       types.String `tfsdk:"last_name" json:"last_name,computed"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled" json:"two_factor_authentication_enabled,computed"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_role

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountRoleResultDataSourceEnvelope struct {
	Result AccountRoleDataSourceModel `json:"result,computed"`
}

type AccountRoleDataSourceModel struct {
	AccountID   types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	RoleID      types.String                                                    `tfsdk:"role_id" path:"role_id,required"`
	Description types.String                                                    `tfsdk:"description" json:"description,computed"`
	ID          types.String                                                    `tfsdk:"id" json:"id,computed"`
	Name        types.String                                                    `tfsdk:"name" json:"name,computed"`
	Permissions customfield.NestedObject[AccountRolePermissionsDataSourceModel] `tfsdk:"permissions" json:"permissions,computed"`
}

func (m *AccountRoleDataSourceModel) toReadParams(_ context.Context) (params accounts.RoleGetParams, diags diag.Diagnostics) {
	params = accounts.RoleGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AccountRolePermissionsDataSourceModel struct {
	Analytics    customfield.NestedObject[AccountRolePermissionsAnalyticsDataSourceModel]    `tfsdk:"analytics" json:"analytics,computed"`
	Billing      customfield.NestedObject[AccountRolePermissionsBillingDataSourceModel]      `tfsdk:"billing" json:"billing,computed"`
	CachePurge   customfield.NestedObject[AccountRolePermissionsCachePurgeDataSourceModel]   `tfsdk:"cache_purge" json:"cache_purge,computed"`
	DNS          customfield.NestedObject[AccountRolePermissionsDNSDataSourceModel]          `tfsdk:"dns" json:"dns,computed"`
	DNSRecords   customfield.NestedObject[AccountRolePermissionsDNSRecordsDataSourceModel]   `tfsdk:"dns_records" json:"dns_records,computed"`
	LB           customfield.NestedObject[AccountRolePermissionsLBDataSourceModel]           `tfsdk:"lb" json:"lb,computed"`
	Logs         customfield.NestedObject[AccountRolePermissionsLogsDataSourceModel]         `tfsdk:"logs" json:"logs,computed"`
	Organization customfield.NestedObject[AccountRolePermissionsOrganizationDataSourceModel] `tfsdk:"organization" json:"organization,computed"`
	SSL          customfield.NestedObject[AccountRolePermissionsSSLDataSourceModel]          `tfsdk:"ssl" json:"ssl,computed"`
	WAF          customfield.NestedObject[AccountRolePermissionsWAFDataSourceModel]          `tfsdk:"waf" json:"waf,computed"`
	ZoneSettings customfield.NestedObject[AccountRolePermissionsZoneSettingsDataSourceModel] `tfsdk:"zone_settings" json:"zone_settings,computed"`
	Zones        customfield.NestedObject[AccountRolePermissionsZonesDataSourceModel]        `tfsdk:"zones" json:"zones,computed"`
}

type AccountRolePermissionsAnalyticsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsBillingDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsCachePurgeDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsDNSDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsDNSRecordsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsLBDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsLogsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsOrganizationDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsSSLDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsWAFDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsZoneSettingsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolePermissionsZonesDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

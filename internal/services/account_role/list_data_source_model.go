// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_role

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountRolesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountRolesResultDataSourceModel] `json:"result,computed"`
}

type AccountRolesDataSourceModel struct {
	AccountID types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                     `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AccountRolesResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccountRolesDataSourceModel) toListParams(_ context.Context) (params accounts.RoleListParams, diags diag.Diagnostics) {
	params = accounts.RoleListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AccountRolesResultDataSourceModel struct {
	ID          types.String                                                     `tfsdk:"id" json:"id,computed"`
	Description types.String                                                     `tfsdk:"description" json:"description,computed"`
	Name        types.String                                                     `tfsdk:"name" json:"name,computed"`
	Permissions customfield.NestedObject[AccountRolesPermissionsDataSourceModel] `tfsdk:"permissions" json:"permissions,computed"`
}

type AccountRolesPermissionsDataSourceModel struct {
	Analytics    customfield.NestedObject[AccountRolesPermissionsAnalyticsDataSourceModel]    `tfsdk:"analytics" json:"analytics,computed"`
	Billing      customfield.NestedObject[AccountRolesPermissionsBillingDataSourceModel]      `tfsdk:"billing" json:"billing,computed"`
	CachePurge   customfield.NestedObject[AccountRolesPermissionsCachePurgeDataSourceModel]   `tfsdk:"cache_purge" json:"cache_purge,computed"`
	DNS          customfield.NestedObject[AccountRolesPermissionsDNSDataSourceModel]          `tfsdk:"dns" json:"dns,computed"`
	DNSRecords   customfield.NestedObject[AccountRolesPermissionsDNSRecordsDataSourceModel]   `tfsdk:"dns_records" json:"dns_records,computed"`
	LB           customfield.NestedObject[AccountRolesPermissionsLBDataSourceModel]           `tfsdk:"lb" json:"lb,computed"`
	Logs         customfield.NestedObject[AccountRolesPermissionsLogsDataSourceModel]         `tfsdk:"logs" json:"logs,computed"`
	Organization customfield.NestedObject[AccountRolesPermissionsOrganizationDataSourceModel] `tfsdk:"organization" json:"organization,computed"`
	SSL          customfield.NestedObject[AccountRolesPermissionsSSLDataSourceModel]          `tfsdk:"ssl" json:"ssl,computed"`
	WAF          customfield.NestedObject[AccountRolesPermissionsWAFDataSourceModel]          `tfsdk:"waf" json:"waf,computed"`
	ZoneSettings customfield.NestedObject[AccountRolesPermissionsZoneSettingsDataSourceModel] `tfsdk:"zone_settings" json:"zone_settings,computed"`
	Zones        customfield.NestedObject[AccountRolesPermissionsZonesDataSourceModel]        `tfsdk:"zones" json:"zones,computed"`
}

type AccountRolesPermissionsAnalyticsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsBillingDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsCachePurgeDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsDNSDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsDNSRecordsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsLBDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsLogsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsOrganizationDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsSSLDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsWAFDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsZoneSettingsDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

type AccountRolesPermissionsZonesDataSourceModel struct {
	Read  types.Bool `tfsdk:"read" json:"read,computed"`
	Write types.Bool `tfsdk:"write" json:"write,computed"`
}

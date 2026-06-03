// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor_group

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/load_balancers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerMonitorGroupsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[LoadBalancerMonitorGroupsResultDataSourceModel] `json:"result,computed"`
}

type LoadBalancerMonitorGroupsDataSourceModel struct {
	AccountID types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                  `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[LoadBalancerMonitorGroupsResultDataSourceModel] `tfsdk:"result"`
}

func (m *LoadBalancerMonitorGroupsDataSourceModel) toListParams(_ context.Context) (params load_balancers.MonitorGroupListParams, diags diag.Diagnostics) {
	params = load_balancers.MonitorGroupListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type LoadBalancerMonitorGroupsResultDataSourceModel struct {
	ID          types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Description types.String                                                                 `tfsdk:"description" json:"description,computed"`
	Members     customfield.NestedObjectSet[LoadBalancerMonitorGroupsMembersDataSourceModel] `tfsdk:"members" json:"members,computed"`
	CreatedOn   timetypes.RFC3339                                                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339                                                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

type LoadBalancerMonitorGroupsMembersDataSourceModel struct {
	Enabled        types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	MonitorID      types.String      `tfsdk:"monitor_id" json:"monitor_id,computed"`
	MonitoringOnly types.Bool        `tfsdk:"monitoring_only" json:"monitoring_only,computed"`
	MustBeHealthy  types.Bool        `tfsdk:"must_be_healthy" json:"must_be_healthy,computed"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt      timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

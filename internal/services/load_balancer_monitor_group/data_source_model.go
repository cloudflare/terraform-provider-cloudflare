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

type LoadBalancerMonitorGroupResultDataSourceEnvelope struct {
	Result LoadBalancerMonitorGroupDataSourceModel `json:"result,computed"`
}

type LoadBalancerMonitorGroupDataSourceModel struct {
	ID             types.String                                                                `tfsdk:"id" path:"monitor_group_id,computed"`
	MonitorGroupID types.String                                                                `tfsdk:"monitor_group_id" path:"monitor_group_id,required"`
	AccountID      types.String                                                                `tfsdk:"account_id" path:"account_id,required"`
	CreatedOn      timetypes.RFC3339                                                           `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description    types.String                                                                `tfsdk:"description" json:"description,computed"`
	ModifiedOn     timetypes.RFC3339                                                           `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Members        customfield.NestedObjectSet[LoadBalancerMonitorGroupMembersDataSourceModel] `tfsdk:"members" json:"members,computed"`
}

func (m *LoadBalancerMonitorGroupDataSourceModel) toReadParams(_ context.Context) (params load_balancers.MonitorGroupGetParams, diags diag.Diagnostics) {
	params = load_balancers.MonitorGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type LoadBalancerMonitorGroupMembersDataSourceModel struct {
	Enabled        types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	MonitorID      types.String      `tfsdk:"monitor_id" json:"monitor_id,computed"`
	MonitoringOnly types.Bool        `tfsdk:"monitoring_only" json:"monitoring_only,computed"`
	MustBeHealthy  types.Bool        `tfsdk:"must_be_healthy" json:"must_be_healthy,computed"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt      timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

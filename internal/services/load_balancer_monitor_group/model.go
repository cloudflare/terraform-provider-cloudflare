// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor_group

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerMonitorGroupResultEnvelope struct {
	Result LoadBalancerMonitorGroupModel `json:"result"`
}

type LoadBalancerMonitorGroupModel struct {
	ID          types.String                             `tfsdk:"id" json:"id,required"`
	AccountID   types.String                             `tfsdk:"account_id" path:"account_id,required"`
	Description types.String                             `tfsdk:"description" json:"description,required"`
	Members     *[]*LoadBalancerMonitorGroupMembersModel `tfsdk:"members" json:"members,required"`
	CreatedAt   timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt   timetypes.RFC3339                        `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m LoadBalancerMonitorGroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m LoadBalancerMonitorGroupModel) MarshalJSONForUpdate(state LoadBalancerMonitorGroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type LoadBalancerMonitorGroupMembersModel struct {
	Enabled        types.Bool        `tfsdk:"enabled" json:"enabled,required"`
	MonitorID      types.String      `tfsdk:"monitor_id" json:"monitor_id,required"`
	MonitoringOnly types.Bool        `tfsdk:"monitoring_only" json:"monitoring_only,required"`
	MustBeHealthy  types.Bool        `tfsdk:"must_be_healthy" json:"must_be_healthy,required"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt      timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

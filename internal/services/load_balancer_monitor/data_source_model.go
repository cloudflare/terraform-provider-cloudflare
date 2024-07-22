// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerMonitorResultDataSourceEnvelope struct {
	Result LoadBalancerMonitorDataSourceModel `json:"result,computed"`
}

type LoadBalancerMonitorResultListDataSourceEnvelope struct {
	Result *[]*LoadBalancerMonitorDataSourceModel `json:"result,computed"`
}

type LoadBalancerMonitorDataSourceModel struct {
	AccountID       types.String                                 `tfsdk:"account_id" path:"account_id"`
	MonitorID       types.String                                 `tfsdk:"monitor_id" path:"monitor_id"`
	ID              types.String                                 `tfsdk:"id" json:"id"`
	AllowInsecure   types.Bool                                   `tfsdk:"allow_insecure" json:"allow_insecure,computed"`
	ConsecutiveDown types.Int64                                  `tfsdk:"consecutive_down" json:"consecutive_down,computed"`
	ConsecutiveUp   types.Int64                                  `tfsdk:"consecutive_up" json:"consecutive_up,computed"`
	CreatedOn       types.String                                 `tfsdk:"created_on" json:"created_on,computed"`
	Description     types.String                                 `tfsdk:"description" json:"description"`
	ExpectedBody    types.String                                 `tfsdk:"expected_body" json:"expected_body"`
	ExpectedCodes   types.String                                 `tfsdk:"expected_codes" json:"expected_codes,computed"`
	FollowRedirects types.Bool                                   `tfsdk:"follow_redirects" json:"follow_redirects,computed"`
	Header          types.String                                 `tfsdk:"header" json:"header"`
	Interval        types.Int64                                  `tfsdk:"interval" json:"interval,computed"`
	Method          types.String                                 `tfsdk:"method" json:"method,computed"`
	ModifiedOn      types.String                                 `tfsdk:"modified_on" json:"modified_on,computed"`
	Path            types.String                                 `tfsdk:"path" json:"path,computed"`
	Port            types.Int64                                  `tfsdk:"port" json:"port,computed"`
	ProbeZone       types.String                                 `tfsdk:"probe_zone" json:"probe_zone"`
	Retries         types.Int64                                  `tfsdk:"retries" json:"retries,computed"`
	Timeout         types.Int64                                  `tfsdk:"timeout" json:"timeout,computed"`
	Type            types.String                                 `tfsdk:"type" json:"type,computed"`
	FindOneBy       *LoadBalancerMonitorFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type LoadBalancerMonitorFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}

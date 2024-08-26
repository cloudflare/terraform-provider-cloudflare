// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerMonitorResultEnvelope struct {
	Result LoadBalancerMonitorModel `json:"result"`
}

type LoadBalancerMonitorModel struct {
	ID              types.String               `tfsdk:"id" json:"id,computed"`
	AccountID       types.String               `tfsdk:"account_id" path:"account_id"`
	Description     types.String               `tfsdk:"description" json:"description"`
	ExpectedBody    types.String               `tfsdk:"expected_body" json:"expected_body"`
	ProbeZone       types.String               `tfsdk:"probe_zone" json:"probe_zone"`
	Header          map[string]*[]types.String `tfsdk:"header" json:"header"`
	AllowInsecure   types.Bool                 `tfsdk:"allow_insecure" json:"allow_insecure"`
	ConsecutiveDown types.Int64                `tfsdk:"consecutive_down" json:"consecutive_down"`
	ConsecutiveUp   types.Int64                `tfsdk:"consecutive_up" json:"consecutive_up"`
	ExpectedCodes   types.String               `tfsdk:"expected_codes" json:"expected_codes"`
	FollowRedirects types.Bool                 `tfsdk:"follow_redirects" json:"follow_redirects"`
	Interval        types.Int64                `tfsdk:"interval" json:"interval"`
	Method          types.String               `tfsdk:"method" json:"method"`
	Path            types.String               `tfsdk:"path" json:"path"`
	Port            types.Int64                `tfsdk:"port" json:"port"`
	Retries         types.Int64                `tfsdk:"retries" json:"retries"`
	Timeout         types.Int64                `tfsdk:"timeout" json:"timeout"`
	Type            types.String               `tfsdk:"type" json:"type"`
	CreatedOn       timetypes.RFC3339          `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn      timetypes.RFC3339          `tfsdk:"modified_on" json:"modified_on,computed"`
}

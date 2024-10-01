// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LoadBalancerMonitorResultEnvelope struct {
	Result LoadBalancerMonitorModel `json:"result"`
}

type LoadBalancerMonitorModel struct {
	ID              types.String               `tfsdk:"id" json:"id,computed"`
	AccountID       types.String               `tfsdk:"account_id" path:"account_id,required"`
	Description     types.String               `tfsdk:"description" json:"description,optional"`
	ExpectedBody    types.String               `tfsdk:"expected_body" json:"expected_body,optional"`
	ProbeZone       types.String               `tfsdk:"probe_zone" json:"probe_zone,optional"`
	Header          map[string]*[]types.String `tfsdk:"header" json:"header,optional"`
	AllowInsecure   types.Bool                 `tfsdk:"allow_insecure" json:"allow_insecure,computed_optional"`
	ConsecutiveDown types.Int64                `tfsdk:"consecutive_down" json:"consecutive_down,computed_optional"`
	ConsecutiveUp   types.Int64                `tfsdk:"consecutive_up" json:"consecutive_up,computed_optional"`
	ExpectedCodes   types.String               `tfsdk:"expected_codes" json:"expected_codes,computed_optional"`
	FollowRedirects types.Bool                 `tfsdk:"follow_redirects" json:"follow_redirects,computed_optional"`
	Interval        types.Int64                `tfsdk:"interval" json:"interval,computed_optional"`
	Method          types.String               `tfsdk:"method" json:"method,computed_optional"`
	Path            types.String               `tfsdk:"path" json:"path,computed_optional"`
	Port            types.Int64                `tfsdk:"port" json:"port,computed_optional"`
	Retries         types.Int64                `tfsdk:"retries" json:"retries,computed_optional"`
	Timeout         types.Int64                `tfsdk:"timeout" json:"timeout,computed_optional"`
	Type            types.String               `tfsdk:"type" json:"type,computed_optional"`
	CreatedOn       types.String               `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn      types.String               `tfsdk:"modified_on" json:"modified_on,computed"`
}

func (m LoadBalancerMonitorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m LoadBalancerMonitorModel) MarshalJSONForUpdate(state LoadBalancerMonitorModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HealthcheckResultEnvelope struct {
	Result HealthcheckModel `json:"result,computed"`
}

type HealthcheckModel struct {
	ID                   types.String                `tfsdk:"id" json:"id,computed"`
	ZoneID               types.String                `tfsdk:"zone_id" path:"zone_id"`
	Address              types.String                `tfsdk:"address" json:"address"`
	Name                 types.String                `tfsdk:"name" json:"name"`
	CheckRegions         *[]types.String             `tfsdk:"check_regions" json:"check_regions"`
	ConsecutiveFails     types.Int64                 `tfsdk:"consecutive_fails" json:"consecutive_fails"`
	ConsecutiveSuccesses types.Int64                 `tfsdk:"consecutive_successes" json:"consecutive_successes"`
	Description          types.String                `tfsdk:"description" json:"description"`
	HTTPConfig           *HealthcheckHTTPConfigModel `tfsdk:"http_config" json:"http_config"`
	Interval             types.Int64                 `tfsdk:"interval" json:"interval"`
	Retries              types.Int64                 `tfsdk:"retries" json:"retries"`
	Suspended            types.Bool                  `tfsdk:"suspended" json:"suspended"`
	TCPConfig            *HealthcheckTCPConfigModel  `tfsdk:"tcp_config" json:"tcp_config"`
	Timeout              types.Int64                 `tfsdk:"timeout" json:"timeout"`
	Type                 types.String                `tfsdk:"type" json:"type"`
}

type HealthcheckHTTPConfigModel struct {
	AllowInsecure   types.Bool                 `tfsdk:"allow_insecure" json:"allow_insecure"`
	ExpectedBody    types.String               `tfsdk:"expected_body" json:"expected_body"`
	ExpectedCodes   *[]types.String            `tfsdk:"expected_codes" json:"expected_codes"`
	FollowRedirects types.Bool                 `tfsdk:"follow_redirects" json:"follow_redirects"`
	Header          map[string]*[]types.String `tfsdk:"header" json:"header"`
	Method          types.String               `tfsdk:"method" json:"method"`
	Path            types.String               `tfsdk:"path" json:"path"`
	Port            types.Int64                `tfsdk:"port" json:"port"`
}

type HealthcheckTCPConfigModel struct {
	Method types.String `tfsdk:"method" json:"method"`
	Port   types.Int64  `tfsdk:"port" json:"port"`
}

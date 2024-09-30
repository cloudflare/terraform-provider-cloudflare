// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HealthcheckResultEnvelope struct {
	Result HealthcheckModel `json:"result"`
}

type HealthcheckModel struct {
	ID                   types.String                                         `tfsdk:"id" json:"id,computed"`
	ZoneID               types.String                                         `tfsdk:"zone_id" path:"zone_id,required"`
	Address              types.String                                         `tfsdk:"address" json:"address,required"`
	Name                 types.String                                         `tfsdk:"name" json:"name,required"`
	Description          types.String                                         `tfsdk:"description" json:"description,optional"`
	CheckRegions         *[]types.String                                      `tfsdk:"check_regions" json:"check_regions,optional"`
	ConsecutiveFails     types.Int64                                          `tfsdk:"consecutive_fails" json:"consecutive_fails,computed_optional"`
	ConsecutiveSuccesses types.Int64                                          `tfsdk:"consecutive_successes" json:"consecutive_successes,computed_optional"`
	Interval             types.Int64                                          `tfsdk:"interval" json:"interval,computed_optional"`
	Retries              types.Int64                                          `tfsdk:"retries" json:"retries,computed_optional"`
	Suspended            types.Bool                                           `tfsdk:"suspended" json:"suspended,computed_optional"`
	Timeout              types.Int64                                          `tfsdk:"timeout" json:"timeout,computed_optional"`
	Type                 types.String                                         `tfsdk:"type" json:"type,computed_optional"`
	HTTPConfig           customfield.NestedObject[HealthcheckHTTPConfigModel] `tfsdk:"http_config" json:"http_config,computed_optional"`
	TCPConfig            customfield.NestedObject[HealthcheckTCPConfigModel]  `tfsdk:"tcp_config" json:"tcp_config,computed_optional"`
	CreatedOn            timetypes.RFC3339                                    `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	FailureReason        types.String                                         `tfsdk:"failure_reason" json:"failure_reason,computed"`
	ModifiedOn           timetypes.RFC3339                                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Status               types.String                                         `tfsdk:"status" json:"status,computed"`
}

type HealthcheckHTTPConfigModel struct {
	AllowInsecure   types.Bool                 `tfsdk:"allow_insecure" json:"allow_insecure,computed_optional"`
	ExpectedBody    types.String               `tfsdk:"expected_body" json:"expected_body,optional"`
	ExpectedCodes   *[]types.String            `tfsdk:"expected_codes" json:"expected_codes,optional"`
	FollowRedirects types.Bool                 `tfsdk:"follow_redirects" json:"follow_redirects,computed_optional"`
	Header          map[string]*[]types.String `tfsdk:"header" json:"header,optional"`
	Method          types.String               `tfsdk:"method" json:"method,computed_optional"`
	Path            types.String               `tfsdk:"path" json:"path,computed_optional"`
	Port            types.Int64                `tfsdk:"port" json:"port,computed_optional"`
}

type HealthcheckTCPConfigModel struct {
	Method types.String `tfsdk:"method" json:"method,computed_optional"`
	Port   types.Int64  `tfsdk:"port" json:"port,computed_optional"`
}

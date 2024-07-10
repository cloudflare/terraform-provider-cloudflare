// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HealthcheckResultDataSourceEnvelope struct {
	Result HealthcheckDataSourceModel `json:"result,computed"`
}

type HealthcheckResultListDataSourceEnvelope struct {
	Result *[]*HealthcheckDataSourceModel `json:"result,computed"`
}

type HealthcheckDataSourceModel struct {
	ZoneID               types.String                          `tfsdk:"zone_id" path:"zone_id"`
	HealthcheckID        types.String                          `tfsdk:"healthcheck_id" path:"healthcheck_id"`
	ID                   types.String                          `tfsdk:"id" json:"id,computed"`
	Address              types.String                          `tfsdk:"address" json:"address"`
	CheckRegions         types.String                          `tfsdk:"check_regions" json:"check_regions"`
	ConsecutiveFails     types.Int64                           `tfsdk:"consecutive_fails" json:"consecutive_fails,computed"`
	ConsecutiveSuccesses types.Int64                           `tfsdk:"consecutive_successes" json:"consecutive_successes,computed"`
	CreatedOn            types.String                          `tfsdk:"created_on" json:"created_on,computed"`
	Description          types.String                          `tfsdk:"description" json:"description"`
	FailureReason        types.String                          `tfsdk:"failure_reason" json:"failure_reason,computed"`
	HTTPConfig           *HealthcheckHTTPConfigDataSourceModel `tfsdk:"http_config" json:"http_config"`
	Interval             types.Int64                           `tfsdk:"interval" json:"interval,computed"`
	ModifiedOn           types.String                          `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                 types.String                          `tfsdk:"name" json:"name"`
	Retries              types.Int64                           `tfsdk:"retries" json:"retries,computed"`
	Status               types.String                          `tfsdk:"status" json:"status,computed"`
	Suspended            types.Bool                            `tfsdk:"suspended" json:"suspended,computed"`
	TCPConfig            *HealthcheckTCPConfigDataSourceModel  `tfsdk:"tcp_config" json:"tcp_config"`
	Timeout              types.Int64                           `tfsdk:"timeout" json:"timeout,computed"`
	Type                 types.String                          `tfsdk:"type" json:"type,computed"`
	FindOneBy            *HealthcheckFindOneByDataSourceModel  `tfsdk:"find_one_by"`
}

type HealthcheckHTTPConfigDataSourceModel struct {
	AllowInsecure   types.Bool   `tfsdk:"allow_insecure" json:"allow_insecure,computed"`
	ExpectedBody    types.String `tfsdk:"expected_body" json:"expected_body"`
	ExpectedCodes   types.String `tfsdk:"expected_codes" json:"expected_codes"`
	FollowRedirects types.Bool   `tfsdk:"follow_redirects" json:"follow_redirects,computed"`
	Header          types.String `tfsdk:"header" json:"header"`
	Method          types.String `tfsdk:"method" json:"method,computed"`
	Path            types.String `tfsdk:"path" json:"path,computed"`
	Port            types.Int64  `tfsdk:"port" json:"port,computed"`
}

type HealthcheckTCPConfigDataSourceModel struct {
	Method types.String `tfsdk:"method" json:"method,computed"`
	Port   types.Int64  `tfsdk:"port" json:"port,computed"`
}

type HealthcheckFindOneByDataSourceModel struct {
	ZoneID  types.String `tfsdk:"zone_id" path:"zone_id"`
	Page    types.String `tfsdk:"page" query:"page"`
	PerPage types.String `tfsdk:"per_page" query:"per_page"`
}

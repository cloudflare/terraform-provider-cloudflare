// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HealthchecksResultListDataSourceEnvelope struct {
	Result *[]*HealthchecksItemsDataSourceModel `json:"result,computed"`
}

type HealthchecksDataSourceModel struct {
	ZoneID   types.String                         `tfsdk:"zone_id" path:"zone_id"`
	Page     types.String                         `tfsdk:"page" query:"page"`
	PerPage  types.String                         `tfsdk:"per_page" query:"per_page"`
	MaxItems types.Int64                          `tfsdk:"max_items"`
	Items    *[]*HealthchecksItemsDataSourceModel `tfsdk:"items"`
}

type HealthchecksItemsDataSourceModel struct {
	ID                   types.String    `tfsdk:"id" json:"id,computed"`
	Address              types.String    `tfsdk:"address" json:"address,computed"`
	CheckRegions         *[]types.String `tfsdk:"check_regions" json:"check_regions,computed"`
	ConsecutiveFails     types.Int64     `tfsdk:"consecutive_fails" json:"consecutive_fails,computed"`
	ConsecutiveSuccesses types.Int64     `tfsdk:"consecutive_successes" json:"consecutive_successes,computed"`
	CreatedOn            types.String    `tfsdk:"created_on" json:"created_on,computed"`
	Description          types.String    `tfsdk:"description" json:"description,computed"`
	FailureReason        types.String    `tfsdk:"failure_reason" json:"failure_reason,computed"`
	Interval             types.Int64     `tfsdk:"interval" json:"interval,computed"`
	ModifiedOn           types.String    `tfsdk:"modified_on" json:"modified_on,computed"`
	Name                 types.String    `tfsdk:"name" json:"name,computed"`
	Retries              types.Int64     `tfsdk:"retries" json:"retries,computed"`
	Status               types.String    `tfsdk:"status" json:"status,computed"`
	Suspended            types.Bool      `tfsdk:"suspended" json:"suspended,computed"`
	Timeout              types.Int64     `tfsdk:"timeout" json:"timeout,computed"`
	Type                 types.String    `tfsdk:"type" json:"type,computed"`
}

type HealthchecksItemsHTTPConfigDataSourceModel struct {
	AllowInsecure   types.Bool                 `tfsdk:"allow_insecure" json:"allow_insecure,computed"`
	ExpectedBody    types.String               `tfsdk:"expected_body" json:"expected_body,computed"`
	ExpectedCodes   *[]types.String            `tfsdk:"expected_codes" json:"expected_codes,computed"`
	FollowRedirects types.Bool                 `tfsdk:"follow_redirects" json:"follow_redirects,computed"`
	Header          map[string]*[]types.String `tfsdk:"header" json:"header,computed"`
	Method          types.String               `tfsdk:"method" json:"method,computed"`
	Path            types.String               `tfsdk:"path" json:"path,computed"`
	Port            types.Int64                `tfsdk:"port" json:"port,computed"`
}

type HealthchecksItemsTCPConfigDataSourceModel struct {
	Method types.String `tfsdk:"method" json:"method,computed"`
	Port   types.Int64  `tfsdk:"port" json:"port,computed"`
}

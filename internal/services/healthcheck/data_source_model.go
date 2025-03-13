// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/healthchecks"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type HealthcheckResultDataSourceEnvelope struct {
Result HealthcheckDataSourceModel `json:"result,computed"`
}

type HealthcheckDataSourceModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
HealthcheckID types.String `tfsdk:"healthcheck_id" path:"healthcheck_id,optional"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Address types.String `tfsdk:"address" json:"address,computed"`
ConsecutiveFails types.Int64 `tfsdk:"consecutive_fails" json:"consecutive_fails,computed"`
ConsecutiveSuccesses types.Int64 `tfsdk:"consecutive_successes" json:"consecutive_successes,computed"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
Description types.String `tfsdk:"description" json:"description,computed"`
FailureReason types.String `tfsdk:"failure_reason" json:"failure_reason,computed"`
Interval types.Int64 `tfsdk:"interval" json:"interval,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
Retries types.Int64 `tfsdk:"retries" json:"retries,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
Suspended types.Bool `tfsdk:"suspended" json:"suspended,computed"`
Timeout types.Int64 `tfsdk:"timeout" json:"timeout,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
CheckRegions customfield.List[types.String] `tfsdk:"check_regions" json:"check_regions,computed"`
HTTPConfig customfield.NestedObject[HealthcheckHTTPConfigDataSourceModel] `tfsdk:"http_config" json:"http_config,computed"`
TCPConfig customfield.NestedObject[HealthcheckTCPConfigDataSourceModel] `tfsdk:"tcp_config" json:"tcp_config,computed"`
}

func (m *HealthcheckDataSourceModel) toReadParams(_ context.Context) (params healthchecks.HealthcheckGetParams, diags diag.Diagnostics) {
  params = healthchecks.HealthcheckGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

func (m *HealthcheckDataSourceModel) toListParams(_ context.Context) (params healthchecks.HealthcheckListParams, diags diag.Diagnostics) {
  params = healthchecks.HealthcheckListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

type HealthcheckHTTPConfigDataSourceModel struct {
AllowInsecure types.Bool `tfsdk:"allow_insecure" json:"allow_insecure,computed"`
ExpectedBody types.String `tfsdk:"expected_body" json:"expected_body,computed"`
ExpectedCodes customfield.List[types.String] `tfsdk:"expected_codes" json:"expected_codes,computed"`
FollowRedirects types.Bool `tfsdk:"follow_redirects" json:"follow_redirects,computed"`
Header customfield.Map[customfield.List[types.String]] `tfsdk:"header" json:"header,computed"`
Method types.String `tfsdk:"method" json:"method,computed"`
Path types.String `tfsdk:"path" json:"path,computed"`
Port types.Int64 `tfsdk:"port" json:"port,computed"`
}

type HealthcheckTCPConfigDataSourceModel struct {
Method types.String `tfsdk:"method" json:"method,computed"`
Port types.Int64 `tfsdk:"port" json:"port,computed"`
}

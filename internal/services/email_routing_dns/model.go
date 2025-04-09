// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingDNSResultEnvelope struct {
Result EmailRoutingDNSModel `json:"result"`
}

type EmailRoutingDNSModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Name types.String `tfsdk:"name" json:"name,required"`
Created timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
Modified timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
SkipWizard types.Bool `tfsdk:"skip_wizard" json:"skip_wizard,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
Success types.Bool `tfsdk:"success" json:"success,computed"`
Tag types.String `tfsdk:"tag" json:"tag,computed"`
Errors customfield.NestedObjectList[EmailRoutingDNSErrorsModel] `tfsdk:"errors" json:"errors,computed"`
Messages customfield.NestedObjectList[EmailRoutingDNSMessagesModel] `tfsdk:"messages" json:"messages,computed"`
Result customfield.NestedObject[EmailRoutingDNSResultModel] `tfsdk:"result" json:"result,computed"`
ResultInfo customfield.NestedObject[EmailRoutingDNSResultInfoModel] `tfsdk:"result_info" json:"result_info,computed"`
}

func (m EmailRoutingDNSModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m EmailRoutingDNSModel) MarshalJSONForUpdate(state EmailRoutingDNSModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}

type EmailRoutingDNSErrorsModel struct {
Code types.Int64 `tfsdk:"code" json:"code,computed"`
Message types.String `tfsdk:"message" json:"message,computed"`
}

type EmailRoutingDNSMessagesModel struct {
Code types.Int64 `tfsdk:"code" json:"code,computed"`
Message types.String `tfsdk:"message" json:"message,computed"`
}

type EmailRoutingDNSResultModel struct {
Errors customfield.NestedObjectList[EmailRoutingDNSResultErrorsModel] `tfsdk:"errors" json:"errors,computed"`
Record customfield.NestedObjectList[EmailRoutingDNSResultRecordModel] `tfsdk:"record" json:"record,computed"`
Content types.String `tfsdk:"content" json:"content,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Priority types.Float64 `tfsdk:"priority" json:"priority,computed"`
TTL types.Float64 `tfsdk:"ttl" json:"ttl,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type EmailRoutingDNSResultErrorsModel struct {
Code types.String `tfsdk:"code" json:"code,computed"`
Missing customfield.NestedObject[EmailRoutingDNSResultErrorsMissingModel] `tfsdk:"missing" json:"missing,computed"`
}

type EmailRoutingDNSResultErrorsMissingModel struct {
Content types.String `tfsdk:"content" json:"content,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Priority types.Float64 `tfsdk:"priority" json:"priority,computed"`
TTL types.Float64 `tfsdk:"ttl" json:"ttl,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type EmailRoutingDNSResultRecordModel struct {
Content types.String `tfsdk:"content" json:"content,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Priority types.Float64 `tfsdk:"priority" json:"priority,computed"`
TTL types.Float64 `tfsdk:"ttl" json:"ttl,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type EmailRoutingDNSResultInfoModel struct {
Count types.Float64 `tfsdk:"count" json:"count,computed"`
Page types.Float64 `tfsdk:"page" json:"page,computed"`
PerPage types.Float64 `tfsdk:"per_page" json:"per_page,computed"`
TotalCount types.Float64 `tfsdk:"total_count" json:"total_count,computed"`
}

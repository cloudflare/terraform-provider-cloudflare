// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingDNSDataSourceModel struct {
	ZoneID     types.String                                                         `tfsdk:"zone_id" path:"zone_id,required"`
	Subdomain  types.String                                                         `tfsdk:"subdomain" query:"subdomain,optional"`
	Success    types.Bool                                                           `tfsdk:"success" json:"success,computed"`
	Errors     customfield.NestedObjectList[EmailRoutingDNSErrorsDataSourceModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages   customfield.NestedObjectList[EmailRoutingDNSMessagesDataSourceModel] `tfsdk:"messages" json:"messages,computed"`
	Result     customfield.NestedObject[EmailRoutingDNSResultDataSourceModel]       `tfsdk:"result" json:"result,computed"`
	ResultInfo customfield.NestedObject[EmailRoutingDNSResultInfoDataSourceModel]   `tfsdk:"result_info" json:"result_info,computed"`
}

func (m *EmailRoutingDNSDataSourceModel) toReadParams(_ context.Context) (params email_routing.DNSGetParams, diags diag.Diagnostics) {
	params = email_routing.DNSGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Subdomain.IsNull() {
		params.Subdomain = cloudflare.F(m.Subdomain.ValueString())
	}

	return
}

type EmailRoutingDNSErrorsDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type EmailRoutingDNSMessagesDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type EmailRoutingDNSResultDataSourceModel struct {
	Errors   customfield.NestedObjectList[EmailRoutingDNSResultErrorsDataSourceModel] `tfsdk:"errors" json:"errors,computed"`
	Record   customfield.NestedObjectList[EmailRoutingDNSResultRecordDataSourceModel] `tfsdk:"record" json:"record,computed"`
	Content  types.String                                                             `tfsdk:"content" json:"content,computed"`
	Name     types.String                                                             `tfsdk:"name" json:"name,computed"`
	Priority types.Float64                                                            `tfsdk:"priority" json:"priority,computed"`
	TTL      types.Float64                                                            `tfsdk:"ttl" json:"ttl,computed"`
	Type     types.String                                                             `tfsdk:"type" json:"type,computed"`
}

type EmailRoutingDNSResultErrorsDataSourceModel struct {
	Code    types.String                                                                `tfsdk:"code" json:"code,computed"`
	Missing customfield.NestedObject[EmailRoutingDNSResultErrorsMissingDataSourceModel] `tfsdk:"missing" json:"missing,computed"`
}

type EmailRoutingDNSResultErrorsMissingDataSourceModel struct {
	Content  types.String  `tfsdk:"content" json:"content,computed"`
	Name     types.String  `tfsdk:"name" json:"name,computed"`
	Priority types.Float64 `tfsdk:"priority" json:"priority,computed"`
	TTL      types.Float64 `tfsdk:"ttl" json:"ttl,computed"`
	Type     types.String  `tfsdk:"type" json:"type,computed"`
}

type EmailRoutingDNSResultRecordDataSourceModel struct {
	Content  types.String  `tfsdk:"content" json:"content,computed"`
	Name     types.String  `tfsdk:"name" json:"name,computed"`
	Priority types.Float64 `tfsdk:"priority" json:"priority,computed"`
	TTL      types.Float64 `tfsdk:"ttl" json:"ttl,computed"`
	Type     types.String  `tfsdk:"type" json:"type,computed"`
}

type EmailRoutingDNSResultInfoDataSourceModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count,computed"`
	Page       types.Float64 `tfsdk:"page" json:"page,computed"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page,computed"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count,computed"`
}

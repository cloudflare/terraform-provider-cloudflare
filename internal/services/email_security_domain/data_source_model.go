// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_domain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityDomainResultDataSourceEnvelope struct {
	Result EmailSecurityDomainDataSourceModel `json:"result,computed"`
}

type EmailSecurityDomainDataSourceModel struct {
	ID                   types.String                                                                `tfsdk:"id" path:"domain_id,computed"`
	DomainID             types.String                                                                `tfsdk:"domain_id" path:"domain_id,optional"`
	AccountID            types.String                                                                `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt            timetypes.RFC3339                                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DMARCStatus          types.String                                                                `tfsdk:"dmarc_status" json:"dmarc_status,computed"`
	Domain               types.String                                                                `tfsdk:"domain" json:"domain,computed"`
	Folder               types.String                                                                `tfsdk:"folder" json:"folder,computed"`
	InboxProvider        types.String                                                                `tfsdk:"inbox_provider" json:"inbox_provider,computed"`
	IntegrationID        types.String                                                                `tfsdk:"integration_id" json:"integration_id,computed"`
	LastModified         timetypes.RFC3339                                                           `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	LookbackHops         types.Int64                                                                 `tfsdk:"lookback_hops" json:"lookback_hops,computed"`
	ModifiedAt           timetypes.RFC3339                                                           `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	O365TenantID         types.String                                                                `tfsdk:"o365_tenant_id" json:"o365_tenant_id,computed"`
	RequireTLSInbound    types.Bool                                                                  `tfsdk:"require_tls_inbound" json:"require_tls_inbound,computed"`
	RequireTLSOutbound   types.Bool                                                                  `tfsdk:"require_tls_outbound" json:"require_tls_outbound,computed"`
	SPFStatus            types.String                                                                `tfsdk:"spf_status" json:"spf_status,computed"`
	Status               types.String                                                                `tfsdk:"status" json:"status,computed"`
	Transport            types.String                                                                `tfsdk:"transport" json:"transport,computed"`
	AllowedDeliveryModes customfield.Set[types.String]                                               `tfsdk:"allowed_delivery_modes" json:"allowed_delivery_modes,computed"`
	DropDispositions     customfield.Set[types.String]                                               `tfsdk:"drop_dispositions" json:"drop_dispositions,computed"`
	IPRestrictions       customfield.Set[types.String]                                               `tfsdk:"ip_restrictions" json:"ip_restrictions,computed"`
	Regions              customfield.Set[types.String]                                               `tfsdk:"regions" json:"regions,computed"`
	Authorization        customfield.NestedObject[EmailSecurityDomainAuthorizationDataSourceModel]   `tfsdk:"authorization" json:"authorization,computed"`
	EmailsProcessed      customfield.NestedObject[EmailSecurityDomainEmailsProcessedDataSourceModel] `tfsdk:"emails_processed" json:"emails_processed,computed"`
	Filter               *EmailSecurityDomainFindOneByDataSourceModel                                `tfsdk:"filter"`
}

func (m *EmailSecurityDomainDataSourceModel) toReadParams(_ context.Context) (params email_security.SettingDomainGetParams, diags diag.Diagnostics) {
	params = email_security.SettingDomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *EmailSecurityDomainDataSourceModel) toListParams(_ context.Context) (params email_security.SettingDomainListParams, diags diag.Diagnostics) {
	mFilterDomain := []string{}
	if m.Filter.Domain != nil {
		for _, item := range *m.Filter.Domain {
			mFilterDomain = append(mFilterDomain, item.ValueString())
		}
	}

	params = email_security.SettingDomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
		Domain:    cloudflare.F(mFilterDomain),
	}

	if !m.Filter.ActiveDeliveryMode.IsNull() {
		params.ActiveDeliveryMode = cloudflare.F(email_security.SettingDomainListParamsActiveDeliveryMode(m.Filter.ActiveDeliveryMode.ValueString()))
	}
	if !m.Filter.AllowedDeliveryMode.IsNull() {
		params.AllowedDeliveryMode = cloudflare.F(email_security.SettingDomainListParamsAllowedDeliveryMode(m.Filter.AllowedDeliveryMode.ValueString()))
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingDomainListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.IntegrationID.IsNull() {
		params.IntegrationID = cloudflare.F(m.Filter.IntegrationID.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingDomainListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(email_security.SettingDomainListParamsStatus(m.Filter.Status.ValueString()))
	}

	return
}

type EmailSecurityDomainAuthorizationDataSourceModel struct {
	Authorized    types.Bool        `tfsdk:"authorized" json:"authorized,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	StatusMessage types.String      `tfsdk:"status_message" json:"status_message,computed"`
}

type EmailSecurityDomainEmailsProcessedDataSourceModel struct {
	Timestamp                    timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	TotalEmailsProcessed         types.Int64       `tfsdk:"total_emails_processed" json:"total_emails_processed,computed"`
	TotalEmailsProcessedPrevious types.Int64       `tfsdk:"total_emails_processed_previous" json:"total_emails_processed_previous,computed"`
}

type EmailSecurityDomainFindOneByDataSourceModel struct {
	ActiveDeliveryMode  types.String    `tfsdk:"active_delivery_mode" query:"active_delivery_mode,optional"`
	AllowedDeliveryMode types.String    `tfsdk:"allowed_delivery_mode" query:"allowed_delivery_mode,optional"`
	Direction           types.String    `tfsdk:"direction" query:"direction,optional"`
	Domain              *[]types.String `tfsdk:"domain" query:"domain,optional"`
	IntegrationID       types.String    `tfsdk:"integration_id" query:"integration_id,optional"`
	Order               types.String    `tfsdk:"order" query:"order,optional"`
	Search              types.String    `tfsdk:"search" query:"search,optional"`
	Status              types.String    `tfsdk:"status" query:"status,optional"`
}

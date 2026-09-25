// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_domain

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityDomainResultEnvelope struct {
	Result EmailSecurityDomainModel `json:"result"`
}

type EmailSecurityDomainModel struct {
	ID                   types.String                                                      `tfsdk:"id" json:"id,computed"`
	AccountID            types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	Domain               types.String                                                      `tfsdk:"domain" json:"domain,required"`
	AllowedDeliveryModes *[]types.String                                                   `tfsdk:"allowed_delivery_modes" json:"allowed_delivery_modes,required"`
	DropDispositions     *[]types.String                                                   `tfsdk:"drop_dispositions" json:"drop_dispositions,required"`
	IPRestrictions       *[]types.String                                                   `tfsdk:"ip_restrictions" json:"ip_restrictions,required"`
	Regions              *[]types.String                                                   `tfsdk:"regions" json:"regions,required"`
	IntegrationID        types.String                                                      `tfsdk:"integration_id" json:"integration_id,optional"`
	Transport            types.String                                                      `tfsdk:"transport" json:"transport,optional"`
	Folder               types.String                                                      `tfsdk:"folder" json:"folder,computed_optional"`
	LookbackHops         types.Int64                                                       `tfsdk:"lookback_hops" json:"lookback_hops,computed_optional"`
	RequireTLSInbound    types.Bool                                                        `tfsdk:"require_tls_inbound" json:"require_tls_inbound,computed_optional"`
	RequireTLSOutbound   types.Bool                                                        `tfsdk:"require_tls_outbound" json:"require_tls_outbound,computed_optional"`
	CreatedAt            timetypes.RFC3339                                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DMARCStatus          types.String                                                      `tfsdk:"dmarc_status" json:"dmarc_status,computed"`
	InboxProvider        types.String                                                      `tfsdk:"inbox_provider" json:"inbox_provider,computed"`
	LastModified         timetypes.RFC3339                                                 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	ModifiedAt           timetypes.RFC3339                                                 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	O365TenantID         types.String                                                      `tfsdk:"o365_tenant_id" json:"o365_tenant_id,computed"`
	SPFStatus            types.String                                                      `tfsdk:"spf_status" json:"spf_status,computed"`
	Status               types.String                                                      `tfsdk:"status" json:"status,computed"`
	Authorization        customfield.NestedObject[EmailSecurityDomainAuthorizationModel]   `tfsdk:"authorization" json:"authorization,computed"`
	EmailsProcessed      customfield.NestedObject[EmailSecurityDomainEmailsProcessedModel] `tfsdk:"emails_processed" json:"emails_processed,computed"`
}

func (m EmailSecurityDomainModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailSecurityDomainModel) MarshalJSONForUpdate(state EmailSecurityDomainModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type EmailSecurityDomainAuthorizationModel struct {
	Authorized    types.Bool        `tfsdk:"authorized" json:"authorized,computed"`
	Timestamp     timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	StatusMessage types.String      `tfsdk:"status_message" json:"status_message,computed"`
}

type EmailSecurityDomainEmailsProcessedModel struct {
	Timestamp                    timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
	TotalEmailsProcessed         types.Int64       `tfsdk:"total_emails_processed" json:"total_emails_processed,computed"`
	TotalEmailsProcessedPrevious types.Int64       `tfsdk:"total_emails_processed_previous" json:"total_emails_processed_previous,computed"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/alerting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyWebhooksResultDataSourceEnvelope struct {
	Result NotificationPolicyWebhooksDataSourceModel `json:"result,computed"`
}

type NotificationPolicyWebhooksResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[NotificationPolicyWebhooksDataSourceModel] `json:"result,computed"`
}

type NotificationPolicyWebhooksDataSourceModel struct {
	AccountID   types.String                                        `tfsdk:"account_id" path:"account_id"`
	WebhookID   types.String                                        `tfsdk:"webhook_id" path:"webhook_id"`
	CreatedAt   timetypes.RFC3339                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID          types.String                                        `tfsdk:"id" json:"id,computed"`
	LastFailure timetypes.RFC3339                                   `tfsdk:"last_failure" json:"last_failure,computed" format:"date-time"`
	LastSuccess timetypes.RFC3339                                   `tfsdk:"last_success" json:"last_success,computed" format:"date-time"`
	Name        types.String                                        `tfsdk:"name" json:"name,computed_optional"`
	Secret      types.String                                        `tfsdk:"secret" json:"secret,computed_optional"`
	Type        types.String                                        `tfsdk:"type" json:"type,computed_optional"`
	URL         types.String                                        `tfsdk:"url" json:"url,computed_optional"`
	Filter      *NotificationPolicyWebhooksFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *NotificationPolicyWebhooksDataSourceModel) toReadParams() (params alerting.DestinationWebhookGetParams, diags diag.Diagnostics) {
	params = alerting.DestinationWebhookGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *NotificationPolicyWebhooksDataSourceModel) toListParams() (params alerting.DestinationWebhookListParams, diags diag.Diagnostics) {
	params = alerting.DestinationWebhookListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type NotificationPolicyWebhooksFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}

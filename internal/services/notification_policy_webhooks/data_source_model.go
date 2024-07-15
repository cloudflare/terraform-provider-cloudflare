// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyWebhooksResultDataSourceEnvelope struct {
	Result NotificationPolicyWebhooksDataSourceModel `json:"result,computed"`
}

type NotificationPolicyWebhooksResultListDataSourceEnvelope struct {
	Result *[]*NotificationPolicyWebhooksDataSourceModel `json:"result,computed"`
}

type NotificationPolicyWebhooksDataSourceModel struct {
	AccountID   types.String                                        `tfsdk:"account_id" path:"account_id"`
	WebhookID   types.String                                        `tfsdk:"webhook_id" path:"webhook_id"`
	ID          types.String                                        `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339                                   `tfsdk:"created_at" json:"created_at,computed"`
	LastFailure timetypes.RFC3339                                   `tfsdk:"last_failure" json:"last_failure,computed"`
	LastSuccess timetypes.RFC3339                                   `tfsdk:"last_success" json:"last_success,computed"`
	Name        types.String                                        `tfsdk:"name" json:"name"`
	Secret      types.String                                        `tfsdk:"secret" json:"secret"`
	Type        types.String                                        `tfsdk:"type" json:"type"`
	URL         types.String                                        `tfsdk:"url" json:"url"`
	FindOneBy   *NotificationPolicyWebhooksFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type NotificationPolicyWebhooksFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}

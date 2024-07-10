// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyWebhooksResultEnvelope struct {
	Result NotificationPolicyWebhooksModel `json:"result,computed"`
}

type NotificationPolicyWebhooksModel struct {
	ID          types.String                                `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                                `tfsdk:"account_id" path:"account_id"`
	Name        types.String                                `tfsdk:"name" json:"name"`
	URL         types.String                                `tfsdk:"url" json:"url"`
	Secret      types.String                                `tfsdk:"secret" json:"secret"`
	Errors      *[]*NotificationPolicyWebhooksErrorsModel   `tfsdk:"errors" json:"errors,computed"`
	Messages    *[]*NotificationPolicyWebhooksMessagesModel `tfsdk:"messages" json:"messages,computed"`
	Success     types.Bool                                  `tfsdk:"success" json:"success,computed"`
	CreatedAt   types.String                                `tfsdk:"created_at" json:"created_at,computed"`
	LastFailure types.String                                `tfsdk:"last_failure" json:"last_failure,computed"`
	LastSuccess types.String                                `tfsdk:"last_success" json:"last_success,computed"`
	Type        types.String                                `tfsdk:"type" json:"type,computed"`
}

type NotificationPolicyWebhooksErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type NotificationPolicyWebhooksMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

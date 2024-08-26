// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyWebhooksResultEnvelope struct {
	Result NotificationPolicyWebhooksModel `json:"result"`
}

type NotificationPolicyWebhooksModel struct {
	ID          types.String                                                          `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                                                          `tfsdk:"account_id" path:"account_id"`
	Name        types.String                                                          `tfsdk:"name" json:"name"`
	URL         types.String                                                          `tfsdk:"url" json:"url"`
	Secret      types.String                                                          `tfsdk:"secret" json:"secret"`
	CreatedAt   timetypes.RFC3339                                                     `tfsdk:"created_at" json:"created_at,computed"`
	LastFailure timetypes.RFC3339                                                     `tfsdk:"last_failure" json:"last_failure,computed"`
	LastSuccess timetypes.RFC3339                                                     `tfsdk:"last_success" json:"last_success,computed"`
	Success     types.Bool                                                            `tfsdk:"success" json:"success,computed"`
	Type        types.String                                                          `tfsdk:"type" json:"type,computed"`
	Errors      customfield.NestedObjectList[NotificationPolicyWebhooksErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages    customfield.NestedObjectList[NotificationPolicyWebhooksMessagesModel] `tfsdk:"messages" json:"messages,computed"`
	ResultInfo  customfield.NestedObject[NotificationPolicyWebhooksResultInfoModel]   `tfsdk:"result_info" json:"result_info,computed"`
}

type NotificationPolicyWebhooksErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type NotificationPolicyWebhooksMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type NotificationPolicyWebhooksResultInfoModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count"`
	Page       types.Float64 `tfsdk:"page" json:"page"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count"`
}

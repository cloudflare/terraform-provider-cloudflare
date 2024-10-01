// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyWebhooksResultEnvelope struct {
	Result NotificationPolicyWebhooksModel `json:"result"`
}

type NotificationPolicyWebhooksModel struct {
	ID          types.String                                                          `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	Name        types.String                                                          `tfsdk:"name" json:"name,required"`
	URL         types.String                                                          `tfsdk:"url" json:"url,required"`
	Secret      types.String                                                          `tfsdk:"secret" json:"secret,optional"`
	CreatedAt   timetypes.RFC3339                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastFailure timetypes.RFC3339                                                     `tfsdk:"last_failure" json:"last_failure,computed" format:"date-time"`
	LastSuccess timetypes.RFC3339                                                     `tfsdk:"last_success" json:"last_success,computed" format:"date-time"`
	Success     types.Bool                                                            `tfsdk:"success" json:"success,computed"`
	Type        types.String                                                          `tfsdk:"type" json:"type,computed"`
	Errors      customfield.NestedObjectList[NotificationPolicyWebhooksErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages    customfield.NestedObjectList[NotificationPolicyWebhooksMessagesModel] `tfsdk:"messages" json:"messages,computed"`
	ResultInfo  customfield.NestedObject[NotificationPolicyWebhooksResultInfoModel]   `tfsdk:"result_info" json:"result_info,computed"`
}

func (m NotificationPolicyWebhooksModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m NotificationPolicyWebhooksModel) MarshalJSONForUpdate(state NotificationPolicyWebhooksModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type NotificationPolicyWebhooksErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type NotificationPolicyWebhooksMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type NotificationPolicyWebhooksResultInfoModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count,computed"`
	Page       types.Float64 `tfsdk:"page" json:"page,computed"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page,computed"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count,computed"`
}

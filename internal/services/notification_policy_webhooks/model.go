// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyWebhooksResultEnvelope struct {
Result NotificationPolicyWebhooksModel `json:"result"`
}

type NotificationPolicyWebhooksModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Name types.String `tfsdk:"name" json:"name,required"`
URL types.String `tfsdk:"url" json:"url,required"`
Secret types.String `tfsdk:"secret" json:"secret,optional"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
LastFailure timetypes.RFC3339 `tfsdk:"last_failure" json:"last_failure,computed" format:"date-time"`
LastSuccess timetypes.RFC3339 `tfsdk:"last_success" json:"last_success,computed" format:"date-time"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

func (m NotificationPolicyWebhooksModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m NotificationPolicyWebhooksModel) MarshalJSONForUpdate(state NotificationPolicyWebhooksModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

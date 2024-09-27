// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_webhook

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamWebhookResultEnvelope struct {
	Result StreamWebhookModel `json:"result"`
}

type StreamWebhookModel struct {
	AccountID       types.String `tfsdk:"account_id" path:"account_id,required"`
	NotificationURL types.String `tfsdk:"notification_url" json:"notificationUrl,required"`
}

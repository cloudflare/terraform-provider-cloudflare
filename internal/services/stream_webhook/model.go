// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_webhook

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamWebhookResultEnvelope struct {
	Result StreamWebhookModel `json:"result"`
}

type StreamWebhookModel struct {
	AccountID       types.String      `tfsdk:"account_id" path:"account_id,required"`
	NotificationURL types.String      `tfsdk:"notification_url" json:"notificationUrl,optional"`
	Modified        timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Secret          types.String      `tfsdk:"secret" json:"secret,computed"`
}

func (m StreamWebhookModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StreamWebhookModel) MarshalJSONForUpdate(state StreamWebhookModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

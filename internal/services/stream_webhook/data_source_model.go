// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_webhook

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/stream"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamWebhookResultDataSourceEnvelope struct {
	Result StreamWebhookDataSourceModel `json:"result,computed"`
}

type StreamWebhookDataSourceModel struct {
	AccountID       types.String      `tfsdk:"account_id" path:"account_id,optional"`
	Modified        timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	NotificationURL types.String      `tfsdk:"notification_url" json:"notificationUrl,computed"`
	Secret          types.String      `tfsdk:"secret" json:"secret,computed"`
}

func (m *StreamWebhookDataSourceModel) toReadParams(_ context.Context) (params stream.WebhookGetParams, diags diag.Diagnostics) {
	params = stream.WebhookGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	}

	return
}

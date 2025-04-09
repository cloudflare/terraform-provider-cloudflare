// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/alerting"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NotificationPolicyWebhooksResultDataSourceEnvelope struct {
	Result NotificationPolicyWebhooksDataSourceModel `json:"result,computed"`
}

type NotificationPolicyWebhooksDataSourceModel struct {
	ID          types.String      `tfsdk:"id" path:"webhook_id,computed"`
	WebhookID   types.String      `tfsdk:"webhook_id" path:"webhook_id,optional"`
	AccountID   types.String      `tfsdk:"account_id" path:"account_id,required"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastFailure timetypes.RFC3339 `tfsdk:"last_failure" json:"last_failure,computed" format:"date-time"`
	LastSuccess timetypes.RFC3339 `tfsdk:"last_success" json:"last_success,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	Secret      types.String      `tfsdk:"secret" json:"secret,computed"`
	Type        types.String      `tfsdk:"type" json:"type,computed"`
	URL         types.String      `tfsdk:"url" json:"url,computed"`
}

func (m *NotificationPolicyWebhooksDataSourceModel) toReadParams(_ context.Context) (params alerting.DestinationWebhookGetParams, diags diag.Diagnostics) {
	params = alerting.DestinationWebhookGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

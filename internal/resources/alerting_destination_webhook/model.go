// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting_destination_webhook

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AlertingDestinationWebhookResultEnvelope struct {
	Result AlertingDestinationWebhookModel `json:"result,computed"`
}

type AlertingDestinationWebhookModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	Name      types.String `tfsdk:"name" json:"name"`
	URL       types.String `tfsdk:"url" json:"url"`
	Secret    types.String `tfsdk:"secret" json:"secret"`
}

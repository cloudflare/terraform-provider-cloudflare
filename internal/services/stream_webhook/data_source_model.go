// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_webhook

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/stream"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamWebhookResultDataSourceEnvelope struct {
	Result StreamWebhookDataSourceModel `json:"result,computed"`
}

type StreamWebhookDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

func (m *StreamWebhookDataSourceModel) toReadParams(_ context.Context) (params stream.WebhookGetParams, diags diag.Diagnostics) {
	params = stream.WebhookGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/queues"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueConsumerResultDataSourceEnvelope struct {
	Result QueueConsumerDataSourceModel `json:"result,computed"`
}

type QueueConsumerDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	QueueID   types.String `tfsdk:"queue_id" path:"queue_id,required"`
}

func (m *QueueConsumerDataSourceModel) toReadParams(_ context.Context) (params queues.ConsumerGetParams, diags diag.Diagnostics) {
	params = queues.ConsumerGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

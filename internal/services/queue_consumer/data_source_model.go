// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/queues"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type QueueConsumerResultDataSourceEnvelope struct {
	Result QueueConsumerDataSourceModel `json:"result,computed"`
}

type QueueConsumerDataSourceModel struct {
	AccountID  types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	QueueID    types.String                                                   `tfsdk:"queue_id" path:"queue_id,required"`
	ConsumerID types.String                                                   `tfsdk:"consumer_id" json:"consumer_id,computed"`
	CreatedOn  types.String                                                   `tfsdk:"created_on" json:"created_on,computed"`
	Script     types.String                                                   `tfsdk:"script" json:"script,computed"`
	ScriptName types.String                                                   `tfsdk:"script_name" json:"script_name,computed"`
	Type       types.String                                                   `tfsdk:"type" json:"type,computed"`
	Settings   customfield.NestedObject[QueueConsumerSettingsDataSourceModel] `tfsdk:"settings" json:"settings,computed"`
}

func (m *QueueConsumerDataSourceModel) toReadParams(_ context.Context) (params queues.ConsumerGetParams, diags diag.Diagnostics) {
	params = queues.ConsumerGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type QueueConsumerSettingsDataSourceModel struct {
	BatchSize           types.Float64 `tfsdk:"batch_size" json:"batch_size,computed"`
	MaxConcurrency      types.Float64 `tfsdk:"max_concurrency" json:"max_concurrency,computed"`
	MaxRetries          types.Float64 `tfsdk:"max_retries" json:"max_retries,computed"`
	MaxWaitTimeMs       types.Float64 `tfsdk:"max_wait_time_ms" json:"max_wait_time_ms,computed"`
	RetryDelay          types.Float64 `tfsdk:"retry_delay" json:"retry_delay,computed"`
	VisibilityTimeoutMs types.Float64 `tfsdk:"visibility_timeout_ms" json:"visibility_timeout_ms,computed"`
}

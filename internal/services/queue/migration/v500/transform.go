package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a source cloudflare_queue state (v4) to target cloudflare_queue state (v500).
func Transform(ctx context.Context, source SourceCloudflareQueueModel) (*TargetQueueModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError("Missing required field", "account_id is required for queue migration")
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError("Missing required field", "name is required for queue migration")
		return nil, diags
	}

	target := &TargetQueueModel{
		// Direct copies
		ID:        source.ID,
		AccountID: source.AccountID,

		// Rename: name → queue_name
		QueueName: source.Name,

		// Copy id to the new queue_id field (v5 requires both)
		QueueID: source.ID,

		// Computed fields: initialize as null, API will populate on next plan/apply
		ConsumersTotalCount: types.Float64Null(),
		ProducersTotalCount: types.Float64Null(),
		CreatedOn:           types.StringNull(),
		ModifiedOn:          types.StringNull(),
		Settings:            customfield.NullObject[TargetQueueSettingsModel](ctx),
		Consumers:           customfield.NullObjectList[TargetQueueConsumersModel](ctx),
		Producers:           customfield.NullObjectList[TargetQueueProducersModel](ctx),
	}

	return target, diags
}

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema_version=0 to schema_version=500.
// This is a no-op upgrade: cloudflare_queue_consumer is a v5-native resource with no v4
// predecessor. States at version=0 were created before explicit schema versioning was
// applied to this resource. The schema is fully compatible — just copy state through.
func UpgradeFromV0(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading queue_consumer state from schema_version=0")
	// No-op: schema is compatible, just copy raw state through.
	// Using raw state avoids issues with custom field type serialization.
	resp.State.Raw = req.State.Raw
}

// UpgradeFromV1 handles state upgrades from schema_version=1 to schema_version=500.
// This is a no-op upgrade: version 1 is the dormant production version of the v5 schema.
// No schema changes occurred between version 1 and version 500.
func UpgradeFromV1(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading queue_consumer state from schema_version=1")
	// No-op: schema is compatible, just copy raw state through.
	resp.State.Raw = req.State.Raw
}

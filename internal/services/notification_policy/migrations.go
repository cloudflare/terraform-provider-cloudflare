// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/notification_policy/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*NotificationPolicyResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) -> v5 (version=500): Full transformation
//   - filters: MaxItems:1 list -> SingleNestedAttribute object
//   - Three integration Sets -> single mechanisms nested object
//   - Integration items: drop "name" field, keep only "id"
//   - Filter fields: Set -> List conversion (~35 fields)
//
// 2. v5 state (version=1) -> v5 (version=500): No-op upgrade
//
// The separation of schema versions (v4=0, v5=1/500) eliminates the need for
// dual-format detection that was required in earlier implementations.
func (r *NotificationPolicyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// PriorSchema: nil — v4 SDKv2 state encoding is incompatible with
		// Plugin Framework schema types. The upgrader reads raw JSON directly.
		0: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider with version=1
		// This is a no-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}

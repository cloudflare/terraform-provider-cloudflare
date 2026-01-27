// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_rule/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*PageRuleResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
func (r *PageRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflarePageRuleSchema()
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handles schema_version=0 from BOTH v4 and early v5:
		// - v4 SDKv2 (no SchemaVersion set, defaults to 0): actions as array
		// - Early v5 Plugin Framework (no Version set, defaults to 0): actions as object
		// PriorSchema is nil to bypass Terraform's validation - handler detects format and parses manually
		0: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Migration: legacy v4.x SDKv2 provider with schema_version=3 (if explicitly set)
		3: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV3,
		},
		// No-op: released version bump (v5.0.0-v5.16.0 releases that set Version: 4)
		4: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
	}
}

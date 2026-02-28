// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"
	"os"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*RulesetResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=1) -> v5 (version=500): Full transformation
//   - MaxItems:1 ListNestedBlock arrays -> SingleNestedAttribute objects
//   - headers: List[{name,...}] -> Map[name -> {...}]
//   - cookie_fields/request_fields/response_fields: Set[string] -> List[{name:string}]
//   - query_string include/exclude: Set[string] -> {list:[...]} or {all:true}
//   - products/phases/rulesets: TypeSet -> ListAttribute
//   - disable_railgun: removed in v5
//   - rules map: map[string]string (CSV) -> map[string][]string
//
// 2. v5 state (version=1) -> v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
//   - Just bumps version number, no transformation
//
// The v4 Plugin Framework provider used schema_version=1 (after its internal V0->V1
// migration for ratelimit field rename). Both v4 and v5 state are at version=1.
// GetSchemaVersion(1, 500) ensures controlled rollout via TF_MIG_TEST.
func (r *RulesetResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v5Schema := ResourceSchema(ctx)

	if os.Getenv("TF_MIG_TEST") == "" {
		// Production mode: preserve existing upgraders only
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &v5Schema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	// Test mode (TF_MIG_TEST=1): full StateUpgrader migration
	v4Schema := v500.SourceV4RulesetSchema()

	return map[int64]resource.StateUpgrader{
		// Handle fresh v5 resources (version 0 -> 500)
		0: {
			PriorSchema:   &v5Schema,
			StateUpgrader: v500.UpgradeFromV5,
		},

		// Handle state from v4 Plugin Framework provider (schema_version=1)
		1: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV4,
		},
	}
}

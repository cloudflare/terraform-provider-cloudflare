// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_monitor/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*LoadBalancerMonitorResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) -> v5 (version=500): Full transformation
//    - Header field: TypeSet (nested) -> MapAttribute
//    - Default value additions for v5 fields
//    - Direct pass-through for compatible fields
//
// 2. v5 state (version=1) -> v5 (version=500): No-op upgrade
//    - Just bumps version number, no data transformation
//
// The separation of schema versions (v4=0, v5=1/500) eliminates the need for
// dual-format detection that was required in earlier implementations.
func (r *LoadBalancerMonitorResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	sourceSchema := v500.SourceLoadBalancerMonitorSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// This performs full transformation from v4 to v5 format
		0: {
			PriorSchema:   &sourceSchema,
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

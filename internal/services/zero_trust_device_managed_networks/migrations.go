// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_managed_networks/migration/v500"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustDeviceManagedNetworksResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustDeviceManagedNetworksResource)(nil)

// MoveState handles moves from cloudflare_device_managed_networks (v0) to cloudflare_zero_trust_device_managed_networks (v500).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_device_managed_networks.example
//	    to   = cloudflare_zero_trust_device_managed_networks.example
//	}
func (r *ZeroTrustDeviceManagedNetworksResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceCloudflareDeviceManagedNetworksSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState handles schema version upgrades for cloudflare_zero_trust_device_managed_networks.
// This is triggered when users manually run `terraform state mv` (Terraform < 1.8).
func (r *ZeroTrustDeviceManagedNetworksResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	// Use source schema for v0 - the legacy SDKv2 provider used config as a list block (array in JSON),
	// not a SingleNestedAttribute (object in JSON) like the v5 schema.
	sourceSchema := v500.SourceCloudflareDeviceManagedNetworksSchema()

	return map[int64]resource.StateUpgrader{
		// Handle state moved from legacy cloudflare_device_managed_networks (schema_version=0 from the SDKv2 provider)
		// When users run `terraform state mv cloudflare_device_managed_networks.x cloudflare_zero_trust_device_managed_networks.x`,
		// the schema_version=0 is preserved, triggering this upgrader.
		// Must use sourceSchema because v0 state has config as a list block, not a SingleNestedAttribute.
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromLegacyV0,
		},
		// Handle upgrades from earlier v500 versions (no schema changes, just version bump)
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}

package zero_trust_access_identity_provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_identity_provider/migration/v500"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustAccessIdentityProviderResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessIdentityProviderResource)(nil)

// MoveState handles moves from cloudflare_access_identity_provider to cloudflare_zero_trust_access_identity_provider.
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_identity_provider.example
//	    to   = cloudflare_zero_trust_access_identity_provider.example
//	}
func (r *ZeroTrustAccessIdentityProviderResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceAccessIdentityProviderSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade (when TF_MIG_TEST=1)
func (r *ZeroTrustAccessIdentityProviderResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceAccessIdentityProviderSchema()
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// Handle state from v5 Plugin Framework provider with version=1
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}

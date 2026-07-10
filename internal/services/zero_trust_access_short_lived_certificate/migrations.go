package zero_trust_access_short_lived_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_short_lived_certificate/migration/v500"
)

func init() {
	// Provide target schema to migration package (avoids circular import).
	v500.V5TargetSchema = func(ctx context.Context) schema.Schema {
		return ResourceSchema(ctx)
	}
}

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessShortLivedCertificateResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustAccessShortLivedCertificateResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration
// from cloudflare_access_ca_certificate to cloudflare_zero_trust_access_short_lived_certificate.
func (r *ZeroTrustAccessShortLivedCertificateResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceAccessCACertificateSchema()

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// Two upgrade paths share schema_version=0:
// - v4 SDKv2 provider: uses "application_id" attribute
// - v5.16.0 (dormant) provider: uses "app_id" attribute
//
// Because the two source schemas are incompatible, PriorSchema for version 0 is
// nil and the handler inspects raw JSON to decide which path to take.
//
// - v5 Plugin Framework provider version=1: no-op copy-through.
func (r *ZeroTrustAccessShortLivedCertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle BOTH v4 (application_id) AND v5.16.0 (app_id) states.
		0: {
			PriorSchema:   nil, // Use RawState — schemas are incompatible
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}

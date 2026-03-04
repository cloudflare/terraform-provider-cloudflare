package zero_trust_access_mtls_certificate

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_mtls_certificate/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessMTLSCertificateResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustAccessMTLSCertificateResource)(nil)

// MoveState registers state movers for resource renames.
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration
// from cloudflare_access_mutual_tls_certificate to cloudflare_zero_trust_access_mtls_certificate.
func (r *ZeroTrustAccessMTLSCertificateResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceAccessMutualTLSCertificateSchema()

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// Clear schema version separation:
// - v4 SDKv2 provider: schema_version=0
// - v5 Plugin Framework provider: version=1 (production) or version=500 (test)
func (r *ZeroTrustAccessMTLSCertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

	v4Schema := v500.SourceAccessMutualTLSCertificateSchema()

	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_domain/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*PagesDomainResource)(nil)

// UpgradeState returns state upgraders for handling schema version migrations.
// Version 0: v4 provider schema (pre-5.x) - "domain" field
// Version 1/500: v5 provider schema - "name" field
func (r *PagesDomainResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourcePagesDomainSchema()

	return map[int64]resource.StateUpgrader{
		// Handle upgrades from v4 provider (schema_version=0)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		// Handle upgrades within v5 series (schema_version=1+) - no-op
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}

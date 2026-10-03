// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation/migration/v501"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*APIShieldOperationResource)(nil)

func (r *APIShieldOperationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	sourceSchema := v500.SourceCloudflareAPIShieldOperationSchema()

	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV4,
		},
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
		// v5.26.0 incorrectly stored feature and with_schemas in state (schema version 500).
		// Strip them so they don't cause unknown-attribute errors under the v501 schema.
		500: {
			StateUpgrader: v501.UpgradeFromV500,
		},
	}
}

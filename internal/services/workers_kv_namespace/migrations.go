// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv_namespace

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv_namespace/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*WorkersKVNamespaceResource)(nil)

func (r *WorkersKVNamespaceResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceWorkersKVNamespaceSchema()
	targetSchema := ResourceSchema(ctx)
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}

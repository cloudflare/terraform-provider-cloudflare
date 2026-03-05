// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostname

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_hostname/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.ResourceWithUpgradeState = (*CustomHostnameResource)(nil)

func (r *CustomHostnameResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)
	unionSchema := v500.UnionV0Schema(ResourceSchema, ctx)

	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   unionSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}

func init() {
	v500.V5SchemaV0 = func(ctx context.Context) *schema.Schema {
		s := ResourceSchema(ctx)
		s.Version = 0
		return &s
	}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*D1DatabaseResource)(nil)

func (r *D1DatabaseResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

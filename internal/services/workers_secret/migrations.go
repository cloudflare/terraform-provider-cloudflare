// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = &WorkersSecretResource{}

func (r *WorkersSecretResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

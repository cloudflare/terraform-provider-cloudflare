// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*APIShieldOperationResource)(nil)

func (r *APIShieldOperationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

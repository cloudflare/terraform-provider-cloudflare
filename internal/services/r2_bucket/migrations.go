// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = &R2BucketResource{}

func (r *R2BucketResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

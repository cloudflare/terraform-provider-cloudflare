// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustListResource)(nil)

func (r *ZeroTrustListResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

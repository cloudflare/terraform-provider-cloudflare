// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = &ZeroTrustLocalDomainFallbackResource{}

func (r *ZeroTrustLocalDomainFallbackResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

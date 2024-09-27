// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_incoming

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*SecondaryDNSIncomingResource)(nil)

func (r *SecondaryDNSIncomingResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

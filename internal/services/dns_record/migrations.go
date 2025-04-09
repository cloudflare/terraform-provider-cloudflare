// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_record

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*DNSRecordResource)(nil)

func (r *DNSRecordResource) UpgradeState(ctx context.Context) (map[int64]resource.StateUpgrader) {
  return map[int64]resource.StateUpgrader{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*MTLSCertificateResource)(nil)

func (r *MTLSCertificateResource) UpgradeState(ctx context.Context) (map[int64]resource.StateUpgrader) {
  return map[int64]resource.StateUpgrader{}
}

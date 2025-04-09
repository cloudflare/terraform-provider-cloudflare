// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*RegistrarDomainResource)(nil)

func (r *RegistrarDomainResource) UpgradeState(ctx context.Context) (map[int64]resource.StateUpgrader) {
  return map[int64]resource.StateUpgrader{}
}

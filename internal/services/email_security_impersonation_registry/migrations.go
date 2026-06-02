package email_security_impersonation_registry

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_impersonation_registry/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*EmailSecurityImpersonationRegistryResource)(nil)

func (r *EmailSecurityImpersonationRegistryResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	ps := v500.PriorSchema()
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &ps,
			StateUpgrader: v500.UpgradeFromV0,
		},
	}
}

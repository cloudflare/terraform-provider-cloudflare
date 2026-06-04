package email_security_trusted_domains

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_trusted_domains/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*EmailSecurityTrustedDomainsResource)(nil)

func (r *EmailSecurityTrustedDomainsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	ps := v500.PriorSchema()
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &ps,
			StateUpgrader: v500.UpgradeFromV0,
		},
	}
}

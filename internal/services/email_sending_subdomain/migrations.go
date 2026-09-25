// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_sending_subdomain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*EmailSendingSubdomainResource)(nil)

func (r *EmailSendingSubdomainResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

// This file is intentionally NOT code-generated.

package internal

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_service_token"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/provider"
)

var _ provider.ProviderWithActions = (*CloudflareProvider)(nil)

// Actions lists the provider's Terraform actions (Terraform 1.14 and later).
// Codegen doesn't produce actions, so they live here. Each action reads the
// client from ConfigureResponse.ActionData, which Configure in provider.go sets.
func (p *CloudflareProvider) Actions(ctx context.Context) []func() action.Action {
	return []func() action.Action{
		zero_trust_access_service_token.NewRefreshAction,
	}
}

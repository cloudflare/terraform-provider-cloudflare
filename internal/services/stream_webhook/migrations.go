// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_webhook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*StreamWebhookResource)(nil)

func (r *StreamWebhookResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

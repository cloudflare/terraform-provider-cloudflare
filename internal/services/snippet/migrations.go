// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*SnippetResource)(nil)

func (r *SnippetResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

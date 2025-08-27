// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*StreamDownloadResource)(nil)

func (r *StreamDownloadResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

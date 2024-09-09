// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_datasets

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*DLPDatasetsResource)(nil)

func (r *DLPDatasetsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

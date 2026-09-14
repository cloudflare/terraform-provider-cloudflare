// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package field_extractor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*FieldExtractorResource)(nil)

func (r *FieldExtractorResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

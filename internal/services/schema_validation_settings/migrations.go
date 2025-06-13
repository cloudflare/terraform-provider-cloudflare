// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*SchemaValidationSettingsResource)(nil)

func (r *SchemaValidationSettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &DevicePostureIntegrationsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DevicePostureIntegrationsDataSource{}

func (r DevicePostureIntegrationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (r *DevicePostureIntegrationsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DevicePostureIntegrationsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

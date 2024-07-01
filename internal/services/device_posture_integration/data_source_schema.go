// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_posture_integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &DevicePostureIntegrationDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DevicePostureIntegrationDataSource{}

func (r DevicePostureIntegrationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (r *DevicePostureIntegrationDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DevicePostureIntegrationDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

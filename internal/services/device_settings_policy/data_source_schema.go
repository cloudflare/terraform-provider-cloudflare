// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_settings_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &DeviceSettingsPolicyDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DeviceSettingsPolicyDataSource{}

func (r DeviceSettingsPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (r *DeviceSettingsPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DeviceSettingsPolicyDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

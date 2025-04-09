// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation_schema_validation_settings

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*APIShieldOperationSchemaValidationSettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "operation_id": schema.StringAttribute{
        Description: "UUID.",
        Required: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier.",
        Required: true,
      },
      "mitigation_action": schema.StringAttribute{
        Description: "When set, this applies a mitigation action to this operation\n\n  - `log` log request when request does not conform to schema for this operation\n  - `block` deny access to the site when request does not conform to schema for this operation\n  - `none` will skip mitigation for this operation\n  - `null` indicates that no operation level mitigation is in place, see Zone Level Schema Validation Settings for mitigation action that will be applied\nAvailable values: \"log\", \"block\", \"none\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "log",
          "block",
          "none",
        ),
        },
      },
    },
  }
}

func (d *APIShieldOperationSchemaValidationSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *APIShieldOperationSchemaValidationSettingsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

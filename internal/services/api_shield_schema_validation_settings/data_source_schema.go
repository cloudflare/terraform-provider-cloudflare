// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_schema_validation_settings

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*APIShieldSchemaValidationSettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "zone_id": schema.StringAttribute{
        Description: "Identifier.",
        Required: true,
      },
      "validation_default_mitigation_action": schema.StringAttribute{
        Description: "The default mitigation action used when there is no mitigation action defined on the operation\n\nMitigation actions are as follows:\n\n  * `log` - log request when request does not conform to schema\n  * `block` - deny access to the site when request does not conform to schema\n\nA special value of of `none` will skip running schema validation entirely for the request when there is no mitigation action defined on the operation\nAvailable values: \"none\", \"log\", \"block\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "none",
          "log",
          "block",
        ),
        },
      },
      "validation_override_mitigation_action": schema.StringAttribute{
        Description: "When set, this overrides both zone level and operation level mitigation actions.\n\n  - `none` will skip running schema validation entirely for the request\n  - `null` indicates that no override is in place\nAvailable values: \"none\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive("none"),
        },
      },
    },
  }
}

func (d *APIShieldSchemaValidationSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *APIShieldSchemaValidationSettingsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

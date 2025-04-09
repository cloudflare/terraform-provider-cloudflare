// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_behavior

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustRiskBehaviorDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Required: true,
      },
      "behaviors": schema.MapNestedAttribute{
        Computed: true,
        CustomType: customfield.NewNestedObjectMapType[ZeroTrustRiskBehaviorBehaviorsDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "description": schema.StringAttribute{
              Computed: true,
            },
            "enabled": schema.BoolAttribute{
              Computed: true,
            },
            "name": schema.StringAttribute{
              Computed: true,
            },
            "risk_level": schema.StringAttribute{
              Description: `Available values: "low", "medium", "high".`,
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "low",
                "medium",
                "high",
              ),
              },
            },
          },
        },
      },
    },
  }
}

func (d *ZeroTrustRiskBehaviorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustRiskBehaviorDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

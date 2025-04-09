// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PageShieldPoliciesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
      },
      "max_items": schema.Int64Attribute{
        Description: "Max items to fetch, default: 1000",
        Optional: true,
        Validators: []validator.Int64{
        int64validator.AtLeast(0),
        },
      },
      "result": schema.ListNestedAttribute{
        Description: "The items returned by the data source",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[PageShieldPoliciesResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "Identifier",
              Computed: true,
            },
            "action": schema.StringAttribute{
              Description: "The action to take if the expression matches\nAvailable values: \"allow\", \"log\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive("allow", "log"),
              },
            },
            "description": schema.StringAttribute{
              Description: "A description for the policy",
              Computed: true,
            },
            "enabled": schema.BoolAttribute{
              Description: "Whether the policy is enabled",
              Computed: true,
            },
            "expression": schema.StringAttribute{
              Description: "The expression which must match for the policy to be applied, using the Cloudflare Firewall rule expression syntax",
              Computed: true,
            },
            "value": schema.StringAttribute{
              Description: "The policy which will be applied",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *PageShieldPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PageShieldPoliciesDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

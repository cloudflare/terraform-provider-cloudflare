// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_categories

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustGatewayCategoriesListDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
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
        CustomType: customfield.NewNestedObjectListType[ZeroTrustGatewayCategoriesListResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
              Description: "The identifier for this category. There is only one category per ID.",
              Computed: true,
            },
            "beta": schema.BoolAttribute{
              Description: "True if the category is in beta and subject to change.",
              Computed: true,
            },
            "class": schema.StringAttribute{
              Description: "Which account types are allowed to create policies based on this category. `blocked` categories are blocked unconditionally for all accounts. `removalPending` categories can be removed from policies but not added. `noBlock` categories cannot be blocked.\nAvailable values: \"free\", \"premium\", \"blocked\", \"removalPending\", \"noBlock\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "free",
                "premium",
                "blocked",
                "removalPending",
                "noBlock",
              ),
              },
            },
            "description": schema.StringAttribute{
              Description: "A short summary of domains in the category.",
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "The name of the category.",
              Computed: true,
            },
            "subcategories": schema.ListNestedAttribute{
              Description: "All subcategories for this category.",
              Computed: true,
              CustomType: customfield.NewNestedObjectListType[ZeroTrustGatewayCategoriesListSubcategoriesDataSourceModel](ctx),
              NestedObject: schema.NestedAttributeObject{
                Attributes: map[string]schema.Attribute{
                  "id": schema.Int64Attribute{
                    Description: "The identifier for this category. There is only one category per ID.",
                    Computed: true,
                  },
                  "beta": schema.BoolAttribute{
                    Description: "True if the category is in beta and subject to change.",
                    Computed: true,
                  },
                  "class": schema.StringAttribute{
                    Description: "Which account types are allowed to create policies based on this category. `blocked` categories are blocked unconditionally for all accounts. `removalPending` categories can be removed from policies but not added. `noBlock` categories cannot be blocked.\nAvailable values: \"free\", \"premium\", \"blocked\", \"removalPending\", \"noBlock\".",
                    Computed: true,
                    Validators: []validator.String{
                    stringvalidator.OneOfCaseInsensitive(
                      "free",
                      "premium",
                      "blocked",
                      "removalPending",
                      "noBlock",
                    ),
                    },
                  },
                  "description": schema.StringAttribute{
                    Description: "A short summary of domains in the category.",
                    Computed: true,
                  },
                  "name": schema.StringAttribute{
                    Description: "The name of the category.",
                    Computed: true,
                  },
                },
              },
            },
          },
        },
      },
    },
  }
}

func (d *ZeroTrustGatewayCategoriesListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustGatewayCategoriesListDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/path"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustListDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "API Resource UUID tag.",
        Computed: true,
      },
      "list_id": schema.StringAttribute{
        Description: "API Resource UUID tag.",
        Optional: true,
      },
      "account_id": schema.StringAttribute{
        Required: true,
      },
      "created_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "description": schema.StringAttribute{
        Description: "The description of the list.",
        Computed: true,
      },
      "list_count": schema.Float64Attribute{
        Description: "The number of items in the list.",
        Computed: true,
      },
      "name": schema.StringAttribute{
        Description: "The name of the list.",
        Computed: true,
      },
      "type": schema.StringAttribute{
        Description: "The type of list.\nAvailable values: \"SERIAL\", \"URL\", \"DOMAIN\", \"EMAIL\", \"IP\".",
        Computed: true,
        Validators: []validator.String{
        stringvalidator.OneOfCaseInsensitive(
          "SERIAL",
          "URL",
          "DOMAIN",
          "EMAIL",
          "IP",
        ),
        },
      },
      "updated_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "items": schema.ListNestedAttribute{
        Description: "The items in the list.",
        Computed: true,
        CustomType: customfield.NewNestedObjectListType[ZeroTrustListItemsDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "created_at": schema.StringAttribute{
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "description": schema.StringAttribute{
              Description: "The description of the list item, if present",
              Computed: true,
            },
            "value": schema.StringAttribute{
              Description: "The value of the item in a list.",
              Computed: true,
            },
          },
        },
      },
      "filter": schema.SingleNestedAttribute{
        Optional: true,
        Attributes: map[string]schema.Attribute{
          "type": schema.StringAttribute{
            Description: "The type of list.\nAvailable values: \"SERIAL\", \"URL\", \"DOMAIN\", \"EMAIL\", \"IP\".",
            Optional: true,
            Validators: []validator.String{
            stringvalidator.OneOfCaseInsensitive(
              "SERIAL",
              "URL",
              "DOMAIN",
              "EMAIL",
              "IP",
            ),
            },
          },
        },
      },
    },
  }
}

func (d *ZeroTrustListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustListDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  datasourcevalidator.ExactlyOneOf(path.MatchRoot("list_id"), path.MatchRoot("filter")),
  }
}

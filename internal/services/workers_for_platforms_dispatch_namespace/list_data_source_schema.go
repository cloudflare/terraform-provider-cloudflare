// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersForPlatformsDispatchNamespacesDataSource)(nil)

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
        CustomType: customfield.NewNestedObjectListType[WorkersForPlatformsDispatchNamespacesResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "created_by": schema.StringAttribute{
              Description: "Identifier",
              Computed: true,
            },
            "created_on": schema.StringAttribute{
              Description: "When the script was created.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "modified_by": schema.StringAttribute{
              Description: "Identifier",
              Computed: true,
            },
            "modified_on": schema.StringAttribute{
              Description: "When the script was last modified.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "namespace_id": schema.StringAttribute{
              Description: "API Resource UUID tag.",
              Computed: true,
            },
            "namespace_name": schema.StringAttribute{
              Description: "Name of the Workers for Platforms dispatch namespace.",
              Computed: true,
            },
            "script_count": schema.Int64Attribute{
              Description: "The current number of scripts in this Dispatch Namespace",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *WorkersForPlatformsDispatchNamespacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkersForPlatformsDispatchNamespacesDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

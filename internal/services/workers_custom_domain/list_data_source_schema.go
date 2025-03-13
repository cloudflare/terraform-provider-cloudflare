// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersCustomDomainsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "Identifer of the account.",
        Required: true,
      },
      "environment": schema.StringAttribute{
        Description: "Worker environment associated with the zone and hostname.",
        Optional: true,
      },
      "hostname": schema.StringAttribute{
        Description: "Hostname of the Worker Domain.",
        Optional: true,
      },
      "service": schema.StringAttribute{
        Description: "Worker service associated with the zone and hostname.",
        Optional: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier of the zone.",
        Optional: true,
      },
      "zone_name": schema.StringAttribute{
        Description: "Name of the zone.",
        Optional: true,
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
        CustomType: customfield.NewNestedObjectListType[WorkersCustomDomainsResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "Identifer of the Worker Domain.",
              Computed: true,
            },
            "environment": schema.StringAttribute{
              Description: "Worker environment associated with the zone and hostname.",
              Computed: true,
            },
            "hostname": schema.StringAttribute{
              Description: "Hostname of the Worker Domain.",
              Computed: true,
            },
            "service": schema.StringAttribute{
              Description: "Worker service associated with the zone and hostname.",
              Computed: true,
            },
            "zone_id": schema.StringAttribute{
              Description: "Identifier of the zone.",
              Computed: true,
            },
            "zone_name": schema.StringAttribute{
              Description: "Name of the zone.",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *WorkersCustomDomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkersCustomDomainsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

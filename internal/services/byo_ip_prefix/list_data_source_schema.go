// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package byo_ip_prefix

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ByoIPPrefixesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
        Description: "Identifier of a Cloudflare account.",
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
        CustomType: customfield.NewNestedObjectListType[ByoIPPrefixesResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "Identifier of an IP Prefix.",
              Computed: true,
            },
            "account_id": schema.StringAttribute{
              Description: "Identifier of a Cloudflare account.",
              Computed: true,
            },
            "advertised": schema.BoolAttribute{
              Description: "Prefix advertisement status to the Internet. This field is only not 'null' if on demand is enabled.",
              Computed: true,
            },
            "advertised_modified_at": schema.StringAttribute{
              Description: "Last time the advertisement status was changed. This field is only not 'null' if on demand is enabled.",
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "approved": schema.StringAttribute{
              Description: "Approval state of the prefix (P = pending, V = active).",
              Computed: true,
            },
            "asn": schema.Int64Attribute{
              Description: "Autonomous System Number (ASN) the prefix will be advertised under.",
              Computed: true,
            },
            "cidr": schema.StringAttribute{
              Description: "IP Prefix in Classless Inter-Domain Routing format.",
              Computed: true,
            },
            "created_at": schema.StringAttribute{
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "description": schema.StringAttribute{
              Description: "Description of the prefix.",
              Computed: true,
            },
            "loa_document_id": schema.StringAttribute{
              Description: "Identifier for the uploaded LOA document.",
              Computed: true,
            },
            "modified_at": schema.StringAttribute{
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "on_demand_enabled": schema.BoolAttribute{
              Description: "Whether advertisement of the prefix to the Internet may be dynamically enabled or disabled.",
              Computed: true,
            },
            "on_demand_locked": schema.BoolAttribute{
              Description: "Whether advertisement status of the prefix is locked, meaning it cannot be changed.",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *ByoIPPrefixesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ByoIPPrefixesDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

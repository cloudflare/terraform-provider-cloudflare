// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_hostname

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*Web3HostnamesDataSource)(nil)

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
        CustomType: customfield.NewNestedObjectListType[Web3HostnamesResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Description: "Identifier",
              Computed: true,
            },
            "created_on": schema.StringAttribute{
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "description": schema.StringAttribute{
              Description: "An optional description of the hostname.",
              Computed: true,
            },
            "dnslink": schema.StringAttribute{
              Description: "DNSLink value used if the target is ipfs.",
              Computed: true,
            },
            "modified_on": schema.StringAttribute{
              Computed: true,
              CustomType: timetypes.RFC3339Type{

              },
            },
            "name": schema.StringAttribute{
              Description: "The hostname that will point to the target gateway via CNAME.",
              Computed: true,
            },
            "status": schema.StringAttribute{
              Description: "Status of the hostname's activation.\nAvailable values: \"active\", \"pending\", \"deleting\", \"error\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "active",
                "pending",
                "deleting",
                "error",
              ),
              },
            },
            "target": schema.StringAttribute{
              Description: "Target gateway of the hostname.\nAvailable values: \"ethereum\", \"ipfs\", \"ipfs_universal_path\".",
              Computed: true,
              Validators: []validator.String{
              stringvalidator.OneOfCaseInsensitive(
                "ethereum",
                "ipfs",
                "ipfs_universal_path",
              ),
              },
            },
          },
        },
      },
    },
  }
}

func (d *Web3HostnamesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *Web3HostnamesDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

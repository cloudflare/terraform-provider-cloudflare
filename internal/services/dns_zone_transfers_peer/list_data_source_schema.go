// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_peer

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSZoneTransfersPeersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "account_id": schema.StringAttribute{
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
        CustomType: customfield.NewNestedObjectListType[DNSZoneTransfersPeersResultDataSourceModel](ctx),
        NestedObject: schema.NestedAttributeObject{
          Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
              Computed: true,
            },
            "name": schema.StringAttribute{
              Description: "The name of the peer.",
              Computed: true,
            },
            "ip": schema.StringAttribute{
              Description: "IPv4/IPv6 address of primary or secondary nameserver, depending on what zone this peer is linked to. For primary zones this IP defines the IP of the secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary zones this IP defines the IP of the primary nameserver Cloudflare will send AXFR/IXFR requests to.",
              Computed: true,
            },
            "ixfr_enable": schema.BoolAttribute{
              Description: "Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary zones.",
              Computed: true,
            },
            "port": schema.Float64Attribute{
              Description: "DNS port of primary or secondary nameserver, depending on what zone this peer is linked to.",
              Computed: true,
            },
            "tsig_id": schema.StringAttribute{
              Description: "TSIG authentication will be used for zone transfer if configured.",
              Computed: true,
            },
          },
        },
      },
    },
  }
}

func (d *DNSZoneTransfersPeersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = ListDataSourceSchema(ctx)
}

func (d *DNSZoneTransfersPeersDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

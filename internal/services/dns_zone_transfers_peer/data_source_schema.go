// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_peer

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSZoneTransfersPeerDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Computed: true,
      },
      "peer_id": schema.StringAttribute{
        Optional: true,
      },
      "account_id": schema.StringAttribute{
        Required: true,
      },
      "ip": schema.StringAttribute{
        Description: "IPv4/IPv6 address of primary or secondary nameserver, depending on what zone this peer is linked to. For primary zones this IP defines the IP of the secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary zones this IP defines the IP of the primary nameserver Cloudflare will send AXFR/IXFR requests to.",
        Computed: true,
      },
      "ixfr_enable": schema.BoolAttribute{
        Description: "Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary zones.",
        Computed: true,
      },
      "name": schema.StringAttribute{
        Description: "The name of the peer.",
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
  }
}

func (d *DNSZoneTransfersPeerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *DNSZoneTransfersPeerDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

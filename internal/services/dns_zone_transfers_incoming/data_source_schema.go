// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_incoming

import (
  "context"

  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSZoneTransfersIncomingDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "zone_id": schema.StringAttribute{
        Required: true,
      },
      "auto_refresh_seconds": schema.Float64Attribute{
        Description: "How often should a secondary zone auto refresh regardless of DNS NOTIFY.\nNot applicable for primary zones.",
        Computed: true,
      },
      "checked_time": schema.StringAttribute{
        Description: "The time for a specific event.",
        Computed: true,
      },
      "created_time": schema.StringAttribute{
        Description: "The time for a specific event.",
        Computed: true,
      },
      "id": schema.StringAttribute{
        Computed: true,
      },
      "modified_time": schema.StringAttribute{
        Description: "The time for a specific event.",
        Computed: true,
      },
      "name": schema.StringAttribute{
        Description: "Zone name.",
        Computed: true,
      },
      "soa_serial": schema.Float64Attribute{
        Description: "The serial number of the SOA for the given zone.",
        Computed: true,
      },
      "peers": schema.ListAttribute{
        Description: "A list of peer tags.",
        Computed: true,
        CustomType: customfield.NewListType[types.String](ctx),
        ElementType: types.StringType,
      },
    },
  }
}

func (d *DNSZoneTransfersIncomingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *DNSZoneTransfersIncomingDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

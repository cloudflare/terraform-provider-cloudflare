// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_service_token

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessServiceTokenDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "UUID.",
        Computed: true,
      },
      "service_token_id": schema.StringAttribute{
        Description: "UUID.",
        Optional: true,
      },
      "account_id": schema.StringAttribute{
        Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
        Optional: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
        Optional: true,
      },
      "client_id": schema.StringAttribute{
        Description: "The Client ID for the service token. Access will check for this value in the `CF-Access-Client-ID` request header.",
        Computed: true,
      },
      "created_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "duration": schema.StringAttribute{
        Description: "The duration for how long the service token will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The default is 1 year in hours (8760h).",
        Computed: true,
      },
      "expires_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "last_seen_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "name": schema.StringAttribute{
        Description: "The name of the service token.",
        Computed: true,
      },
      "updated_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "filter": schema.SingleNestedAttribute{
        Optional: true,
        Attributes: map[string]schema.Attribute{
          "name": schema.StringAttribute{
            Description: "The name of the service token.",
            Optional: true,
          },
          "search": schema.StringAttribute{
            Description: "Search for service tokens by other listed query parameters.",
            Optional: true,
          },
        },
      },
    },
  }
}

func (d *ZeroTrustAccessServiceTokenDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessServiceTokenDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  datasourcevalidator.ExactlyOneOf(path.MatchRoot("service_token_id"), path.MatchRoot("filter")),
  datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
  }
}

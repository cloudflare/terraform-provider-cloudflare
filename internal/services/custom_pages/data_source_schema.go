// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
  "github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomPagesDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "identifier": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
      },
      "account_id": schema.StringAttribute{
        Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
        Optional: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
        Optional: true,
      },
    },
  }
}

func (d *CustomPagesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomPagesDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
  }
}

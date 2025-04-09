// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessTagDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "id": schema.StringAttribute{
        Description: "The name of the tag",
        Computed: true,
      },
      "tag_name": schema.StringAttribute{
        Description: "The name of the tag",
        Optional: true,
      },
      "account_id": schema.StringAttribute{
        Description: "Identifier.",
        Required: true,
      },
      "app_count": schema.Int64Attribute{
        Description: "The number of applications that have this tag",
        Computed: true,
      },
      "created_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
      "name": schema.StringAttribute{
        Description: "The name of the tag",
        Computed: true,
      },
      "updated_at": schema.StringAttribute{
        Computed: true,
        CustomType: timetypes.RFC3339Type{

        },
      },
    },
  }
}

func (d *ZeroTrustAccessTagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessTagDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

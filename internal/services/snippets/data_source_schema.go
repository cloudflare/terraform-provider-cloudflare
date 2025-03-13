// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/datasource"
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*SnippetsDataSource)(nil)

func DataSourceSchema(ctx context.Context) (schema.Schema) {
  return schema.Schema{
    Attributes: map[string]schema.Attribute{
      "snippet_name": schema.StringAttribute{
        Description: "Snippet identifying name",
        Required: true,
      },
      "zone_id": schema.StringAttribute{
        Description: "Identifier",
        Required: true,
      },
      "created_on": schema.StringAttribute{
        Description: "Creation time of the snippet",
        Computed: true,
      },
      "modified_on": schema.StringAttribute{
        Description: "Modification time of the snippet",
        Computed: true,
      },
    },
  }
}

func (d *SnippetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
  resp.Schema = DataSourceSchema(ctx)
}

func (d *SnippetsDataSource) ConfigValidators(_ context.Context) ([]datasource.ConfigValidator) {
  return []datasource.ConfigValidator{
  }
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_kv

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersKVDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"key_name": schema.StringAttribute{
				Description: "A key's name. The name may be at most 512 bytes. All printable, non-whitespace characters are valid. Use percent-encoding to define key names as part of a URL.",
				Required:    true,
			},
			"namespace_id": schema.StringAttribute{
				Description: "Namespace identifier tag.",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "A byte sequence that was stored, up to 25 MiB in length.",
				Computed:    true,
			},
			"metadata": schema.StringAttribute{
				Description: "Arbitrary JSON associated with the key/value pair.",
				Computed:    true,
			},
		},
	}
}

func (d *WorkersKVDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersKVDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

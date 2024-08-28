// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersScriptsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkersScriptsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the script in the Workers system. Usually the script name.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the script was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"etag": schema.StringAttribute{
							Description: "Hashed script content, can be used in a If-None-Match header when updating.",
							Computed:    true,
						},
						"logpush": schema.BoolAttribute{
							Description: "Whether Logpush is turned on for the Worker.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the script was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"placement_mode": schema.StringAttribute{
							Description: "Specifies the placement mode for the Worker (e.g. 'smart').",
							Computed:    true,
						},
						"tail_consumers": schema.ListNestedAttribute{
							Description: "List of Workers that will consume logs from the attached Worker.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[WorkersScriptsTailConsumersDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"service": schema.StringAttribute{
										Description: "Name of Worker that is to be the consumer.",
										Computed:    true,
									},
									"environment": schema.StringAttribute{
										Description: "Optional environment if the Worker utilizes one.",
										Computed:    true,
									},
									"namespace": schema.StringAttribute{
										Description: "Optional dispatch namespace the script belongs to.",
										Computed:    true,
									},
								},
							},
						},
						"usage_model": schema.StringAttribute{
							Description: "Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *WorkersScriptsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkersScriptsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

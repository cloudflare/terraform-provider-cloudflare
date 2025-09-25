// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkflowsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"search": schema.StringAttribute{
				Description: "Allows filtering workflows` name.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkflowsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"class_name": schema.StringAttribute{
							Computed: true,
						},
						"created_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"instances": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[WorkflowsInstancesDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"complete": schema.Float64Attribute{
									Computed: true,
								},
								"errored": schema.Float64Attribute{
									Computed: true,
								},
								"paused": schema.Float64Attribute{
									Computed: true,
								},
								"queued": schema.Float64Attribute{
									Computed: true,
								},
								"running": schema.Float64Attribute{
									Computed: true,
								},
								"terminated": schema.Float64Attribute{
									Computed: true,
								},
								"waiting": schema.Float64Attribute{
									Computed: true,
								},
								"waiting_for_pause": schema.Float64Attribute{
									Computed: true,
								},
							},
						},
						"modified_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"script_name": schema.StringAttribute{
							Computed: true,
						},
						"triggered_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *WorkflowsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkflowsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

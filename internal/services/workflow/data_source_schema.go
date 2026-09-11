// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkflowDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"workflow_name": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Optional:    true,
			},
			"class_name": schema.StringAttribute{
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
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
			"instances": schema.MapAttribute{
				Computed:    true,
				CustomType:  customfield.NewMapType[types.Float64](ctx),
				ElementType: types.Float64Type,
			},
			"schedules": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[WorkflowSchedulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cron": schema.StringAttribute{
							Computed: true,
						},
						"next_instance": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"search": schema.StringAttribute{
						Description: "Allows filtering workflows` name.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *WorkflowDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkflowDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("workflow_name"), path.MatchRoot("filter")),
	}
}

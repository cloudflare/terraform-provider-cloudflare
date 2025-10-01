// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflow

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkflowDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"workflow_name": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
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
			"instances": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[WorkflowInstancesDataSourceModel](ctx),
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

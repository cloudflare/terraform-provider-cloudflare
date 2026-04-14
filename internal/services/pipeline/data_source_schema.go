// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*PipelineDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Specifies the public ID of the pipeline.",
				Computed:    true,
			},
			"pipeline_id": schema.StringAttribute{
				Description: "Specifies the public ID of the pipeline.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Specifies the public ID of the account.",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"failure_reason": schema.StringAttribute{
				Description: "Indicates the reason for the failure of the Pipeline.",
				Computed:    true,
			},
			"modified_at": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Description: "Indicates the name of the Pipeline.",
				Computed:    true,
			},
			"sql": schema.StringAttribute{
				Description: "Specifies SQL for the Pipeline processing flow.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Indicates the current status of the Pipeline.",
				Computed:    true,
			},
			"tables": schema.ListNestedAttribute{
				Description: "List of streams and sinks used by this pipeline.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PipelineTablesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the connection (stream or sink).",
							Computed:    true,
						},
						"latest": schema.Int64Attribute{
							Description: "Latest available version of the connection.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name of the connection.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Type of the connection.\nAvailable values: \"stream\", \"sink\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("stream", "sink"),
							},
						},
						"version": schema.Int64Attribute{
							Description: "Current version of the connection used by this pipeline.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *PipelineDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PipelineDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

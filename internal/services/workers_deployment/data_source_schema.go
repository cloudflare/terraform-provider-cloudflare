// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersDeploymentDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Individual Workers Admin",
				"Individual Workers Content Read-Only",
				"Individual Workers Editor",
				"Individual Workers Metadata Read-Only",
				"Workers Admin",
				"Workers Content Read-Only",
				"Workers Editor",
				"Workers Metadata Read-Only",
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"deployment_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script.",
				Required:    true,
			},
			"author_email": schema.StringAttribute{
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"source": schema.StringAttribute{
				Computed: true,
			},
			"strategy": schema.StringAttribute{
				Description: "Strategy for routing traffic across the listed versions. `percentage` splits traffic according to each version's `percentage`.\nAvailable values: \"percentage\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("percentage"),
				},
			},
			"annotations": schema.SingleNestedAttribute{
				Description: "Metadata attached to the deployment.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[WorkersDeploymentAnnotationsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"workers_message": schema.StringAttribute{
						Description: "Human-readable message about the deployment. Truncated to 1000 bytes if longer.",
						Computed:    true,
					},
					"workers_triggered_by": schema.StringAttribute{
						Description: "Operation that triggered the creation of the deployment.",
						Computed:    true,
					},
				},
			},
			"code_update_strategy": schema.SingleNestedAttribute{
				Description: "Choose whether to apply Durable Object code updates immediately or defer them. When omitted, the default strategy applies.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[WorkersDeploymentCodeUpdateStrategyDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "For `mode: immediate`: Apply Durable Object code updates immediately.\n\nFor `mode: deferred`: Allow active Durable Objects to hibernate before updating their code.\nAvailable values: \"immediate\", \"deferred\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("immediate", "deferred"),
						},
					},
					"max_delay": schema.Float64Attribute{
						Description: "For `mode: deferred`: Set the maximum number of seconds to wait before forcing the update. When omitted, the default maximum delay applies.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 86400),
						},
					},
				},
			},
			"versions": schema.ListNestedAttribute{
				Description: "Worker versions included in this deployment. Each object must contain a `version_id` UUID and a `percentage`; percentages across all objects must total 100. In the `cf` CLI, pass the entire array as one JSON value to `--versions`, either inline, for example `--versions '[{\"version_id\":\"023e105f-2a42-4f8b-a1c1-73f6a2a30c0f\",\"percentage\":100}]'`, or from a JSON file with `--versions @versions.json`.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkersDeploymentVersionsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"percentage": schema.Float64Attribute{
							Description: "Percentage of traffic served by this version.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0.01, 100),
							},
						},
						"version_id": schema.StringAttribute{
							Description: "Identifier of the Worker Version.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *WorkersDeploymentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersDeploymentDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

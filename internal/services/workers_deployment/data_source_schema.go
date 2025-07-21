// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_deployment

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"deployment_id": schema.StringAttribute{
				Required: true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script, used in URLs and route configuration.",
				Required:    true,
			},
			"author_email": schema.StringAttribute{
				Computed: true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"source": schema.StringAttribute{
				Computed: true,
			},
			"strategy": schema.StringAttribute{
				Description: `Available values: "percentage".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("percentage"),
				},
			},
			"annotations": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[WorkersDeploymentAnnotationsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"workers_message": schema.StringAttribute{
						Description: "Human-readable message about the deployment. Truncated to 100 bytes.",
						Computed:    true,
					},
				},
			},
			"versions": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[WorkersDeploymentVersionsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"percentage": schema.Float64Attribute{
							Computed: true,
							Validators: []validator.Float64{
								float64validator.Between(0.01, 100),
							},
						},
						"version_id": schema.StringAttribute{
							Computed: true,
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

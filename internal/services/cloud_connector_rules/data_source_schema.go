// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*CloudConnectorRulesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"rules": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"expression": schema.StringAttribute{
							Computed: true,
						},
						"parameters": schema.SingleNestedAttribute{
							Description: "Parameters of Cloud Connector Rule",
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Description: "Host to perform Cloud Connection to",
									Computed:    true,
								},
							},
						},
						"provider": schema.StringAttribute{
							Description: "Cloud Provider type\nAvailable values: \"aws_s3\", \"r2\", \"gcp_storage\", \"azure_storage\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"aws_s3",
									"r2",
									"gcp_storage",
									"azure_storage",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *CloudConnectorRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CloudConnectorRulesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

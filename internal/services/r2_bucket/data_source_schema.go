// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*R2BucketDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account ID",
				Required:    true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the bucket",
				Required:    true,
			},
			"creation_date": schema.StringAttribute{
				Description: "Creation timestamp",
				Optional:    true,
			},
			"location": schema.StringAttribute{
				Description: "Location of the bucket",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"apac",
						"eeur",
						"enam",
						"weur",
						"wnam",
					),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the bucket",
				Optional:    true,
			},
			"storage_class": schema.StringAttribute{
				Description: "Storage class for newly uploaded objects, unless specified otherwise.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Standard", "InfrequentAccess"),
				},
			},
		},
	}
}

func (d *R2BucketDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *R2BucketDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

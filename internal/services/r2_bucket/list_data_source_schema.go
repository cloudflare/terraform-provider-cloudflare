// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &R2BucketsDataSource{}

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account ID",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Direction to order buckets",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"name_contains": schema.StringAttribute{
				Description: "Bucket names to filter by. Only buckets with this phrase in their name will be returned.",
				Optional:    true,
			},
			"order": schema.StringAttribute{
				Description: "Field to order buckets by",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("name"),
				},
			},
			"start_after": schema.StringAttribute{
				Description: "Bucket name to start searching after. Buckets are ordered lexicographically.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"creation_date": schema.StringAttribute{
							Description: "Creation timestamp",
							Computed:    true,
							Optional:    true,
						},
						"location": schema.StringAttribute{
							Description: "Location of the bucket",
							Computed:    true,
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
							Computed:    true,
							Optional:    true,
						},
						"storage_class": schema.StringAttribute{
							Description: "Storage class for newly uploaded objects, unless specified otherwise.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("Standard", "InfrequentAccess"),
							},
						},
					},
				},
			},
		},
	}
}

func (d *R2BucketsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *R2BucketsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

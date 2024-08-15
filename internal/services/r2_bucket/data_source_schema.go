// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &R2BucketDataSource{}

func (d *R2BucketDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account ID",
				Optional:    true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the bucket",
				Optional:    true,
			},
			"storage_class": schema.StringAttribute{
				Description: "Storage class for newly uploaded objects, unless specified otherwise.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Standard", "InfrequentAccess"),
				},
			},
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
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
				},
			},
		},
	}
}

func (d *R2BucketDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

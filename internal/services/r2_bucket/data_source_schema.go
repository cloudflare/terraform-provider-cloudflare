// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &R2BucketDataSource{}
var _ datasource.DataSourceWithValidateConfig = &R2BucketDataSource{}

func (r R2BucketDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
					stringvalidator.OneOfCaseInsensitive("apac", "eeur", "enam", "weur", "wnam"),
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account ID",
						Required:    true,
					},
					"cursor": schema.StringAttribute{
						Description: "Pagination cursor received during the last List Buckets call. R2 buckets are paginated using cursors instead of page numbers.",
						Optional:    true,
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
					"per_page": schema.Float64Attribute{
						Description: "Maximum number of buckets to return in a single call",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(1, 1000),
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

func (r *R2BucketDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *R2BucketDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_sippy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*R2BucketSippyDataSource)(nil)

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
			"enabled": schema.BoolAttribute{
				Description: "State of Sippy for this bucket",
				Computed:    true,
			},
			"destination": schema.SingleNestedAttribute{
				Description: "Details about the configured destination bucket",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[R2BucketSippyDestinationDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"access_key_id": schema.StringAttribute{
						Description: "ID of the Cloudflare API token used when writing objects to this\nbucket",
						Computed:    true,
					},
					"account": schema.StringAttribute{
						Computed: true,
					},
					"bucket": schema.StringAttribute{
						Description: "Name of the bucket on the provider",
						Computed:    true,
					},
					"provider": schema.StringAttribute{
						Description: "available values: \"r2\"",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("r2"),
						},
					},
				},
			},
			"source": schema.SingleNestedAttribute{
				Description: "Details about the configured source bucket",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[R2BucketSippySourceDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"bucket": schema.StringAttribute{
						Description: "Name of the bucket on the provider",
						Computed:    true,
					},
					"provider": schema.StringAttribute{
						Description: "available values: \"aws\", \"gcs\"",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("aws", "gcs"),
						},
					},
					"region": schema.StringAttribute{
						Description: "Region where the bucket resides (AWS only)",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *R2BucketSippyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *R2BucketSippyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

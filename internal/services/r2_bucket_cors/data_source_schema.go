// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_cors

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*R2BucketCORSDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account ID.",
				Required:    true,
			},
			"bucket_name": schema.StringAttribute{
				Description: "Name of the bucket.",
				Required:    true,
			},
			"rules": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[R2BucketCORSRulesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allowed": schema.SingleNestedAttribute{
							Description: "Object specifying allowed origins, methods and headers for this CORS rule.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[R2BucketCORSRulesAllowedDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"methods": schema.ListAttribute{
									Description: "Specifies the value for the Access-Control-Allow-Methods header R2 sets when requesting objects in a bucket from a browser.",
									Computed:    true,
									Validators: []validator.List{
										listvalidator.ValueStringsAre(
											stringvalidator.OneOfCaseInsensitive(
												"GET",
												"PUT",
												"POST",
												"DELETE",
												"HEAD",
											),
										),
									},
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"origins": schema.ListAttribute{
									Description: "Specifies the value for the Access-Control-Allow-Origin header R2 sets when requesting objects in a bucket from a browser.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"headers": schema.ListAttribute{
									Description: "Specifies the value for the Access-Control-Allow-Headers header R2 sets when requesting objects in this bucket from a browser. Cross-origin requests that include custom headers (e.g. x-user-id) should specify these headers as AllowedHeaders.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
						"id": schema.StringAttribute{
							Description: "Identifier for this rule.",
							Computed:    true,
						},
						"expose_headers": schema.ListAttribute{
							Description: "Specifies the headers that can be exposed back, and accessed by, the JavaScript making the cross-origin request. If you need to access headers beyond the safelisted response headers, such as Content-Encoding or cf-cache-status, you must specify it here.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"max_age_seconds": schema.Float64Attribute{
							Description: "Specifies the amount of time (in seconds) browsers are allowed to cache CORS preflight responses. Browsers may limit this to 2 hours or less, even if the maximum value (86400) is specified.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *R2BucketCORSDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *R2BucketCORSDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

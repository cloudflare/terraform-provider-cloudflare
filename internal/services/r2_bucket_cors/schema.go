// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_cors

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*R2BucketCORSResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Account ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"bucket_name": schema.StringAttribute{
				Description:   "Name of the bucket.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"jurisdiction": schema.StringAttribute{
				Description: "Jurisdiction of the bucket",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("default"),
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"default",
						"eu",
						"fedramp",
					),
				},
			},
			"rules": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[R2BucketCORSRulesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allowed": schema.SingleNestedAttribute{
							Description: "Object specifying allowed origins, methods and headers for this CORS rule.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"methods": schema.ListAttribute{
									Description: "Specifies the value for the Access-Control-Allow-Methods header R2 sets when requesting objects in a bucket from a browser.",
									Required:    true,
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
									ElementType: types.StringType,
								},
								"origins": schema.ListAttribute{
									Description: "Specifies the value for the Access-Control-Allow-Origin header R2 sets when requesting objects in a bucket from a browser.",
									Required:    true,
									ElementType: types.StringType,
								},
								"headers": schema.ListAttribute{
									Description: "Specifies the value for the Access-Control-Allow-Headers header R2 sets when requesting objects in this bucket from a browser. Cross-origin requests that include custom headers (e.g. x-user-id) should specify these headers as AllowedHeaders.",
									Optional:    true,
									ElementType: types.StringType,
								},
							},
						},
						"id": schema.StringAttribute{
							Description: "Identifier for this rule.",
							Optional:    true,
						},
						"expose_headers": schema.ListAttribute{
							Description: "Specifies the headers that can be exposed back, and accessed by, the JavaScript making the cross-origin request. If you need to access headers beyond the safelisted response headers, such as Content-Encoding or cf-cache-status, you must specify it here.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"max_age_seconds": schema.Float64Attribute{
							Description: "Specifies the amount of time (in seconds) browsers are allowed to cache CORS preflight responses. Browsers may limit this to 2 hours or less, even if the maximum value (86400) is specified.",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (r *R2BucketCORSResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *R2BucketCORSResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

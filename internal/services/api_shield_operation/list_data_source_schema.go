// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &APIShieldOperationsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &APIShieldOperationsDataSource{}

func (r APIShieldOperationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"diff": schema.BoolAttribute{
				Description: "When `true`, only return API Discovery results that are not saved into API Shield Endpoint Management",
				Optional:    true,
			},
			"direction": schema.StringAttribute{
				Description: "Direction to order results.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"endpoint": schema.StringAttribute{
				Description: "Filter results to only include endpoints containing this pattern.",
				Optional:    true,
			},
			"host": schema.ListAttribute{
				Description: "Filter results to only include the specified hosts.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"method": schema.ListAttribute{
				Description: "Filter results to only include the specified HTTP methods.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"order": schema.StringAttribute{
				Description: "Field to order by",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("host", "method", "endpoint", "traffic_stats.requests", "traffic_stats.last_updated"),
				},
			},
			"origin": schema.StringAttribute{
				Description: "Filter results to only include discovery results sourced from a particular discovery engine\n  * `ML` - Discovered operations that were sourced using ML API Discovery\n  * `SessionIdentifier` - Discovered operations that were sourced using Session Identifier API Discovery\n",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("ML", "SessionIdentifier"),
				},
			},
			"page": schema.Int64Attribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"per_page": schema.Int64Attribute{
				Description: "Maximum number of results per page.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(5, 50),
				},
			},
			"state": schema.StringAttribute{
				Description: "Filter results to only include discovery results in a particular state. States are as follows\n  * `review` - Discovered operations that are not saved into API Shield Endpoint Management\n  * `saved` - Discovered operations that are already saved into API Shield Endpoint Management\n  * `ignored` - Discovered operations that have been marked as ignored\n",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("review", "saved", "ignored"),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID",
							Computed:    true,
						},
						"endpoint": schema.StringAttribute{
							Description: "The endpoint which can contain path parameter templates in curly braces, each will be replaced from left to right with {varN}, starting with {var1}, during insertion. This will further be Cloudflare-normalized upon insertion. See: https://developers.cloudflare.com/rules/normalization/how-it-works/.",
							Computed:    true,
						},
						"host": schema.StringAttribute{
							Description: "RFC3986-compliant host.",
							Computed:    true,
						},
						"last_updated": schema.StringAttribute{
							Computed: true,
						},
						"method": schema.StringAttribute{
							Description: "The HTTP method used to access the endpoint.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("GET", "POST", "HEAD", "OPTIONS", "PUT", "DELETE", "CONNECT", "PATCH", "TRACE"),
							},
						},
						"origin": schema.ListAttribute{
							Description: "API discovery engine(s) that discovered this operation",
							Computed:    true,
							ElementType: types.StringType,
						},
						"state": schema.StringAttribute{
							Description: "State of operation in API Discovery\n  * `review` - Operation is not saved into API Shield Endpoint Management\n  * `saved` - Operation is saved into API Shield Endpoint Management\n  * `ignored` - Operation is marked as ignored\n",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("review", "saved", "ignored"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *APIShieldOperationsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *APIShieldOperationsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

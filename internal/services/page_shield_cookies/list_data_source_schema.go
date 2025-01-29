// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_cookies

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*PageShieldCookiesListDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "The direction used to sort returned cookies.'",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"domain": schema.StringAttribute{
				Description: "Filters the returned cookies that match the specified domain attribute",
				Optional:    true,
			},
			"export": schema.StringAttribute{
				Description: "Export the list of cookies as a file.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("csv"),
				},
			},
			"hosts": schema.StringAttribute{
				Description: "Includes cookies that match one or more URL-encoded hostnames separated by commas.\n\nWildcards are supported at the start and end of each hostname to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match\n",
				Optional:    true,
			},
			"http_only": schema.BoolAttribute{
				Description: "Filters the returned cookies that are set with HttpOnly",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Filters the returned cookies that match the specified name.\nWildcards are supported at the start and end to support starts with, ends with\nand contains. e.g. session*\n",
				Optional:    true,
			},
			"order_by": schema.StringAttribute{
				Description: "The field used to sort returned cookies.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("first_seen_at", "last_seen_at"),
				},
			},
			"page": schema.StringAttribute{
				Description: "The current page number of the paginated results.\n\nWe additionally support a special value \"all\". When \"all\" is used, the API will return all the cookies\nwith the applied filters in a single page. This feature is best-effort and it may only work for zones with \na low number of cookies\n",
				Optional:    true,
			},
			"page_url": schema.StringAttribute{
				Description: "Includes connections that match one or more page URLs (separated by commas) where they were last seen\n\nWildcards are supported at the start and end of each page URL to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match\n",
				Optional:    true,
			},
			"path": schema.StringAttribute{
				Description: "Filters the returned cookies that match the specified path attribute",
				Optional:    true,
			},
			"per_page": schema.Float64Attribute{
				Description: "The number of results per page.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(1, 100),
				},
			},
			"same_site": schema.StringAttribute{
				Description: "Filters the returned cookies that match the specified same_site attribute",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"lax",
						"strict",
						"none",
					),
				},
			},
			"secure": schema.BoolAttribute{
				Description: "Filters the returned cookies that are set with Secure",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "Filters the returned cookies that match the specified type attribute",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("first_party", "unknown"),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[PageShieldCookiesListResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"first_seen_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"host": schema.StringAttribute{
							Computed: true,
						},
						"last_seen_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("first_party", "unknown"),
							},
						},
						"domain_attribute": schema.StringAttribute{
							Computed: true,
						},
						"expires_attribute": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"http_only_attribute": schema.BoolAttribute{
							Computed: true,
						},
						"max_age_attribute": schema.Int64Attribute{
							Computed: true,
						},
						"page_urls": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"path_attribute": schema.StringAttribute{
							Computed: true,
						},
						"same_site_attribute": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"lax",
									"strict",
									"none",
								),
							},
						},
						"secure_attribute": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *PageShieldCookiesListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PageShieldCookiesListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

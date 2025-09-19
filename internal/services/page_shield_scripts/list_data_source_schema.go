// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_scripts

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

var _ datasource.DataSourceWithConfigValidators = (*PageShieldScriptsListDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"direction": schema.StringAttribute{
				Description: "The direction used to sort returned scripts.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"exclude_urls": schema.StringAttribute{
				Description: "Excludes scripts whose URL contains one of the URL-encoded URLs separated by commas.",
				Optional:    true,
			},
			"export": schema.StringAttribute{
				Description: "Export the list of scripts as a file.\nAvailable values: \"csv\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("csv"),
				},
			},
			"hosts": schema.StringAttribute{
				Description: "Includes scripts that match one or more URL-encoded hostnames separated by commas.\n\nWildcards are supported at the start and end of each hostname to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match",
				Optional:    true,
			},
			"order_by": schema.StringAttribute{
				Description: "The field used to sort returned scripts.\nAvailable values: \"first_seen_at\", \"last_seen_at\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("first_seen_at", "last_seen_at"),
				},
			},
			"page": schema.StringAttribute{
				Description: "The current page number of the paginated results.\n\nWe additionally support a special value \"all\". When \"all\" is used, the API will return all the scripts\nwith the applied filters in a single page. This feature is best-effort and it may only work for zones with \na low number of scripts",
				Optional:    true,
			},
			"page_url": schema.StringAttribute{
				Description: "Includes scripts that match one or more page URLs (separated by commas) where they were last seen\n\nWildcards are supported at the start and end of each page URL to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match",
				Optional:    true,
			},
			"per_page": schema.Float64Attribute{
				Description: "The number of results per page.",
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(1, 100),
				},
			},
			"prioritize_malicious": schema.BoolAttribute{
				Description: "When true, malicious scripts appear first in the returned scripts.",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Filters the returned scripts using a comma-separated list of scripts statuses. Accepted values: `active`, `infrequent`, and `inactive`. The default value is `active`.",
				Optional:    true,
			},
			"urls": schema.StringAttribute{
				Description: "Includes scripts whose URL contain one or more URL-encoded URLs separated by commas.",
				Optional:    true,
			},
			"exclude_cdn_cgi": schema.BoolAttribute{
				Description: "When true, excludes scripts seen in a `/cdn-cgi` path from the returned scripts. The default value is true.",
				Computed:    true,
				Optional:    true,
			},
			"exclude_duplicates": schema.BoolAttribute{
				Description: "When true, excludes duplicate scripts. We consider a script duplicate of another if their javascript\ncontent matches and they share the same url host and zone hostname. In such case, we return the most\nrecent script for the URL host and zone hostname combination.",
				Computed:    true,
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[PageShieldScriptsListResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"added_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
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
						"url": schema.StringAttribute{
							Computed: true,
						},
						"url_contains_cdn_cgi_path": schema.BoolAttribute{
							Computed: true,
						},
						"cryptomining_score": schema.Int64Attribute{
							Description: "The cryptomining score of the JavaScript content.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
							},
						},
						"dataflow_score": schema.Int64Attribute{
							Description: "The dataflow score of the JavaScript content.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
							},
						},
						"domain_reported_malicious": schema.BoolAttribute{
							Computed: true,
						},
						"fetched_at": schema.StringAttribute{
							Description: "The timestamp of when the script was last fetched.",
							Computed:    true,
						},
						"first_page_url": schema.StringAttribute{
							Computed: true,
						},
						"hash": schema.StringAttribute{
							Description: "The computed hash of the analyzed script.",
							Computed:    true,
						},
						"js_integrity_score": schema.Int64Attribute{
							Description: "The integrity score of the JavaScript content.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
							},
						},
						"magecart_score": schema.Int64Attribute{
							Description: "The magecart score of the JavaScript content.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
							},
						},
						"malicious_domain_categories": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"malicious_url_categories": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"malware_score": schema.Int64Attribute{
							Description: "The malware score of the JavaScript content.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
							},
						},
						"obfuscation_score": schema.Int64Attribute{
							Description: "The obfuscation score of the JavaScript content.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 99),
							},
						},
						"page_urls": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"url_reported_malicious": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *PageShieldScriptsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PageShieldScriptsListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

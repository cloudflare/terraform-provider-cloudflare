// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_scripts

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*PageShieldScriptsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"script_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"versions": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
						"fetched_at": schema.StringAttribute{
							Description: "The timestamp of when the script was last fetched.",
							Computed:    true,
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
					},
				},
			},
			"added_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
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
			"first_seen_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"hash": schema.StringAttribute{
				Description: "The computed hash of the analyzed script.",
				Computed:    true,
			},
			"host": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"js_integrity_score": schema.Int64Attribute{
				Description: "The integrity score of the JavaScript content.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 99),
				},
			},
			"last_seen_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"magecart_score": schema.Int64Attribute{
				Description: "The magecart score of the JavaScript content.",
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 99),
				},
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
			"url": schema.StringAttribute{
				Computed: true,
			},
			"url_contains_cdn_cgi_path": schema.BoolAttribute{
				Computed: true,
			},
			"url_reported_malicious": schema.BoolAttribute{
				Computed: true,
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
			"page_urls": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"direction": schema.StringAttribute{
						Description: "The direction used to sort returned scripts.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"exclude_cdn_cgi": schema.BoolAttribute{
						Description: "When true, excludes scripts seen in a `/cdn-cgi` path from the returned scripts. The default value is true.",
						Computed:    true,
						Optional:    true,
					},
					"exclude_duplicates": schema.BoolAttribute{
						Description: "When true, excludes duplicate scripts. We consider a script duplicate of another if their javascript\ncontent matches and they share the same url host and zone hostname. In such case, we return the most\nrecent script for the URL host and zone hostname combination.\n",
						Computed:    true,
						Optional:    true,
					},
					"exclude_urls": schema.StringAttribute{
						Description: "Excludes scripts whose URL contains one of the URL-encoded URLs separated by commas.\n",
						Optional:    true,
					},
					"export": schema.StringAttribute{
						Description: "Export the list of scripts as a file.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("csv"),
						},
					},
					"hosts": schema.StringAttribute{
						Description: "Includes scripts that match one or more URL-encoded hostnames separated by commas.\n\nWildcards are supported at the start and end of each hostname to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match\n",
						Optional:    true,
					},
					"order_by": schema.StringAttribute{
						Description: "The field used to sort returned scripts.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("first_seen_at", "last_seen_at"),
						},
					},
					"page": schema.StringAttribute{
						Description: "The current page number of the paginated results.\n\nWe additionally support a special value \"all\". When \"all\" is used, the API will return all the scripts\nwith the applied filters in a single page. This feature is best-effort and it may only work for zones with \na low number of scripts\n",
						Optional:    true,
					},
					"page_url": schema.StringAttribute{
						Description: "Includes scripts that match one or more page URLs (separated by commas) where they were last seen\n\nWildcards are supported at the start and end of each page URL to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match\n",
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
						Description: "Includes scripts whose URL contain one or more URL-encoded URLs separated by commas.\n",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *PageShieldScriptsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PageShieldScriptsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("script_id"), path.MatchRoot("zone_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("script_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("zone_id")),
	}
}

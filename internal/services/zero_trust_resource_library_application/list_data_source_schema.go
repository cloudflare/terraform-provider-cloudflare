// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustResourceLibraryApplicationsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"fields": schema.StringAttribute{
				Description: "Return only the listed properties on each application, as a comma-separated list.\nUse this to keep responses small when you only need part of each application, for\nexample populating a picker with `fields=id,name` instead of downloading every\nhostname and IP subnet.\n\nOmit this parameter to receive the full application object, including\nwhen lookup is used.\n\n`id` is always returned.\n\nSelectable properties: `id`, `name`, `human_id`, `version`, `hostnames`,\n`support_domains`, `ip_subnets`, `port_protocols`, `supported`, `gen_ai_score`,\n`application_confidence_score`, `created_at`, `updated_at`, `review_status`.\n\nUnknown or empty property names return `400`.",
				Optional:    true,
			},
			"filter": schema.StringAttribute{
				Description: "Filter applications using `key:value` format. Supported filter keys:\n- `name`: application name (for example, `name:HR`).\n- `id`: application ID (for example, `id:498`).\n- `human_id`: human-readable ID (for example, `human_id:HR`).\n- `hostname`: hostname or support domain (for example, `hostname:portal.example.com`).\n- `source`: application source name (for example, `source:cloudflare`).\n- `ip_subnet`: CIDR containment; returns applications where a stored subnet contains the search value (for example, `ip_subnet:10.0.1.5/32` matches `10.0.0.0/16`).\n- `category_id`: category ID (for example, `category_id:12`).\n- `category_name`: category name (for example, `category_name:HR`).\n- `supported`: supported Cloudflare product. Values: `GATEWAY`, `ACCESS`, `CASB`.\n- `review_status`: account-specific Gateway review status. Values: `approved`, `unapproved`, `in_review`, `unreviewed`.",
				Optional:    true,
			},
			"lookup": schema.StringAttribute{
				Description: "Find complete, case-insensitive application name mentions in a phrase or\nrule sentence. Misspelled names do not match. Returns matching applications\nranked by relevance. Must be between 2 and 1000 characters. Uses the same\nfilters, fields, and pagination as listing. `search` cannot be combined\nwith lookup. Results are ranked by relevance, so `order_by` is ignored.",
				Optional:    true,
			},
			"order_by": schema.StringAttribute{
				Description: "Order results using field:direction format. Supported fields are name, id, human_id,\ncategory_id, application_type, application_confidence_score, and gen_ai_score.\nSupported directions are asc and desc. Ignored when search is provided; results are\nranked by relevance instead. Also ignored when lookup is provided.",
				Optional:    true,
			},
			"search": schema.StringAttribute{
				Description: "Fuzzy search across application name and hostnames. Results are ranked by relevance. Must be between 2 and 200 characters. Can be combined with filter parameters.",
				Optional:    true,
			},
			"limit": schema.Int64Attribute{
				Description: "Number of results to return. Defaults to 25; values are clamped to 1–250.",
				Computed:    true,
				Optional:    true,
			},
			"offset": schema.Int64Attribute{
				Description: "Number of results to skip. Defaults to 0; negative values are clamped to 0.",
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustResourceLibraryApplicationsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "Numeric identifier for the application.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"application_confidence_score": schema.Float64Attribute{
							Description: "Confidence score for the application. Returns -1 when no score is available.",
							Computed:    true,
						},
						"application_score_composition": schema.StringAttribute{
							Description: "Returns the score composition breakdown for the application.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"application_source": schema.StringAttribute{
							Description: "Returns the application source.",
							Computed:    true,
						},
						"application_type": schema.StringAttribute{
							Description: "Returns the application type.",
							Computed:    true,
						},
						"application_type_description": schema.StringAttribute{
							Description: "Returns the application type description.",
							Computed:    true,
						},
						"category_id": schema.Int64Attribute{
							Description: "Numeric identifier for an application category.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"created_at": schema.StringAttribute{
							Description: "Returns the application creation time.",
							Computed:    true,
						},
						"gen_ai_score": schema.Float64Attribute{
							Description: "GenAI score for the application. Returns -1 when no score is available.",
							Computed:    true,
						},
						"hostnames": schema.SetAttribute{
							Description: "Hostnames matched by the application.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"human_id": schema.StringAttribute{
							Description: "Human-readable identifier for the application.",
							Computed:    true,
						},
						"ip_subnets": schema.SetAttribute{
							Description: "IP subnets for this application. Custom application create and update requests accept IPv4 prefix lengths /8 through /32 and IPv6 prefix lengths /32 through /128.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"name": schema.StringAttribute{
							Description: "Application name.",
							Computed:    true,
						},
						"port_protocols": schema.SetAttribute{
							Description: "Ports matched by the application, in `protocol/port` or inclusive `protocol/start-end` format (for example, `tcp/443` or `udp/10000-20000`). TCP and UDP ports must be between 0 and 65535.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"review_status": schema.StringAttribute{
							Description: "The account-specific Gateway review status. Applications with no assigned review status are returned as `unreviewed`.\nAvailable values: \"approved\", \"unapproved\", \"in_review\", \"unreviewed\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"approved",
									"unapproved",
									"in_review",
									"unreviewed",
								),
							},
						},
						"support_domains": schema.SetAttribute{
							Description: "Support domains matched by the application.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"supported": schema.SetAttribute{
							Description: "Cloudflare products that support this application.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"updated_at": schema.StringAttribute{
							Description: "Returns the application update time.",
							Computed:    true,
						},
						"version": schema.StringAttribute{
							Description: "Returns the application version.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustResourceLibraryApplicationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustResourceLibraryApplicationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

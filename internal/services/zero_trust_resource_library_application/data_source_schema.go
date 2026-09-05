// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustResourceLibraryApplicationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Returns the application ID.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"application_confidence_score": schema.Float64Attribute{
				Description: "Confidence score for the application. Returns -1 when no score is available.",
				Computed:    true,
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
				Description: "Returns the category ID.",
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
			"human_id": schema.StringAttribute{
				Description: "Returns the human readable ID.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Returns the application name.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "Returns the application update time.",
				Computed:    true,
			},
			"version": schema.StringAttribute{
				Description: "Returns the application version.",
				Computed:    true,
			},
			"hostnames": schema.SetAttribute{
				Description: "Hostnames matched by the application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"ip_subnets": schema.SetAttribute{
				Description: "IP subnets matched by the application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
			},
			"port_protocols": schema.SetAttribute{
				Description: "Port and protocol pairs matched by the application.",
				Computed:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
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
			"application_score_composition": schema.StringAttribute{
				Description: "Returns the score composition breakdown for the application.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"fields": schema.StringAttribute{
						Description: "Return only the listed properties on each application, as a comma-separated list.\nUse this to keep responses small when you only need part of each application — for\nexample populating a picker with `fields=id,name` instead of downloading every\nhostname and IP subnet.\n\nOmit this parameter to receive the full application object.\n\n`id` is always returned.\n\nSelectable properties: `id`, `name`, `human_id`, `version`, `hostnames`,\n`support_domains`, `ip_subnets`, `port_protocols`, `supported`, `gen_ai_score`,\n`application_confidence_score`, `created_at`, `updated_at`, `review_status`.\n\nUnknown or empty property names return `400`.",
						Optional:    true,
					},
					"filter": schema.StringAttribute{
						Description: "Filter applications using key:value format. Supported filter keys:\n- name: Filter by application name (e.g., name:HR)\n- id: Filter by application ID (e.g., id:498)\n- human_id: Filter by human-readable ID (e.g., human_id:HR)\n- hostname: Filter by hostname or support domain (e.g., hostname:portal.example.com)\n- source: Filter by application source name (e.g., source:cloudflare)\n- ip_subnet: Filter by IP subnet using CIDR containment — returns applications where any stored subnet contains the search value (e.g., ip_subnet:10.0.1.5/32 matches apps with 10.0.0.0/16)\n- category_id: Filter by category ID (e.g., category_id:12).\n- category_name: Filter by category name (e.g., category_name:HR).\n- supported: Filter by supported Cloudflare product (e.g., supported:ACCESS). Values: GATEWAY, ACCESS, CASB.\n- review_status: Filter by the account's Gateway review status. Values: approved, unapproved, in_review, unreviewed.\n.",
						Optional:    true,
					},
					"limit": schema.Int64Attribute{
						Description: "Limit of number of results to return (max 250).",
						Computed:    true,
						Optional:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "Offset of results to return.",
						Computed:    true,
						Optional:    true,
					},
					"order_by": schema.StringAttribute{
						Description: "Order results using field:direction format. Supported fields are name, id, human_id,\ncategory_id, application_type, application_confidence_score, and gen_ai_score.\nSupported directions are asc and desc. Ignored when search is provided; results are\nranked by relevance instead.",
						Optional:    true,
					},
					"search": schema.StringAttribute{
						Description: "Fuzzy search across application name and hostnames. Results are ranked by relevance. Must be between 2 and 200 characters. Can be combined with filter parameters.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustResourceLibraryApplicationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustResourceLibraryApplicationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("id"), path.MatchRoot("filter")),
	}
}

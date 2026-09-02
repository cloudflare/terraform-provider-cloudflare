// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_resource_library_application

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
			"filter": schema.StringAttribute{
				Description: "Filter applications using key:value format. Supported filter keys:\n- name: Filter by application name (e.g., name:HR)\n- id: Filter by application ID (e.g., id:498)\n- human_id: Filter by human-readable ID (e.g., human_id:HR)\n- hostname: Filter by hostname or support domain (e.g., hostname:portal.example.com)\n- source: Filter by application source name (e.g., source:cloudflare)\n- ip_subnet: Filter by IP subnet using CIDR containment — returns applications where any stored subnet contains the search value (e.g., ip_subnet:10.0.1.5/32 matches apps with 10.0.0.0/16)\n- category_id: Filter by category ID (e.g., category_id:12).\n- category_name: Filter by category name (e.g., category_name:HR).\n- supported: Filter by supported Cloudflare product (e.g., supported:ACCESS). Values: GATEWAY, ACCESS, CASB.\n.",
				Optional:    true,
			},
			"order_by": schema.StringAttribute{
				Description: "Order results by field name and direction (e.g., name:asc). Ignored when search is provided; results are ranked by relevance instead.",
				Optional:    true,
			},
			"search": schema.StringAttribute{
				Description: "Fuzzy search across application name and hostnames. Results are ranked by relevance. Must be between 2 and 200 characters. Can be combined with filter parameters.",
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
							Description: "Returns the application ID.",
							Computed:    true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
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
						"hostnames": schema.SetAttribute{
							Description: "Hostnames matched by the application.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"human_id": schema.StringAttribute{
							Description: "Returns the human readable ID.",
							Computed:    true,
						},
						"ip_subnets": schema.SetAttribute{
							Description: "IP subnets matched by the application.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"name": schema.StringAttribute{
							Description: "Returns the application name.",
							Computed:    true,
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
						"updated_at": schema.StringAttribute{
							Description: "Returns the application update time.",
							Computed:    true,
						},
						"version": schema.StringAttribute{
							Description: "Returns the application version.",
							Computed:    true,
						},
						"application_score_composition": schema.StringAttribute{
							Description: "Returns the score composition breakdown for the application.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
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

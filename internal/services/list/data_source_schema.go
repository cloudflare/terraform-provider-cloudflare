// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ListDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID for this resource.",
				Required:    true,
			},
			"list_id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the list was created.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the list.",
				Computed:    true,
			},
			"kind": schema.StringAttribute{
				Description: "The type of the list. Each type supports specific list items (IP addresses, ASNs, hostnames or redirects).\nAvailable values: \"ip\", \"redirect\", \"hostname\", \"asn\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ip",
						"redirect",
						"hostname",
						"asn",
					),
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the list was last modified.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "An informative name for the list. Use this name in filter and rule expressions.",
				Computed:    true,
			},
			"num_items": schema.Float64Attribute{
				Description: "The number of items in the list.",
				Computed:    true,
			},
			"num_referencing_filters": schema.Float64Attribute{
				Description: "The number of [filters](/api/resources/filters/) referencing the list.",
				Computed:    true,
			},
			"search": schema.StringAttribute{
				Description: " A search query to filter returned items. Its meaning depends on the list type: IP addresses must start with the provided string, hostnames and bulk redirects must contain the string, and ASNs must match the string exactly.",
				Optional:    true,
			},
			"items": schema.SetNestedAttribute{
				Description: "The items in the list. If set, this overwrites all items in the list. Do not use with `cloudflare_list_item`.",
				CustomType:  customfield.NewNestedObjectSetType[ListItemDataSourceModel](ctx),
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"asn": schema.Int64Attribute{
							Description: "A non-negative 32 bit integer",
							Computed:    true,
						},
						"comment": schema.StringAttribute{
							Description: "An informative summary of the list item.",
							Computed:    true,
						},
						"hostname": schema.SingleNestedAttribute{
							Description: "Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).",
							CustomType:  customfield.NewNestedObjectType[ListItemHostnameDataSourceModel](ctx),
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"url_hostname": schema.StringAttribute{
									Computed: true,
								},
								"exclude_exact_hostname": schema.BoolAttribute{
									Description: "Only applies to wildcard hostnames (e.g., *.example.com). When true (default), only subdomains are blocked. When false, both the root domain and subdomains are blocked.",
									Computed:    true,
								},
							},
						},
						"ip": schema.StringAttribute{
							Description: "An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.",
							Computed:    true,
						},
						"redirect": schema.SingleNestedAttribute{
							Description: "The definition of the redirect.",
							CustomType:  customfield.NewNestedObjectType[ListItemRedirectDataSourceModel](ctx),
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"source_url": schema.StringAttribute{
									Computed: true,
								},
								"target_url": schema.StringAttribute{
									Computed: true,
								},
								"include_subdomains": schema.BoolAttribute{
									Computed: true,
								},
								"preserve_path_suffix": schema.BoolAttribute{
									Computed: true,
								},
								"preserve_query_string": schema.BoolAttribute{
									Computed: true,
								},
								"status_code": schema.Int64Attribute{
									Description: "Available values: 301, 302, 307, 308.",
									Computed:    true,
								},
								"subpath_matching": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

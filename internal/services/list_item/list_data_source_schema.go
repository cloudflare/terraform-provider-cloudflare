// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ListItemsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"list_id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Required:    true,
			},
			"search": schema.StringAttribute{
				Description: "A search query to filter returned items. Its meaning depends on the list type: IP addresses must start with the provided string, hostnames and bulk redirects must contain the string, and ASNs must match the string exactly.",
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
				CustomType:  customfield.NewNestedObjectListType[ListItemsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique ID of the list.",
							Computed:    true,
						},
						"asn": schema.Int64Attribute{
							Description: "A non-negative 32 bit integer",
							Computed:    true,
						},
						"comment": schema.StringAttribute{
							Description: "An informative summary of the list item.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "The RFC 3339 timestamp of when the item was created.",
							Computed:    true,
						},
						"hostname": schema.SingleNestedAttribute{
							Description: "Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ListItemsHostnameDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"url_hostname": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"ip": schema.StringAttribute{
							Description: "An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "The RFC 3339 timestamp of when the item was last modified.",
							Computed:    true,
						},
						"redirect": schema.SingleNestedAttribute{
							Description: "The definition of the redirect.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ListItemsRedirectDataSourceModel](ctx),
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
									Computed: true,
									Validators: []validator.Int64{
										int64validator.OneOf(
											301,
											302,
											307,
											308,
										),
									},
								},
								"subpath_matching": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
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
							Computed: true,
							Validators: []validator.Int64{
								int64validator.OneOf(
									301,
									302,
									307,
									308,
								),
							},
						},
						"subpath_matching": schema.BoolAttribute{
							Computed: true,
						},
						"url_hostname": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *ListItemsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ListItemsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

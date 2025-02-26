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

var _ datasource.DataSourceWithConfigValidators = (*ListItemDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"item_id": schema.StringAttribute{
				Description: "The unique ID of the item in the List.",
				Required:    true,
			},
			"list_id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Required:    true,
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
			"id": schema.StringAttribute{
				Description: "The unique ID of the list.",
				Computed:    true,
			},
			"ip": schema.StringAttribute{
				Description: "An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The RFC 3339 timestamp of when the item was last modified.",
				Computed:    true,
			},
			"hostname": schema.SingleNestedAttribute{
				Description: "Valid characters for hostnames are ASCII(7) letters from a to z, the digits from 0 to 9, wildcards (*), and the hyphen (-).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ListItemHostnameDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"url_hostname": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"redirect": schema.SingleNestedAttribute{
				Description: "The definition of the redirect.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ListItemRedirectDataSourceModel](ctx),
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
		},
	}
}

func (d *ListItemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ListItemDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

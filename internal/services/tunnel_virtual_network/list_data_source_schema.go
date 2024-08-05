// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_virtual_network

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &TunnelVirtualNetworksDataSource{}

func (d *TunnelVirtualNetworksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "UUID of the virtual network.",
				Optional:    true,
			},
			"is_default": schema.BoolAttribute{
				Description: "If `true`, only include the default virtual network. If `false`, exclude the default virtual network. If empty, all virtual networks will be included.",
				Optional:    true,
			},
			"is_deleted": schema.BoolAttribute{
				Description: "If `true`, only include deleted virtual networks. If `false`, exclude deleted virtual networks. If empty, all virtual networks will be included.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for the virtual network.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID of the virtual network.",
							Computed:    true,
						},
						"comment": schema.StringAttribute{
							Description: "Optional remark describing the virtual network.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Timestamp of when the resource was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"is_default_network": schema.BoolAttribute{
							Description: "If `true`, this virtual network is the default for the account.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "A user-friendly name for the virtual network.",
							Computed:    true,
						},
						"deleted_at": schema.StringAttribute{
							Description: "Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *TunnelVirtualNetworksDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

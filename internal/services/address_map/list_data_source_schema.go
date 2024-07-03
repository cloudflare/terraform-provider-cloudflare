// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &AddressMapsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AddressMapsDataSource{}

func (r AddressMapsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"can_delete": schema.BoolAttribute{
							Description: "If set to false, then the Address Map cannot be deleted via API. This is true for Cloudflare-managed maps.",
							Computed:    true,
						},
						"can_modify_ips": schema.BoolAttribute{
							Description: "If set to false, then the IPs on the Address Map cannot be modified via the API. This is true for Cloudflare-managed maps.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"default_sni": schema.StringAttribute{
							Description: "If you have legacy TLS clients which do not send the TLS server name indicator, then you can specify one default SNI on the map. If Cloudflare receives a TLS handshake from a client without an SNI, it will respond with the default SNI on those IPs. The default SNI can be any valid zone or subdomain owned by the account.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An optional description field which may be used to describe the types of IPs or zones on the map.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Address Map is enabled or not. Cloudflare's DNS will not respond with IP addresses on an Address Map until the map is enabled.",
							Computed:    true,
						},
						"modified_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *AddressMapsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AddressMapsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &AddressMapDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AddressMapDataSource{}

func (r AddressMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"address_map_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"can_delete": schema.BoolAttribute{
				Description: "If set to false, then the Address Map cannot be deleted via API. This is true for Cloudflare-managed maps.",
				Optional:    true,
			},
			"can_modify_ips": schema.BoolAttribute{
				Description: "If set to false, then the IPs on the Address Map cannot be modified via the API. This is true for Cloudflare-managed maps.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Optional: true,
			},
			"default_sni": schema.StringAttribute{
				Description: "If you have legacy TLS clients which do not send the TLS server name indicator, then you can specify one default SNI on the map. If Cloudflare receives a TLS handshake from a client without an SNI, it will respond with the default SNI on those IPs. The default SNI can be any valid zone or subdomain owned by the account.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "An optional description field which may be used to describe the types of IPs or zones on the map.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the Address Map is enabled or not. Cloudflare's DNS will not respond with IP addresses on an Address Map until the map is enabled.",
				Computed:    true,
				Optional:    true,
			},
			"ips": schema.ListNestedAttribute{
				Description: "The set of IPs on the Address Map.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Optional: true,
						},
						"ip": schema.StringAttribute{
							Description: "An IPv4 or IPv6 address.",
							Optional:    true,
						},
					},
				},
			},
			"memberships": schema.ListNestedAttribute{
				Description: "Zones and Accounts which will be assigned IPs on this Address Map. A zone membership will take priority over an account membership.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"can_delete": schema.BoolAttribute{
							Description: "Controls whether the membership can be deleted via the API or not.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Optional: true,
						},
						"identifier": schema.StringAttribute{
							Description: "The identifier for the membership (eg. a zone or account tag).",
							Optional:    true,
						},
						"kind": schema.StringAttribute{
							Description: "The type of the membership.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("zone", "account"),
							},
						},
					},
				},
			},
			"modified_at": schema.StringAttribute{
				Optional: true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *AddressMapDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AddressMapDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

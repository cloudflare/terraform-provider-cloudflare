// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package address_map

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AddressMapDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier of an Address Map.",
				Computed:    true,
			},
			"address_map_id": schema.StringAttribute{
				Description: "Identifier of an Address Map.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier of a Cloudflare account.",
				Required:    true,
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
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
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
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"ips": schema.ListNestedAttribute{
				Description: "The set of IPs on the Address Map.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AddressMapIPsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"ip": schema.StringAttribute{
							Description: "An IPv4 or IPv6 address.",
							Computed:    true,
						},
					},
				},
			},
			"memberships": schema.ListNestedAttribute{
				Description: "Zones and Accounts which will be assigned IPs on this Address Map. A zone membership will take priority over an account membership.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AddressMapMembershipsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"can_delete": schema.BoolAttribute{
							Description: "Controls whether the membership can be deleted via the API or not.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"identifier": schema.StringAttribute{
							Description: "The identifier for the membership (eg. a zone or account tag).",
							Computed:    true,
						},
						"kind": schema.StringAttribute{
							Description: "The type of the membership.\navailable values: \"zone\", \"account\"",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("zone", "account"),
							},
						},
					},
				},
			},
		},
	}
}

func (d *AddressMapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AddressMapDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_acl

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitSiteACLsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
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
				CustomType:  customfield.NewNestedObjectListType[MagicTransitSiteACLsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Description for the ACL.",
							Computed:    true,
						},
						"forward_locally": schema.BoolAttribute{
							Description: "The desired forwarding action for this ACL policy. If set to \"false\", the policy will forward traffic to Cloudflare. If set to \"true\", the policy will forward traffic locally on the Magic Connector. If not included in request, will default to false.",
							Computed:    true,
						},
						"lan_1": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[MagicTransitSiteACLsLAN1DataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"lan_id": schema.StringAttribute{
									Description: "The identifier for the LAN you want to create an ACL policy with.",
									Computed:    true,
								},
								"lan_name": schema.StringAttribute{
									Description: "The name of the LAN based on the provided lan_id.",
									Computed:    true,
								},
								"ports": schema.ListAttribute{
									Description: "Array of ports on the provided LAN that will be included in the ACL. If no ports are provided, communication on any port on this LAN is allowed.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.Int64](ctx),
									ElementType: types.Int64Type,
								},
								"subnets": schema.ListAttribute{
									Description: "Array of subnet IPs within the LAN that will be included in the ACL. If no subnets are provided, communication on any subnets on this LAN are allowed.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
						"lan_2": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[MagicTransitSiteACLsLAN2DataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"lan_id": schema.StringAttribute{
									Description: "The identifier for the LAN you want to create an ACL policy with.",
									Computed:    true,
								},
								"lan_name": schema.StringAttribute{
									Description: "The name of the LAN based on the provided lan_id.",
									Computed:    true,
								},
								"ports": schema.ListAttribute{
									Description: "Array of ports on the provided LAN that will be included in the ACL. If no ports are provided, communication on any port on this LAN is allowed.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.Int64](ctx),
									ElementType: types.Int64Type,
								},
								"subnets": schema.ListAttribute{
									Description: "Array of subnet IPs within the LAN that will be included in the ACL. If no subnets are provided, communication on any subnets on this LAN are allowed.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the ACL.",
							Computed:    true,
						},
						"protocols": schema.ListAttribute{
							Computed: true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"tcp",
										"udp",
										"icmp",
									),
								),
							},
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"unidirectional": schema.BoolAttribute{
							Description: "The desired traffic direction for this ACL policy. If set to \"false\", the policy will allow bidirectional traffic. If set to \"true\", the policy will only allow traffic in one direction. If not included in request, will default to false.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *MagicTransitSiteACLsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MagicTransitSiteACLsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_acl

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitSiteACLDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"acl_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"site_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description for the ACL.",
				Computed:    true,
			},
			"forward_locally": schema.BoolAttribute{
				Description: "The desired forwarding action for this ACL policy. If set to \"false\", the policy will forward traffic to Cloudflare. If set to \"true\", the policy will forward traffic locally on the Magic Connector. If not included in request, will default to false.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the ACL.",
				Computed:    true,
			},
			"unidirectional": schema.BoolAttribute{
				Description: "The desired traffic direction for this ACL policy. If set to \"false\", the policy will allow bidirectional traffic. If set to \"true\", the policy will only allow traffic in one direction. If not included in request, will default to false.",
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
			"lan_1": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicTransitSiteACLLAN1DataSourceModel](ctx),
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
				CustomType: customfield.NewNestedObjectType[MagicTransitSiteACLLAN2DataSourceModel](ctx),
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"site_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *MagicTransitSiteACLDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicTransitSiteACLDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(
			path.MatchRoot("account_id"),
			path.MatchRoot("acl_id"),
			path.MatchRoot("site_id"),
		),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("acl_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("site_id")),
	}
}

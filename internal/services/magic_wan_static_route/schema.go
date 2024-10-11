// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MagicWANStaticRouteResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"route_id": schema.StringAttribute{
				Description:   "Identifier",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "An optional human provided description of the static route.",
				Optional:    true,
			},
			"nexthop": schema.StringAttribute{
				Description: "The next-hop IP Address for the static route.",
				Optional:    true,
			},
			"prefix": schema.StringAttribute{
				Description: "IP Prefix in Classless Inter-Domain Routing format.",
				Optional:    true,
			},
			"priority": schema.Int64Attribute{
				Description: "Priority of the static route.",
				Optional:    true,
			},
			"weight": schema.Int64Attribute{
				Description: "Optional weight of the ECMP scope - if provided.",
				Optional:    true,
			},
			"scope": schema.SingleNestedAttribute{
				Description: "Used only for ECMP routes.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[MagicWANStaticRouteScopeModel](ctx),
				Attributes: map[string]schema.Attribute{
					"colo_names": schema.ListAttribute{
						Description: "List of colo names for the ECMP scope.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"colo_regions": schema.ListAttribute{
						Description: "List of colo regions for the ECMP scope.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"modified": schema.BoolAttribute{
				Computed: true,
			},
			"modified_route": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicWANStaticRouteModifiedRouteModel](ctx),
				Attributes: map[string]schema.Attribute{
					"nexthop": schema.StringAttribute{
						Description: "The next-hop IP Address for the static route.",
						Computed:    true,
					},
					"prefix": schema.StringAttribute{
						Description: "IP Prefix in Classless Inter-Domain Routing format.",
						Computed:    true,
					},
					"priority": schema.Int64Attribute{
						Description: "Priority of the static route.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "When the route was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"description": schema.StringAttribute{
						Description: "An optional human provided description of the static route.",
						Computed:    true,
					},
					"modified_on": schema.StringAttribute{
						Description: "When the route was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"scope": schema.SingleNestedAttribute{
						Description: "Used only for ECMP routes.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[MagicWANStaticRouteModifiedRouteScopeModel](ctx),
						Attributes: map[string]schema.Attribute{
							"colo_names": schema.ListAttribute{
								Description: "List of colo names for the ECMP scope.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"colo_regions": schema.ListAttribute{
								Description: "List of colo regions for the ECMP scope.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"weight": schema.Int64Attribute{
						Description: "Optional weight of the ECMP scope - if provided.",
						Computed:    true,
					},
				},
			},
			"route": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[MagicWANStaticRouteRouteModel](ctx),
				Attributes: map[string]schema.Attribute{
					"nexthop": schema.StringAttribute{
						Description: "The next-hop IP Address for the static route.",
						Computed:    true,
					},
					"prefix": schema.StringAttribute{
						Description: "IP Prefix in Classless Inter-Domain Routing format.",
						Computed:    true,
					},
					"priority": schema.Int64Attribute{
						Description: "Priority of the static route.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "When the route was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"description": schema.StringAttribute{
						Description: "An optional human provided description of the static route.",
						Computed:    true,
					},
					"modified_on": schema.StringAttribute{
						Description: "When the route was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"scope": schema.SingleNestedAttribute{
						Description: "Used only for ECMP routes.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[MagicWANStaticRouteRouteScopeModel](ctx),
						Attributes: map[string]schema.Attribute{
							"colo_names": schema.ListAttribute{
								Description: "List of colo names for the ECMP scope.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"colo_regions": schema.ListAttribute{
								Description: "List of colo regions for the ECMP scope.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
						},
					},
					"weight": schema.Int64Attribute{
						Description: "Optional weight of the ECMP scope - if provided.",
						Computed:    true,
					},
				},
			},
			"routes": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[MagicWANStaticRouteRoutesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nexthop": schema.StringAttribute{
							Description: "The next-hop IP Address for the static route.",
							Computed:    true,
						},
						"prefix": schema.StringAttribute{
							Description: "IP Prefix in Classless Inter-Domain Routing format.",
							Computed:    true,
						},
						"priority": schema.Int64Attribute{
							Description: "Priority of the static route.",
							Computed:    true,
						},
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the route was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "An optional human provided description of the static route.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the route was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"scope": schema.SingleNestedAttribute{
							Description: "Used only for ECMP routes.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[MagicWANStaticRouteRoutesScopeModel](ctx),
							Attributes: map[string]schema.Attribute{
								"colo_names": schema.ListAttribute{
									Description: "List of colo names for the ECMP scope.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"colo_regions": schema.ListAttribute{
									Description: "List of colo regions for the ECMP scope.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
						"weight": schema.Int64Attribute{
							Description: "Optional weight of the ECMP scope - if provided.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *MagicWANStaticRouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicWANStaticRouteResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

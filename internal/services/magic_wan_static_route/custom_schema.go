package magic_wan_static_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func customResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"nexthop": schema.StringAttribute{
				Description: "The next-hop IP Address for the static route.",
				Required:    true,
			},
			"prefix": schema.StringAttribute{
				Description: "IP Prefix in Classless Inter-Domain Routing format.",
				Required:    true,
			},
			"priority": schema.Int64Attribute{
				Description: "Priority of the static route.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "An optional human provided description of the static route.",
				Optional:    true,
			},
			"weight": schema.Int64Attribute{
				Description: "Optional weight of the ECMP scope - if provided.",
				Optional:    true,
			},
			"scope": schema.SingleNestedAttribute{
				Description: "Used only for ECMP routes.",
				Optional:    true,
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
			"created_on": schema.StringAttribute{
				Description: "When the route was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the route was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

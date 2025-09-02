package magic_transit_connector

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitConnectorResource)(nil)

func CustomResourceSchema(_ context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"device": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
					},
					"serial_number": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
					},
				},
			},
			"activated": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"interrupt_window_duration_hours": schema.Float64Attribute{
				Optional: true,
				Computed: true,
			},
			"interrupt_window_hour_of_day": schema.Float64Attribute{
				Optional: true,
				Computed: true,
			},
			"notes": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"timezone": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

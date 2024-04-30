// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r MagicTransitSiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the site.",
				Required:    true,
			},
			"connector_id": schema.StringAttribute{
				Description: "Magic WAN Connector identifier tag.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"ha_mode": schema.BoolAttribute{
				Description: "Site high availability mode. If set to true, the site can have two connectors and runs in high availability mode.",
				Optional:    true,
			},
			"location": schema.SingleNestedAttribute{
				Description: "Location of site in latitude and longitude.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"lat": schema.StringAttribute{
						Description: "Latitude",
						Optional:    true,
					},
					"lon": schema.StringAttribute{
						Description: "Longitude",
						Optional:    true,
					},
				},
			},
			"secondary_connector_id": schema.StringAttribute{
				Description: "Magic WAN Connector identifier tag. Used when high availability mode is on.",
				Optional:    true,
			},
		},
	}
}

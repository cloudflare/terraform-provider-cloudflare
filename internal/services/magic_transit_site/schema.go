// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitSiteResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
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
			"ha_mode": schema.BoolAttribute{
				Description:   "Site high availability mode. If set to true, the site can have two connectors and runs in high availability mode.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the site.",
				Required:    true,
			},
			"connector_id": schema.StringAttribute{
				Description: "Magic Connector identifier tag.",
				Computed:    true,
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"secondary_connector_id": schema.StringAttribute{
				Description: "Magic Connector identifier tag. Used when high availability mode is on.",
				Computed:    true,
				Optional:    true,
			},
			"location": schema.SingleNestedAttribute{
				Description: "Location of site in latitude and longitude.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[MagicTransitSiteLocationModel](ctx),
				Attributes: map[string]schema.Attribute{
					"lat": schema.StringAttribute{
						Description: "Latitude",
						Computed:    true,
						Optional:    true,
					},
					"lon": schema.StringAttribute{
						Description: "Longitude",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *MagicTransitSiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicTransitSiteResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

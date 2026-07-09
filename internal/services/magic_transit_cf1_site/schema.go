// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_cf1_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitCf1SiteResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Magic Transit Read",
				"Magic Transit Write",
				"Magic WAN Read",
				"Magic WAN Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"body": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "A human-provided name describing the CF1 Site that should be unique within the account.",
							Required:    true,
						},
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "A human-provided description of the CF1 Site.",
							Optional:    true,
						},
						"location": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"lat": schema.Float64Attribute{
									Description: "Latitude of the CF1 Site.",
									Optional:    true,
								},
								"long": schema.Float64Attribute{
									Description: "Longitude of the CF1 Site.",
									Optional:    true,
								},
								"name": schema.StringAttribute{
									Description: "Name of nearest town, city, or village.",
									Optional:    true,
								},
							},
						},
						"modified_on": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "A human-provided description of the CF1 Site.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A human-provided name describing the CF1 Site that should be unique within the account.",
				Optional:    true,
			},
			"location": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"lat": schema.Float64Attribute{
						Description: "Latitude of the CF1 Site.",
						Optional:    true,
					},
					"long": schema.Float64Attribute{
						Description: "Longitude of the CF1 Site.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of nearest town, city, or village.",
						Optional:    true,
					},
				},
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *MagicTransitCf1SiteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicTransitCf1SiteResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

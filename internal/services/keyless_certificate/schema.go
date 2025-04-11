// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*KeylessCertificateResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Keyless certificate identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"certificate": schema.StringAttribute{
				Description:   "The zone's SSL certificate or SSL certificate and intermediate(s).",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"bundle_method": schema.StringAttribute{
				Description: "A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it.\nAvailable values: \"ubiquitous\", \"optimal\", \"force\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ubiquitous",
						"optimal",
						"force",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Default:       stringdefault.StaticString("ubiquitous"),
			},
			"host": schema.StringAttribute{
				Description: "The keyless SSL name.",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description:        "Whether or not the Keyless SSL is on or off.",
				Optional:           true,
				DeprecationMessage: "This attribute is deprecated.",
			},
			"name": schema.StringAttribute{
				Description: "The keyless SSL name.",
				Optional:    true,
			},
			"port": schema.Float64Attribute{
				Description: "The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.",
				Computed:    true,
				Optional:    true,
				Default:     float64default.StaticFloat64(24008),
			},
			"tunnel": schema.SingleNestedAttribute{
				Description: "Configuration for using Keyless SSL through a Cloudflare Tunnel",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[KeylessCertificateTunnelModel](ctx),
				Attributes: map[string]schema.Attribute{
					"private_ip": schema.StringAttribute{
						Description: "Private IP of the Key Server Host",
						Required:    true,
					},
					"vnet_id": schema.StringAttribute{
						Description: "Cloudflare Tunnel Virtual Network ID",
						Required:    true,
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "When the Keyless SSL was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the Keyless SSL was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"status": schema.StringAttribute{
				Description: "Status of the Keyless SSL.\nAvailable values: \"active\", \"deleted\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "deleted"),
				},
			},
			"permissions": schema.ListAttribute{
				Description: "Available permissions for the Keyless SSL for the current user requesting the item.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (r *KeylessCertificateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *KeylessCertificateResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

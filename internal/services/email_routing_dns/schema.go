// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*EmailRoutingDNSResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"created": schema.StringAttribute{
				Description: "The date and time the settings have been created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"enabled": schema.BoolAttribute{
				Description: "State of the zone settings for Email Routing.",
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the settings have been modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Domain of your zone.",
				Computed:    true,
			},
			"skip_wizard": schema.BoolAttribute{
				Description: "Flag to check if the user skipped the configuration wizard.",
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"status": schema.StringAttribute{
				Description: "Show the state of your account, and the type or configuration error.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ready",
						"unconfigured",
						"misconfigured",
						"misconfigured/locked",
						"unlocked",
					),
				},
			},
			"tag": schema.StringAttribute{
				Description: "Email Routing settings tag. (Deprecated, replaced by Email Routing settings identifier)",
				Computed:    true,
			},
		},
	}
}

func (r *EmailRoutingDNSResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmailRoutingDNSResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

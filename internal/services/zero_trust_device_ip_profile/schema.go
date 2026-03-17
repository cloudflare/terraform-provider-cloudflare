// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_ip_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDeviceIPProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the Device IP profile.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"match": schema.StringAttribute{
				Description: `The wirefilter expression to match registrations. Available values: "identity.name", "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.saml_attributes".`,
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for the Device IP profile.",
				Required:    true,
			},
			"precedence": schema.Int64Attribute{
				Description: "The precedence of the Device IP profile. Lower values indicate higher precedence. Device IP profile will be evaluated in ascending order of this field.",
				Required:    true,
			},
			"subnet_id": schema.StringAttribute{
				Description: "The ID of the Subnet.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "An optional description of the Device IP profile.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the Device IP profile will be applied to matching devices.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"created_at": schema.StringAttribute{
				Description: "The RFC3339Nano timestamp when the Device IP profile was created.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "The RFC3339Nano timestamp when the Device IP profile was last updated.",
				Computed:    true,
			},
		},
	}
}

func (r *ZeroTrustDeviceIPProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDeviceIPProfileResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_address

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*EmailRoutingAddressResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Email Routing Addresses Read",
				"Email Routing Addresses Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Destination address identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"email": schema.StringAttribute{
				Description:   "The contact email address of the user.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Description: "Destination address status. Non-admin callers may only set verified addresses back to unverified; setting to verified requires admin privileges.\nAvailable values: \"unverified\", \"verified\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("unverified", "verified"),
				},
			},
			"created": schema.StringAttribute{
				Description: "The date and time the destination address has been created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the destination address was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"tag": schema.StringAttribute{
				Description:        "Destination address tag. (Deprecated, replaced by destination address identifier)",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
			},
			"verified": schema.StringAttribute{
				Description: "The date and time the destination address has been verified. Null means not verified yet.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *EmailRoutingAddressResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmailRoutingAddressResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

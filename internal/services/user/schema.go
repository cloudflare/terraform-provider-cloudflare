// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*UserResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier of the user.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"country": schema.StringAttribute{
				Description: "The country in which the user lives.",
				Optional:    true,
			},
			"first_name": schema.StringAttribute{
				Description: "User's first name",
				Optional:    true,
			},
			"last_name": schema.StringAttribute{
				Description: "User's last name",
				Optional:    true,
			},
			"telephone": schema.StringAttribute{
				Description: "User's telephone number",
				Optional:    true,
			},
			"zipcode": schema.StringAttribute{
				Description: "The zipcode or postal code where the user lives.",
				Optional:    true,
			},
			"has_business_zones": schema.BoolAttribute{
				Description: "Indicates whether user has any business zones",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"has_enterprise_zones": schema.BoolAttribute{
				Description: "Indicates whether user has any enterprise zones",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"has_pro_zones": schema.BoolAttribute{
				Description: "Indicates whether user has any pro zones",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"suspended": schema.BoolAttribute{
				Description: "Indicates whether user has been suspended",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"two_factor_authentication_enabled": schema.BoolAttribute{
				Description: "Indicates whether two-factor authentication is enabled for the user account. Does not apply to API authentication.",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"two_factor_authentication_locked": schema.BoolAttribute{
				Description: "Indicates whether two-factor authentication is required by one of the accounts that the user is a member of.",
				Computed:    true,
				Default:     booldefault.StaticBool(false),
			},
			"betas": schema.ListAttribute{
				Description: "Lists the betas that the user is participating in.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"organizations": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[UserOrganizationsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Organization name.",
							Computed:    true,
						},
						"permissions": schema.ListAttribute{
							Description: "Access permissions for this User.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"roles": schema.ListAttribute{
							Description: "List of roles that a user has within an organization.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"status": schema.StringAttribute{
							Description: "Whether the user is a member of the organization or has an invitation pending.\nAvailable values: \"member\", \"invited\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("member", "invited"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *UserResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
